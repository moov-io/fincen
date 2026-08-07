package currency_transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotalAmount_NilCurrencyTransactionActivity(t *testing.T) {
	// Empty Activity unmarshaled from XML leaves CurrencyTransactionActivity nil.
	// Validate/generateAttrs must not panic when summing amounts.
	act := ActivityType{}
	require.NotPanics(t, func() {
		require.Equal(t, 0.0, act.TotalAmount())
	})
}
