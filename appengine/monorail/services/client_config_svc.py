# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is govered by a BSD-style
# license that can be found in the LICENSE file or at
# https://developers.google.com/open-source/licenses/bsd

import base64
import json
import logging
import os
import time
import urllib

from google import protobuf
from google.appengine.api import app_identity
from google.appengine.api import urlfetch

import settings
from framework import framework_constants
from proto import api_clients_config_pb2


CONFIG_FILE_PATH = os.path.join(
    os.path.dirname(os.path.dirname(os.path.realpath(__file__))),
    'testing', 'api_clients.cfg')
MONORAIL_CONFIG_SET = urllib.quote(
    'services/%s' % app_identity.get_application_id(), safe='')
LUCI_CONFIG_URL = (
    'https://luci-config.appspot.com/_ah/api/config/v1/config_sets'
    '/%s/config/api_clients.cfg') % MONORAIL_CONFIG_SET


client_config_svc = None
service_account_map = None


class ClientConfigService(object):
  """The persistence layer for client config data."""

  # One hour
  EXPIRES_IN = 3600

  def __init__(self):
    self.client_configs = None
    self.load_time = 0

  def GetConfigs(self, use_cache=True, cur_time=None):
    """Read client configs."""

    cur_time = cur_time or int(time.time())
    force_load = False
    if not self.client_configs:
      force_load = True
    elif not use_cache:
      force_load = True
    elif cur_time - self.load_time > self.EXPIRES_IN:
      force_load = True

    if force_load:
      if settings.dev_mode or settings.unit_test_mode:
        self._ReadFromLocal()
      else:
        self._ReadFromLuciConfig()
      
    return self.client_configs

  def _ReadFromLocal(self):
    try:
      with open(CONFIG_FILE_PATH, 'r') as f:
        content_text = f.read()
      logging.info('Read client configs from local file.')
      cfg = api_clients_config_pb2.ClientCfg()
      protobuf.text_format.Merge(content_text, cfg)
      self.client_configs = cfg
      self.load_time = int(time.time())
    except Exception as ex:
      logging.exception(
          'Failed to read client configs: %s',
          str(ex))

  def _ReadFromLuciConfig(self):
    try:
      authorization_token, _ = app_identity.get_access_token(
          framework_constants.OAUTH_SCOPE)
      response = urlfetch.fetch(
          LUCI_CONFIG_URL,
          method=urlfetch.GET,
          follow_redirects=False,
          headers={'Content-Type': 'application/json; charset=UTF-8',
                   'Authorization': 'Bearer ' + authorization_token})
      if response.status_code == 200:
        content = json.loads(response.content)
        config_content = content['content']
        content_text = base64.b64decode(config_content)
        logging.info('luci-config content decoded: %r.', content_text)
        cfg = api_clients_config_pb2.ClientCfg()
        protobuf.text_format.Merge(content_text, cfg)
        self.client_configs = cfg
        self.load_time = int(time.time())
      else:
        logging.error('Invalid response from luci-config: %r', response)
    except Exception as ex:
      logging.exception(
          'Failed to retrieve client configs from luci-config: %s',
          str(ex))

  def GetClientIDEmails(self):
    """Get client IDs and Emails."""
    self.GetConfigs(use_cache=True)
    client_ids = [c.client_id for c in self.client_configs.clients]
    client_emails = [c.client_email for c in self.client_configs.clients]
    return client_ids, client_emails

  def GetDisplayNames(self):
    """Get client display names."""
    self.GetConfigs(use_cache=True)
    names_dict = {}
    for client in self.client_configs.clients:
      if client.display_name:
        names_dict[client.client_email] = client.display_name
    return names_dict


def GetClientConfigSvc():
  global client_config_svc
  if client_config_svc is None:
    client_config_svc = ClientConfigService()
  return client_config_svc

def GetServiceAccountMap():
  global service_account_map
  if service_account_map is None:
    service_account_map = GetClientConfigSvc().GetDisplayNames()
  return service_account_map
  