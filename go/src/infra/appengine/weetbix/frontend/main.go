// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"net/http"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/config/server/cfgmodule"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/cron"
	"go.chromium.org/luci/server/gaeemulation"
	"go.chromium.org/luci/server/module"
	"go.chromium.org/luci/server/router"

	"infra/appengine/weetbix/internal/bugs"
	"infra/appengine/weetbix/internal/config"
)

func init() {
	datastore.EnableSafeGet()
}

func main() {
	modules := []module.Module{
		cfgmodule.NewModuleFromFlags(),
		cron.NewModuleFromFlags(),
		gaeemulation.NewModuleFromFlags(),
	}
	server.Main(nil, modules, func(srv *server.Server) error {
		srv.Routes.GET("/", nil, func(c *router.Context) {
			logging.Debugf(c.Context, "Hello world")
			cfg, err := config.Get(c.Context)
			if err != nil {
				http.Error(c.Writer, "Internal server error.", http.StatusInternalServerError)
				return
			}
			cl, err := bugs.NewMonorailClient(c.Context, cfg.GetMonorailHostname())
			if err != nil {
				logging.Errorf(c.Context, "Getting Monorail client: %v", err)
				http.Error(c.Writer, "Internal server error.", http.StatusInternalServerError)
				return
			}
			issue, err := cl.GetIssue(c.Context, "projects/chromium/issues/2")
			if err != nil {
				logging.Errorf(c.Context, "Getting Monorail issue: %v", err)
				http.Error(c.Writer, "Internal server error.", http.StatusInternalServerError)
				return
			}
			c.Writer.Write([]byte("Hello, world! This is Weetbix.\n"))
			c.Writer.Write([]byte("Issue title: " + issue.GetSummary()))
		})

		// GAE crons.
		cron.RegisterHandler("read-config", config.Update)

		return nil
	})
}