// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package reviewer

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/common/proto"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/signing"
	"go.chromium.org/luci/server/auth/signing/signingtest"

	"infra/appengine/rubber-stamper/config"
	"infra/appengine/rubber-stamper/internal/gerrit"
	"infra/appengine/rubber-stamper/tasks/taskspb"
)

func TestReviewChange(t *testing.T) {
	Convey("review change", t, func() {
		ctx := memory.Use(context.Background())
		cfg := &config.Config{
			HostConfigs: map[string]*config.HostConfig{
				"test-host": {
					BenignFilePattern: &config.BenignFilePattern{
						FileExtensionMap: map[string]*config.Paths{
							"": {
								Paths: []string{"a/x", "a/q/y"},
							},
							".txt": {
								Paths: []string{"a/b.txt", "a/c.txt", "a/e/*", "a/f*"},
							},
						},
					},
				},
			},
		}
		ctx = gerrit.Setup(ctx)
		So(config.SetTestConfig(ctx, cfg), ShouldBeNil)
		ctx = auth.ModifyConfig(ctx, func(cfg auth.Config) auth.Config {
			cfg.Signer = signingtest.NewSigner(&signing.ServiceInfo{
				ServiceAccountName: "srv-account@example.com",
			})
			return cfg
		})

		ctl := gomock.NewController(t)
		defer ctl.Finish()
		gerritMock := gerritpb.NewMockGerritClient(ctl)
		clientMap := map[string]gerrit.Client{
			"test-host-review.googlesource.com": gerritMock,
		}
		ctx = gerrit.SetTestClientFactory(ctx, clientMap)

		t := &taskspb.ChangeReviewTask{
			Host:     "test-host",
			Number:   12345,
			Revision: "123abc",
		}
		Convey("invalid BenignFileChange", func() {
			gerritMock.EXPECT().ListFiles(gomock.Any(), proto.MatcherEqual(&gerritpb.ListFilesRequest{
				Number:     t.Number,
				RevisionId: t.Revision,
			})).Return(&gerritpb.ListFilesResponse{
				Files: map[string]*gerritpb.FileInfo{
					"a/d.txt":     nil,
					"a/p":         nil,
					"a/e/p/p.txt": nil,
					"a/f/z.txt":   nil,
					"a/fz.txt":    nil,
				},
			}, nil)
			gerritMock.EXPECT().SetReview(gomock.Any(), proto.MatcherEqual(&gerritpb.SetReviewRequest{
				Number:     t.Number,
				RevisionId: t.Revision,
				Message:    "The change cannot be auto-reviewed. The following files do not match the benign file configuration: a/d.txt, a/e/p/p.txt, a/f/z.txt, a/p",
			})).Return(&gerritpb.ReviewResult{}, nil)
			gerritMock.EXPECT().DeleteReviewer(gomock.Any(), proto.MatcherEqual(&gerritpb.DeleteReviewerRequest{
				Number:    t.Number,
				AccountId: "srv-account@example.com",
			})).Return(nil, nil)

			err := ReviewChange(ctx, t)
			So(err, ShouldBeNil)
		})
	})
}
