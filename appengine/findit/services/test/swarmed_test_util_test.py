# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import json
import mock
import time

from dto import swarming_task_error
from dto.swarming_task_error import SwarmingTaskError
from dto.test_location import TestLocation
from infra_api_clients.swarming import swarming_util
from libs import analysis_status
from libs.test_results import test_results_util
from libs.test_results.gtest_test_results import GtestTestResults
from libs.test_results.blink_web_test_results import BlinkWebTestResults
from libs.test_results.resultdb_test_results import ResultDBTestResults
from go.chromium.org.luci.resultdb.proto.v1 import test_result_pb2
from model.wf_swarming_task import WfSwarmingTask
from services import constants
from services import resultdb
from services import swarmed_test_util
from waterfall import waterfall_config
from waterfall.test import wf_testcase

_GTEST_RESULTS = GtestTestResults(None)


class _MockedTask(object):

  def __init__(self, task_id):
    self.task_id = task_id


class SwarmedTestUtilTest(wf_testcase.WaterfallTestCase):

  @mock.patch.object(
      swarmed_test_util, 'GetTestResultForSwarmingTask', return_value={})
  def testGetTestLocationNoTestLocations(self, _):
    self.assertIsNone(swarmed_test_util.GetTestLocation('task', 'test'))

  @mock.patch.object(
      GtestTestResults, 'IsTestResultsInExpectedFormat', return_value=True)
  @mock.patch.object(swarmed_test_util, 'GetTestResultForSwarmingTask')
  def testGetTestLocation(self, mock_get_isolated_output, _):
    test_name = 'test'
    expected_test_location = {
        'line': 123,
        'file': '/path/to/test_file.cc',
    }
    mock_get_isolated_output.return_value = GtestTestResults(
        {'test_locations': {
            test_name: expected_test_location,
        }})

    self.assertEqual(
        TestLocation.FromSerializable(expected_test_location),
        swarmed_test_util.GetTestLocation('task', test_name))

  def testGetTaskIdFromSwarmingTaskEntity(self):
    swarming_task = WfSwarmingTask.Create('m', 'b', 123, 's')
    swarming_task.task_id = 'task_id'
    swarming_task.put()

    self.assertEqual(
        'task_id',
        swarmed_test_util.GetTaskIdFromSwarmingTaskEntity(
            swarming_task.key.urlsafe()))

  def testGetTaskIdFromSwarmingTaskEntityNoTask(self):
    swarming_task = WfSwarmingTask.Create('m', 'b', 200, 's')
    swarming_task.put()
    key = swarming_task.key.urlsafe()
    swarming_task.key.delete()
    with self.assertRaises(Exception):
      swarmed_test_util.GetTaskIdFromSwarmingTaskEntity(key)

  @mock.patch.object(
      waterfall_config,
      'GetSwarmingSettings',
      return_value={
          'get_swarming_task_id_wait_seconds': 0,
          'get_swarming_task_id_timeout_seconds': -1
      })
  def testGetTaskIdFromSwarmingTaskEntityTimeOut(self, _):
    swarming_task = WfSwarmingTask.Create('m', 'b', 123, 's')
    swarming_task.put()
    with self.assertRaises(Exception):
      swarmed_test_util.GetTaskIdFromSwarmingTaskEntity(
          swarming_task.key.urlsafe())

  def testWaitingForTheTaskId(self):
    master_name = 'm'
    builder_name = 'b'
    build_number = 1
    step_name = 's'

    swarming_task = WfSwarmingTask.Create(master_name, builder_name,
                                          build_number, step_name)
    swarming_task.status = analysis_status.PENDING
    swarming_task.put()

    def MockedSleep(*_):
      swarming_task = WfSwarmingTask.Get(master_name, builder_name,
                                         build_number, step_name)
      self.assertEqual(analysis_status.PENDING, swarming_task.status)
      swarming_task.status = analysis_status.RUNNING
      swarming_task.task_id = 'task_id'
      swarming_task.put()

    self.mock(time, 'sleep', MockedSleep)

    self.assertEqual(
        'task_id',
        swarmed_test_util.GetTaskIdFromSwarmingTaskEntity(
            swarming_task.key.urlsafe()))

  @mock.patch.object(swarming_util, 'GetSwarmingTaskResultById')
  def testGetSwarmingTaskDataAndResultNoData(self, mock_fn):
    error = {'code': 1, 'message': 'error'}
    mock_fn.return_value = (None, error)
    self.assertEqual(
        (None, None, SwarmingTaskError.FromSerializable(error)),
        swarmed_test_util.GetSwarmingTaskDataAndResult('task_id', None))

  @mock.patch.object(swarming_util, 'GetSwarmingTaskResultById')
  def testGetSwarmingTaskDataAndResultFailedState(self, mock_fn):
    data = {'state': 'BOT_DIED', 'outputs_ref': 'outputs_ref'}
    mock_fn.return_value = (data, None)
    error = SwarmingTaskError.FromSerializable({
        'code': swarming_task_error.BOT_DIED,
        'message': 'BOT_DIED'
    })
    self.assertEqual(
        (data, None, error),
        swarmed_test_util.GetSwarmingTaskDataAndResult('task_id', None))

  @mock.patch.object(swarming_util, 'GetSwarmingTaskResultById')
  def testGetSwarmingTaskDataAndResultRunning(self, mock_fn):
    data = {'state': constants.STATE_RUNNING, 'outputs_ref': 'outputs_ref'}
    mock_fn.return_value = (data, None)
    self.assertEqual(
        (data, None, None),
        swarmed_test_util.GetSwarmingTaskDataAndResult('task_id', None))

  @mock.patch.object(
      swarming_util,
      'GetSwarmingTaskResultById',
      return_value=({
          'state': constants.STATE_COMPLETED
      }, None))
  @mock.patch.object(
      swarming_util,
      'GetInvocationNameForSwarmingTask',
      return_value="inv_name")
  @mock.patch.object(resultdb, 'query_resultdb')
  def testGetTestResultForSwarmingTaskWithResultDBEnabled(
      self, resultdb_mock, *_):
    resultdb_mock.return_value = [
        test_result_pb2.TestResult(
            test_id="ninja://gpu:gl_tests/SharedImageTest.Basic")
    ]
    data, result, error = swarmed_test_util.GetSwarmingTaskDataAndResult(
        'task_id')
    self.assertEqual({'state': constants.STATE_COMPLETED}, data)
    self.assertIsInstance(result, ResultDBTestResults)
    self.assertIsNone(error)

  @mock.patch.object(
      swarmed_test_util,
      'GetSwarmingTaskDataAndResult',
      return_value=('data', 'content', 'error'))
  def testGetTestResultForSwarmingTask(self, mock_fn):
    self.assertEqual(
        'content',
        swarmed_test_util.GetTestResultForSwarmingTask('task_id', None))
    mock_fn.assert_called_once_with('task_id', None)
