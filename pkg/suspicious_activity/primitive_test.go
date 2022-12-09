// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package suspicious_activity

import (
	"testing"

	"github.com/moov-io/fincen"
	"github.com/stretchr/testify/require"
)

// validating elements
var (
	_ fincen.ElementActivity = (*ActivityType)(nil)
	_ fincen.Element         = (*AccountType)(nil)
	_ fincen.Element         = (*AccountType)(nil)
	_ fincen.Element         = (*ActivityAssociationType)(nil)
	_ fincen.Element         = (*ActivityIPAddressType)(nil)
	_ fincen.Element         = (*ActivityNarrativeInformationType)(nil)
	_ fincen.Element         = (*ActivitySupportDocumentType)(nil)
	_ fincen.Element         = (*ActivityType)(nil)
	_ fincen.Element         = (*AddressType)(nil)
	_ fincen.Element         = (*AssociationParty)(nil)
	_ fincen.Element         = (*AccountAssociationParty)(nil)
	_ fincen.Element         = (*AssetsAttributeType)(nil)
	_ fincen.Element         = (*AssetsType)(nil)
	_ fincen.Element         = (*CyberEventIndicatorsType)(nil)
	_ fincen.Element         = (*ElectronicAddressType)(nil)
	_ fincen.Element         = (*OrganizationClassificationTypeSubtypeType)(nil)
	_ fincen.Element         = (*PartyAccountAssociationType)(nil)
	_ fincen.Element         = (*PartyAssociationType)(nil)
	_ fincen.Element         = (*PartyIdentificationType)(nil)
	_ fincen.Element         = (*PartyNameType)(nil)
	_ fincen.Element         = (*PartyOccupationBusinessType)(nil)
	_ fincen.Element         = (*PartyType)(nil)
	_ fincen.Element         = (*PhoneNumberType)(nil)
	_ fincen.Element         = (*SuspiciousActivityClassificationType)(nil)
	_ fincen.Element         = (*SuspiciousActivityType)(nil)
)

func TestPrimitives(t *testing.T) {

	t.Run("ValidateActivityNarrativeSequenceNumber", func(t *testing.T) {
		var sample ValidateActivityNarrativeSequenceNumber
		require.Equal(t, "The ActivityNarrativeSequenceNumber has invalid value", sample.Validate().Error())

		sample = 5
		require.NoError(t, sample.Validate())
	})

	t.Run("ValidateActivityPartyCodeType", func(t *testing.T) {
		var sample ValidateActivityPartyCodeType
		require.Equal(t, "The ActivityPartyCode has invalid value", sample.Validate().Error())

		sample = PartyTransmitter
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

		sample = fincen.IndicateLegalName
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
	})

	t.Run("ValidatePartyAccountAssociationCodeType", func(t *testing.T) {
		var sample ValidatePartyAccountAssociationCodeType
		require.Equal(t, "The PartyAccountAssociationCode has invalid value", sample.Validate().Error())

		sample = "7"
		require.NoError(t, sample.Validate())
	})

}
