# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
"""This module is for swarming related logics that only for Findit.
"""

from collections import defaultdict
import copy
import json
import logging

from google.appengine.api import app_identity

from gae_libs import token
from infra_api_clients.swarming import swarming_util
from infra_api_clients.swarming.swarming_task_request import CIPDPackages
from libs import time_util
from libs.list_of_basestring import ListOfBasestring
from waterfall import waterfall_config

_PUBSUB_TOPIC = 'projects/%s/topics/swarming'


def SwarmingHost():
  return waterfall_config.GetSwarmingSettings().get('server_host')


def GetSwarmingTaskUrl(task_id):
  return 'https://{}/task?id={}'.format(SwarmingHost(), task_id)


def GetReferredSwarmingTaskRequestInfo(master_name, builder_name, build_number,
                                       step_name, http_client):
  """Gets referred swarming task request.

  Returns:
    (ref_task_id, ref_request): Referred swarming task id and request.
  """
  swarming_task_items = ListSwarmingTasksDataByTags(http_client, builder_name,
                                                    build_number, step_name)

  if not swarming_task_items:
    raise Exception('Cannot find referred swarming task for %s/%s/%d/%s' %
                    (master_name, builder_name, build_number, step_name))

  ref_task_id = swarming_task_items[0].task_id
  ref_request = swarming_util.GetSwarmingTaskRequest(SwarmingHost(),
                                                     ref_task_id, http_client)
  return ref_task_id, ref_request


def _UpdateRequestWithPubSubCallback(request, runner_id):
  request.pubsub_topic = _PUBSUB_TOPIC % app_identity.get_application_id()
  request.pubsub_auth_token = token.GenerateAuthToken('pubsub', 'swarming',
                                                      runner_id)
  request.pubsub_userdata = json.dumps({'runner_id': runner_id})


def _IsTestFilter(arg):
  test_filter_args = [
      '--gtest_filter', '--test-launcher-filter-file',
      '--isolated-script-test-filter'
  ]
  return any(
      arg.startswith(test_filter_arg) for test_filter_arg in test_filter_args)


def _SetTestFilters(args, tests, iterations):
  """Given a test command, sets parameters to run only the specified tests.

  Args:
    args (ListOfBasestring): A list of arguments to pass to the test command.
    tests (list): A list of tests to run.
    iterations (int): Number of iterations each test should run.

  Returns:
    (list):  A modified list of arguments to pass to the test command.
  """
  # Remove existing test filter first.
  res = ListOfBasestring.FromSerializable(
      [a for a in args if not _IsTestFilter(a)])

  res.append('--isolated-script-test-filter=%s' % '::'.join(tests))

  res.append('--isolated-script-test-repeat=%s' % iterations)

  res.append('--isolated-script-test-launcher-retry-limit=0')

  # Also rerun disabled tests. Scenario: the test was disabled before Findit
  # runs any analysis. One possible case:
  #   1. A gtest became flaky on CQ, but Findit was not automatically
  #      triggered to run any analysis because:
  #      * the test is not flaky enough
  #   2. The test got disabled, but no culprit was identified.
  #   3. Some developer starts the investigation and requests Findit to
  #      analyze the flaky test.
  #   4. Findit picks the latest Waterfall build of the matching configuration
  #      for the CQ build in which the flaky test is found.
  #   5. In the picked Waterfall build, the test is already disabled.
  #
  # Note: test runner on Android ignores this flag because it is not supported
  # yet even though it exists.
  res.append('--isolated-script-test-also-run-disabled-tests')
  return res


