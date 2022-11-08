// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package exempt_designation

import (
	"testing"

	"github.com/moov-io/fincen"
	"github.com/stretchr/testify/require"
)

// validating elements
var (
	_ fincen.ElementActivity = (*ActivityType)(nil)
	_ fincen.Element         = (*ActivityAssociationType)(nil)
	_ fincen.Element         = (*ActivityType)(nil)
	_ fincen.Element         = (*AddressType)(nil)
	_ fincen.Element         = (*DesignationExemptActivityType)(nil)
	_ fincen.Element         = (*ElectronicAddressType)(nil)
	_ fincen.Element         = (*PartyIdentificationType)(nil)
	_ fincen.Element         = (*PartyNameType)(nil)
	_ fincen.Element         = (*PartyOccupationBusinessType)(nil)
	_ fincen.Element         = (*PartyType)(nil)
	_ fincen.Element         = (*PhoneNumberType)(nil)
)

func TestPrimitives(t *testing.T) {

	t.Run("ValidateActivityPartyCodeType", func(t *testing.T) {
		var sample ValidateActivityPartyCodeType
		require.Equal(t, "The ActivityPartyCode has invalid value", sample.Validate().Error())

		sample = "35"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidatePartyIdentificationCodeType", func(t *testing.T) {
		var sample ValidatePartyIdentificationCodeType
		require.Equal(t, "The PartyIdentificationCode has invalid value", sample.Validate().Error())

		sample = "1"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidatePartyNameCodeType", func(t *testing.T) {
		var sample ValidatePartyNameCodeType
		require.Equal(t, "The PartyNameCode has invalid value", sample.Validate().Error())

		sample = "L"
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidateFederalRegulatorCodeType", func(t *testing.T) {
		var sample ValidateFederalRegulatorCodeType
		require.Equal(t, "The FederalRegulatorCode has invalid value", sample.Validate().Error())

		sample = "1"
		require.NoError(t, sample.Validate())
	})

	t.Run("EFilingPriorDocumentNumberType", func(t *testing.T) {
		var sample EFilingPriorDocumentNumberType
		require.Equal(t, "The EFilingPriorDocumentNumber has invalid value", sample.Validate().Error())

		sample = "00000000000000"
		require.NoError(t, sample.Validate())
		require.Equal(t, "00000000000000", sample.String())
	})

}
