// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package currency_transaction

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/fincen"
)

func mocParties() map[string][]byte {
	parties := make(map[string][]byte)

	parties["35"] = []byte(`<Party SeqNum="3">
	<ActivityPartyTypeCode>35</ActivityPartyTypeCode>
	<EFilingCoverageBeginningDateText>20170101</EFilingCoverageBeginningDateText>
	<EFilingCoverageEndDateText>20170131</EFilingCoverageEndDateText>
	<PartyName SeqNum="4">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Transmitter Legal Name</RawPartyFullName>
	</PartyName>
	<Address SeqNum="5">
		<RawCityText>Vienna</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<RawStreetAddress1Text>123 Address Road</RawStreetAddress1Text>
		<RawZIPCode>22102</RawZIPCode>
	</Address>
	<PhoneNumber SeqNum="6">
		<PhoneNumberText>7039991234</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="7">
		<PartyIdentificationNumberText>547898569</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>4</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="8">
		<PartyIdentificationNumberText>PTCC1234</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>28</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties["37"] = []byte(`<Party SeqNum="9">
	<ActivityPartyTypeCode>37</ActivityPartyTypeCode>
	<EFilingCoverageBeginningDateText>20170101</EFilingCoverageBeginningDateText>
	<EFilingCoverageEndDateText>20170131</EFilingCoverageEndDateText>
	<PartyName SeqNum="10">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Transmitter Contact Name</RawPartyFullName>
	</PartyName>
</Party>`)

	parties["30"] = []byte(`<Party SeqNum="11">
	<ActivityPartyTypeCode>30</ActivityPartyTypeCode>
	<PrimaryRegulatorTypeCode>9</PrimaryRegulatorTypeCode>
	<PartyName SeqNum="12">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Filing Institution Legal Name</RawPartyFullName>
	</PartyName>
	<PartyName SeqNum="13">
		<PartyNameTypeCode>DBA</PartyNameTypeCode>
		<RawPartyFullName>Filing Institution Alternate DBA Name</RawPartyFullName>
	</PartyName>
	<Address SeqNum="14">
		<RawCityText>Vienna</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<RawStreetAddress1Text>456 Address Way</RawStreetAddress1Text>
		<RawZIPCode>554789985</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="15">
		<PartyIdentificationNumberText>554785215</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="16">
		<PartyIdentificationNumberText>15478564</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>10</PartyIdentificationTypeCode>
	</PartyIdentification>
	<OrganizationClassificationTypeSubtype SeqNum="17">
		<OrganizationTypeID>2</OrganizationTypeID>
	</OrganizationClassificationTypeSubtype>
</Party>`)

	parties["8"] = []byte(`<Party SeqNum="18">
	<ActivityPartyTypeCode>8</ActivityPartyTypeCode>
	<PartyName SeqNum="19">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Contact Office Legal Name</RawPartyFullName>
	</PartyName>
	<PhoneNumber SeqNum="20">
		<PhoneNumberExtensionText>2210</PhoneNumberExtensionText>
		<PhoneNumberText>7039874589</PhoneNumberText>
	</PhoneNumber>
</Party>`)

	parties["34"] = []byte(`<Party SeqNum="21">
	<ActivityPartyTypeCode>34</ActivityPartyTypeCode>
	<IndividualEntityCashInAmountText>15000</IndividualEntityCashInAmountText>
	<IndividualEntityCashOutAmountText></IndividualEntityCashOutAmountText>
	<PrimaryRegulatorTypeCode>9</PrimaryRegulatorTypeCode>
	<PartyName SeqNum="22">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Transaction Location Legal Name</RawPartyFullName>
	</PartyName>
	<PartyName SeqNum="23">
		<PartyNameTypeCode>DBA</PartyNameTypeCode>
		<RawPartyFullName>Location Alternate DBA Name</RawPartyFullName>
	</PartyName>
	<Address SeqNum="24">
		<RawCityText>Vienna</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<RawStreetAddress1Text>789 Address Court</RawStreetAddress1Text>
		<RawZIPCode>55478</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="25">
		<PartyIdentificationNumberText>145878965</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="26">
		<PartyIdentificationNumberText>45899856</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>11</PartyIdentificationTypeCode>
	</PartyIdentification>
	<OrganizationClassificationTypeSubtype SeqNum="27">
		<OrganizationSubtypeID>1999</OrganizationSubtypeID>
		<OrganizationTypeID>1</OrganizationTypeID>
		<OtherOrganizationSubTypeText>Other casino</OtherOrganizationSubTypeText>
	</OrganizationClassificationTypeSubtype>
</Party>`)

	parties["50"] = []byte(`<Party SeqNum="28">
	<ActivityPartyTypeCode>50</ActivityPartyTypeCode>
	<FemaleGenderIndicator>Y</FemaleGenderIndicator>
	<IndividualBirthDateText>19750120</IndividualBirthDateText>
	<IndividualEntityCashInAmountText>20000</IndividualEntityCashInAmountText>
	<MultipleTransactionsPersonsIndividualsIndicator>Y</MultipleTransactionsPersonsIndividualsIndicator>
	<PartyName SeqNum="29">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Doe</RawEntityIndividualLastName>
		<RawIndividualFirstName>John</RawIndividualFirstName>
		<RawIndividualMiddleName>Johnson</RawIndividualMiddleName>
		<RawIndividualNameSuffixText>Jr.</RawIndividualNameSuffixText>
	</PartyName>
	<PartyName SeqNum="30">
		<PartyNameTypeCode>AKA</PartyNameTypeCode>
		<RawEntityIndividualLastName>JJ</RawEntityIndividualLastName>
	</PartyName>
	<Address SeqNum="31">
		<RawCityText>Vienna</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<StreetAddressUnknownIndicator>Y</StreetAddressUnknownIndicator>
		<ZIPCodeUnknownIndicator>Y</ZIPCodeUnknownIndicator>
	</Address>
	<PhoneNumber SeqNum="32">
		<PhoneNumberText>1458987456</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="33">
		<TINUnknownIndicator>Y</TINUnknownIndicator>
	</PartyIdentification>
	<PartyIdentification SeqNum="34">
		<OtherIssuerCountryText>US</OtherIssuerCountryText>
		<OtherIssuerStateText>TX</OtherIssuerStateText>
		<PartyIdentificationNumberText>44589774512</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>5</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="35">
		<NAICSCode>6214</NAICSCode>
		<OccupationBusinessText>Outpatient Care Centers</OccupationBusinessText>
	</PartyOccupationBusiness>
	<ElectronicAddress SeqNum="36">
		<ElectronicAddressText>email123@server.gov</ElectronicAddressText>
	</ElectronicAddress>
	<Account SeqNum="37">
		<AccountNumberText>1115478569</AccountNumberText>
		<PartyAccountAssociation SeqNum="38">
			<PartyAccountAssociationTypeCode>8</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="39">
		<AccountNumberText>3365998541</AccountNumberText>
		<PartyAccountAssociation SeqNum="40">
			<PartyAccountAssociationTypeCode>8</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="41">
		<AccountNumberText>4857985691</AccountNumberText>
		<PartyAccountAssociation SeqNum="42">
			<PartyAccountAssociationTypeCode>9</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
</Party>`)

	parties["17"] = []byte(`<Party SeqNum="28">
	<ActivityPartyTypeCode>17</ActivityPartyTypeCode>
	<FemaleGenderIndicator>Y</FemaleGenderIndicator>
	<IndividualBirthDateText>19750120</IndividualBirthDateText>
	<IndividualEntityCashInAmountText>20000</IndividualEntityCashInAmountText>
	<MultipleTransactionsPersonsIndividualsIndicator>Y</MultipleTransactionsPersonsIndividualsIndicator>
	<PartyName SeqNum="29">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Doe</RawEntityIndividualLastName>
		<RawIndividualFirstName>John</RawIndividualFirstName>
		<RawIndividualMiddleName>Johnson</RawIndividualMiddleName>
		<RawIndividualNameSuffixText>Jr.</RawIndividualNameSuffixText>
	</PartyName>
	<PartyName SeqNum="30">
		<PartyNameTypeCode>AKA</PartyNameTypeCode>
		<RawEntityIndividualLastName>JJ</RawEntityIndividualLastName>
	</PartyName>
	<Address SeqNum="31">
		<RawCityText>Vienna</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<StreetAddressUnknownIndicator>Y</StreetAddressUnknownIndicator>
		<ZIPCodeUnknownIndicator>Y</ZIPCodeUnknownIndicator>
	</Address>
	<PhoneNumber SeqNum="32">
		<PhoneNumberText>1458987456</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="33">
		<TINUnknownIndicator>Y</TINUnknownIndicator>
	</PartyIdentification>
	<PartyIdentification SeqNum="34">
		<OtherIssuerCountryText>US</OtherIssuerCountryText>
		<OtherIssuerStateText>TX</OtherIssuerStateText>
		<PartyIdentificationNumberText>44589774512</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>5</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="35">
		<NAICSCode>6214</NAICSCode>
		<OccupationBusinessText>Outpatient Care Centers</OccupationBusinessText>
	</PartyOccupationBusiness>
	<ElectronicAddress SeqNum="36">
		<ElectronicAddressText>email123@server.gov</ElectronicAddressText>
	</ElectronicAddress>
	<Account SeqNum="37">
		<AccountNumberText>1115478569</AccountNumberText>
		<PartyAccountAssociation SeqNum="38">
			<PartyAccountAssociationTypeCode>8</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="39">
		<AccountNumberText>3365998541</AccountNumberText>
		<PartyAccountAssociation SeqNum="40">
			<PartyAccountAssociationTypeCode>8</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="41">
		<AccountNumberText>4857985691</AccountNumberText>
		<PartyAccountAssociation SeqNum="42">
			<PartyAccountAssociationTypeCode>9</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
</Party>`)

	parties["23"] = []byte(`<Party SeqNum="28">
	<ActivityPartyTypeCode>23</ActivityPartyTypeCode>
	<FemaleGenderIndicator>Y</FemaleGenderIndicator>
	<IndividualBirthDateText>19750120</IndividualBirthDateText>
	<IndividualEntityCashInAmountText>20000</IndividualEntityCashInAmountText>
	<MultipleTransactionsPersonsIndividualsIndicator>Y</MultipleTransactionsPersonsIndividualsIndicator>
	<PartyName SeqNum="29">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Doe</RawEntityIndividualLastName>
		<RawIndividualFirstName>John</RawIndividualFirstName>
		<RawIndividualMiddleName>Johnson</RawIndividualMiddleName>
		<RawIndividualNameSuffixText>Jr.</RawIndividualNameSuffixText>
	</PartyName>
	<PartyName SeqNum="30">
		<PartyNameTypeCode>AKA</PartyNameTypeCode>
		<RawEntityIndividualLastName>JJ</RawEntityIndividualLastName>
	</PartyName>
	<Address SeqNum="31">
		<RawCityText>Vienna</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<StreetAddressUnknownIndicator>Y</StreetAddressUnknownIndicator>
		<ZIPCodeUnknownIndicator>Y</ZIPCodeUnknownIndicator>
	</Address>
	<PhoneNumber SeqNum="32">
		<PhoneNumberText>1458987456</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="33">
		<TINUnknownIndicator>Y</TINUnknownIndicator>
	</PartyIdentification>
	<PartyIdentification SeqNum="34">
		<OtherIssuerCountryText>US</OtherIssuerCountryText>
		<OtherIssuerStateText>TX</OtherIssuerStateText>
		<PartyIdentificationNumberText>44589774512</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>5</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="35">
		<NAICSCode>6214</NAICSCode>
		<OccupationBusinessText>Outpatient Care Centers</OccupationBusinessText>
	</PartyOccupationBusiness>
	<ElectronicAddress SeqNum="36">
		<ElectronicAddressText>email123@server.gov</ElectronicAddressText>
	</ElectronicAddress>
	<Account SeqNum="37">
		<AccountNumberText>1115478569</AccountNumberText>
		<PartyAccountAssociation SeqNum="38">
			<PartyAccountAssociationTypeCode>8</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="39">
		<AccountNumberText>3365998541</AccountNumberText>
		<PartyAccountAssociation SeqNum="40">
			<PartyAccountAssociationTypeCode>8</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="41">
		<AccountNumberText>4857985691</AccountNumberText>
		<PartyAccountAssociation SeqNum="42">
			<PartyAccountAssociationTypeCode>9</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
</Party>`)

	parties["58"] = []byte(`<Party SeqNum="28">
	<ActivityPartyTypeCode>58</ActivityPartyTypeCode>
	<FemaleGenderIndicator>Y</FemaleGenderIndicator>
	<IndividualBirthDateText>19750120</IndividualBirthDateText>
	<IndividualEntityCashInAmountText>20000</IndividualEntityCashInAmountText>
	<MultipleTransactionsPersonsIndividualsIndicator>Y</MultipleTransactionsPersonsIndividualsIndicator>
	<PartyName SeqNum="29">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Doe</RawEntityIndividualLastName>
		<RawIndividualFirstName>John</RawIndividualFirstName>
		<RawIndividualMiddleName>Johnson</RawIndividualMiddleName>
		<RawIndividualNameSuffixText>Jr.</RawIndividualNameSuffixText>
	</PartyName>
	<PartyName SeqNum="30">
		<PartyNameTypeCode>AKA</PartyNameTypeCode>
		<RawEntityIndividualLastName>JJ</RawEntityIndividualLastName>
	</PartyName>
	<Address SeqNum="31">
		<RawCityText>Vienna</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<StreetAddressUnknownIndicator>Y</StreetAddressUnknownIndicator>
		<ZIPCodeUnknownIndicator>Y</ZIPCodeUnknownIndicator>
	</Address>
	<PhoneNumber SeqNum="32">
		<PhoneNumberText>1458987456</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="33">
		<TINUnknownIndicator>Y</TINUnknownIndicator>
	</PartyIdentification>
	<PartyIdentification SeqNum="34">
		<OtherIssuerCountryText>US</OtherIssuerCountryText>
		<OtherIssuerStateText>TX</OtherIssuerStateText>
		<PartyIdentificationNumberText>44589774512</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>5</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="35">
		<NAICSCode>6214</NAICSCode>
		<OccupationBusinessText>Outpatient Care Centers</OccupationBusinessText>
	</PartyOccupationBusiness>
	<ElectronicAddress SeqNum="36">
		<ElectronicAddressText>email123@server.gov</ElectronicAddressText>
	</ElectronicAddress>
	<Account SeqNum="37">
		<AccountNumberText>1115478569</AccountNumberText>
		<PartyAccountAssociation SeqNum="38">
			<PartyAccountAssociationTypeCode>8</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="39">
		<AccountNumberText>3365998541</AccountNumberText>
		<PartyAccountAssociation SeqNum="40">
			<PartyAccountAssociationTypeCode>8</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="41">
		<AccountNumberText>4857985691</AccountNumberText>
		<PartyAccountAssociation SeqNum="42">
			<PartyAccountAssociationTypeCode>9</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
</Party>`)

	return parties
}

func TestParty(t *testing.T) {
	t.Run("Transmitter Party", func(t *testing.T) {
		sample := mocParties()["35"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(3))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("35"))
		require.Equal(t, *party.EFilingCoverageBeginningDateText, fincen.DateYYYYMMDDType("20170101"))
		require.Equal(t, *party.EFilingCoverageEndDateText, fincen.DateYYYYMMDDType("20170131"))
		require.Equal(t, len(party.PartyName), 1)
		require.Equal(t, len(party.PartyIdentification), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(4))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter Legal Name"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(5))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Vienna"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("123 Address Road"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("22102"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(6))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("7039991234"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(7))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("547898569"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("4"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(8))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("PTCC1234"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("28"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Transmitter Contact Party", func(t *testing.T) {
		sample := mocParties()["37"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(9))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("37"))
		require.Equal(t, *party.EFilingCoverageBeginningDateText, fincen.DateYYYYMMDDType("20170101"))
		require.Equal(t, *party.EFilingCoverageEndDateText, fincen.DateYYYYMMDDType("20170131"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(10))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter Contact Name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Reporting Financial Institution Party", func(t *testing.T) {
		sample := mocParties()["30"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(11))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("30"))
		require.Equal(t, len(party.PartyName), 2)
		require.Equal(t, len(party.PartyIdentification), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(12))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Filing Institution Legal Name"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(13))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("DBA"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Filing Institution Alternate DBA Name"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(14))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Vienna"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("456 Address Way"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("554789985"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(15))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("554785215"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(16))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("15478564"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("10"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Contact for Assistance Party", func(t *testing.T) {
		sample := mocParties()["8"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(18))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("8"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(19))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Contact Office Legal Name"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(20))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("7039874589"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Transaction Location Party", func(t *testing.T) {
		sample := mocParties()["34"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(21))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("34"))
		require.Equal(t, *party.IndividualEntityCashInAmountText, fincen.RestrictString15("15000"))
		require.Equal(t, *party.PrimaryRegulatorTypeCode, ValidateFederalRegulatorCodeType("9"))
		require.Equal(t, len(party.PartyName), 2)
		require.Equal(t, len(party.PartyIdentification), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(22))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transaction Location Legal Name"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(23))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("DBA"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Location Alternate DBA Name"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(24))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Vienna"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("789 Address Court"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("55478"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(25))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("145878965"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(26))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("45899856"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("11"))

		org := party.OrganizationClassificationTypeSubtype
		require.Equal(t, org.SeqNum, fincen.SeqNumber(27))
		require.Equal(t, *org.OrganizationSubtypeID, fincen.ValidateOrganizationSubtypeCodeCtrType(1999))
		require.Equal(t, org.OrganizationTypeID, fincen.ValidateOrganizationCodeType(1))
		require.Equal(t, *org.OtherOrganizationSubTypeText, fincen.RestrictString50("Other casino"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Person involved in the transaction", func(t *testing.T) {
		sample := mocParties()["50"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(28))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("50"))
		require.Equal(t, *party.FemaleGenderIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.IndividualBirthDateText, fincen.DateYYYYMMDDOrBlankTypeDOB("19750120"))
		require.Equal(t, *party.IndividualEntityCashInAmountText, fincen.RestrictString15("20000"))
		require.Equal(t, *party.MultipleTransactionsPersonsIndividualsIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, len(party.PartyName), 2)
		require.Equal(t, len(party.PartyIdentification), 2)
		require.Equal(t, len(party.Account), 3)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(29))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Doe"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("John"))
		require.Equal(t, *name.RawIndividualMiddleName, fincen.RestrictString35("Johnson"))
		require.Equal(t, *name.RawIndividualNameSuffixText, fincen.RestrictString35("Jr."))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(30))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("AKA"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("JJ"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(31))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Vienna"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.StreetAddressUnknownIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *address.ZIPCodeUnknownIndicator, fincen.ValidateIndicatorType("Y"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(32))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("1458987456"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(33))
		require.Equal(t, *identification.TINUnknownIndicator, fincen.ValidateIndicatorType("Y"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(34))
		require.Equal(t, *identification.OtherIssuerCountryText, fincen.RestrictString2("US"))
		require.Equal(t, *identification.OtherIssuerStateText, fincen.RestrictString3("TX"))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("44589774512"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("5"))

		business := party.PartyOccupationBusiness
		require.Equal(t, business.SeqNum, fincen.SeqNumber(35))
		require.Equal(t, *business.NAICSCode, fincen.RestrictString6("6214"))
		require.Equal(t, *business.OccupationBusinessText, fincen.RestrictString50("Outpatient Care Centers"))

		eAddress := party.ElectronicAddress
		require.Equal(t, eAddress.SeqNum, fincen.SeqNumber(36))
		require.Equal(t, *eAddress.ElectronicAddressText, fincen.RestrictString517("email123@server.gov"))

		acc := party.Account[0]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(37))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("1115478569"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(38))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("8"))

		acc = party.Account[1]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(39))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("3365998541"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(40))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("8"))

		acc = party.Account[2]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(41))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("4857985691"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(42))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("9"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Person conducting transaction for another", func(t *testing.T) {
		sample := mocParties()["17"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(28))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("17"))
		require.Equal(t, *party.FemaleGenderIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.IndividualBirthDateText, fincen.DateYYYYMMDDOrBlankTypeDOB("19750120"))
		require.Equal(t, *party.IndividualEntityCashInAmountText, fincen.RestrictString15("20000"))
		require.Equal(t, *party.MultipleTransactionsPersonsIndividualsIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, len(party.PartyName), 2)
		require.Equal(t, len(party.PartyIdentification), 2)
		require.Equal(t, len(party.Account), 3)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(29))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Doe"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("John"))
		require.Equal(t, *name.RawIndividualMiddleName, fincen.RestrictString35("Johnson"))
		require.Equal(t, *name.RawIndividualNameSuffixText, fincen.RestrictString35("Jr."))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(30))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("AKA"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("JJ"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(31))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Vienna"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.StreetAddressUnknownIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *address.ZIPCodeUnknownIndicator, fincen.ValidateIndicatorType("Y"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(32))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("1458987456"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(33))
		require.Equal(t, *identification.TINUnknownIndicator, fincen.ValidateIndicatorType("Y"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(34))
		require.Equal(t, *identification.OtherIssuerCountryText, fincen.RestrictString2("US"))
		require.Equal(t, *identification.OtherIssuerStateText, fincen.RestrictString3("TX"))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("44589774512"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("5"))

		business := party.PartyOccupationBusiness
		require.Equal(t, business.SeqNum, fincen.SeqNumber(35))
		require.Equal(t, *business.NAICSCode, fincen.RestrictString6("6214"))
		require.Equal(t, *business.OccupationBusinessText, fincen.RestrictString50("Outpatient Care Centers"))

		eAddress := party.ElectronicAddress
		require.Equal(t, eAddress.SeqNum, fincen.SeqNumber(36))
		require.Equal(t, *eAddress.ElectronicAddressText, fincen.RestrictString517("email123@server.gov"))

		acc := party.Account[0]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(37))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("1115478569"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(38))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("8"))

		acc = party.Account[1]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(39))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("3365998541"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(40))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("8"))

		acc = party.Account[2]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(41))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("4857985691"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(42))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("9"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Person on whose behalf this transaction was conducted", func(t *testing.T) {
		sample := mocParties()["23"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(28))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("23"))
		require.Equal(t, *party.FemaleGenderIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.IndividualBirthDateText, fincen.DateYYYYMMDDOrBlankTypeDOB("19750120"))
		require.Equal(t, *party.IndividualEntityCashInAmountText, fincen.RestrictString15("20000"))
		require.Equal(t, *party.MultipleTransactionsPersonsIndividualsIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, len(party.PartyName), 2)
		require.Equal(t, len(party.PartyIdentification), 2)
		require.Equal(t, len(party.Account), 3)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(29))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Doe"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("John"))
		require.Equal(t, *name.RawIndividualMiddleName, fincen.RestrictString35("Johnson"))
		require.Equal(t, *name.RawIndividualNameSuffixText, fincen.RestrictString35("Jr."))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(30))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("AKA"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("JJ"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(31))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Vienna"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.StreetAddressUnknownIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *address.ZIPCodeUnknownIndicator, fincen.ValidateIndicatorType("Y"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(32))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("1458987456"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(33))
		require.Equal(t, *identification.TINUnknownIndicator, fincen.ValidateIndicatorType("Y"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(34))
		require.Equal(t, *identification.OtherIssuerCountryText, fincen.RestrictString2("US"))
		require.Equal(t, *identification.OtherIssuerStateText, fincen.RestrictString3("TX"))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("44589774512"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("5"))

		business := party.PartyOccupationBusiness
		require.Equal(t, business.SeqNum, fincen.SeqNumber(35))
		require.Equal(t, *business.NAICSCode, fincen.RestrictString6("6214"))
		require.Equal(t, *business.OccupationBusinessText, fincen.RestrictString50("Outpatient Care Centers"))

		eAddress := party.ElectronicAddress
		require.Equal(t, eAddress.SeqNum, fincen.SeqNumber(36))
		require.Equal(t, *eAddress.ElectronicAddressText, fincen.RestrictString517("email123@server.gov"))

		acc := party.Account[0]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(37))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("1115478569"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(38))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("8"))

		acc = party.Account[1]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(39))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("3365998541"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(40))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("8"))

		acc = party.Account[2]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(41))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("4857985691"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(42))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("9"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Common carrier", func(t *testing.T) {
		sample := mocParties()["58"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(28))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("58"))
		require.Equal(t, *party.FemaleGenderIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.IndividualBirthDateText, fincen.DateYYYYMMDDOrBlankTypeDOB("19750120"))
		require.Equal(t, *party.IndividualEntityCashInAmountText, fincen.RestrictString15("20000"))
		require.Equal(t, *party.MultipleTransactionsPersonsIndividualsIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, len(party.PartyName), 2)
		require.Equal(t, len(party.PartyIdentification), 2)
		require.Equal(t, len(party.Account), 3)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(29))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Doe"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("John"))
		require.Equal(t, *name.RawIndividualMiddleName, fincen.RestrictString35("Johnson"))
		require.Equal(t, *name.RawIndividualNameSuffixText, fincen.RestrictString35("Jr."))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(30))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("AKA"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("JJ"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(31))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Vienna"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.StreetAddressUnknownIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *address.ZIPCodeUnknownIndicator, fincen.ValidateIndicatorType("Y"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(32))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("1458987456"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(33))
		require.Equal(t, *identification.TINUnknownIndicator, fincen.ValidateIndicatorType("Y"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(34))
		require.Equal(t, *identification.OtherIssuerCountryText, fincen.RestrictString2("US"))
		require.Equal(t, *identification.OtherIssuerStateText, fincen.RestrictString3("TX"))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("44589774512"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("5"))

		business := party.PartyOccupationBusiness
		require.Equal(t, business.SeqNum, fincen.SeqNumber(35))
		require.Equal(t, *business.NAICSCode, fincen.RestrictString6("6214"))
		require.Equal(t, *business.OccupationBusinessText, fincen.RestrictString50("Outpatient Care Centers"))

		eAddress := party.ElectronicAddress
		require.Equal(t, eAddress.SeqNum, fincen.SeqNumber(36))
		require.Equal(t, *eAddress.ElectronicAddressText, fincen.RestrictString517("email123@server.gov"))

		acc := party.Account[0]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(37))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("1115478569"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(38))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("8"))

		acc = party.Account[1]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(39))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("3365998541"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(40))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("8"))

		acc = party.Account[2]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(41))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("4857985691"))
		require.Equal(t, acc.PartyAccountAssociation.SeqNum, fincen.SeqNumber(42))
		require.Equal(t, acc.PartyAccountAssociation.PartyAccountAssociationTypeCode, ValidatePartyAccountAssociationCodeType("9"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

}

func TestElements(t *testing.T) {

	t.Run("ActivityType", func(t *testing.T) {
		activity := NewActivity()

		require.Equal(t, "CTRX", activity.FormTypeCode())
		require.Equal(t, "The Party has invalid min & max range", activity.Validate().Error())

		for i := 0; i < 6; i++ {
			activity.Party = append(activity.Party, PartyType{})
		}
		require.Equal(t, "The Party(type 35) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, PartyType{ActivityPartyTypeCode: "35"})
		require.Equal(t, "The Party(type 37) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, PartyType{ActivityPartyTypeCode: "37"})
		require.Equal(t, "The Party(type 30) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, PartyType{ActivityPartyTypeCode: "30"})
		require.Equal(t, "The Party(type 8) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, PartyType{ActivityPartyTypeCode: "8"})
		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", activity.Validate().Error())
	})

	t.Run("ActivityAssociationType", func(t *testing.T) {
		var sample ActivityAssociationType

		require.Equal(t, "The ActivityAssociation has invalid value", sample.Validate().Error())

		indicator := fincen.ValidateIndicatorType("Y")

		sample.CorrectsAmendsPriorReportIndicator = &indicator
		sample.FinCENDirectBackFileIndicator = &indicator
		sample.InitialReportIndicator = &indicator
		require.Equal(t, "The ActivityAssociation has invalid value", sample.Validate().Error())
	})

	t.Run("PartyType", func(t *testing.T) {
		var sample PartyType

		require.Equal(t, "The PartyName has invalid min & max range", sample.Validate().Error())

		indicator := fincen.ValidateIndicatorType("Y")

		sample.BirthDateUnknownIndicator = &indicator
		require.Equal(t, "The BirthDateUnknownIndicator should be omitted", sample.Validate().Error())

		sample.BirthDateUnknownIndicator = nil
		mmDD := fincen.DateYYYYMMDDType("20170101")
		sample.EFilingCoverageBeginningDateText = &mmDD
		require.Equal(t, "The EFilingCoverageBeginningDateText should be omitted", sample.Validate().Error())

		sample.EFilingCoverageBeginningDateText = nil
		sample.EFilingCoverageEndDateText = &mmDD
		require.Equal(t, "The EFilingCoverageEndDateText should be omitted", sample.Validate().Error())

		sample.EFilingCoverageEndDateText = nil
		sample.FemaleGenderIndicator = &indicator
		require.Equal(t, "The FemaleGenderIndicator should be omitted", sample.Validate().Error())

		sample.FemaleGenderIndicator = nil
		dob := fincen.DateYYYYMMDDOrBlankTypeDOB("20170101")
		sample.IndividualBirthDateText = &dob
		require.Equal(t, "The IndividualBirthDateText should be omitted", sample.Validate().Error())

		sample.IndividualBirthDateText = nil
		v15 := fincen.RestrictString15("SA")
		sample.IndividualEntityCashInAmountText = &v15
		require.Equal(t, "The IndividualEntityCashInAmountText should be omitted", sample.Validate().Error())

		sample.IndividualEntityCashInAmountText = nil
		sample.IndividualEntityCashOutAmountText = &v15
		require.Equal(t, "The IndividualEntityCashOutAmountText should be omitted", sample.Validate().Error())

		sample.IndividualEntityCashOutAmountText = nil
		sample.MaleGenderIndicator = &indicator
		require.Equal(t, "The MaleGenderIndicator should be omitted", sample.Validate().Error())

		sample.MaleGenderIndicator = nil
		sample.MultipleTransactionsPersonsIndividualsIndicator = &indicator
		require.Equal(t, "The MultipleTransactionsPersonsIndividualsIndicator should be omitted", sample.Validate().Error())

		sample.MultipleTransactionsPersonsIndividualsIndicator = nil
		sample.PartyAsEntityOrganizationIndicator = &indicator
		require.Equal(t, "The PartyAsEntityOrganizationIndicator should be omitted", sample.Validate().Error())

		sample.PartyAsEntityOrganizationIndicator = nil
		sample.UnknownGenderIndicator = &indicator
		require.Equal(t, "The UnknownGenderIndicator should be omitted", sample.Validate().Error())

		sample.UnknownGenderIndicator = nil
		c1 := ValidateFederalRegulatorCodeType("1")
		sample.PrimaryRegulatorTypeCode = &c1
		require.Equal(t, "The PrimaryRegulatorTypeCode should be omitted", sample.Validate().Error())

		sample.PrimaryRegulatorTypeCode = nil
		require.Equal(t, "The PartyName has invalid min & max range", sample.Validate().Error())

		sample.PartyName = append(sample.PartyName, PartyNameType{})
		a1 := AddressType{}
		sample.Address = &a1
		require.Equal(t, "The Address should be omitted", sample.Validate().Error())

		sample.Address = nil
		p1 := PhoneNumberType{}
		sample.PhoneNumber = &p1
		require.Equal(t, "The PhoneNumber should be omitted", sample.Validate().Error())

		sample.PhoneNumber = nil
		sample.PartyIdentification = append(sample.PartyIdentification, PartyIdentificationType{})
		require.Equal(t, "The PartyIdentification should be omitted", sample.Validate().Error())

		sample.PartyIdentification = nil
		sample.OrganizationClassificationTypeSubtype = &OrganizationClassificationTypeSubtypeType{}
		require.Equal(t, "The OrganizationClassificationTypeSubtype should be omitted", sample.Validate().Error())

		sample.OrganizationClassificationTypeSubtype = nil
		sample.PartyOccupationBusiness = &PartyOccupationBusinessType{}
		require.Equal(t, "The PartyOccupationBusiness should be omitted", sample.Validate().Error())

		sample.PartyOccupationBusiness = nil
		sample.ElectronicAddress = &ElectronicAddressType{}
		require.Equal(t, "The ElectronicAddress should be omitted", sample.Validate().Error())

		sample.ElectronicAddress = nil
		sample.Account = append(sample.Account, AccountType{})
		require.Equal(t, "The Account should be omitted", sample.Validate("35").Error())

		sample.Account = nil
		sample.PartyIdentification = append(sample.PartyIdentification, PartyIdentificationType{}, PartyIdentificationType{}, PartyIdentificationType{})
		sample.ActivityPartyTypeCode = "35"
		require.Equal(t, "The PartyIdentification has invalid min & max range", sample.Validate("35").Error())

	})

	t.Run("PartyNameType", func(t *testing.T) {
		var sample PartyNameType

		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate().Error())
		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate("INVALID").Error())

		indicator := fincen.ValidateIndicatorType("Y")

		sample.EntityLastNameUnknownIndicator = &indicator
		require.Equal(t, "The EntityLastNameUnknownIndicator should be omitted", sample.Validate("INVALID").Error())

		sample.EntityLastNameUnknownIndicator = nil
		sample.FirstNameUnknownIndicator = &indicator
		require.Equal(t, "The FirstNameUnknownIndicator should be omitted", sample.Validate("INVALID").Error())

		sample.FirstNameUnknownIndicator = nil
		partyName := ValidatePartyNameCodeType("DBA")
		sample.PartyNameTypeCode = &partyName
		require.Equal(t, "The PartyNameTypeCode has invalid value", sample.Validate("INVALID").Error())

		sample.PartyNameTypeCode = nil
		v150 := fincen.RestrictString150("SA")
		sample.RawEntityIndividualLastName = &v150
		require.Equal(t, "The RawEntityIndividualLastName should be omitted", sample.Validate("INVALID").Error())

		sample.RawEntityIndividualLastName = nil
		v35 := fincen.RestrictString35("SA")
		sample.RawIndividualFirstName = &v35
		require.Equal(t, "The RawIndividualFirstName should be omitted", sample.Validate("INVALID").Error())

		sample.RawIndividualFirstName = nil
		sample.RawIndividualMiddleName = &v35
		require.Equal(t, "The RawIndividualMiddleName should be omitted", sample.Validate("INVALID").Error())

		sample.RawIndividualMiddleName = nil
		sample.RawIndividualNameSuffixText = &v35
		require.Equal(t, "The RawIndividualNameSuffixText should be omitted", sample.Validate("INVALID").Error())

		sample.RawIndividualNameSuffixText = nil
		sample.RawPartyFullName = &v150
		require.Equal(t, "The RawPartyFullName should be omitted", sample.Validate("50").Error())
	})

	t.Run("AddressType", func(t *testing.T) {
		var sample AddressType

		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate().Error())
		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate("INVALID").Error())

		indicator := fincen.ValidateIndicatorType("Y")

		sample.CityUnknownIndicator = &indicator
		require.Equal(t, "The CityUnknownIndicator should be omitted", sample.Validate("INVALID").Error())

		sample.CityUnknownIndicator = nil
		sample.CountryCodeUnknownIndicator = &indicator
		require.Equal(t, "The CountryCodeUnknownIndicator should be omitted", sample.Validate("INVALID").Error())

		sample.CountryCodeUnknownIndicator = nil
		v50 := fincen.RestrictString50("SA")
		sample.RawCityText = &v50
		require.Equal(t, "The RawCityText should be omitted", sample.Validate("INVALID").Error())

		sample.RawCityText = nil
		v2 := fincen.RestrictString2("SA")
		sample.RawCountryCodeText = &v2
		require.Equal(t, "The RawCountryCodeText should be omitted", sample.Validate("INVALID").Error())

		sample.RawCountryCodeText = nil
		v3 := fincen.RestrictString3("SA")
		sample.RawStateCodeText = &v3
		require.Equal(t, "The RawStateCodeText should be omitted", sample.Validate("INVALID").Error())

		sample.RawStateCodeText = nil
		v100 := fincen.RestrictString100("SA")
		sample.RawStreetAddress1Text = &v100
		require.Equal(t, "The RawStreetAddress1Text should be omitted", sample.Validate("INVALID").Error())

		sample.RawStreetAddress1Text = nil
		v9 := fincen.RestrictString9("SA")
		sample.RawZIPCode = &v9
		require.Equal(t, "The RawZIPCode should be omitted", sample.Validate("INVALID").Error())

		sample.RawZIPCode = nil
		sample.StateCodeUnknownIndicator = &indicator
		require.Equal(t, "The StateCodeUnknownIndicator should be omitted", sample.Validate("INVALID").Error())

		sample.StateCodeUnknownIndicator = nil
		sample.StreetAddressUnknownIndicator = &indicator
		require.Equal(t, "The StreetAddressUnknownIndicator should be omitted", sample.Validate("INVALID").Error())

		sample.StreetAddressUnknownIndicator = nil
		sample.ZIPCodeUnknownIndicator = &indicator
		require.Equal(t, "The ZIPCodeUnknownIndicator should be omitted", sample.Validate("INVALID").Error())
	})

	t.Run("PhoneNumberType", func(t *testing.T) {
		var sample PhoneNumberType

		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate().Error())
		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate("INVALID").Error())

		v6 := fincen.RestrictString6("SA")
		sample.PhoneNumberExtensionText = &v6
		require.Equal(t, "The PhoneNumberExtensionText should be omitted", sample.Validate("INVALID").Error())

		sample.PhoneNumberExtensionText = nil
		v16 := fincen.RestrictString16("SA")
		sample.PhoneNumberText = &v16
		require.Equal(t, "The PhoneNumberText should be omitted", sample.Validate("INVALID").Error())
	})

	t.Run("PartyIdentificationType", func(t *testing.T) {
		var sample PartyIdentificationType

		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate().Error())
		require.Equal(t, "The SeqNumber has invalid value (SeqNumber)", sample.Validate("INVALID").Error())

		indicator := fincen.ValidateIndicatorType("Y")
		sample.IdentificationPresentUnknownIndicator = &indicator
		require.Equal(t, "The IdentificationPresentUnknownIndicator should be omitted", sample.Validate("INVALID").Error())

		sample.IdentificationPresentUnknownIndicator = nil
		v2 := fincen.RestrictString2("SA")
		sample.OtherIssuerCountryText = &v2
		require.Equal(t, "The OtherIssuerCountryText should be omitted", sample.Validate("INVALID").Error())

		sample.OtherIssuerCountryText = nil
		v3 := fincen.RestrictString3("SA")
		sample.OtherIssuerStateText = &v3
		require.Equal(t, "The OtherIssuerStateText should be omitted", sample.Validate("INVALID").Error())

		sample.OtherIssuerStateText = nil
		v50 := fincen.RestrictString50("SA")
		sample.OtherPartyIdentificationTypeText = &v50
		require.Equal(t, "The OtherPartyIdentificationTypeText should be omitted", sample.Validate("INVALID").Error())

		sample.OtherPartyIdentificationTypeText = nil
		v25 := fincen.RestrictString25("SA")
		sample.PartyIdentificationNumberText = &v25
		require.Equal(t, "The PartyIdentificationNumberText should be omitted", sample.Validate("INVALID").Error())

		sample.PartyIdentificationNumberText = nil
		identificationCode := ValidatePartyIdentificationCodeType("1")
		sample.PartyIdentificationTypeCode = &identificationCode
		require.Equal(t, "The PartyIdentificationTypeCode should be omitted", sample.Validate("INVALID").Error())

		sample.PartyIdentificationTypeCode = nil
		sample.TINUnknownIndicator = &indicator
		require.Equal(t, "The TINUnknownIndicator should be omitted", sample.Validate("INVALID").Error())

	})

}
