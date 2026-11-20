---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Docker

We publish a [public Docker image `moov/fincen`](https://hub.docker.com/r/moov/fincen/) from Docker Hub or use this repository.
<!--
We also have Docker images for [OpenShift](https://quay.io/repository/moov/fincen?tab=tags) published as `quay.io/moov/fincen`.
-->

Pull & start the Docker image:
```
docker pull moov/fincen:latest
docker run -p 8206:8206 -p 8207:8207 moov/fincen:latest
```

Check http server alive:
```
curl localhost:8206/ping
```
```
PONG
```

The `/validator` and `/reformat` endpoints accept `multipart/form-data` with a file field named `input`.

Validate a file on the HTTP server:
```
curl -F "input=@./data/samples/sar_batch.xml" http://localhost:8206/validator
```
```
{"status":"valid file"}
```

The same command works with the `ctr_batch.xml`, `doep_batch.xml`, `fbar_batch.xml`, and `8300_batch.xml` samples.

Reformat the file with generated attributes. Attributes are regenerated only when you send `generate_attrs=true`. You can also send `format=json` for JSON output. `/validator` accepts `disable_attrs=true` to skip attribute checks.

```
curl -F "input=@./data/samples/ctr_batch.xml" -F "generate_attrs=true" http://localhost:8206/reformat
```
```
<fc2:EFilingBatchXML TotalAmount="47000" PartyCount="6" ActivityCount="1" xsi:schemaLocation="www.fincen.gov/base https://www.fincen.gov/sites/default/files/schema/base/EFL_CTRXBatchSchema.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:fc2="www.fincen.gov/base">
  <fc2:FormTypeCode>CTRX</fc2:FormTypeCode>
  <fc2:Activity SeqNum="1">
    ...
```
