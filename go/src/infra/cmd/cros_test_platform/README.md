# ChromeOS Test Platform recipe step implementations.

This binary's subcommands implement the various phases of the the cros_test_platform recipe, which handles cros_test_platform [requests](https://chromium.googlesource.com/chromiumos/infra/proto/+/refs/heads/master/src/test_platform/request.proto).

A cros_test_platform execution runs through these steps and subcommands, in
this order:

#### `scheduler-traffic-split`
Inspect the parameters of the request, and the current traffic split configuration,
to determine which scheduling backend (autotest or skylab) this request should
be routed to. Also, do any necessary request munging (such as remapping of capacity
pools or accounts) such that clients can be insulated from backend migrations.

#### `enumerate`
Inspect the request's test plan, and the test artifacts that were created for
the build-under-test, to determine which tests to run, with what arguments.

#### `autotest-execute` or `skylab-execute`
Execute the enumerated tests, in the correct backend; wait for them to complete, and collect and summarize their results.


## How to run unittests

Make sure you run set up the environment properly first by running `eval` step
in [the setup guide](https://chromium.googlesource.com/infra/infra/+/master/go/#get-the-code)
first. Note that the `eval` command setup temporary environment for the shell,
so it need to be run every time a new terminal spawned.

After that, invoke `go test` from the [`go/src/infra`](../../) folder, e.g:
`$ go test infra/cmd/cros_test_platform/internal/execution`
