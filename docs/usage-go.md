---
layout: page
title: Go Library Usage
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go Library

See [Golang's install instructions](https://golang.org/doc/install) for help in setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/fincen/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/fincen.git

$ go get -u github.com/moov-io/fincen

$ go doc github.com/moov-io/fincen BatchHeader
```

The package [`github.com/moov-io/fincen`](https://pkg.go.dev/github.com/moov-io/fincen) offers a Go-based reader, writer, and validator for Fincen BSA forms.

<ul>
<li><a href="https://moov-io.github.io/fincen/cash-payments/">Cash Payments over $10k</a></li>
<li><a href="https://moov-io.github.io/fincen/ctr/">Currency Transaction Report</a></li>
<li><a href="https://moov-io.github.io/fincen/doep/">Designation of Exempt Person</a></li>
<li><a href="https://moov-io.github.io/fincen/fbar/">Foreign Bank and Financial Accounts</a></li>
<li><a href="https://moov-io.github.io/fincen/sar/">Suspicious Activity Report</a></li>
</ul>
