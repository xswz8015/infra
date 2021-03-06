// Copyright 2020 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package plugsupport

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/encoding/prototext"

	"go.chromium.org/luci/common/errors"

	"infra/tools/migrator/internal/migratorpb"
)

// ProjectDir is an absolute path to a migrator project directory.
type ProjectDir string

// PluginDir returns the absolute path of the migrator project's plugin code
// directory.
func (p ProjectDir) PluginDir() string {
	return filepath.Join(string(p), "_plugin")
}

// ConfigDir returns the absolute path of the migrator project's config
// directory.
func (p ProjectDir) ConfigDir() string {
	return filepath.Join(string(p), ".migration")
}

// ConfigFile returns the absolute path of the migrator project's main config
// file.
//
// The existence of this file is used to determine if a folder is a migrator
// project.
func (p ProjectDir) ConfigFile() string {
	return filepath.Join(p.ConfigDir(), "config.cfg")
}

// CommitMessageFile is the absolute path to the file with the commit message to
// use in `upload` subcommand.
func (p ProjectDir) CommitMessageFile() string {
	return filepath.Join(string(p), "commit-message.txt")
}

// TrashDir returns the absolute path of the migrator project's trash
// directory.
//
// The trash directory is used to compile the plugin; New runs of migrator will
// make best-effort attempts to clean up this directory using CleanTrash().
func (p ProjectDir) TrashDir() string {
	return filepath.Join(string(p), ".trash")
}

// ScanReportPath returns the absolute path of the migrator project's CSV scan
// report file.
func (p ProjectDir) ScanReportPath() string {
	return filepath.Join(string(p), "scan.csv")
}

// StatusReportPath returns the absolute path of the migrator project's CSV
// status report file.
func (p ProjectDir) StatusReportPath() string {
	return filepath.Join(string(p), "status.csv")
}

// CommitReportPath returns the absolute path of the migrator project's CSV
// commit report file.
func (p ProjectDir) CommitReportPath() string {
	return filepath.Join(string(p), "commit.csv")
}

// UploadReportPath returns the absolute path of the migrator project's CSV
// upload report file.
func (p ProjectDir) UploadReportPath() string {
	return filepath.Join(string(p), "upload.csv")
}

// RebaseReportPath returns the absolute path of the migrator project's CSV
// rebase report file.
func (p ProjectDir) RebaseReportPath() string {
	return filepath.Join(string(p), "rebase.csv")
}

// ProjectLog returns the absolute path of the scan log for a given LUCI
// project within this migrator project.
func (p ProjectDir) ProjectLog(projectID string) string {
	return filepath.Join(string(p), projectID+".scan.log")
}

// CheckoutTemp returns a temporary checkout directory.
//
// During repo creation, the initial git repo is cloned here and then moved to
// its CheckoutDir() path on success.
func (p ProjectDir) CheckoutTemp(checkoutID string) string {
	return filepath.Join(p.TrashDir(), checkoutID)
}

// CheckoutDir returns the path for a git checkout.
func (p ProjectDir) CheckoutDir(checkoutID string) string {
	return filepath.Join(string(p), checkoutID)
}

// MkTempDir generates a new temporary directory within TrashDir().
func (p ProjectDir) MkTempDir() (string, error) {
	if err := os.Mkdir(p.TrashDir(), 0777); err != nil {
		return "", err
	}
	return ioutil.TempDir(p.TrashDir(), "")
}

// CleanTrash removes TrashDir().
func (p ProjectDir) CleanTrash() error {
	return os.RemoveAll(p.TrashDir())
}

// LoadConfigFile loads the migration project config.
func (p ProjectDir) LoadConfigFile() (*migratorpb.Config, error) {
	blob, err := ioutil.ReadFile(p.ConfigFile())
	if err != nil {
		return nil, errors.Annotate(err, "failed to load the migration project config").Err()
	}
	var cfg migratorpb.Config
	if err := (prototext.UnmarshalOptions{}).Unmarshal(blob, &cfg); err != nil {
		return nil, errors.Annotate(err, "failed to unmarshal the migration project config %q", p.ConfigFile()).Err()
	}
	return &cfg, nil
}

// LoadProjectFilter loads the config and parses `projects_re` field there.
func (p ProjectDir) LoadProjectFilter() (Filter, error) {
	cfg, err := p.LoadConfigFile()
	if err != nil {
		return nil, err
	}
	if len(cfg.ProjectsRe) == 0 {
		return func(string) bool { return true }, nil
	}
	filter, err := NewFilter(cfg.ProjectsRe)
	if err != nil {
		return nil, errors.Annotate(err, "in projects_re").Err()
	}
	return filter, nil
}

// FindProjectRoot finds a migrator ProjectDir starting from `abspath` and
// working up towards the filesystem root.
func FindProjectRoot(abspath string) (ProjectDir, error) {
	curPath := ProjectDir(abspath)
	for {
		if st, err := os.Stat(curPath.ConfigFile()); err == nil {
			if st.Mode().IsRegular() {
				if _, err := curPath.LoadConfigFile(); err != nil {
					return "", errors.Annotate(err, "bad migration project: %q", curPath).Err()
				}
				return curPath, nil
			}
		}
		newPath := ProjectDir(filepath.Dir(string(curPath)))
		if newPath == curPath {
			break
		}
		curPath = newPath
	}
	return "", errors.Reason("not in a migrator project: %q", abspath).Err()
}
