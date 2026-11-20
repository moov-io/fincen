// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"

	"github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	buf, err := os.ReadFile(path.Join("data", "samples", "doep_batch.xml"))
	if err != nil {
		return err
	}

	r := batch.NewReport(fincen.Report110)
	err = xml.Unmarshal(buf, r)
	if err != nil {
		return err
	}

	buf, err = xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(buf))
	return nil
}
