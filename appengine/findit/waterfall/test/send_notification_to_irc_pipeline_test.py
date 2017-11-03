# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import mock

from pipelines.pipeline_inputs_and_outputs import CLKey
from pipelines.pipeline_inputs_and_outputs import (
    SendNotificationToIrcPipelineInput)
from services import gerrit
from services import irc
from waterfall.send_notification_to_irc_pipeline import (
    SendNotificationToIrcPipeline)
from waterfall.test import wf_testcase


class SendNotificationToIrcPipelineTest(wf_testcase.WaterfallTestCase):

  @mock.patch.object(irc, 'SendMessageToIrc', return_value=True)
  def testSendNotification(self, _):
    repo_name = 'chromium'
    revision = 'rev'
    revert_status = gerrit.CREATED_BY_FINDIT
    submitted = True
    pipeline_input = SendNotificationToIrcPipelineInput(
        cl_key=CLKey(
            repo_name=repo_name.decode('utf-8'),
            revision=revision.decode('utf-8')),
        revert_status=revert_status,
        submitted=submitted)
    pipeline = SendNotificationToIrcPipeline(pipeline_input)
    self.assertTrue(pipeline.run(pipeline_input))
