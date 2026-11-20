// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReformatWritesBannerToStderr(t *testing.T) {
	sample := filepath.Join("..", "..", "data", "samples", "ctr_batch.xml")
	out, err := Reformat([]string{sample}, true, "xml")
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(strings.TrimSpace(string(out)), "<"))
	require.NotContains(t, string(out), "Formatting report file")
	require.NotContains(t, string(out), "Formating report file")
}

func TestReformatCLIRedirectIsXML(t *testing.T) {
	// Build and run the command so we cover the process stdout/stderr split.
	bin := filepath.Join(t.TempDir(), "fincen")
	build := exec.Command("go", "build", "-o", bin, ".")
	build.Dir = "."
	out, err := build.CombinedOutput()
	require.NoError(t, err, string(out))

	sample, err := filepath.Abs(filepath.Join("..", "..", "data", "samples", "ctr_batch.xml"))
	require.NoError(t, err)

	cmd := exec.Command(bin, "reformat", "-generate-attrs", sample)
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	require.NoError(t, cmd.Run())
	require.Contains(t, stderr.String(), "Formatting report file")
	require.True(t, strings.HasPrefix(strings.TrimSpace(stdout.String()), "<"))
	require.NotContains(t, stdout.String(), "Formatting report file")
}
