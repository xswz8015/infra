// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package provision run provisioning for DUT.
package provision

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/golang/protobuf/jsonpb"
	build_api "go.chromium.org/chromiumos/config/go/build/api"
	"go.chromium.org/chromiumos/config/go/test/api"
	lab_api "go.chromium.org/chromiumos/config/go/test/lab/api"
	"go.chromium.org/luci/common/errors"

	"infra/cros/cmd/cros-tool-runner/internal/common"
	"infra/cros/cmd/cros-tool-runner/internal/docker"
	"infra/cros/cmd/cros-tool-runner/internal/services"
)

const (
	// Cros-dut result temp dir.
	crosDutResultsTempDir = "cros-dut-results"

	// Cros-provision result temp dir.
	crosProvisionResultsTempDir = "cros-provision-results"
)

// Result holds result data.
type Result struct {
	Out *api.CrosProvisionResponse
	Err error
}

// Run runs provisioning software dependencies per DUT.
func Run(ctx context.Context, device *api.CrosToolRunnerProvisionRequest_Device, crosDutContainer, crosProvisionContainer *build_api.ContainerImageInfo, token string) *Result {
	res := &Result{
		Out: &api.CrosProvisionResponse{
			Id: device.GetDut().GetId(),
			Outcome: &api.CrosProvisionResponse_Failure{
				Failure: &api.InstallFailure{
					Reason: api.InstallFailure_REASON_PROVISIONING_FAILED,
				},
			},
		},
	}
	if device == nil || device.GetProvisionState() == nil {
		res.Err = errors.Reason("run provision: DUT input is empty").Err()
		return res
	}
	dutName := res.Out.Id.GetValue()
	cacheServerInfo := device.GetDut().GetCacheServer()
	dutSshInfo := device.GetDut().GetChromeos().GetSsh()
	log.Printf("Preparing for provisioning of %q, with: %s", dutName, device.GetProvisionState())

	// Create separate network to run docker independent.
	log.Printf("--> Creating the network for %q ...", dutName)
	networkName := fmt.Sprintf("%s_network", dutName)

	if err := docker.CreateNetwork(ctx, networkName); err != nil {
		res.Err = errors.Annotate(err, "run provision").Err()
		return res
	}
	defer func() {
		docker.RemoveNetwork(ctx, networkName)
	}()
	log.Printf("--> Network was created for %q ...", dutName)

	// Create temp results dir for cros-dut
	crosDutResultsDir, err := ioutil.TempDir("", crosDutResultsTempDir)
	if err != nil {
		log.Printf("cros-dut results temp directory creation failed with error: %s", err)
		res.Err = errors.Annotate(err, "create dut service: create temp dir").Err()
		return res
	}

	log.Printf("--> Starting cros-dut service for %q ...", dutName)
	dutService, err := services.CreateDutService(ctx, crosDutContainer, dutName, networkName, cacheServerInfo, dutSshInfo, crosDutResultsDir, token)
	if err != nil {
		res.Err = errors.Annotate(err, "run provision").Err()
		return res
	}
	defer func() {
		dutService.Remove(ctx)
		common.AddContentsToLog("log.txt", crosDutResultsDir, "Reading cros-dut log file.")
	}()
	log.Printf("--> Container of cros-dut was started for %q", dutName)

	// Create temp results dir for cros-provision
	crosProvisionResultsDir, err := ioutil.TempDir("", crosProvisionResultsTempDir)
	if err != nil {
		res.Err = errors.Annotate(err, "run provision: create temp dir").Err()
		return res
	}

	log.Printf("--> Starting cros-provision service for %q ...", dutName)
	provisionReq := &api.CrosProvisionRequest{
		Dut:            device.GetDut(),
		ProvisionState: device.GetProvisionState(),
		DutServer: &lab_api.IpEndpoint{
			Address: dutService.Name,
			Port:    int32(dutService.ServicePort),
		},
	}

	provisionService, err := services.RunProvisionCLI(ctx, crosProvisionContainer, networkName, provisionReq, crosProvisionResultsDir, token)
	if err != nil {
		res.Err = errors.Annotate(err, "run provision").Err()
		return res
	}
	defer func() {
		if provisionService != nil {
			provisionService.Remove(ctx)
		}
		common.AddContentsToLog("log.txt", crosProvisionResultsDir, "Reading cros-provision log file.")
	}()
	log.Printf("--> Started cros-provision service for %q", dutName)

	resultFileName := path.Join(crosProvisionResultsDir, services.OutputFileName)
	if _, err := os.Stat(resultFileName); os.IsNotExist(err) {
		res.Err = errors.Reason("run provision: result not found").Err()
		return res
	}
	out, err := readProvisionOutput(resultFileName)
	if err != nil {
		res.Err = errors.Annotate(err, "run provision").Err()
		return res
	}
	log.Printf("result file %s: found. %s", dutName, out)
	if f := out.GetFailure(); f != nil {
		res.Out.Outcome = &api.CrosProvisionResponse_Failure{
			Failure: f,
		}
		res.Err = errors.Annotate(err, "run provision").Err()
	} else {
		res.Out.Outcome = &api.CrosProvisionResponse_Success{
			Success: &api.InstallSuccess{},
		}
		res.Err = nil
	}
	return res
}

// readProvisionOutput reads output file generated by cros-provision.
func readProvisionOutput(filePath string) (*api.CrosProvisionResponse, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Annotate(err, "read output").Err()
	}
	out := &api.CrosProvisionResponse{}
	err = jsonpb.Unmarshal(r, out)

	log.Printf("cros-provision response:" + out.String())
	return out, errors.Annotate(err, "read output").Err()
}
