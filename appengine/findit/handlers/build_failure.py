# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
import os

from google.appengine.api import users

from base_handler import BaseHandler
from base_handler import Permission
from waterfall import buildbot
from waterfall import build_failure_analysis_pipelines
from waterfall import masters


BUILD_FAILURE_ANALYSIS_TASKQUEUE = 'build-failure-analysis-queue'


class BuildFailure(BaseHandler):
  PERMISSION_LEVEL = Permission.ANYONE

  def _IsPipelineAccessible(self):
    # Pipeline is accessible if the app is run locally during development or if
    # the currently logged-in user is an admin.
    return (os.environ['SERVER_SOFTWARE'].startswith('Development') or
            users.is_current_user_admin())

  def HandleGet(self):
    """Triggers analysis of a build failure on demand and return current result.

    If the final analysis result is available, set cache-control to 1 day to
    avoid overload by unnecessary and frequent query from clients; otherwise
    set cache-control to 5 seconds to allow repeated query.

    Serve HTML page or JSON result as requested.
    """
    url = self.request.get('url', '').strip()
    build_info = buildbot.ParseBuildUrl(url)
    if not build_info:
      return BaseHandler.CreateError(
          'Url "%s" is not pointing to a build.' % url, 501)
    master_name, builder_name, build_number = build_info

    if not masters.MasterIsSupported(master_name):
      return BaseHandler.CreateError(
          'Master "%s" is not supported yet.' % master_name, 501)

    force = self.request.get('force') == '1'
    analysis = build_failure_analysis_pipelines.ScheduleAnalysisIfNeeded(
        master_name, builder_name, build_number, force,
        BUILD_FAILURE_ANALYSIS_TASKQUEUE)

    def FormatDatetime(datetime):
      if not datetime:
        return None
      else:
        return datetime.strftime('%Y-%m-%d %H:%M:%S')

    data = {
        'master_name': analysis.master_name,
        'builder_name': analysis.builder_name,
        'build_number': analysis.build_number,
        'build_url': buildbot.CreateBuildUrl(
            analysis.master_name, analysis.builder_name, analysis.build_number),
        'pipeline_url': analysis.pipeline_url,
        'pipeline_accessible': self._IsPipelineAccessible(),
        'analysis_started': FormatDatetime(analysis.start_time),
        'analysis_updated': FormatDatetime(analysis.updated_time),
        'analysis_completed': analysis.completed,
        'analysis_failed': analysis.failed,
        'analysis_result': analysis.result,
    }

    return {'template': 'build_failure.html', 'data': data}

  def HandlePost(self):  # pragma: no cover
    return self.HandleGet()
