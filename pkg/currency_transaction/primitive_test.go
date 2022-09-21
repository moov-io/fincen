// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package currency_transaction

import (
	"github.com/moov-io/fincen"
	"github.com/stretchr/testify/require"
	"testing"
)

// validating elements
var (
	_ fincen.ElementActivity = (*ActivityType)(nil)
	_ fincen.Element         = (*AccountType)(nil)
	_ fincen.Element         = (*ActivityAssociationType)(nil)
	_ fincen.Element         = (*ActivityType)(nil)
	_ fincen.Element         = (*AddressType)(nil)
	_ fincen.Element         = (*CurrencyTransactionActivityDetailType)(nil)
	_ fincen.Element         = (*CurrencyTransactionActivityType)(nil)
	_ fincen.Element         = (*ElectronicAddressType)(nil)
	_ fincen.Element         = (*OrganizationClassificationTypeSubtypeType)(nil)
	_ fincen.Element         = (*PartyAccountAssociationType)(nil)
	_ fincen.Element         = (*PartyIdentificationType)(nil)
	_ fincen.Element         = (*PartyNameType)(nil)
	_ fincen.Element         = (*PartyOccupationBusinessType)(nil)
	_ fincen.Element         = (*PartyType)(nil)
	_ fincen.Element         = (*PhoneNumberType)(nil)
)

func TestPrimitives(t *testing.T) {

	t.Run("ValidateActivityPartyCodeType", func(t *testing.T) {
		var sample ValidateActivityPartyCodeType
		require.Equal(t, "The ValidateActivityPartyCode has invalid value", sample.Validate().Error())

		sample = "35"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidatePartyIdentificationCodeType", func(t *testing.T) {
		var sample ValidatePartyIdentificationCodeType
		require.Equal(t, "The ValidatePartyIdentificationCode has invalid value", sample.Validate().Error())

		sample = "1"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidatePartyNameCodeType", func(t *testing.T) {
		var sample ValidatePartyNameCodeType
		require.Equal(t, "The ValidatePartyNameCode has invalid value", sample.Validate().Error())

		sample = "L"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidateFederalRegulatorCodeType", func(t *testing.T) {
		var sample ValidateFederalRegulatorCodeType
		require.Equal(t, "The ValidateFederalRegulatorCode has invalid value", sample.Validate().Error())

		sample = "1"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidatePartyAccountAssociationCodeType", func(t *testing.T) {
		var sample ValidatePartyAccountAssociationCodeType
		require.Equal(t, "The ValidatePartyAccountAssociationCode has invalid value", sample.Validate().Error())

		sample = "8"
		require.NoError(t, sample.Validate())
	})

}
