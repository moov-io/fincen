---
layout: page
title: Cash Payments over $10k (FinCEN Form 8300) | Moov FinCEN
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Overview

Cash Payments over $10k (FinCEN Form 8300)


Cash payments report can create using fincen go library

1. Create a [EFilingBatchXML](https://godoc.org/github.com/moov-io/fincen/pkg/batch#EFilingBatchXML) with `batch.NewReport(fincen.Form8300)`.
2. Create available [ActivityType](https://godoc.org/github.com/moov-io/pkg/cash_payments#ActivityType) records with `cash_payments.NewActivity()`.
3. Append created activities into Batch XML report with `batch.AppendActivity(activity)`.
4. Validate Batch XML report with `Validate()` and figure out report problems.
5. Generate Batch XML report attributes with `GenerateAttrs()`
6. Generate Batch XML report attributes with `GenerateSeqNumbers()`
7. Getting xml contents from Batch XML report.

# Create an acknowledgement

FinCEN SAR XML batch acknowledgement can create using fincen go library

1. Create a [EFilingBatchXML](https://godoc.org/github.com/moov-io/fincen/pkg/batch#EFilingBatchXML) with `batch.NewReport(fincen.Form8300)`.
2. Create a [EFilingSubmissionXML](https://godoc.org/github.com/moov-io/pkg/batch#EFilingSubmissionXML).
3. Validate Batch XML report with `Validate()` and figure out report problems.
4. Generate Batch XML report attributes with `GenerateAttrs()`
5. Getting xml contents from Batch XML report.

# Examples
1. [Read](https://github.com/moov-io/fincen/tree/master/examples/cash_payment_read/main.go)
2. [Write](https://github.com/moov-io/fincen/tree/master/examples/cash_payment_write/main.go)
