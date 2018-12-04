// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inventory

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/mock/gomock"

	fleet "infra/appengine/crosskylabadmin/api/fleet/v1"
	"infra/appengine/crosskylabadmin/app/config"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestEnsurePoolHealthyFailsWithoutDryrun(t *testing.T) {
	Convey("EnsurePoolHealthy fails without dryrun", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		_, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{})
		So(err, ShouldErrLike, status.Errorf(codes.Unimplemented, ""))
	})
}

func TestEnsurePoolHealthyDryrun(t *testing.T) {
	Convey("EnsurePoolHealthy(dryrun) fails with no DutSelector", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		_, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{
			Options: &fleet.EnsurePoolHealthyRequest_Options{Dryrun: true},
		})
		So(err, ShouldErrLike, status.Errorf(codes.InvalidArgument, ""))
	})

	Convey("EnsurePoolHealthy succeeds with no changes for empty inventory", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		err := setupLabInventoryArchive(tf.C, tf.FakeGitiles, []testInventoryDut{})
		So(err, ShouldBeNil)

		resp, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{
			DutSelector: &fleet.DutSelector{
				Model: "link",
			},
			SparePool:  "suites",
			TargetPool: "cq",
			Options:    &fleet.EnsurePoolHealthyRequest_Options{Dryrun: true},
		})
		So(err, ShouldBeNil)
		So(resp.GetSparePoolStatus().GetSize(), ShouldEqual, 0)
		So(resp.GetSparePoolStatus().GetHealthyCount(), ShouldEqual, 0)
		So(resp.GetTargetPoolStatus().GetSize(), ShouldEqual, 0)
		So(resp.GetTargetPoolStatus().GetHealthyCount(), ShouldEqual, 0)
		So(resp.Changes, ShouldHaveLength, 0)
		So(resp.Failures, ShouldHaveLength, 0)
	})

	Convey("EnsurePoolHealthy swaps no DUT with all DUTs healthy", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		err := setupLabInventoryArchive(tf.C, tf.FakeGitiles, []testInventoryDut{
			{"link_cq_healthy", "link", "DUT_POOL_CQ"},
			{"link_suites_healthy", "link", "DUT_POOL_SUITES"},
		})
		So(err, ShouldBeNil)
		expectDutsWithHealth(tf.MockTracker, map[string]fleet.Health{
			"link_cq_healthy":     fleet.Health_Healthy,
			"link_suites_healthy": fleet.Health_Healthy,
		})

		resp, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{
			DutSelector: &fleet.DutSelector{
				Model: "link",
			},
			SparePool:  "DUT_POOL_SUITES",
			TargetPool: "DUT_POOL_CQ",
			Options:    &fleet.EnsurePoolHealthyRequest_Options{Dryrun: true},
		})
		So(err, ShouldBeNil)
		So(resp.GetSparePoolStatus().GetSize(), ShouldEqual, 1)
		So(resp.GetSparePoolStatus().GetHealthyCount(), ShouldEqual, 1)
		So(resp.GetTargetPoolStatus().GetSize(), ShouldEqual, 1)
		So(resp.GetTargetPoolStatus().GetHealthyCount(), ShouldEqual, 1)
		So(resp.Changes, ShouldHaveLength, 0)
		So(resp.Failures, ShouldHaveLength, 0)
	})

	Convey("EnsurePoolHealthy swaps one DUT with one DUT needed and one available", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		err := setupLabInventoryArchive(tf.C, tf.FakeGitiles, []testInventoryDut{
			{"link_cq_unhealthy", "link", "DUT_POOL_CQ"},
			{"link_suites_healthy", "link", "DUT_POOL_SUITES"},
		})
		So(err, ShouldBeNil)
		expectDutsWithHealth(tf.MockTracker, map[string]fleet.Health{
			"link_cq_unhealthy":   fleet.Health_Unhealthy,
			"link_suites_healthy": fleet.Health_Healthy,
		})

		resp, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{
			DutSelector: &fleet.DutSelector{
				Model: "link",
			},
			SparePool:  "DUT_POOL_SUITES",
			TargetPool: "DUT_POOL_CQ",
			Options:    &fleet.EnsurePoolHealthyRequest_Options{Dryrun: true},
		})

		So(err, ShouldBeNil)
		So(resp.GetSparePoolStatus().GetSize(), ShouldEqual, 1)
		So(resp.GetSparePoolStatus().GetHealthyCount(), ShouldEqual, 0)
		So(resp.GetTargetPoolStatus().GetSize(), ShouldEqual, 1)
		So(resp.GetTargetPoolStatus().GetHealthyCount(), ShouldEqual, 1)

		mc := poolChangeMap(resp.Changes)
		So(mc, ShouldResemble, map[string]*fleet.PoolChange{
			"link_cq_unhealthy": {
				DutId:   "link_cq_unhealthy",
				OldPool: "DUT_POOL_CQ",
				NewPool: "DUT_POOL_SUITES",
			},
			"link_suites_healthy": {
				DutId:   "link_suites_healthy",
				OldPool: "DUT_POOL_SUITES",
				NewPool: "DUT_POOL_CQ",
			},
		})

		So(resp.Failures, ShouldHaveLength, 0)
	})

	Convey("EnsurePoolHealthy swaps one DUT and reports failure with two DUTs needed but one available", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		err := setupLabInventoryArchive(tf.C, tf.FakeGitiles, []testInventoryDut{
			{"link_cq_unhealthy_1", "link", "DUT_POOL_CQ"},
			{"link_cq_unhealthy_2", "link", "DUT_POOL_CQ"},
			{"link_suites_healthy", "link", "DUT_POOL_SUITES"},
		})
		So(err, ShouldBeNil)

		expectDutsWithHealth(tf.MockTracker, map[string]fleet.Health{
			"link_cq_unhealthy_1": fleet.Health_Unhealthy,
			"link_cq_unhealthy_2": fleet.Health_Unhealthy,
			"link_suites_healthy": fleet.Health_Healthy,
		})

		resp, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{
			DutSelector: &fleet.DutSelector{
				Model: "link",
			},
			SparePool:  "DUT_POOL_SUITES",
			TargetPool: "DUT_POOL_CQ",
			Options:    &fleet.EnsurePoolHealthyRequest_Options{Dryrun: true},
		})

		So(err, ShouldBeNil)
		So(resp.GetSparePoolStatus().GetSize(), ShouldEqual, 1)
		So(resp.GetSparePoolStatus().GetHealthyCount(), ShouldEqual, 0)
		So(resp.GetTargetPoolStatus().GetSize(), ShouldEqual, 2)
		So(resp.GetTargetPoolStatus().GetHealthyCount(), ShouldEqual, 1)

		So(resp.Changes, ShouldHaveLength, 2)
		mc := poolChangeMap(resp.Changes)
		So(mc["link_suites_healthy"], ShouldResemble, &fleet.PoolChange{
			DutId:   "link_suites_healthy",
			OldPool: "DUT_POOL_SUITES",
			NewPool: "DUT_POOL_CQ",
		})
		if d, ok := mc["link_cq_unhealthy_1"]; ok {
			So(d, ShouldResemble, &fleet.PoolChange{
				DutId:   "link_cq_unhealthy_1",
				OldPool: "DUT_POOL_CQ",
				NewPool: "DUT_POOL_SUITES",
			})
		} else if d, ok := mc["link_cq_unhealthy_2"]; ok {
			So(d, ShouldResemble, &fleet.PoolChange{
				DutId:   "link_cq_unhealthy_2",
				OldPool: "DUT_POOL_CQ",
				NewPool: "DUT_POOL_SUITES",
			})
		} else {
			t.Error("no DUT swapped out of target pool")
		}

		So(resp.Failures, ShouldResemble, []fleet.EnsurePoolHealthyResponse_Failure{fleet.EnsurePoolHealthyResponse_NOT_ENOUGH_HEALTHY_SPARES})
	})

	Convey("EnsurePoolHealthy treats target DUT with unknown health as unhealthy", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		err := setupLabInventoryArchive(tf.C, tf.FakeGitiles, []testInventoryDut{
			{"link_cq_unknown", "link", "DUT_POOL_CQ"},
			{"link_suites_healthy", "link", "DUT_POOL_SUITES"},
		})
		So(err, ShouldBeNil)
		expectDutsWithHealth(tf.MockTracker, map[string]fleet.Health{
			"link_suites_healthy": fleet.Health_Healthy,
		})

		resp, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{
			DutSelector: &fleet.DutSelector{
				Model: "link",
			},
			SparePool:  "DUT_POOL_SUITES",
			TargetPool: "DUT_POOL_CQ",
			Options:    &fleet.EnsurePoolHealthyRequest_Options{Dryrun: true},
		})

		So(err, ShouldBeNil)
		So(resp.GetSparePoolStatus().GetSize(), ShouldEqual, 1)
		So(resp.GetSparePoolStatus().GetHealthyCount(), ShouldEqual, 0)
		So(resp.GetTargetPoolStatus().GetSize(), ShouldEqual, 1)
		So(resp.GetTargetPoolStatus().GetHealthyCount(), ShouldEqual, 1)

		mc := poolChangeMap(resp.Changes)
		So(mc, ShouldResemble, map[string]*fleet.PoolChange{
			"link_cq_unknown": {
				DutId:   "link_cq_unknown",
				OldPool: "DUT_POOL_CQ",
				NewPool: "DUT_POOL_SUITES",
			},
			"link_suites_healthy": {
				DutId:   "link_suites_healthy",
				OldPool: "DUT_POOL_SUITES",
				NewPool: "DUT_POOL_CQ",
			},
		})

		So(resp.Failures, ShouldHaveLength, 0)
	})

	Convey("EnsurePoolHealthy filters DUTs by environment", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		ptext := `
			duts {
				common {
					id: "dut_in_env"
					hostname: "dut_in_env"
					labels {
						model: "link"
						critical_pools: DUT_POOL_CQ
					}
					environment: ENVIRONMENT_STAGING
				}
			}
			duts {
				common {
					id: "dut_not_in_env"
					hostname: "dut_not_in_env"
					labels {
						model: "link"
						critical_pools: DUT_POOL_CQ
					}
					environment: ENVIRONMENT_PROD
				}
			}
		`
		So(tf.FakeGitiles.addArchive(config.Get(tf.C).Inventory, []byte(ptext)), ShouldBeNil)
		expectDutsWithHealth(tf.MockTracker, map[string]fleet.Health{
			"dut_in_env":    fleet.Health_Healthy,
			"dut_no_in_env": fleet.Health_Healthy,
		})

		resp, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{
			DutSelector: &fleet.DutSelector{
				Model: "link",
			},
			SparePool:  "DUT_POOL_SUITES",
			TargetPool: "DUT_POOL_CQ",
			Options:    &fleet.EnsurePoolHealthyRequest_Options{Dryrun: true},
		})
		So(err, ShouldBeNil)
		So(resp.GetSparePoolStatus().GetSize(), ShouldEqual, 0)
		So(resp.GetSparePoolStatus().GetHealthyCount(), ShouldEqual, 0)
		So(resp.GetTargetPoolStatus().GetSize(), ShouldEqual, 1)
		So(resp.GetTargetPoolStatus().GetHealthyCount(), ShouldEqual, 1)
		So(resp.Changes, ShouldHaveLength, 0)
		So(resp.Failures, ShouldHaveLength, 0)
	})

	Convey("EnsurePoolHealthy swaps no DUTs and reports failure with too many unhealthy DUTs", t, func() {
		tf, validate := newTestFixture(t)
		defer validate()

		err := setupLabInventoryArchive(tf.C, tf.FakeGitiles, []testInventoryDut{
			{"link_cq_unhealthy_1", "link", "DUT_POOL_CQ"},
			{"link_cq_unhealthy_2", "link", "DUT_POOL_CQ"},
		})
		So(err, ShouldBeNil)

		expectDutsWithHealth(tf.MockTracker, map[string]fleet.Health{
			"link_cq_unhealthy_1": fleet.Health_Unhealthy,
			"link_cq_unhealthy_2": fleet.Health_Unhealthy,
		})

		resp, err := tf.Inventory.EnsurePoolHealthy(tf.C, &fleet.EnsurePoolHealthyRequest{
			DutSelector: &fleet.DutSelector{
				Model: "link",
			},
			SparePool:        "DUT_POOL_SUITES",
			TargetPool:       "DUT_POOL_CQ",
			MaxUnhealthyDuts: 1,
			Options:          &fleet.EnsurePoolHealthyRequest_Options{Dryrun: true},
		})

		So(err, ShouldBeNil)
		So(resp.GetTargetPoolStatus().GetSize(), ShouldEqual, 2)
		So(resp.GetTargetPoolStatus().GetHealthyCount(), ShouldEqual, 0)
		So(resp.Changes, ShouldHaveLength, 0)
		So(resp.Failures, ShouldResemble, []fleet.EnsurePoolHealthyResponse_Failure{fleet.EnsurePoolHealthyResponse_TOO_MANY_UNHEALTHY_DUTS})
	})

}

type testInventoryDut struct {
	id    string
	model string
	pool  string
}

func setupLabInventoryArchive(c context.Context, g *fakeGitilesClient, duts []testInventoryDut) error {
	ptext := ""
	for _, dut := range duts {
		ptext = fmt.Sprintf(`%s
			duts {
				common {
					id: "%s"
					hostname: "%s"
					labels {
						model: "%s"
						critical_pools: %s
					}
					environment: ENVIRONMENT_STAGING
				}
			}`,
			ptext,
			dut.id, dut.id, dut.model, dut.pool,
		)
	}
	return g.addArchive(config.Get(c).Inventory, []byte(ptext))
}

func expectDutsWithHealth(t *fleet.MockTrackerServer, dutHealths map[string]fleet.Health) {
	ft := &trackerPartialFake{dutHealths}
	t.EXPECT().SummarizeBots(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(ft.SummarizeBots)
}

func poolChangeMap(pcs []*fleet.PoolChange) map[string]*fleet.PoolChange {
	mc := make(map[string]*fleet.PoolChange)
	for _, c := range pcs {
		mc[c.DutId] = c
	}
	return mc
}
