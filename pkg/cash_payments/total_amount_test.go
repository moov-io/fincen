package cash_payments

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotalAmount_NilCurrencyTransactionActivity(t *testing.T) {
	act := ActivityType{}
	require.NotPanics(t, func() {
		require.Equal(t, 0.0, act.TotalAmount())
	})
}
