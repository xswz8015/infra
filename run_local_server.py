#!/bin/bash
# Copyright (c) 2010 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
../google_appengine/dev_appserver.py . --debug --use_sqlite --require_indexes --datastore_path=./tmp.db
