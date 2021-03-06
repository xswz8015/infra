# Monorail pRPC API

This directory holds all the source for the Monorail pRPC API. This API is
implemented using `.proto` files to describe a `gRPC` interface (services,
methods, and request/response messages). It then uses a shim which
converts the
[`gRPC` server](http://www.grpc.io/docs/tutorials/basic/python.html)
(which doesn't work on AppEngine, due to lack of support for HTTP/2) into a
[`pRPC` server](https://godoc.org/github.com/luci/luci-go/grpc/prpc) which
supports communication over HTTP/1.1, as well as text and JSON IO.

## Getting Started

In order to make API requests, your client needs to either:

- Present an OAuth token generated by an allowed client ID or email.
- Provide a XSRF token.
- [For local dev only] Send a test account header, as used by `test_call`
  described below.

## Making requests

You can make anonymous requests to a server running locally like this:

```bash
$ ./api/test_call monorail.Users GetUser '{"email": "test@example.com"}'
```

Requests that require a signed-in user can be tested locally like this:

```bash
$ ./api/test_call monorail.Issues GetIssue \
  '{"issue_ref": {"project_name": "rutabaga", "local_id": 1}}' \
  --test-account=test@example.com
```

## API Documentation

All methods, request parameters, and responses are documented in
[./api_proto](./api_proto).

# Development

## Regenerating Python from Protocol Buffers

In order to regenerate the python server and client stubs from the `.proto`
files, run this command:

```bash
$ make prpc_proto_v0
```
