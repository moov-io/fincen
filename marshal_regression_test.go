package fincen_test

import (
	"testing"

	"github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
	"github.com/stretchr/testify/require"
)

func TestMarshal_EmptyBatchNoPanic(t *testing.T) {
	report, err := batch.CreateReportWithBuffer([]byte("<EFilingBatchXML></EFilingBatchXML>"))
	require.NoError(t, err)
	require.NotPanics(t, func() {
		_, _ = fincen.Marshal(report)
	})
}
