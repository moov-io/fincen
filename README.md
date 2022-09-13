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
Fincen (Financial crimes enforcment network) BSA data transmission methods for the the BSA E-Filing System

A go library for reading and writing Fincen BSA forms. It is capable of generating, validating, and batching submissions.

The project is Fincen which is a financial reporting for the United States. There are many forms that are required in this project but we will start with SAR(Suspicious activity report). I would like a go library that creates the structs to build and validate the SAR form. It also needs to be able to export the file to XML.

In the future we will wrap the service with json/http services so that it can work as a web service and not just a go library.
New project
https://github.com/moov-io/fincen/blob/master/README.md

We will end up with a sub project for each of the form parsers. This is all of the forms.
https://bsaefiling.fincen.treas.gov/FilingInformation.html

The final phase of the project will be a service that collects fillings and then on a scheduled time will batch the fillings and submit them electronically.
https://bsaefiling.fincen.treas.gov/docs/SDTMRequirements.pdf


[About the BSA E-Filing System](https://bsaefiling.fincen.treas.gov/AboutBsa.html)

## Supported Forms

- FinCEN Currency Transaction Report (FinCEN Report 112)
- FinCEN Designation of Exempt Person (FinCEN Report 110)
- FinCEN Suspicious Activity Report (FinCEN Report 111)
- FinCEN Registration of Money Services Business (FinCEN Report 107)
- Report of Foreign Bank and Financial Accounts (FinCEN Report 114)
- Report of Cash Payments Over $10,000 Received in a Trade or Business (FinCEN Form 8300)

## Filing information user guides

The [Filing Information](https://bsaefiling.fincen.treas.gov/FilingInformation.html) website contains PDF's of each of the XML formates. The PDF's guides contain links to XSD's of the formts.

## Become a BSA E-Filer

[Secure Direct Transfer Mode](https://bsaefiling.fincen.treas.gov/SDTMInfo.html)

[Supervisory User Registration](https://bsaefiling1.fincen.treas.gov/AddUser)

## Getting help

 channel | info
 ------- | -------
[Project Documentation](https://moov-io.github.io/fincen/) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/fincen/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
