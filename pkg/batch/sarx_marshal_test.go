// Copyright 2022 The Moov Authors
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

func TestSARXMarshalEmitsZeroAttachmentCounts(t *testing.T) {
	r := NewReport(fincen.Report111)
	r.ActivityCount = 1
	r.PartyCount = 1
	r.TotalAmount = 0
	r.ActivityAttachmentCount = 0
	r.AttachmentCount = 0

	out, err := xml.Marshal(r)
	require.NoError(t, err)
	body := string(out)
	require.Contains(t, body, `ActivityAttachmentCount="0"`)
	require.Contains(t, body, `AttachmentCount="0"`)
}

func TestSARXReformatSampleEmitsAttachmentCounts(t *testing.T) {
	buf, err := os.ReadFile(path.Join("..", "..", "data", "samples", "sar_batch.xml"))
	require.NoError(t, err)

	r, err := CreateReportWithBuffer(buf)
	require.NoError(t, err)
	require.Equal(t, fincen.Report111, r.FormTypeCode)
	require.NoError(t, r.GenerateAttrs())

	out, err := xml.Marshal(r)
	require.NoError(t, err)
	body := string(out)
	require.Contains(t, body, "ActivityAttachmentCount=")
	require.Contains(t, body, "AttachmentCount=")
	require.NotContains(t, body, `ActivityAttachmentCount=""`)
	// The sample has no attachments, so both values are zero and must still appear.
	require.True(t, strings.Contains(body, `ActivityAttachmentCount="0"`) || r.ActivityAttachmentCount > 0)
	require.True(t, strings.Contains(body, `AttachmentCount="0"`) || r.AttachmentCount > 0)
}
