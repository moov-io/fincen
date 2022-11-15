// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package financial_accounts

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
	_ fincen.Element         = (*ActivityNarrativeInformationType)(nil)
	_ fincen.Element         = (*ActivityType)(nil)
	_ fincen.Element         = (*AddressType)(nil)
	_ fincen.Element         = (*ForeignAccountActivityType)(nil)
	_ fincen.Element         = (*PartyIdentificationType)(nil)
	_ fincen.Element         = (*PartyNameType)(nil)
	_ fincen.Element         = (*PartyType)(nil)
	_ fincen.Element         = (*PhoneNumberType)(nil)
)

func TestPrimitives(t *testing.T) {

	t.Run("ValidateActivityPartyCodeType", func(t *testing.T) {
		var sample ValidateActivityPartyCodeType
		require.Equal(t, "The ActivityPartyCode has invalid value", sample.Validate().Error())

		sample = PartyTransmitter
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidateAccountPartyCodeType", func(t *testing.T) {
		var sample ValidateAccountPartyCodeType
		require.Equal(t, "The AccountPartyCode has invalid value", sample.Validate().Error())

		sample = "41"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidateActivityPartyIdentificationCodeType", func(t *testing.T) {
		var sample ValidateActivityPartyIdentificationCodeType
		require.Equal(t, "The ActivityPartyIdentificationCode has invalid value", sample.Validate().Error())

		sample = "1"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidateAccountPartyIdentificationCodeType", func(t *testing.T) {
		var sample ValidateAccountPartyIdentificationCodeType
		require.Equal(t, "The AccountPartyIdentificationCode has invalid value", sample.Validate().Error())

		sample = "1"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidatePartyNameCodeType", func(t *testing.T) {
		var sample ValidatePartyNameCodeType
		require.Equal(t, "The PartyNameCode has invalid value", sample.Validate().Error())

		sample = fincen.IndicateLegalName
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidateIndicatorYNType", func(t *testing.T) {
		var sample ValidateIndicatorYNType
		require.Equal(t, "The IndicatorYNType has invalid value", sample.Validate().Error())

		sample = "Y"
		require.NoError(t, sample.Validate())

		sample = "N"
		require.NoError(t, sample.Validate())
	})

}
