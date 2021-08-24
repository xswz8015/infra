// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.chromium.org/luci/appengine/gaetesting"
	"go.chromium.org/luci/gae/service/datastore"
	"google.golang.org/protobuf/testing/protocmp"

	kartepb "infra/cros/karte/api"
)

// TestCreateAction makes sure that CreateAction returns the action it created and that the action is present in datastore.
func TestCreateAction(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	k := NewKarteFrontend()
	resp, err := k.CreateAction(ctx, &kartepb.CreateActionRequest{
		Action: &kartepb.Action{
			Kind: "ssh-attempt",
		},
	})
	expected := &kartepb.Action{
		Name: "",
		Kind: "ssh-attempt",
	}
	if err != nil {
		t.Error(err)
	}
	if diff := cmp.Diff(expected, resp, protocmp.Transform()); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	// Here we inspect the contents of datastore.
	q, err := newActionEntitiesQuery("", "")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	datastoreActionEntities, err := q.Next(ctx, 0)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if datastoreActionEntities == nil {
		t.Errorf("action entities should not be nil")
	}
	switch len(datastoreActionEntities) {
	case 0:
		t.Errorf("datastore should not be empty")
	case 1:
	default:
		t.Errorf("datastore should not have more than 1 item")
	}
}

// TestCreateObservation makes sure that that CreateObservation fails because
// it isn't implemented.
func TestCreateObservation(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	k := NewKarteFrontend()
	_, err := k.CreateObservation(ctx, &kartepb.CreateObservationRequest{})
	if err == nil {
		t.Error("expected Create Observation to fail")
	}
}

// TestListActionsSmokeTest tests that ListActions does not error.
func TestListActionsSmokeTest(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	k := NewKarteFrontend()
	resp, err := k.ListActions(ctx, &kartepb.ListActionsRequest{})
	if resp == nil {
		t.Errorf("expected resp to not be nil")
	}
	if len(resp.GetActions()) != 0 {
		t.Errorf("expected actions to be trivial")
	}
	if err != nil {
		t.Errorf("expected error to be nil not %s", err)
	}
}

// TestListActions tests that ListActions errors.
func TestListActions(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	if err := PutActionEntities(
		ctx,
		&ActionEntity{
			ID: "aaaa",
		},
	); err != nil {
		t.Error(err)
	}
	k := NewKarteFrontend()
	resp, err := k.ListActions(ctx, &kartepb.ListActionsRequest{})
	if err != nil {
		t.Errorf("expected error to be nil not %s", err)
	}
	if resp == nil {
		t.Errorf("expected resp to not be nil")
	}
	if resp.GetActions() == nil {
		t.Errorf("expected actions to not be nil")
	}
	if len(resp.GetActions()) != 1 {
		t.Errorf("expected len(actions) to be 1 not %d", len(resp.GetActions()))
	}
}

// TestListObservations tests that ListObservations errors.
func TestListObservations(t *testing.T) {
	t.Parallel()
	k := NewKarteFrontend()
	ctx := gaetesting.TestingContext()
	datastore.GetTestable(ctx).Consistent(true)
	resp, err := k.ListObservations(ctx, &kartepb.ListObservationsRequest{})
	if resp == nil {
		t.Errorf("expected resp to not be nil")
	}
	if err != nil {
		t.Errorf("expected error to be nil not %s", err)
	}
}
