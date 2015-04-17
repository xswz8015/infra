# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import Queue
import logging
import threading
import time

import httplib2
import oauth2client

from infra.libs.authentication import get_authenticated_http
from infra.libs.event_mon.log_request_lite_pb2 import LogRequestLite
from infra.libs.event_mon.chrome_infra_log_pb2 import ChromeInfraEvent

def time_ms():
  """Return current timestamp in milliseconds."""
  return int(1000 * time.time())


class _Router(object):
  """Route events to the right destination.

  This object is meant to be a singleton, and is not part of the API.

  Usage:
  router = _Router()
  event = ChromeInfraEvent.LogEventLite(...)
  ... fill in event ...
  router.push_event(event)
  """
  def __init__(self, cache, endpoint=None):
    # cache is defined in config.py. Passed as a parameter to avoid
    # a circular import.

    # endpoint == None means 'dry run'. No data is sent.
    self.endpoint = endpoint
    self.http = httplib2.Http()
    self.cache = cache

    if self.endpoint and self.cache['service_account_creds']:
      logging.debug('Activating OAuth2 authentication.')
      self.http = get_authenticated_http(
        self.cache['service_account_creds'],
        service_accounts_creds_root=self.cache['service_accounts_creds_root'],
        scope='https://www.googleapis.com/auth/cclog'
      )

    self.event_queue = Queue.Queue()
    self._thread = threading.Thread(target=self._router)
    self._thread.daemon = True
    logging.debug('event_mon: starting router thread')
    self._thread.start()

  def _router(self):
    while(True):  # pragma: no branch
      events = self.event_queue.get()
      if events is None:
        break

      # Set this time at the very last moment
      events.request_time_ms = time_ms()
      if self.endpoint:  # pragma: no cover
        logging.info('event_mon: POSTing events to %s', self.endpoint)
        response, _ = self.http.request(
          uri=self.endpoint,
          method='POST',
          headers={'Content-Type': 'application/octet-stream'},
          body=events.SerializeToString()
        )

        if response.status != 200:
          # TODO(pgervais): implement retry / local storage when this
          # happens.
          logging.error('failed to POST data to %s', self.endpoint)
          logging.error('data: %s', str(events)[:1000])
          logging.error(response)
      else:
        infra_events = [str(ChromeInfraEvent.FromString(
          ev.source_extension)) for ev in events.log_event]
        logging.info('Fake post request. Sending:\n%s',
                     '\n'.join(infra_events))

  def close(self, timeout=None):
    """
    Returns:
      success (bool): True if everything went well. Otherwise, there is no
        guarantee that all events have been properly sent to the remote.
    """
    timeout = timeout or 5
    logging.debug('event_mon: trying to close')
    self.event_queue.put(None)
    self._thread.join(timeout)
    # If the thread is still alive at this point, we can't but wait for a call
    # to sys.exit. Since we expect this function to be called at the end of the
    # program, it should come soon.
    success = not self._thread.is_alive()
    if success:
      logging.debug('event_mon: successfully closed.')
    else:  # pragma: no cover
      logging.debug('event_mon: timeout waiting for thread to finish.')
    return success

  def push_event(self, event):
    """Enqueue event to push to the collection service.

    This method offers no guarantee on return that the event have been pushed
    externally, as some buffering can take place.

    Args:
      event (LogRequestLite.LogEventLite): one single event.
    Returns:
      success (bool): False if an error happened. True means 'event accepted',
        but NOT 'event successfully pushed to the remote'.
    """
    if not isinstance(event, LogRequestLite.LogEventLite):
      logging.error('Invalid type for "event": %s (should be LogEventLite)'
                    % str(type(event)))
      return False

    # TODO(pgervais): implement batching.
    request_p = LogRequestLite()
    request_p.log_source_name = 'CHROME_INFRA'
    request_p.log_event.extend((event,))  # copies the protobuf
    self.event_queue.put(request_p)
    return True
