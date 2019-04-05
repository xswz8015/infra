// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import {assert} from 'chai';
import {ChopsSignin} from './index.js';

let element;

suite('chops-signin');

beforeEach(() => {
  element = document.createElement('chops-signin');
  document.body.appendChild(element);
});

afterEach(() => {
  document.body.removeChild(element);
});

test('initializes', () => {
  assert.instanceOf(element, ChopsSignin);
});

test('lack of clientId results in error message', async () => {
  await element.updateComplete;
  assert.isDefined(element.errorMsg);
});

test('clientId set, no error message', async () => {
  element.setAttribute('client-id', 'foobar');
  await element.updateComplete;
  assert.isUndefined(element.errorMsg);
});
