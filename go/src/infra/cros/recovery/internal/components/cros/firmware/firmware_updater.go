// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package firmware

import (
	"context"
	"fmt"
	"time"

	"go.chromium.org/luci/common/errors"

	"infra/cros/recovery/internal/components"
	"infra/cros/recovery/logger"
)

// FirmwareUpdaterRequest holds request data for running firmware updater.
type FirmwareUpdaterRequest struct {
	// Mode used for updater.
	// Possible values is: autoupdate, recovery, factory.
	Mode string
	// Run updater with force option.
	Force bool
}

// RunFirmwareUpdater run chromeos-firmwareupdate to update firmware on the host.
func RunFirmwareUpdater(ctx context.Context, req *FirmwareUpdaterRequest, run components.Runner, log logger.Logger) error {
	switch req.Mode {
	case "autoupdate":
	case "recovery":
	case "factory":
	default:
		return errors.Reason("run firmware updater: mode %q is not supported", req.Mode).Err()
	}
	log.Debug("Run firmware updater: use %q mode.", req.Mode)
	args := []string{
		fmt.Sprintf("--mode=%s", req.Mode),
	}
	if req.Force {
		log.Debug("Run firmware updater: request to run with force.")
		args = append(args, "--force")
	}
	out, err := run(ctx, 5*time.Minute, "chromeos-firmwareupdate", args...)
	log.Debug("Run firmware updater stdout:\n%s", out)
	return errors.Annotate(err, "run firmware update").Err()
}

// ReadFirmwareKeysFromHost read AP keys from the host.
func ReadFirmwareKeysFromHost(ctx context.Context, run components.Runner, log logger.Logger) ([]string, error) {
	const extractImagePath = "/tmp/bios.bin"
	if out, err := run(ctx, 5*time.Minute, "flashrom", "-p", "host", "-r", extractImagePath); err != nil {
		return nil, errors.Annotate(err, "has dev signed firmware").Err()
	} else {
		log.Debug("Extract bios to the host: %s", out)
	}
	if keys, err := readAPKeysFromFile(ctx, extractImagePath, run, log); err != nil {
		return nil, errors.Annotate(err, "read ap info").Err()
	} else {
		return keys, nil
	}
}
