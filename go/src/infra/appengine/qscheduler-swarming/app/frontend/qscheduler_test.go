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

package frontend

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/appengine/gaetesting"

	qscheduler "infra/appengine/qscheduler-swarming/api/qscheduler/v1"
	"infra/qscheduler/qslib/tutils"
	swarming "infra/swarming"
)

func TestAssignTasks(t *testing.T) {
	Convey("Given a testing context with a scheduler pool", t, func() {
		ctx := gaetesting.TestingContext()
		poolID := "Pool1"
		admin := &QSchedulerAdminServerImpl{}
		sch := &QSchedulerServerImpl{}
		view := &QSchedulerViewServerImpl{}
		_, err := admin.CreateSchedulerPool(ctx, &qscheduler.CreateSchedulerPoolRequest{
			PoolId: poolID,
			Config: &qscheduler.SchedulerPoolConfig{
				Labels: []string{"label1"},
			},
		})
		So(err, ShouldBeNil)

		Convey("with an idle task that has been notified", func() {
			taskID := "Task1"
			req := swarming.NotifyTasksRequest{
				SchedulerId: poolID,
				Notifications: []*swarming.NotifyTasksItem{
					{
						Time: tutils.TimestampProto(time.Now()),
						Task: &swarming.TaskSpec{
							Id:    taskID,
							State: swarming.TaskState_PENDING,
							Slices: []*swarming.SliceSpec{
								{Dimensions: []string{"label1"}},
							},
							EnqueuedTime: tutils.TimestampProto(time.Now()),
						},
					},
				},
			}
			_, err := sch.NotifyTasks(ctx, &req)
			So(err, ShouldBeNil)

			resp, err := view.InspectPool(ctx, &qscheduler.InspectPoolRequest{PoolId: poolID})
			So(err, ShouldBeNil)
			So(resp.NumWaitingTasks, ShouldEqual, 1)

			Convey("when AssignTasks is called with an idle bot", func() {
				botID := "Bot1"
				req := swarming.AssignTasksRequest{
					SchedulerId: poolID,
					Time:        tutils.TimestampProto(time.Now()),
					IdleBots: []*swarming.IdleBot{
						{BotId: botID, Dimensions: []string{"label1"}},
					},
				}
				resp, err := sch.AssignTasks(ctx, &req)
				Convey("then the task is assigned to the bot.", func() {
					So(err, ShouldBeNil)
					So(resp.Assignments, ShouldHaveLength, 1)
					So(resp.Assignments[0].BotId, ShouldEqual, botID)
					So(resp.Assignments[0].TaskId, ShouldEqual, taskID)
				})
			})
		})

		Convey("when AssignTasks is called with an idle bot that doesn't have needed dimensions", func() {
			botID := "Bot1"
			req := swarming.AssignTasksRequest{
				SchedulerId: poolID,
				Time:        tutils.TimestampProto(time.Now()),
				IdleBots: []*swarming.IdleBot{
					{BotId: botID},
				},
			}
			resp, err := sch.AssignTasks(ctx, &req)
			Convey("then an error is returned.", func() {
				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "does not have all scheduler dimensions")
			})
		})
	})
}
