// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/maruel/subcommands"
	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/client/authcli"
	swarming "go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.chromium.org/luci/common/data/strpair"
	"go.chromium.org/luci/common/errors"

	"infra/cmd/skylab/internal/site"
)

const progName = "skylab"

type taskPriority struct {
	name  string
	level int
}

var taskPriorityMap = map[string]int{
	"Weekly":    230,
	"CTS":       215,
	"Daily":     200,
	"PostBuild": 170,
	"Default":   140,
	"Build":     110,
	"PFQ":       80,
	"CQ":        50,
	"Super":     49,
}
var defaultTaskPriorityKey = "Default"
var defaultTaskPriority = taskPriorityMap[defaultTaskPriorityKey]

type commonFlags struct {
	debug bool
}

func (f *commonFlags) Register(fl *flag.FlagSet) {
	fl.BoolVar(&f.debug, "debug", false, "Enable debug output")
}

func (f commonFlags) DebugLogger(a subcommands.Application) *log.Logger {
	out := ioutil.Discard
	if f.debug {
		out = a.GetErr()
	}
	return log.New(out, progName, log.LstdFlags|log.Lshortfile)
}

type envFlags struct {
	dev bool
}

func (f *envFlags) Register(fl *flag.FlagSet) {
	fl.BoolVar(&f.dev, "dev", false, "Run in dev environment")
}

func (f envFlags) Env() site.Environment {
	if f.dev {
		return site.Dev
	}
	return site.Prod
}

// httpClient returns an HTTP client with authentication set up.
func httpClient(ctx context.Context, f *authcli.Flags) (*http.Client, error) {
	o, err := f.Options()
	if err != nil {
		return nil, errors.Annotate(err, "failed to get auth options").Err()
	}
	a := auth.NewAuthenticator(ctx, auth.OptionalLogin, o)
	c, err := a.Client()
	if err != nil {
		return nil, errors.Annotate(err, "failed to create HTTP client").Err()
	}
	return c, nil

}

const swarmingAPISuffix = "_ah/api/swarming/v1/"

func newSwarmingService(ctx context.Context, auth authcli.Flags, env site.Environment) (*swarming.Service, error) {
	cl, err := httpClient(ctx, &auth)
	if err != nil {
		return nil, errors.Annotate(err, "create swarming client").Err()
	}

	s, err := swarming.New(cl)
	if err != nil {
		return nil, errors.Annotate(err, "create swarming client").Err()
	}

	s.BasePath = env.SwarmingService + swarmingAPISuffix
	return s, nil
}

func swarmingTaskURL(e site.Environment, taskID string) string {
	return fmt.Sprintf("%stask?id=%s", e.SwarmingService, taskID)
}

// UserErrorReporter reports a detailed error message to the user.
//
// PrintError() uses a UserErrorReporter to print multi-line user error details
// along with the actual error.
type UserErrorReporter interface {
	// Report a user-friendly error through w.
	ReportUserError(w io.Writer)
}

// PrintError reports errors back to the user.
//
// Detailed error information is printed if err is a UserErrorReporter.
func PrintError(w io.Writer, err error) {
	if u, ok := err.(UserErrorReporter); ok {
		u.ReportUserError(w)
	} else {
		fmt.Fprintf(w, "%s\n", err)
	}
}

// NewUsageError creates a new error that also reports flags usage error
// details.
func NewUsageError(flags flag.FlagSet, format string, a ...interface{}) error {
	return &usageError{
		error: fmt.Errorf(format, a...),
		flags: flags,
	}
}

type usageError struct {
	error
	flags flag.FlagSet
}

func (e *usageError) ReportUserError(w io.Writer) {
	fmt.Fprintf(w, "%s\n", e.error)
	e.flags.Usage()
}

// generateAnnotationURL generates a unique logdog url for use as a the logdog annotation endpoint
// of a skylab swarming task (i.e. the -logdog-annotation-url argument to skylab_swarming_worker).
func generateAnnotationURL(e site.Environment) string {
	u := uuid.New()
	return fmt.Sprintf("logdog://%s/%s/skylab/%s/+/annotations",
		e.LogDogHost, e.LUCIProject, u.String())
}

// newTaskRequest creates a new swarming task request for a skylab task.
func newTaskRequest(taskName string, tags []string, slices []*swarming.SwarmingRpcsTaskSlice,
	priority int64) *swarming.SwarmingRpcsNewTaskRequest {
	return &swarming.SwarmingRpcsNewTaskRequest{
		Name:       taskName,
		Tags:       tags,
		TaskSlices: slices,
		Priority:   priority,
	}
}

// toPairs converts a slice of strings in foo:bar form to a slice of swarming rpc string pairs.
func toPairs(dimensions []string) ([]*swarming.SwarmingRpcsStringPair, error) {
	pairs := make([]*swarming.SwarmingRpcsStringPair, len(dimensions))
	for i, d := range dimensions {
		k, v := strpair.Parse(d)
		if v == "" {
			return nil, fmt.Errorf("malformed dimension with key '%s' has no value", k)
		}
		pairs[i] = &swarming.SwarmingRpcsStringPair{Key: k, Value: v}
	}
	return pairs, nil
}

func toKeyvalMap(keyvals []string) (map[string]string, error) {
	m := make(map[string]string, len(keyvals))
	for _, s := range keyvals {
		k, v := strpair.Parse(s)
		if v == "" {
			return nil, fmt.Errorf("malformed keyval with key '%s' has no value", k)
		}
		if _, ok := m[k]; ok {
			return nil, fmt.Errorf("keyval with key %s specified more than once", k)
		}
		m[k] = v
	}
	return m, nil
}

func sortedPriorities() []taskPriority {
	s := make([]taskPriority, 0, len(taskPriorityMap))
	for k, v := range taskPriorityMap {
		s = append(s, taskPriority{k, v})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].level < s[j].level
	})
	return s
}

func sortedPriorityKeys() []string {
	sp := sortedPriorities()
	k := make([]string, 0, len(sp))
	for _, p := range sp {
		k = append(k, p.name)
	}
	return k
}
