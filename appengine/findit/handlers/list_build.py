# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from base_handler import BaseHandler
from base_handler import Permission


class ListBuild(BaseHandler):
  PERMISSION_LEVEL = Permission.ANYONE

  def HandleGet(self):  #pylint: disable=R0201
    """Shows a list of analyzed builds in HTML page."""
    return BaseHandler.CreateError('Not implemented yet!', 501)