def CreateNewSwarmingTaskRequestTemplate(runner_id, ref_task_id, ref_request,
                                         master_name, builder_name, step_name,
                                         tests, iterations):
  """Returns a SwarmingTaskRequest instance to run the given tests only.

  Args:
    ref_task_id (str): Id of the referred swarming task.
    ref_request (SwarmingTaskRequest): Request of the referred swarming task.
    master_name (str): The name of the main waterfall master for a build.
    builder_name (str): The name of the main waterfall builder for a build.
    step_name (str): The name of a failed step in the build.
    tests (list): A list of tests in the step that we want to rerun in task.
    iterations (int): Number of iterations each test should run.
  """
  # Make a copy of the referred request and drop or overwrite some fields.
  new_request = copy.deepcopy(ref_request)
  new_request.name = 'findit/ref_task_id/%s/%s' % (
      ref_task_id, time_util.GetUTCNow().strftime('%Y-%m-%d %H:%M:%S %f'))
  new_request.parent_task_id = ''
  new_request.user = ''

  _UpdateRequestWithPubSubCallback(new_request, runner_id)

  # To force a fresh re-run and ignore cached result of any equivalent run.
  new_request.properties.idempotent = False

  # Set the gtest_filter to run the given tests only.
  # Only one of command or extra_args will be populated.
  if len(new_request.properties.command) > 0:
    assert len(new_request.properties.extra_args) == 0
    new_request.properties.command = _SetTestFilters(
        new_request.properties.command, tests, iterations)
  else:
    new_request.properties.extra_args = _SetTestFilters(
        new_request.properties.extra_args, tests, iterations)

  # Remove the env setting for sharding.
  sharding_settings = ['GTEST_SHARD_INDEX', 'GTEST_TOTAL_SHARDS']
  new_request.properties.env = [
      e for e in new_request.properties.env if e['key'] not in sharding_settings
  ]

  # Remove environment prefixes, caches, and CIPD packages used by the
  # chromium.tests pool task template. Otherwise, swarming complains about a
  # collision.
  # This should be kept in sync with the task template at
  # http://shortn/_rOkMZ6ANDo.
  # See https://crbug.com/1128541#c14 for more details.
  pool_env_prefixes = ['PATH', 'VPYTHON_VIRTUALENV_ROOT']
  new_request.properties.env_prefixes = [
      e for e in new_request.properties.env_prefixes
      if e['key'] not in pool_env_prefixes
  ]
  new_request.properties.caches = [
      c for c in new_request.properties.caches
      if not c['name'].startswith('task_template')
  ]
  filtered_packages = CIPDPackages()
  filtered_packages.extend([
      p for p in new_request.properties.cipd_input.packages
      if not p.path.startswith('.task_template')
  ])
  new_request.properties.cipd_input.packages = filtered_packages

  # Reset tags for searching and monitoring.
  ref_name = swarming_util.GetTagValue(ref_request.tags, 'name')
  new_request.tags = ListOfBasestring()
  new_request.tags.append('ref_master:%s' % master_name)
  new_request.tags.append('ref_buildername:%s' % builder_name)

  new_request.tags.append('ref_stepname:%s' % step_name)
  new_request.tags.append('ref_name:%s' % ref_name)
  new_request.tags.extend(
      ['findit:1', 'project:Chromium', 'purpose:post-commit'])

  # Use a priority much lower than CQ for now (CQ's priority is 30).
  # Later we might use a higher priority -- a lower value here.
  # Note: the smaller value, the higher priority.
  swarming_settings = waterfall_config.GetSwarmingSettings()
  request_expiration_hours = swarming_settings.get('request_expiration_hours')
  new_request.priority = str(
      max(100, swarming_settings.get('default_request_priority')))
  new_request.expiration_secs = str(request_expiration_hours * 60 * 60)

  return new_request


def ListSwarmingTasksDataByTags(http_client,
                                builder_name,
                                build_number,
                                step_name=None,
                                additional_tag_filters=None):
  """Downloads tasks data from swarming server.

  Args:
    http_client (RetryHttpClient): The http client to send HTTPs requests.
    builder_name (str): Value of the buildername tag.
    build_number (int): Value of the buildnumber tag.
    step_name (str): Value of the stepname tag.
    additional_tag_filters (dict): Additional tags.

  Returns:
    (list):  A list of SwarmingTaskData for all tasks with queried tags.
  """
  tag_filters = {
      # Findit v1 only operates on builders in the chromium.ci bucket, so
      # hardcoding this here is fine.
      'project': 'chromium',
      'bucket': 'ci',
      'buildername': builder_name,
      'buildnumber': build_number
  }
  if step_name:
    tag_filters['stepname'] = step_name

  additional_tag_filters = additional_tag_filters or {}
  tag_filters.update(additional_tag_filters)

  return swarming_util.ListTasks(SwarmingHost(), tag_filters, http_client)


