// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzz

import (
	"encoding/xml"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
)

func FuzzCreateReport(f *testing.F) {
	populateCorpus(f)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		report, err := batch.CreateReportWithBuffer([]byte(contents))
		if err != nil || report == nil {
			return
		}

		_ = report.Validate()
		_ = report.GenerateAttrs()
		_ = report.GenerateSeqNumbers()

		// Marshal back out — must not panic
		_, _ = xml.Marshal(report)
		_, _ = fincen.Marshal(report)
	})
}

func FuzzReportTypes(f *testing.F) {
	populateCorpus(f)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 256*1024 {
			t.Skip()
		}

		// Unmarshal into a generic report; type-specific validation is covered by samples.
		r := batch.NewReport(fincen.Report111)
		if err := xml.Unmarshal([]byte(contents), r); err != nil {
			return
		}
		_ = r.Validate()
		_, _ = xml.Marshal(r)
		_, _ = fincen.Marshal(r)
	})
}

func populateCorpus(f *testing.F) {
	f.Helper()

	f.Add("")
	f.Add("<EFilingBatchXML></EFilingBatchXML>")
	f.Add("{}")

	roots := []string{
		filepath.Join("..", "..", "data", "samples"),
		filepath.Join("..", "..", "examples"),
	}
	for _, root := range roots {
		_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil || info == nil || info.IsDir() {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".xml" || ext == ".json" {
				bs, err := os.ReadFile(path)
				if err != nil || len(bs) > 512*1024 {
					return nil
				}
				f.Add(string(bs))
			}
			return nil
		})
	}
}
