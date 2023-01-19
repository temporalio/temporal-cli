---
id: describe
title: temporal workflow describe
sidebar_label: describe
description: words words words
tags:
	- cli
---

### describe

Show information about a Workflow Execution.

>This information can be used to locate a Workflow Execution that failed.

**--address**
host:port for Temporal frontend service

**--codec-auth**
Authorization header to set for requests to Codec Server

**--codec-endpoint**
Remote Codec Server Endpoint

**--color**
when to use color: auto, always, never. (default: auto)

**--context-timeout**
Optional timeout for context of RPC call in seconds (default: 5)

**--env**
Env name to read the client environment variables from (default: default)

**--grpc-meta**
gRPC metadata to send with requests. Format: key=value. Use valid JSON formats for value

**--namespace**
Alias: **-n**
Temporal workflow namespace (default: default)

**--raw**
Print properties as they are stored.

**--reset-points**
Only show auto-reset points.

**--run-id**
Alias: **-r**
Run Id

**--tls-ca-path**
Path to server CA certificate

**--tls-cert-path**
Path to x509 certificate

**--tls-disable-host-verification**
Disable tls host name verification (tls must be enabled)

**--tls-key-path**
Path to private key

**--tls-server-name**
Override for target server name

**--workflow-id**
Alias: **-w**
Workflow Id
