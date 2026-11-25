[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/fincen/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/fincen/api/#get-/ping">API Endpoints</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://pkg.go.dev/badge/github.com/moov-io/fincen.svg)](https://pkg.go.dev/github.com/moov-io/fincen)
[![Build Status](https://github.com/moov-io/fincen/workflows/Go/badge.svg)](https://github.com/moov-io/fincen/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/fincen/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/fincen)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/fincen)](https://goreportcard.com/report/github.com/moov-io/fincen)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/fincen?label=project%20size)](https://github.com/moov-io/fincen)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/fincen/master/LICENSE)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/fincen)](https://hub.docker.com/r/moov/fincen)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/fincen)](https://github.com/moov-io/fincen)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# fincen
Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

Moov's fincen project implements a reader, writer, and validator for Fincen BSA forms in an HTTP server and Go library.

Fincen (Financial Crimes Enforcement Network) BSA data transmission methods for the BSA E-Filing System.

A go library for reading and writing Fincen BSA forms. It is capable of generating, validating, and batching submissions.

The HTTP server is available in a Docker image and the Go package github.com/moov-io/fincen is available.


## Table of contents

- [Project status](#project-status)
- [Supported Forms](#supported-forms)
- [Usage](#usage)
    - As an API
        - [Docker](#docker)
        - [Google Cloud](#google-cloud-run)
        - [HTTP API](https://moov-io.github.io/fincen/usage-docker/)
    - [As a Go module](#go-library)
      - [Build report form](#build-report-form)
    - [As a command line tool](#command-line)
    - [As an in-browser parser](#in-browser-fincen-form-parser)
- [Learn About Fincen](#learn-about-fincen)
- [FAQ](#faq)
- [Getting help](#getting-help)
- [Supported and tested platforms](#supported-and-tested-platforms)
- [Contributing](#contributing)
- [License](#license)

## Project status

The project is Fincen which is a financial reporting for the United States.

There are many forms that are required in this project. **Now implemented XML forms only in following link**.
(Non XML form will implement in another sub project)
https://bsaefiling.fincen.treas.gov/FilingInformation.html

The final phase of the project will be a service that collects fillings and then on a scheduled time will batch the fillings and submit them electronically.
https://bsaefiling.fincen.treas.gov/SDTMInfo.html

## Supported Forms

- FinCEN Currency Transaction Report (FinCEN Report 112)
- FinCEN Designation of Exempt Person (FinCEN Report 110)
- FinCEN Suspicious Activity Report (FinCEN Report 111)
- Report of Foreign Bank and Financial Accounts (FinCEN Report 114)
- Report of Cash Payments Over $10,000 Received in a Trade or Business (FinCEN Form 8300)

## Usage
The Fincen project implements an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/fincen) for creating and modifying Fincen BSA forms.

### Docker

We publish a [public Docker image `moov/fincen`](https://hub.docker.com/r/moov/fincen/tags) on Docker Hub with every tagged release of Fincen.
<!--
We also have Docker images for [OpenShift](https://quay.io/repository/moov/fincen?tab=tags) published as `quay.io/moov/fincen`.
-->

Pull & start the Docker image:
```
docker pull moov/fincen:latest
docker run -p 8206:8206 -p 8207:8207 moov/fincen:latest
```

The HTTP server listens on `:8206`. The admin server (health probes and metrics) listens on `:8207`. The `/validator` and `/reformat` endpoints accept `multipart/form-data` with a file field named `input`.

Validate a file on the HTTP server:
```
curl -F "input=@./data/samples/ctr_batch.xml" http://localhost:8206/validator
```
```
{"status":"valid file"}
```

The same command works with the `sar_batch.xml`, `doep_batch.xml`, `fbar_batch.xml`, and `8300_batch.xml` samples.

Reformat the file with generated attributes. Attributes are regenerated only when you send `generate_attrs=true`. You can also send `format=json` for JSON output. `/validator` accepts `disable_attrs=true` to skip attribute checks.

```
curl -F "input=@./data/samples/ctr_batch.xml" -F "generate_attrs=true" http://localhost:8206/reformat
```
```
<fc2:EFilingBatchXML TotalAmount="47000" PartyCount="6" ActivityCount="1" xsi:schemaLocation="www.fincen.gov/base https://www.fincen.gov/system/files/schema/base/EFL_CTRXBatchSchema.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:fc2="www.fincen.gov/base">
  <fc2:FormTypeCode>CTRX</fc2:FormTypeCode>
  <fc2:Activity SeqNum="1">
    ...
```

The `/validator` and `/reformat` endpoints return these HTTP status codes:

| Status | Meaning |
|--------|---------|
| `200` | The file is valid (`/validator`), or the reformatted report is in the response body (`/reformat`). |
| `400` | The request is not `multipart/form-data`, the `input` field is missing, or the file is not an XML or JSON batch report. |
| `422` | The file was read, but it does not pass validation. The `error` field shows the problem, for example `{"error":"The TotalAmount has invalid value"}`. |

Version v0.4.0 and earlier returned `501` for files that could not be read and for files that did not pass validation.

### Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-fincen-demo
```

Enable the [Container Registry](https://cloud.google.com/container-registry) API for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following Docker commands, substituting your unique project ID:

```
docker pull moov/fincen
docker tag moov/fincen gcr.io/<PROJECT-ID>/fincen
docker push gcr.io/<PROJECT-ID>/fincen
```

Deploy the container to Cloud Run:
```
gcloud run deploy --image gcr.io/<PROJECT-ID>/fincen --port 8206
```

Select your target platform to `1`, service name to `fincen`, and region to the one closest to you (enable Google API service if a prompt appears). Upon a successful build you will be given a URL where the API has been deployed:

```
https://YOUR-FINCEN-APP-URL.a.run.app
```

Now you can check that the service is running:
```
curl https://YOUR-FINCEN-APP-URL.a.run.app/ping
```
You should get this response:
```
PONG
```

The HTTP API provides `/ping`, `/validator`, and `/reformat`. There is no `/files` endpoint.

### Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go 1.26 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/fincen/releases/latest) as well. We highly recommend you use a tagged release for production once a release is cut from current `master`. The GitHub release marked Latest is older than `master`, and `go get` may resolve `v0.3.6`, which is not on the current `master` branch.

```
git clone https://github.com/moov-io/fincen.git

go get github.com/moov-io/fincen
```
The package [`github.com/moov-io/fincen`](https://pkg.go.dev/github.com/moov-io/fincen) offers a Go-based Fincen file reader and writer.

### Build report form

Create a batch with `batch.NewReport` and a form constant:

| Constant | Form type |
| --- | --- |
| `fincen.Report112` | CTRX (Currency Transaction Report) |
| `fincen.Report111` | SARX (Suspicious Activity Report) |
| `fincen.Report110` | DOEPX (Designation of Exempt Person) |
| `fincen.Report114` | FBARX (Foreign Bank and Financial Accounts) |
| `fincen.Form8300` | 8300X (Cash Payments Over $10,000) |
| `fincen.ReportSubmission` | SUBMISSION (acknowledgement) |

Generate attributes, then generate sequence numbers, then validate. `Validate()` fails when counts are stale. `Validate()` checks the library's field and attribute rules only. It does not validate a document against the bundled XSD files in `data/xsd`. A successful `Validate()` result (or an HTTP `/validator` response of `valid file`) does not mean that FinCEN would accept the file.

A complete CTR example is in [`examples/currency_transaction_write`](examples/currency_transaction_write). Run examples from the repository root:

```
go run ./examples/currency_transaction_write
go run ./examples/currency_transaction_read
```

```go
package main

import (
    "encoding/xml"
    "fmt"
    "log"

    "github.com/moov-io/fincen"
    "github.com/moov-io/fincen/pkg/batch"
    "github.com/moov-io/fincen/pkg/currency_transaction"
)

func main() {
    report := batch.NewReport(fincen.Report112)
    activity := currency_transaction.NewActivity()
    // Fill activity fields. See examples/currency_transaction_write for a complete CTR.

    if err := report.AppendActivity(activity); err != nil {
        log.Fatal(err)
    }
    if err := report.GenerateAttrs(); err != nil {
        log.Fatal(err)
    }
    if err := report.GenerateSeqNumbers(); err != nil {
        log.Fatal(err)
    }
    if err := report.Validate(); err != nil {
        log.Fatal(err)
    }

    out, err := xml.MarshalIndent(report, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
}
```

### Command Line

The `fincen-<os>-amd64` asset on each [release](https://github.com/moov-io/fincen/releases/latest) is the HTTP server, not the CLI. The command-line tool lives in `cmd/fincen`. Build it from this repository, or install it with Go from current `master` until a new tagged release includes the CLI:

```
go install github.com/moov-io/fincen/cmd/fincen@master

# or, from a clone of this repository:
go build -o bin/fincen ./cmd/fincen
```

```
Work seamlessly with Fincen BSA forms from the command line.

Usage:
  fincen <command> [flags]

Available commands:
  summary: display form summary
  validate: validate financial report form
  reformat: reformat financial report form

Reformat flags:
  -generate-attrs
        regenerate batch attributes
  -format string
        output format, xml or json (default "xml")
```

```
./bin/fincen validate data/samples/ctr_batch.xml
./bin/fincen reformat -generate-attrs data/samples/ctr_batch.xml
```

### In-browser Fincen form parser
Using our [in-browser utility](https://moov-io.github.io/fincen/webui/), you can instantly verify and reformat Fincen BSA forms.

## Learn about Fincen
- [About the BSA E-Filing System](https://bsaefiling.fincen.treas.gov/AboutBsa.html)
- [Secure Direct Transfer Mode](https://bsaefiling.fincen.treas.gov/SDTMInfo.html)
- [BSA E-Filing System](https://bsaefiling.fincen.treas.gov/)


## FAQ
<details open="true">
<summary ><b>Is there an in-browser tool?</b></summary>
Yes! You can find our browser utility at https://moov-io.github.io/fincen/webui/.
</details>

## Getting help

channel | info
 ------- | -------
[Project Documentation](https://moov-io.github.io/fincen/) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/fincen/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

Note: 32-bit platforms have known issues and are not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) to get started!

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go 1.26 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/fincen/releases/latest) as well. We highly recommend you use a tagged release for production once a release is cut from current `master`.


## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
