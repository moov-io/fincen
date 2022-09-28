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
docker run -p 8080:8080 -p 9090:9090 moov/fincen:latest
```

Check http server alive:
```
curl localhost:8080/ping
```
```
PONG
```

Create a file on the HTTP server:
```
curl -X POST --data-binary "@./data/samples/sar_batch.xml" http://localhost:8080/validator
```
```
{"status":"valid file"}
```

Reformat the file with generated attributes:
```
curl -X POST --data-binary "@./data/samples/ctr_batch.txt" http://localhost:8088/reformat
```
```
<EFilingBatchXML ActivityCount="1" TotalAmount="47000" PartyCount="6" SeqNum="1"
                     xsi:schemaLocation="www.fincen.gov/base https://www.fincen.gov/base https://www.fincen.gov/base/EFL_8300XBatchSchema.xsd"
                     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                     xmlns:fc2="www.fincen.gov/base">
    <Activity SeqNum="1">
...
```