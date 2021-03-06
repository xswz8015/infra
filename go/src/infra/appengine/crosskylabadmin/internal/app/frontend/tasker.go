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
	"context"
	"fmt"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	fleet "infra/appengine/crosskylabadmin/api/fleet/v1"
	"infra/appengine/crosskylabadmin/internal/app/clients"
	"infra/appengine/crosskylabadmin/internal/app/config"
	"infra/appengine/crosskylabadmin/internal/app/frontend/internal/swarming"
	"infra/appengine/crosskylabadmin/internal/app/frontend/internal/worker"
)

// CreateAuditTask kicks off an audit job.
func CreateAuditTask(ctx context.Context, botID, taskname, actions string) (string, error) {
	at := worker.AuditTaskWithActions(ctx, taskname, actions)
	sc, err := clients.NewSwarmingClient(ctx, config.Get(ctx).Swarming.Host)
	if err != nil {
		return "", errors.Annotate(err, "failed to obtain swarming client").Err()
	}
	expSec := int64(24 * 60 * 60)
	execTimeoutSecs := int64(8 * 60 * 60)
	taskURL, err := runTaskByBotID(ctx, at, sc, botID, "", expSec, execTimeoutSecs)
	if err != nil {
		return "", errors.Annotate(err, "fail to create audit task for %s", botID).Err()
	}
	return taskURL, nil
}

func runTaskByBotID(ctx context.Context, at worker.Task, sc clients.SwarmingClient, botID, expectedState string, expirationSecs, executionTimeoutSecs int64) (string, error) {
	cfg := config.Get(ctx)
	tags := swarming.AddCommonTags(
		ctx,
		fmt.Sprintf("%s:%s", at.Name, botID),
		fmt.Sprintf("task:%s", at.Name),
	)
	tags = append(tags, at.Tags...)

	a := swarming.SetCommonTaskArgs(ctx, &clients.SwarmingCreateTaskArgs{
		Cmd:                  at.Cmd,
		BotID:                botID,
		ExecutionTimeoutSecs: executionTimeoutSecs,
		ExpirationSecs:       expirationSecs,
		Priority:             cfg.Cron.FleetAdminTaskPriority,
		Tags:                 tags,
	})
	if expectedState != "" {
		a.DutState = expectedState
	}
	tid, err := sc.CreateTask(ctx, at.Name, a)
	if err != nil {
		return "", errors.Annotate(err, "failed to create task for bot %s", botID).Err()
	}
	logging.Infof(ctx, "successfully kick off task %s for bot %s", tid, botID)
	return swarming.URLForTask(ctx, tid), nil
}

var dutStateForTask = map[fleet.TaskType]string{
	fleet.TaskType_Cleanup: "needs_cleanup",
	fleet.TaskType_Repair:  "needs_repair",
	fleet.TaskType_Reset:   "needs_reset",
}
