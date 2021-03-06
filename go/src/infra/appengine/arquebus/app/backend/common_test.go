// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package backend

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	. "github.com/smartystreets/goconvey/convey"

	"go.chromium.org/luci/appengine/tq"
	"go.chromium.org/luci/appengine/tq/tqtesting"
	"go.chromium.org/luci/common/clock/testclock"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"infra/appengine/arquebus/app/backend/model"
	"infra/appengine/arquebus/app/config"
	"infra/appengine/arquebus/app/util"
	rotationproxy "infra/appengine/rotation-proxy/proto"
	monorail "infra/monorailv2/api/api_proto"
)

var (
	testStart, _    = ptypes.TimestampProto(testclock.TestRecentTimeUTC)
	testEnd, _      = ptypes.TimestampProto(testclock.TestRecentTimeUTC)
	yesterday, _    = ptypes.TimestampProto(time.Now().AddDate(0, 0, -1))
	tomorrow, _     = ptypes.TimestampProto(time.Now().AddDate(0, 0, 1))
	dayAfterNext, _ = ptypes.TimestampProto(time.Now().AddDate(0, 0, 2))

	// sample output from rotation-proxy.
	sampleRotationProxyRotations = map[string]*rotationproxy.Rotation{
		"Rotation 1": {
			Name: "Rotation 1",
			Shifts: []*rotationproxy.Shift{
				{
					Oncalls: []*rotationproxy.OncallPerson{
						{Email: "r1pri@example.com"},
						{Email: "r1sec1@example.com"},
						{Email: "r1sec2@example.com"},
					},
					StartTime: yesterday,
					EndTime:   tomorrow,
				},
			},
		},
		"Rotation 2": {
			Name: "Rotation 2",
			Shifts: []*rotationproxy.Shift{
				{
					Oncalls: []*rotationproxy.OncallPerson{
						{Email: "r2pri@example.com"},
						{Email: "r2sec1@example.com"},
						{Email: "r2sec2@example.com"},
					},
					StartTime: yesterday,
				},
			},
		},
		"Rotation 3": {
			Name: "Rotation 3",
			Shifts: []*rotationproxy.Shift{
				{
					Oncalls: []*rotationproxy.OncallPerson{
						{Email: "r3pri@example.com"},
						{Email: "r3sec1@example.com"},
						{Email: "r3sec2@example.com"},
					},
					StartTime: tomorrow,
					EndTime:   dayAfterNext,
				},
			},
		},
	}
)

// createTestContextWithTQ creates a test context with testable a TaskQueue.
func createTestContextWithTQ() context.Context {
	// create a context with config first.
	c := util.CreateTestContext()
	c = config.SetConfig(c, &config.Config{
		AccessGroup:           "engineers",
		MonorailHostname:      "example.org",
		RotationProxyHostname: "example.net",

		Assigners: []*config.Assigner{},
	}, "config-rev")

	// install TQ handlers
	d := &tq.Dispatcher{}
	registerTaskHandlers(d)
	tq := tqtesting.GetTestable(c, d)
	tq.CreateQueues()
	c = util.SetDispatcher(c, d)

	// install mocked pRPC clients.
	c = setMonorailClient(c, newTestIssueClient())
	c = setRotationProxyClient(c, newTestRotationProxyServiceClient())

	// set sample rotation shifts for rotation-proxy
	for name, rotation := range sampleRotationProxyRotations {
		mockRotation(c, name, rotation)
	}
	return c
}

// createAssigner creates a sample Assigner entity.
func createAssigner(c context.Context, id string) *model.Assigner {
	var cfg config.Assigner
	So(proto.UnmarshalText(util.SampleValidAssignerCfg, &cfg), ShouldBeNil)
	cfg.Id = id

	So(UpdateAssigners(c, []*config.Assigner{&cfg}, "rev-1"), ShouldBeNil)
	assigner, err := GetAssigner(c, id)
	So(assigner.ID, ShouldEqual, id)
	So(err, ShouldBeNil)
	So(assigner, ShouldNotBeNil)

	return assigner
}

func triggerScheduleTaskHandler(c context.Context, id string) []*model.Task {
	req := &ScheduleAssignerTask{AssignerId: id}
	So(scheduleAssignerTaskHandler(c, req), ShouldBeNil)
	_, tasks, err := GetAssignerWithTasks(c, id, 99999, true)
	So(err, ShouldBeNil)
	return tasks
}

func triggerRunTaskHandler(c context.Context, assignerID string, taskID int64) *model.Task {
	req := &RunAssignerTask{AssignerId: assignerID, TaskId: taskID}
	So(runAssignerTaskHandler(c, req), ShouldBeNil)
	assigner, task, err := GetTask(c, assignerID, taskID)
	So(assigner.ID, ShouldEqual, assignerID)
	So(err, ShouldBeNil)
	So(task, ShouldNotBeNil)
	return task
}

func createRawUserSources(sources ...*config.UserSource) [][]byte {
	raw := make([][]byte, len(sources))
	for i, source := range sources {
		raw[i], _ = proto.Marshal(source)
	}
	return raw
}

