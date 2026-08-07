package batch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Regression for CI fuzz crash FuzzCreateReport/474230c329ae08c9:
// empty <Activity/> with FormTypeCode CTRX panicked in TotalAmount via Validate.
func TestValidate_EmptyCTRActivityNoPanic(t *testing.T) {
	input := []byte(`<EFilingBatchXML><Activity></Activity><FormTypeCode>CTRX</FormTypeCode></EFilingBatchXML>`)
	report, err := CreateReportWithBuffer(input)
	require.NoError(t, err)
	require.NotNil(t, report)
	require.NotPanics(t, func() {
		_ = report.Validate()
		_ = report.GenerateAttrs()
		_ = report.GenerateSeqNumbers()
	})
}
