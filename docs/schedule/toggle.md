---
id: toggle
title: temporal schedule toggle
sidebar_label: toggle
description: words words words
tags:
	- cli
---

### toggle

Pauses or unpauses a schedule

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

**--pause**
Pauses the schedule

**--reason**
Free-form text to describe reason for pause/unpause (default: (no reason provided))

**--schedule-id**
Alias: **-s**
Schedule Id

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

**--unpause**
Unpauses the schedule
