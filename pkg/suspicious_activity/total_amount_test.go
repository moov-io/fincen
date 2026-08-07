package suspicious_activity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotalAmount_NilSuspiciousActivity(t *testing.T) {
	act := ActivityType{}
	require.NotPanics(t, func() {
		require.Equal(t, 0.0, act.TotalAmount())
	})
}
