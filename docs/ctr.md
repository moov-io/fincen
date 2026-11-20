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

1. Create a [EFilingBatchXML](https://pkg.go.dev/github.com/moov-io/fincen/pkg/batch#EFilingBatchXML) with `batch.NewReport(fincen.Report112)`.
2. Create available [ActivityType](https://pkg.go.dev/github.com/moov-io/fincen/pkg/currency_transaction#ActivityType) records with `currency_transaction.NewActivity()`.
3. Append created activities into Batch XML report with `batch.AppendActivity(activity)`.
4. Generate batch XML report attributes with `GenerateAttrs()`.
5. Generate sequence numbers with `GenerateSeqNumbers()`.
6. Validate the batch XML report with `Validate()`.
7. Get the XML contents from the batch XML report.

# Create an acknowledgement

FinCEN CTR XML batch acknowledgement can be created using the fincen Go library

1. Create a [EFilingBatchXML](https://pkg.go.dev/github.com/moov-io/fincen/pkg/batch#EFilingBatchXML) with `batch.NewReport(fincen.Report112)`.
2. Create a [EFilingSubmissionXML](https://pkg.go.dev/github.com/moov-io/fincen/pkg/batch#EFilingSubmissionXML).
3. Generate batch XML report attributes with `GenerateAttrs()`.
4. Validate the batch XML report with `Validate()`.
5. Get the XML contents from the batch XML report.

# Examples
1. [Read](https://github.com/moov-io/fincen/tree/master/examples/currency_transaction_read/main.go)
2. [Write](https://github.com/moov-io/fincen/tree/master/examples/currency_transaction_write/main.go)

# Filing

Refer to the [FinCEN.gov website](https://www.fincen.gov/resources/filing-information) for more information.
