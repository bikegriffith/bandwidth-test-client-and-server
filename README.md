# Bandwidth Test Client and Server

Status: Very Early Alpha. Working on upload speed testing prior to download (as the latter
already has more broad support in the community).

## Server

Runs a simple golang HTTP server that will echo the request size. This allows the client to
compute round trip duration as part of the upload speed calculation.

A Makefile and Dockerfile is provided to allow this easily to run in a container.

```
make build
make run
```

## Clients

### node.js

Simple library that hits the go upload server and computes a response

```
# XXX: not yet published on NPM
var tester = require('./nodejs-client/index.js');
tester.simpleUploadTest()
```

```
