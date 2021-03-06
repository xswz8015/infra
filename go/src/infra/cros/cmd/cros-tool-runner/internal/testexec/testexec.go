// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package testexec runs tests.
package testexec

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/jsonpb"

	_go "go.chromium.org/chromiumos/config/go"
	build_api "go.chromium.org/chromiumos/config/go/build/api"
	lab_api "go.chromium.org/chromiumos/config/go/test/lab/api"

	"go.chromium.org/chromiumos/config/go/test/api"
	"go.chromium.org/luci/common/errors"

	"infra/cros/cmd/cros-tool-runner/internal/services"
)

// Run runs tests.
func Run(ctx context.Context, req *api.CrosToolRunnerTestRequest, crosTestContainer, crosDUTContainer *build_api.ContainerImageInfo, token string) (res *api.CrosToolRunnerTestResponse, err error) {
	// Use host network for dev environment which DUT address is in the form localhost:<port>
	const networkName = "host"

	if req.PrimaryDut == nil {
		return nil, errors.Reason("run test: primary DUT is not specified").Err()
	}

	artifactDir, err := filepath.Abs(req.ArtifactDir)
	if err != nil {
		return nil, errors.Annotate(err, "run test: failed to resolve artifact directory %v", req.ArtifactDir).Err()
	}
	// All non-test harness artifacts will be in <artifact_dir>/cros-test/cros-test.
	crosTestDir := path.Join(artifactDir, "cros-test", "cros-test")
	// All test result artifacts will be in <artifact_dir>/cros-test/artifact.
	resultDir := path.Join(artifactDir, "cros-test", "artifact")
	// The input file name.
	inputFileName := path.Join(crosTestDir, "request.json")
	// The directory for cros-dut artifacts.
	crosDUTDir := path.Join(artifactDir, "cros-dut")

	// Setting up directories.
	if err := os.MkdirAll(crosTestDir, 0755); err != nil {
		return nil, errors.Annotate(err, "run test: failed to create directory %s", crosTestDir).Err()
	}
	log.Printf("Run test: created the cros-test directory %s", crosTestDir)
	if err := os.MkdirAll(resultDir, 0755); err != nil {
		return nil, errors.Annotate(err, "run test: failed to create directory %s", resultDir).Err()
	}
	log.Printf("Run test: created the test artifact directory %s", resultDir)

	duts := []*lab_api.Dut{req.PrimaryDut.GetDut()}

	var companions []*api.CrosTestRequest_Device
	for _, c := range req.GetCompanionDuts() {
		companions = append(
			companions, &api.CrosTestRequest_Device{
				Dut: c.GetDut(),
			},
		)
		duts = append(duts, c.GetDut())
	}
	dutServices, err := services.CreateDutServicesForHostNetwork(ctx, crosDUTContainer, duts, crosDUTDir, token)
	if err != nil {
		return nil, errors.Annotate(err, "run test: failed to start DUT servers").Err()
	}
	defer func() {
		for _, d := range dutServices {
			d.Remove(ctx)
		}
	}()
	for i, c := range companions {
		c.DutServer = &lab_api.IpEndpoint{Address: "localhost", Port: int32(dutServices[i+1].ServicePort)}
	}
	testReq := &api.CrosTestRequest{
		TestSuites: req.GetTestSuites(),
		Primary: &api.CrosTestRequest_Device{
			Dut:       req.PrimaryDut.GetDut(),
			DutServer: &lab_api.IpEndpoint{Address: "localhost", Port: int32(dutServices[0].ServicePort)},
		},
		Companions: companions,
	}
	if err := writeTestInput(inputFileName, testReq); err != nil {
		return nil, errors.Annotate(err, "run test: failed to create input file %s", inputFileName).Err()
	}
	if err = services.RunTestCLI(ctx, crosTestContainer, networkName, inputFileName, crosTestDir, resultDir, token); err != nil {
		return nil, errors.Annotate(err, "run test: failed to run test CLI").Err()
	}

	resultFileName := path.Join(crosTestDir, "result.json")
	if _, err := os.Stat(resultFileName); os.IsNotExist(err) {
		return nil, errors.Reason("run test: result not found").Err()
	}
	out, err := readTestOutput(resultFileName)
	if err != nil {
		return nil, errors.Annotate(err, "run test: failed to read test output").Err()
	}

	return prepareTestResponse(artifactDir, out.TestCaseResults)
}

// writeTestInput writes a CrosTestRequest json.
func writeTestInput(file string, req *api.CrosTestRequest) error {
	f, err := os.Create(file)
	if err != nil {
		return errors.Annotate(err, "fail to create file %v", file).Err()
	}
	m := jsonpb.Marshaler{}
	if err := m.Marshal(f, req); err != nil {
		return errors.Annotate(err, "fail to marshal request to file %v", file).Err()
	}
	return nil
}

// readTestOutput reads output file generated by cros-test.
func readTestOutput(filePath string) (*api.CrosTestResponse, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Annotate(err, "read output").Err()
	}
	out := &api.CrosTestResponse{}
	err = jsonpb.Unmarshal(r, out)
	return out, errors.Annotate(err, "read output").Err()
}

// prepareTestResponse prepares a response for test execution.
func prepareTestResponse(resultRootDir string, testCaseResults []*api.TestCaseResult) (res *api.CrosToolRunnerTestResponse, err error) {
	var results []*api.TestCaseResult
	for _, t := range testCaseResults {
		resultDir := strings.Replace(t.ResultDirPath.Path, services.CrosTestRootDirInsideDocker, resultRootDir, 1)
		results = append(results, &api.TestCaseResult{
			TestCaseId:    t.TestCaseId,
			ResultDirPath: &_go.StoragePath{Path: resultDir},
			Verdict:       t.Verdict,
			Reason:        t.Reason,
			TestHarness:   t.TestHarness,
		})
	}
	return &api.CrosToolRunnerTestResponse{
		TestCaseResults: results,
	}, nil
}
