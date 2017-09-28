// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"golang.org/x/net/context"

	"go.chromium.org/luci/hardcoded/chromeinfra"
	"go.chromium.org/luci/vpython/api/vpython"
	"go.chromium.org/luci/vpython/application"
	"go.chromium.org/luci/vpython/cipd"
	"go.chromium.org/luci/vpython/spec"

	"github.com/mitchellh/go-homedir"
	cipdClient "go.chromium.org/luci/cipd/client/cipd"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/common/system/environ"
)

const (
	// BypassENV is an environment variable that is used to detect if we shouldn't
	// do any vpython stuff at all, but should instead directly invoke the next
	// `python` on PATH.
	BypassENV = "VPYTHON_BYPASS"

	// BypassSentinel must be the BypassENV value (verbatim) in order to trigger
	// vpython bypass.
	BypassSentinel = "manually managed python not supported by chrome operations"
)

var cipdPackageLoader = cipd.PackageLoader{
	Options: cipdClient.ClientOptions{
		ServiceURL: chromeinfra.CIPDServiceURL,
		UserAgent:  "vpython",
	},
	Template: func(c context.Context, e *vpython.Environment) (map[string]string, error) {
		tag := pep425TagSelector(runtime.GOOS, e.Pep425Tag)
		if tag == nil {
			return nil, nil
		}
		return getPEP425CIPDTemplateForTag(tag)
	},
}

var defaultConfig = application.Config{
	PackageLoader: &cipdPackageLoader,
	SpecLoader: spec.Loader{
		CommonFilesystemBarriers: []string{
			".gclient",
		},
		CommonSpecNames: []string{
			".vpython",
		},
	},
	VENVPackage: vpython.Spec_Package{
		Name:    "infra/python/virtualenv",
		Version: "version:15.1.0",
	},
	PruneThreshold:    7 * 24 * time.Hour, // One week.
	MaxPrunesPerSweep: 3,
	MaxScriptPathLen:  127, // Maximum POSIX shebang length.
}

func mainImpl(c context.Context, argv []string, env environ.Env) int {
	// Initialize our CIPD package loader from the environment.
	//
	// If we don't have an environment-specific CIPD cache directory, use one
	// relative to the user's home directory.
	if err := cipdPackageLoader.Options.LoadFromEnv(env.GetEmpty); err != nil {
		logging.Errorf(c, "Could not inialize CIPD package loader: %s", err)
		return 1
	}
	if cipdPackageLoader.Options.CacheDir == "" {
		hd, err := homedir.Dir()
		if err == nil {
			cipdPackageLoader.Options.CacheDir = filepath.Join(hd, ".vpython_cipd_cache")
		} else {
			logging.WithError(err).Warningf(c,
				"Failed to resolve user home directory. No CIPD cache will be enabled.")
		}
	}

	// Determine if we're bypassing "vpython".
	defaultConfig.WithVerificationConfig = withVerificationConfig
	defaultConfig.Bypass = env.GetEmpty(BypassENV) == BypassSentinel
	return defaultConfig.Main(c, argv, env)
}

func main() {
	c := context.Background()
	c = gologger.StdConfig.Use(logging.SetLevel(c, logging.Warning))
	os.Exit(mainImpl(c, os.Args, environ.System()))
}
