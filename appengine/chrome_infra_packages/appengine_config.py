# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Custom UI tweaks for components.auth module."""

components_auth_UI_APP_NAME = 'CIPD'

import os

APP_DIR = os.path.dirname(os.path.abspath(__file__))

# AppEngine supplies its own "google" domain. Add our VirtualEnv's "google"
# domain to the list of permissible "google" package paths.
import google
google.__path__.append(os.path.join(
    APP_DIR, 'components', 'third_party', 'protobuf'))

# Assert that "google.protobuf" imports.
import google.protobuf
