---
id: describe
title: temporal batch describe
sidebar_label: describe
description: words words words
tags:
	- cli
---

### describe

Describe a Batch operation job.

>This command shows the progress of an ongoing Batch job.

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

**--fields**
customize fields to print. Set to 'long' to automatically print more of main fields

**--grpc-meta**
gRPC metadata to send with requests. Format: key=value. Use valid JSON formats for value

**--job-id**
Batch Job Id

**--namespace**
Alias: **-n**
Temporal workflow namespace (default: default)

**--output**
Alias: **-o**
format output as: table, json, card. (default: table)

**--time-format**
format time as: relative, iso, raw. (default: relative)

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