func monorailUser(email string) *monorail.UserRef {
	return &monorail.UserRef{DisplayName: email}
}

func emailUserSource(email string) *config.UserSource {
	return &config.UserSource{From: &config.UserSource_Email{Email: email}}
}

func rotationUserSource(name string, position config.Oncall_Position) *config.UserSource {
	return &config.UserSource{
		From: &config.UserSource_Rotation{Rotation: &config.Oncall{
			Name: name, Position: position,
		}},
	}
}

func findPrimaryOncall(shift *rotationproxy.Shift) *monorail.UserRef {
	if len(shift.Oncalls) == 0 {
		return nil
	}
	return monorailUser(shift.Oncalls[0].Email)
}

// ----------------------------------
// test Monorail Issue Client

type testIssueClientStorage struct {
	listIssuesRequest  *monorail.ListIssuesRequest
	listIssuesResponse []*monorail.Issue
	getIssueResponse   map[string]*monorail.Issue

	updateIssueRequestLock sync.Mutex
	updateIssueRequest     map[string]*monorail.UpdateIssueRequest
}

type testIssueClient struct {
	monorail.IssuesClient
	storage *testIssueClientStorage
}

func newTestIssueClient() testIssueClient {
	return testIssueClient{
		storage: &testIssueClientStorage{
			updateIssueRequest: map[string]*monorail.UpdateIssueRequest{},
			getIssueResponse:   map[string]*monorail.Issue{},
		},
	}
}

func (client testIssueClient) UpdateIssue(c context.Context, in *monorail.UpdateIssueRequest, opts ...grpc.CallOption) (*monorail.IssueResponse, error) {
	client.storage.updateIssueRequestLock.Lock()
	defer client.storage.updateIssueRequestLock.Unlock()
	client.storage.updateIssueRequest[genIssueKey(
		in.IssueRef.ProjectName, in.IssueRef.LocalId,
	)] = in
	return &monorail.IssueResponse{}, nil
}

func (client testIssueClient) ListIssues(c context.Context, in *monorail.ListIssuesRequest, opts ...grpc.CallOption) (*monorail.ListIssuesResponse, error) {
	client.storage.listIssuesRequest = in
	return &monorail.ListIssuesResponse{
		Issues: client.storage.listIssuesResponse,
	}, nil
}

func (client testIssueClient) GetIssue(c context.Context, in *monorail.GetIssueRequest, opts ...grpc.CallOption) (*monorail.IssueResponse, error) {
	response := getMonorailClient(c).(testIssueClient).storage.getIssueResponse
	issue, ok := response[genIssueKey(in.IssueRef.ProjectName, in.IssueRef.LocalId)]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "The issue does not exist.")
	}
	return &monorail.IssueResponse{Issue: issue}, nil
}

func mockGetAndListIssues(c context.Context, issues ...*monorail.Issue) {
	mockGetIssues(c, issues...)
	mockListIssues(c, issues...)
}

func mockListIssues(c context.Context, issues ...*monorail.Issue) {
	getMonorailClient(c).(testIssueClient).storage.listIssuesResponse = issues
}

func mockGetIssues(c context.Context, issues ...*monorail.Issue) {
	response := map[string]*monorail.Issue{}
	for _, issue := range issues {
		response[genIssueKey(issue.ProjectName, issue.LocalId)] = issue
	}
	getMonorailClient(c).(testIssueClient).storage.getIssueResponse = response
}

func getIssueUpdateRequest(c context.Context, projectName string, localID uint32) *monorail.UpdateIssueRequest {
	updateIssueRequest := getMonorailClient(c).(testIssueClient).storage.updateIssueRequest
	return updateIssueRequest[genIssueKey(projectName, localID)]
}

func getListIssuesRequest(c context.Context) *monorail.ListIssuesRequest {
	return getMonorailClient(c).(testIssueClient).storage.listIssuesRequest
}

func genIssueKey(projectName string, localID uint32) string {
	return fmt.Sprintf("%s:%d", projectName, localID)
}

// ----------------------------------
// test rotation-proxy RotationProxyService Client

type testRotationProxyServiceClientStorage struct {
	rotationsByName map[string]*rotationproxy.Rotation
}

type testRotationProxyServiceClient struct {
	rotationproxy.RotationProxyServiceClient
	storage *testRotationProxyServiceClientStorage
}

func newTestRotationProxyServiceClient() testRotationProxyServiceClient {
	return testRotationProxyServiceClient{
		storage: &testRotationProxyServiceClientStorage{
			rotationsByName: map[string]*rotationproxy.Rotation{},
		},
	}
}

func (client testRotationProxyServiceClient) GetRotation(c context.Context, req *rotationproxy.GetRotationRequest, opts ...grpc.CallOption) (*rotationproxy.Rotation, error) {
	rotation, exist := client.storage.rotationsByName[req.Name]
	if !exist {
		return nil, status.Errorf(codes.NotFound, `"%s" not found`, req.Name)
	}
	return rotation, nil
}

func mockRotation(c context.Context, name string, rotation *rotationproxy.Rotation) {
	getRotationProxyClient(c).(testRotationProxyServiceClient).storage.rotationsByName[name] = rotation
}
