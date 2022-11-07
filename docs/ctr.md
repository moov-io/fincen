---
layout: page
title: Currency Transaction Report (CTR) - Report 112 | Moov FinCEN
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Overview

Currency Transaction Report (CTR) - Report 112

# Create a report

Currency transaction report can create using fincen go library

1. Create a [EFilingBatchXML](https://godoc.org/github.com/moov-io/fincen/pkg/batch#EFilingBatchXML) with `batch.NewReport(fincen.Report112)`.
2. Create available [ActivityType](https://godoc.org/github.com/moov-io/pkg/currency_transaction#ActivityType) records with `currency_transaction.NewActivity()`.
3. Append created activities into Batch XML report with `batch.AppendActivity(activity)`.
4. Validate Batch XML report with `Validate()` and figure out report problems.
5. Generate Batch XML report attributes with `GenerateAttrs()`
6. Generate Batch XML report attributes with `GenerateSeqNumbers()`
7. Getting xml contents from Batch XML report.

# Create an acknowledgement

FinCEN SAR XML batch acknowledgement can create using fincen go library

1. Create a [EFilingBatchXML](https://godoc.org/github.com/moov-io/fincen/pkg/batch#EFilingBatchXML) with `batch.NewReport(fincen.Report112)`.
2. Create a [EFilingSubmissionXML](https://godoc.org/github.com/moov-io/pkg/batch#EFilingSubmissionXML).
3. Validate Batch XML report with `Validate()` and figure out report problems.
4. Generate Batch XML report attributes with `GenerateAttrs()`
5. Getting xml contents from Batch XML report.

# Examples
1. [Read](https://github.com/moov-io/fincen/tree/master/examples/currency_transaction_read/main.go)
2. [Write](https://github.com/moov-io/fincen/tree/master/examples/currency_transaction_write/main.go)
