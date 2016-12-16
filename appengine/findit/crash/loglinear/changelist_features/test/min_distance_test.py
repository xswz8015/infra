# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import unittest

from crash.loglinear.changelist_features import min_distance
from crash.results import AnalysisInfo
from crash.results import StackInfo
from crash.results import MatchResult
from crash.results import Result
from crash.stacktrace import StackFrame
from libs.gitiles.change_log import ChangeLog
import libs.math.logarithms as lmath


_MAXIMUM = float(min_distance.DEFAULT_MAXIMUM)

_MOCK_FRAME = StackFrame(0, 'src/', 'func', 'f.cc', 'a/b/src/f.cc', [2],
                         repo_url='https://repo_url')

_DUMMY_CHANGELOG = ChangeLog.FromDict({
    'author_name': 'r@chromium.org',
    'message': 'dummy',
    'committer_email': 'r@chromium.org',
    'commit_position': 175900,
    'author_email': 'r@chromium.org',
    'touched_files': [
        {
            'change_type': 'add',
            'new_path': 'a.cc',
            'old_path': None,
        },
    ],
    'author_time': 'Thu Mar 31 21:24:43 2016',
    'committer_time': 'Thu Mar 31 21:28:39 2016',
    'commit_url':
        'https://repo.test/+/1',
    'code_review_url': 'https://codereview.chromium.org/3281',
    'committer_name': 'example@chromium.org',
    'revision': '1',
    'reverted_revision': None
})


class MinDistanceTest(unittest.TestCase):

  def _GetDummyChangeLog(self):
    return _DUMMY_CHANGELOG

  def _GetDummyReport(self):
    return None

  def _GetMockResult(self, mock_min_distance):
    """Returns a ``Result`` with the desired min_distance."""
    result = Result(self._GetDummyChangeLog(), 'src/')
    result.file_to_analysis_info = {
        'file': AnalysisInfo(
            min_distance=mock_min_distance,
            min_distance_frame=_MOCK_FRAME)
    }
    return result

  def testMinDistanceFeatureNone(self):
    """Test that the feature returns log(0) when there are no frames."""
    report = self._GetDummyReport()
    result = Result(self._GetDummyChangeLog(), 'src/')
    self.assertEqual(lmath.LOG_ZERO,
        min_distance.MinDistanceFeature()(report)(result).value)

  def testMinDistanceFeatureIsZero(self):
    """Test that the feature returns log(1) when the min_distance is 0."""
    report = self._GetDummyReport()
    result = self._GetMockResult(0.)
    self.assertEqual(lmath.LOG_ONE,
        min_distance.MinDistanceFeature()(report)(result).value)

  def testMinDistanceFeatureMiddling(self):
    """Test that the feature returns middling scores for middling distances."""
    report = self._GetDummyReport()
    result = self._GetMockResult(42.)
    self.assertEqual(
        lmath.log((_MAXIMUM - 42.) / _MAXIMUM),
        min_distance.MinDistanceFeature()(report)(result).value)

  def testMinDistanceFeatureIsOverMax(self):
    """Test that we return log(0) when the min_distance is too large."""
    report = self._GetDummyReport()
    result = self._GetMockResult(_MAXIMUM + 1)
    self.assertEqual(lmath.LOG_ZERO,
        min_distance.MinDistanceFeature()(report)(result).value)

    result = self._GetMockResult(42.)
    self.assertEqual(lmath.LOG_ZERO,
        min_distance.MinDistanceFeature(10.)(report)(result).value)

  def testMinDistanceChangedFiles(self):
    result = MatchResult(self._GetDummyChangeLog(), 'src/')
    frame = StackFrame(0, 'src/', 'func', 'a.cc', 'src/a.cc', [7],
                       repo_url='https://repo_url')
    result.file_to_stack_infos = {
        'a.cc': [StackInfo(frame, 0)]
    }
    result.file_to_analysis_info = {
        'a.cc': AnalysisInfo(min_distance=0, min_distance_frame=frame)
    }
    changed_files = min_distance.MinDistanceFeature()._ChangedFiles(result)
    self.assertListEqual(
        [changed_file.ToDict() for changed_file in changed_files],
        [{'info': 'Minimum distance (LOC) 0, frame #0',
          'blame_url': 'https://repo_url/+blame/1/a.cc#7',
          'file': 'a.cc'}])
