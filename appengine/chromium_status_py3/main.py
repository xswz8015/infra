# Copyright (c) 2011 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import datetime
import google.cloud.logging

from flask import Flask, make_response, render_template, request
from google.cloud import ndb
from werkzeug.routing import BaseConverter

from appengine_module.chromium_status.static_blobs import ServeHandler
from appengine_module.chromium_status.status import AllStatusPage, MainPage
from appengine_module.chromium_status import auth
from appengine_module.chromium_status import base_page
from appengine_module.chromium_status import git_lkgr
from appengine_module.chromium_status import status
from appengine_module.chromium_status import utils

client = ndb.Client()


class FlaskApp(ndb.Model):  # pylint: disable=W0232
  secret_key = ndb.StringProperty()


# RegexConverter is to support regrex for app.route
class RegexConverter(BaseConverter):

  def __init__(self, url_map, *items):
    super(RegexConverter, self).__init__(url_map)
    self.regex = items[0]


# ndb_wsgi_middleware is for passing ndb context to requests
def ndb_wsgi_middleware(wsgi_app):

  def middleware(environ, start_response):
    with client.context():
      return wsgi_app(environ, start_response)

  return middleware


# If `entrypoint` is not defined in app.yaml, App Engine will look for an app
# called `app` in `main.py`.
app = Flask(__name__)
app.wsgi_app = ndb_wsgi_middleware(app.wsgi_app)  # Wrap the app in middleware.
app.url_map.converters['regex'] = RegexConverter

with client.context():
  app.secret_key = FlaskApp.query().get().secret_key


@app.route('/', methods=['POST', 'GET'])
def main_page():
  if request.method == 'GET':
    return MainPage().get()
  return MainPage().post()


@app.route('/code')
def auth_code():
  return auth.AuthHandler().handle_code()


@app.route('/allstatus')
def all_status():
  return AllStatusPage().get()


@app.route('/current', methods=['POST', 'GET'])
def current():
  if request.method == 'GET':
    return status.CurrentPage().get()
  return status.CurrentPage().post()


@app.route('/revisions', methods=['POST', 'GET'])
@app.route('/commits', methods=['POST', 'GET'])
def commits():
  if request.method == 'GET':
    return git_lkgr.Commits().get()
  git_lkgr.Commits().post()


@app.route('/lkgr')
@app.route('/git-lkgr')
def lkgr():
  return git_lkgr.LastKnownGoodRevisionGIT().get()


@app.route('/status', methods=['POST', 'GET'])
def status_page():
  if request.method == 'GET':
    return status.StatusPage().get()
  return status.StatusPage().post()


@app.route('/status_viewer')
def status_viewer_page():
  return status.StatusViewerPage().get()


@app.route('/<regex("([^/]+\.(?:gif|png|jpg|ico))"):resource>')
def get_resource(resource):
  return ServeHandler().get(resource)


@app.route('/_ah/warmup')
def warmup():
  # Handle your warmup logic here, e.g. set up a database connection pool.
  return '', 200, {}

# Set up cloud logging
logging_client = google.cloud.logging.Client()
logging_client.get_default_handler()
logging_client.setup_logging()

auth.AuthHandler.bootstrap()
base_page.bootstrap()
git_lkgr.bootstrap()
utils.bootstrap()
status.bootstrap()
