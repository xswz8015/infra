// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cv

import (
	"context"
	"testing"

	cvv0 "go.chromium.org/luci/cv/api/v0"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRun(t *testing.T) {
	t.Parallel()

	Convey("Get run", t, func() {
		rID := "projects/chromium/runs/run-id"
		runs := map[string]*cvv0.Run{
			rID: {},
		}
		ctx := UseFakeClient(context.Background(), runs)
		client, err := NewClient(ctx, "host")
		So(err, ShouldBeNil)
		req := &cvv0.GetRunRequest{
			Id: rID,
		}
		_, err = client.GetRun(ctx, req)
		So(err, ShouldBeNil)
	})
}
