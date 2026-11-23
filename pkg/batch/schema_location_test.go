// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package batch

import (
	"encoding/xml"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/moov-io/fincen"
	"github.com/stretchr/testify/require"
)

var expectedSchemaLocations = map[string]string{
	fincen.Form8300:  "www.fincen.gov/base https://www.fincen.gov/system/files/schema/base/EFL_8300XBatchSchema.xsd",
	fincen.Report110: "www.fincen.gov/base https://www.fincen.gov/system/files/schema/base/EFL_DOEPXBatchSchema.xsd",
	fincen.Report111: "www.fincen.gov/base https://www.fincen.gov/system/files/schema/base/EFL_SARXBatchSchema.xsd",
	fincen.Report112: "www.fincen.gov/base https://www.fincen.gov/system/files/schema/base/EFL_CTRXBatchSchema.xsd",
	fincen.Report114: "www.fincen.gov/base https://www.fincen.gov/system/files/schema/base/EFL_FBARXBatchSchema.xsd",
}

func TestSchemaLocation(t *testing.T) {
	for form, want := range expectedSchemaLocations {
		require.Equal(t, want, SchemaLocation(form), form)
	}
	require.Empty(t, SchemaLocation("UNKNOWN"))
}

func TestMarshalEmitsSchemaLocation(t *testing.T) {
	for form, want := range expectedSchemaLocations {
		t.Run(form, func(t *testing.T) {
			out, err := xml.Marshal(NewReport(form))
			require.NoError(t, err)

			body := string(out)
			require.Equal(t, 1, strings.Count(body, "xsi:schemaLocation="))
			require.Contains(t, body, `xsi:schemaLocation="`+want+`"`)
			require.NotContains(t, body, "www.server.gov")
			require.NotContains(t, body, "/sites/default/files/")
			require.NotContains(t, body, "EFL_FinCENSARXBatchSchema.xsd")
		})
	}
}

func TestReformatSamplesEmitSchemaLocation(t *testing.T) {
	samples := map[string]string{
		"8300_batch.xml": fincen.Form8300,
		"doep_batch.xml": fincen.Report110,
		"sar_batch.xml":  fincen.Report111,
		"ctr_batch.xml":  fincen.Report112,
		"fbar_batch.xml": fincen.Report114,
	}
	for name, form := range samples {
		t.Run(name, func(t *testing.T) {
			buf, err := os.ReadFile(path.Join("..", "..", "data", "samples", name))
			require.NoError(t, err)

			r, err := CreateReportWithBuffer(buf)
			require.NoError(t, err)
			require.Equal(t, form, r.FormTypeCode)

			out, err := xml.Marshal(r)
			require.NoError(t, err)

			body := string(out)
			require.Equal(t, 1, strings.Count(body, "xsi:schemaLocation="))
			require.Contains(t, body, `xsi:schemaLocation="`+expectedSchemaLocations[form]+`"`)
		})
	}
}
