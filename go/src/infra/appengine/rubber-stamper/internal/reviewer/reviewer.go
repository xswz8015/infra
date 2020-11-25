// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package reviewer

import (
	"context"
	"fmt"
	"strings"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"

	"infra/appengine/rubber-stamper/config"
	"infra/appengine/rubber-stamper/internal/gerrit"
	"infra/appengine/rubber-stamper/internal/util"
	"infra/appengine/rubber-stamper/tasks/taskspb"
)

// ReviewChange reviews a CL and then either gives a Bot-Commit +1 label or
// leaves a comment explain why the CL shouldn't be passed and removes itself
// as a reviewer.
func ReviewChange(ctx context.Context, t *taskspb.ChangeReviewTask) error {
	cfg, err := config.Get(ctx)
	if err != nil {
		return err
	}
	hostCfg := cfg.HostConfigs[t.Host]

	gc, err := gerrit.GetCurrentClient(ctx, t.Host+"-review.googlesource.com")
	if err != nil {
		return err
	}

	invalidFiles, err := reviewBenignFileChange(ctx, hostCfg, gc, t)
	if err != nil {
		return err
	}
	if len(invalidFiles) > 0 {
		// Invalid BenignFileChange.
		// TODO: Add a go link in the msg, which tells users what the config
		// looks like, what classes of CLs are safe, and how to send a CL to
		// update the config to add a fileset.
		msg := "The change cannot be auto-reviewed. The following files do not match the benign file configuration: "
		msg = msg + strings.Join(invalidFiles[:], ", ")
		setReviewReq := &gerritpb.SetReviewRequest{
			Number:     t.Number,
			RevisionId: t.Revision,
			Message:    msg,
		}
		_, err := gc.SetReview(ctx, setReviewReq)
		if err != nil {
			return fmt.Errorf("failed to leave comment for host %s, cl %d, revision %s: %v", t.Host, t.Number, t.Revision, err.Error())
		}

		sa, err := util.GetServiceAccountName(ctx)
		if err != nil {
			return err
		}
		deleteReviewerReq := &gerritpb.DeleteReviewerRequest{
			Number:    t.Number,
			AccountId: sa,
		}
		_, err = gc.DeleteReviewer(ctx, deleteReviewerReq)
		if err != nil {
			return fmt.Errorf("failed to delete reviewer for host %s, cl %d, revision %s: %v", t.Host, t.Number, t.Revision, err.Error())
		}

		return nil
	}

	// TODO: Bot-Commit +1 & Owners-Override +1
	return nil
}
