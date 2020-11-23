#!/usr/bin/env lucicfg
# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""LUCI project configuration for the production instance of LUCI.

After modifying this file execute it ('./main.star') to regenerate the configs.
This is also enforced by PRESUBMIT.py script.

Includes CI configs for the following subprojects:
  * Codesearch.
  * Gatekeeper cron.
  * Gsubtreed crons.
  * WPT autoroller crons.
  * Chromium tarball publisher.
  * Chromium LKGR finder cron.
  * CrOS DUT flashing cron job.
  * https://chromium.googlesource.com/chromium/tools/build
  * https://chromium.googlesource.com/chromium/tools/depot_tools
  * https://chromium.googlesource.com/infra/infra
  * https://chromium.googlesource.com/infra/luci/luci-go
  * https://chromium.googlesource.com/infra/luci/luci-py
  * https://chromium.googlesource.com/infra/luci/recipes-py
  * https://chromium.googlesource.com/infra/testing/expect_tests
"""

lucicfg.check_version("1.20.0", "Please update depot_tools")

# Enable luci.tree_closer.
lucicfg.enable_experiment("crbug.com/1054172")

# Enable LUCI Realms support.
lucicfg.enable_experiment("crbug.com/1085650")

# Tell lucicfg what files it is allowed to touch.
lucicfg.config(
    config_dir = "generated",
    tracked_files = [
        "commit-queue.cfg",
        "cr-buildbucket.cfg",
        "luci-logdog.cfg",
        "luci-milo.cfg",
        "luci-notify.cfg",
        "luci-notify/email-templates/*.template",
        "luci-scheduler.cfg",
        "project.cfg",
        "realms.cfg",
    ],
    fail_on_warnings = True,
    lint_checks = ["default"],
)

luci.project(
    name = "infra",
    buildbucket = "cr-buildbucket.appspot.com",
    logdog = "luci-logdog.appspot.com",
    milo = "luci-milo.appspot.com",
    notify = "luci-notify.appspot.com",
    scheduler = "luci-scheduler.appspot.com",
    swarming = "chromium-swarm.appspot.com",
    acls = [
        # Publicly readable.
        acl.entry(
            roles = [
                acl.BUILDBUCKET_READER,
                acl.LOGDOG_READER,
                acl.PROJECT_CONFIGS_READER,
                acl.SCHEDULER_READER,
            ],
            groups = "all",
        ),
        # Allow committers to use CQ and to force-trigger and stop CI builds.
        acl.entry(
            roles = [
                acl.SCHEDULER_OWNER,
                acl.CQ_COMMITTER,
            ],
            groups = "project-infra-committers",
        ),
        # Ability to launch CQ dry runs.
        acl.entry(
            roles = acl.CQ_DRY_RUNNER,
            groups = "project-infra-tryjob-access",
        ),
        # Group with bots that have write access to the Logdog prefix.
        acl.entry(
            roles = acl.LOGDOG_WRITER,
            groups = "luci-logdog-chromium-writers",
        ),
    ],
)

# Per-service tweaks.
luci.logdog(gs_bucket = "chromium-luci-logdog")
luci.milo(
    logo = "https://storage.googleapis.com/chrome-infra-public/logo/chrome-infra-logo-200x200.png",
    favicon = "https://storage.googleapis.com/chrome-infra-public/logo/favicon.ico",
)
luci.cq(status_host = "chromium-cq-status.appspot.com")
luci.notify(tree_closing_enabled = True)

# Global builder defaults.
luci.builder.defaults.execution_timeout.set(45 * time.minute)
luci.builder.defaults.experiments.set({"luci.use_realms": 100})
luci.builder.defaults.properties.set({"$kitchen": {"emulate_gce": True}})

# Resources shared by all subprojects.

luci.bucket(name = "ci")

luci.bucket(
    name = "try",
    acls = [
        acl.entry(
            roles = acl.BUILDBUCKET_TRIGGERER,
            groups = [
                "project-infra-tryjob-access",
                "service-account-cq",
            ],
        ),
    ],
)

luci.bucket(
    name = "cron",
    acls = [
        acl.entry(
            roles = acl.BUILDBUCKET_TRIGGERER,
            users = [
                # Allow the cros-flash-scheduler to schedule other builders in
                # the bucket, see //subprojects/cros_flash.star.
                "cros-flash@chops-service-accounts.iam.gserviceaccount.com",
            ],
        ),
    ],
)

luci.bucket(
    name = "tasks",
    acls = [
        acl.entry(
            roles = acl.BUILDBUCKET_TRIGGERER,
            users = [
                # The refresh-skew-tests builder will be triggered by the chrome-release
                # builder which uses the service account below.
                "chrome-official-brancher@chops-service-accounts.iam.gserviceaccount.com",
            ],
        ),
    ],
)

luci.realm(name = "pools/cron")

luci.notifier_template(
    name = "status",
    body = "{{ stepNames .MatchingFailedSteps }} on {{ buildUrl . }} {{ .Build.Builder.Builder }}{{ if .Build.Output.GitilesCommit }} from {{ .Build.Output.GitilesCommit.Id }}{{end}}",
)

luci.list_view(name = "cron")
luci.list_view(name = "tasks")

# Setup Swarming permissions (in particular for LED).

load("//lib/led.star", "led")

led.users(
    groups = "flex-ci-led-users",
    task_realm = "ci",
)

led.users(
    groups = "flex-try-led-users",
    task_realm = "try",
)

led.users(
    groups = "mdb/chrome-troopers",
    task_realm = "cron",
    pool_realm = "pools/cron",
)

# Per-subproject resources. They may refer to the shared resources defined
# above by name.

exec("//subprojects/build.star")
exec("//subprojects/codesearch.star")
exec("//subprojects/depot_tools.star")
exec("//subprojects/expect_tests.star")
exec("//subprojects/gatekeeper.star")
exec("//subprojects/infra.star")
exec("//subprojects/lkgr.star")
exec("//subprojects/luci-go.star")
exec("//subprojects/luci-py.star")
exec("//subprojects/recipe_engine.star")
exec("//subprojects/tarballs.star")
exec("//subprojects/wpt.star")
exec("//subprojects/weblayer.star")
