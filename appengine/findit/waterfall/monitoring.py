# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import gae_ts_mon

swarming_tasks = gae_ts_mon.CounterMetric(
    'findit/swarmingtasks', description='Swarming tasks triggered')

issues = gae_ts_mon.CounterMetric(
    'findit/issues', description='Bugs updated with findings')

flakes = gae_ts_mon.CounterMetric(
    'findit/flakes', description='Flakes requested or analyzed')
