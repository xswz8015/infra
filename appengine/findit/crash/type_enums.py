# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.


class CallStackFormatType(object):
  JAVA = 1
  SYZYASAN = 2
  DEFAULT = 3


class CallStackLanguageType(object):
  CPP = 1
  JAVA = 2


class CrashClient(object):
  FRACAS = 'fracas'
  CRACAS = 'cracas'
  CLUSTERFUZZ = 'clusterfuzz'
