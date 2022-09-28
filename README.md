[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/fincen/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/fincen/api/#get-/files">API Endpoints</a>
  ·
  <a href="https://moov.io/blog/education/fincen-api-guide/">API Guide</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/fincen?status.svg)](https://godoc.org/github.com/moov-io/fincen)
[![Build Status](https://github.com/moov-io/fincen/workflows/Go/badge.svg)](https://github.com/moov-io/fincen/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/fincen/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/fincen)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/fincen)](https://goreportcard.com/report/github.com/moov-io/fincen)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/fincen?label=project%20size)](https://github.com/moov-io/fincen)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/fincen/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/fincen)](https://hub.docker.com/r/moov/fincen)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/fincen)](https://github.com/moov-io/fincen)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# fincen
Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

Moov's fincen project implements a reader, writer, and validator for Fincen BSA forms in an HTTP server and Go library.

Fincen (Financial crimes enforcment network) BSA data transmission methods for the the BSA E-Filing System.

A go library for reading and writing Fincen BSA forms. It is capable of generating, validating, and batching submissions.

The HTTP server is available in a Docker image and the Go package github.com/moov-io/fincen is available.


## Table of contents

- [Project status](#project-status)
- [Supported Forms](#supported-forms)
- [Usage](#usage)
    - As an API
        - [Docker](#docker)
        - [Google Cloud](#google-cloud-run-button)
        - [HTTP API](#http-api)
        - [Data persistence](#data-persistence)
    - [As a Go module](#go-library)
      - [Build report form](#build-report-form)
    - [As a command line tool](#command-line)
    - [As an in-browser parser](##in-browser-fincen-form-parser)
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
https://bsaefiling.fincen.treas.gov/docs/SDTMRequirements.pdf

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
We also have Docker images for [OpenShift](https://quay.io/repository/moov/fincen?tab=tags) published as `quay.io/moov/fincen`.

Pull & start the Docker image:
```
docker pull moov/fincen:latest
docker run -p 8088:8088 -p 9098:9098 moov/fincen:latest
```

Validate a file on the HTTP server:
```
curl -X POST --data-binary "@./data/samples/ctr_batch.txt" http://localhost:8088/validator
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
gcloud run deploy --image gcr.io/<PROJECT-ID>/fincen --port 8088
```

Select your target platform to `1`, service name to `fincen`, and region to the one closest to you (enable Google API service if a prompt appears). Upon a successful build you will be given a URL where the API has been deployed:

```
https://YOUR-FINCEN-APP-URL.a.run.app
```

Now you can list files stored in-memory:
```
curl https://YOUR-FINCEN-APP-URL.a.run.app/files
```
You should get this response:
```
null
```

### Data persistence

By design, Fincen  **does not persist** (save) any data about the files or entry details created. The only storage occurs in memory of the process and upon restart Fincen will have no files or data saved. Also, no in-memory encryption of the data is performed.

### Go library

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/fincen/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/fincen.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/fincen
```
The package [`github.com/moov-io/fincen`](https://pkg.go.dev/github.com/moov-io/fincen) offers a Go-based Fincen file reader and writer.

### Build report form

**Creating XML Batch Reporting form**

Fincen project used general XML batch reporting form struct.
Available types are "SUBMISSION", "CTRX", "SARX", "DOEPX", "FBARX", "8300X" 

```
// create report with type
newReport := NewReport("SUBMISSION")
```

**Adding activities by each type**

Activity should add into the form struct using sub package 

```
// create activity (ctr)
var newActivity currency_transaction.ActivityType

...
setting newActivity
...

err := newReport.AppendActivity(&newActivity)
```

**Generating new xml attributes**
```
err = newReport.GenerateAttrs()
```

**Validating report form**
```
err = newReport.Validate()
```

### In-browser Fincen form parser
Using our [in-browser utility](http://oss.moov.io/fincen/), you can instantly verify and reformat Fincen BSA forms.

## Learn about Fincen
- [About the BSA E-Filing System](https://bsaefiling.fincen.treas.gov/AboutBsa.html)
- [Secure Direct Transfer Mode](https://bsaefiling.fincen.treas.gov/SDTMInfo.html)
- [Supervisory User Registration](https://bsaefiling1.fincen.treas.gov/AddUser)


## FAQ
<details open="true">
<summary ><b>Is there an in-browser tool?</b></summary>
Yes! You can find our browser utility at http://oss.moov.io/fincen/.
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

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/fincen/releases/latest) as well. We highly recommend you use a tagged release for production.


## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