def GetNeededIsolatedDataFromTaskResults(task_results, only_failure):
  needed_isolated_data = defaultdict(list)
  for item in task_results:
    swarming_step_name = item.tags.get(
        'stepname')[0] if 'stepname' in item.tags else None

    if not item.outputs_ref or not swarming_step_name:
      # Task might time out and no outputs_ref was saved.
      continue

    if only_failure:
      if item.non_internal_failure:
        isolated_data = swarming_util.GenerateIsolatedData(item.outputs_ref)
        needed_isolated_data[swarming_step_name].append(isolated_data)
    else:
      isolated_data = swarming_util.GenerateIsolatedData(item.outputs_ref)
      needed_isolated_data[swarming_step_name].append(isolated_data)
  return needed_isolated_data


def GetFailedSwarmingTasksGroupedByStepName(task_results):
  """ Return a dictionary of failed swarming task id grouped by step name
  Args:
    task_results (list): An array of SwarmingTaskData to check.
  """
  step_to_task_ids_map = defaultdict(list)
  for item in task_results:
    swarming_step_name = item.tags.get(
        'stepname')[0] if 'stepname' in item.tags else None
    if not swarming_step_name:
      continue
    if item.non_internal_failure:
      step_to_task_ids_map[swarming_step_name].append(item.task_id)
  return step_to_task_ids_map


def GetIsolatedDataForStep(builder_name,
                           build_number,
                           step_name,
                           http_client,
                           only_failure=True):
  """Returns the isolated data for a specific step.

  Args:
    master_name (str): Value of the master tag.
    builder_name (str): Value of the buildername tag.
    build_number (int): Value of the buildnumber tag.
    step_name (str): Value of the stepname tag.
    http_client (RetryHttpClient): The http client to send HTTPs requests.
    only_failure (bool): A flag to determine if only failure info is needed.
  """
  step_isolated_data = []
  items = ListSwarmingTasksDataByTags(http_client, builder_name, build_number,
                                      step_name)
  if not items:
    return step_isolated_data

  step_isolated_data = GetNeededIsolatedDataFromTaskResults(items, only_failure)
  return step_isolated_data[step_name]


def GetIsolatedDataForFailedStepsInABuild(builder_name, build_number,
                                          failed_steps, http_client):
  """Gets the isolated data for failed steps for a build.

  Args:
    builder_name (str): The name of the main waterfall builder.
    build_number (int): The build number to retrieve the isolated sha of.
    failed_steps (TestFailedSteps): A dict of failed steps.

  Returns:
    build_isolated_data(dict): A dict of isolated data of failed steps in a
      build.
  """
  items = ListSwarmingTasksDataByTags(http_client, builder_name, build_number)
  if not items:
    return {}

  isolated_data = GetNeededIsolatedDataFromTaskResults(items, True)
  build_isolated_data = {}
  for step, data in isolated_data.iteritems():
    if step in failed_steps:
      build_isolated_data[step] = data
  return build_isolated_data


def GetSwarmingTaskIdsForFailedSteps(builder_name, build_number, failed_steps,
                                     http_client):
  items = ListSwarmingTasksDataByTags(http_client, builder_name, build_number)
  if not items:
    return {}
  step_to_swarming_ids_map = GetFailedSwarmingTasksGroupedByStepName(items)
  result = {}
  for step, swarming_ids in step_to_swarming_ids_map.iteritems():
    if step in failed_steps:
      result[step] = swarming_ids
  return result


def GetIsolatedShaForStep(builder_name, build_number, step_name, http_client):
  """Gets the isolated sha of a master/builder/build/step.

  Args:
    builder_name (str): The name of the main waterfall builder.
    build_number (int): The build number to retrieve the isolated sha of.
    step_name (str): The step name to retrieve the isolated sha of.

  Returns:
    (str): The isolated sha pointing to the compiled binaries at the requested
        configuration.
  """
  items = ListSwarmingTasksDataByTags(http_client, builder_name, build_number,
                                      step_name)
  if not items:
    logging.error('Failed to get swarming task data for %s/%s/%s', builder_name,
                  build_number, step_name)
    return None

  # Each task should have the same sha, so only need to read from the first one.
  if items[0].inputs_ref_sha:
    return items[0].inputs_ref_sha

  logging.error('Isolated sha not found for %s/%s/%s', builder_name,
                build_number, step_name)
  return None


def CanFindSwarmingTaskFromBuildForAStep(http_client, builder_name,
                                         build_number, step_name):
  tasks = ListSwarmingTasksDataByTags(http_client, builder_name, build_number,
                                      step_name)
  return len(tasks) > 0
