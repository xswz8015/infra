// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/server/router"

	"infra/appengine/weetbix/internal/config/compiledcfg"
)

// Handlers provides methods servicing Weetbix HTTP routes.
type Handlers struct {
	cloudProject string
	// prod is set when running in production (not a dev workstation).
	prod bool
}

// NewHandlers initialises a new Handlers instance.
func NewHandlers(cloudProject string, prod bool) *Handlers {
	return &Handlers{cloudProject: cloudProject, prod: prod}
}

func obtainProjectConfigOrError(ctx *router.Context) (project string, cfg *compiledcfg.ProjectConfig, ok bool) {
	projectID := ctx.Params.ByName("project")
	cfg, err := compiledcfg.Project(ctx.Context, projectID, time.Time{})
	if err != nil {
		if err == compiledcfg.NotExistsErr {
			http.Error(ctx.Writer, "Project does not exist in Weetbix.", http.StatusBadRequest)
			return "", nil, false
		}
		logging.Errorf(ctx.Context, "Obtain project config: %v", err)
		http.Error(ctx.Writer, "Internal server error.", http.StatusInternalServerError)
		return "", nil, false
	}
	return projectID, cfg, true
}

func obtainProjectOrError(ctx *router.Context) (project string, ok bool) {
	project, _, ok = obtainProjectConfigOrError(ctx)
	return project, ok
}

func respondWithJSON(ctx *router.Context, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		logging.Errorf(ctx.Context, "Marshalling JSON for response: %s", err)
		http.Error(ctx.Writer, "Internal server error.", http.StatusInternalServerError)
	}
	ctx.Writer.Header().Add("Content-Type", "application/json")
	if _, err := ctx.Writer.Write(bytes); err != nil {
		logging.Errorf(ctx.Context, "Writing JSON response: %s", err)
	}
}
