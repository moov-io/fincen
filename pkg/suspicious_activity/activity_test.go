// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package suspicious_activity

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/moov-io/fincen"

	"github.com/stretchr/testify/require"
)

func mocParties() map[string][]byte {
	parties := make(map[string][]byte)

	parties[PartyTransmitter] = []byte(`<Party SeqNum="4">
	<ActivityPartyTypeCode>35</ActivityPartyTypeCode>
	<PartyName SeqNum="5">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Transmitter legal name</RawPartyFullName>
	</PartyName>
	<Address SeqNum="6">
		<RawCityText>Transmitter city</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<RawStreetAddress1Text>Transmitter street address</RawStreetAddress1Text>
		<RawZIPCode>22113</RawZIPCode>
	</Address>
	<PhoneNumber SeqNum="7">
		<PhoneNumberText>7217894455</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="8">
		<PartyIdentificationNumberText>458985215</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>4</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="9">
		<PartyIdentificationNumberText>458985215</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>28</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties[PartyTransmitterContact] = []byte(`<Party SeqNum="10">
	<ActivityPartyTypeCode>37</ActivityPartyTypeCode>
	<PartyName SeqNum="11">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Transmitter contact legal name</RawPartyFullName>
	</PartyName>
</Party>`)

	parties[PartyFilingInstitution] = []byte(`<Party SeqNum="12">
	<ActivityPartyTypeCode>30</ActivityPartyTypeCode>
	<PrimaryRegulatorTypeCode>7</PrimaryRegulatorTypeCode>
	<PartyName SeqNum="13">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Filer name</RawPartyFullName>
	</PartyName>
	<PartyName SeqNum="14">
		<PartyNameTypeCode>DBA</PartyNameTypeCode>
		<RawPartyFullName>Alternate name</RawPartyFullName>
	</PartyName>
	<Address SeqNum="15">
		<RawCityText>Rockville</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>MD</RawStateCodeText>
		<RawStreetAddress1Text>123 Viers Mill Road</RawStreetAddress1Text>
		<RawZIPCode>20905</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="16">
		<PartyIdentificationNumberText>125478965</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="17">
		<PartyIdentificationNumberText>5558789654</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>10</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="18">
		<PartyIdentificationNumberText>451256558</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>4</PartyIdentificationTypeCode>
	</PartyIdentification>
	<OrganizationClassificationTypeSubtype SeqNum="19">
		<OrganizationSubtypeID>535</OrganizationSubtypeID>
		<OrganizationTypeID>5</OrganizationTypeID>
	</OrganizationClassificationTypeSubtype>
	<OrganizationClassificationTypeSubtype SeqNum="20">
		<OrganizationSubtypeID>529</OrganizationSubtypeID>
		<OrganizationTypeID>5</OrganizationTypeID>
	</OrganizationClassificationTypeSubtype>
</Party>`)

	parties[PartyContactOffice] = []byte(`<Party SeqNum="26">
	<ActivityPartyTypeCode>8</ActivityPartyTypeCode>
	<PartyName SeqNum="27">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Designated contact office</RawPartyFullName>
	</PartyName>
	<PhoneNumber SeqNum="28">
		<PhoneNumberExtensionText>1234</PhoneNumberExtensionText>
		<PhoneNumberText>4157653838</PhoneNumberText>
	</PhoneNumber>
</Party>`)

	parties[PartyLEContactAgency] = []byte(`<Party SeqNum="21">
	<ActivityPartyTypeCode>18</ActivityPartyTypeCode>
	<PartyName SeqNum="22">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>LE Contact Agency</RawPartyFullName>
	</PartyName>
</Party>`)

	parties[PartyLEContactName] = []byte(`<Party SeqNum="23">
	<ActivityPartyTypeCode>19</ActivityPartyTypeCode>
	<ContactDateText>20171105</ContactDateText>
	<PartyName SeqNum="24">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>LE Contact Name</RawPartyFullName>
	</PartyName>
	<PhoneNumber SeqNum="25">
		<PhoneNumberExtensionText>5432</PhoneNumberExtensionText>
		<PhoneNumberText>7039051234</PhoneNumberText>
	</PhoneNumber>
</Party>`)

	parties[PartyFinancialInstitution] = []byte(`<Party SeqNum="29">
	<ActivityPartyTypeCode>34</ActivityPartyTypeCode>
	<LossToFinancialAmountText>15000</LossToFinancialAmountText>
	<NoBranchActivityInvolvedIndicator>Y</NoBranchActivityInvolvedIndicator>
	<PayLocationIndicator>Y</PayLocationIndicator>
	<PrimaryRegulatorTypeCode>4</PrimaryRegulatorTypeCode>
	<PartyName SeqNum="30">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Union Bank of California</RawPartyFullName>
	</PartyName>
	<PartyName SeqNum="31">
		<PartyNameTypeCode>DBA</PartyNameTypeCode>
		<RawPartyFullName>Doing Business As Name</RawPartyFullName>
	</PartyName>
	<Address SeqNum="32">
		<CityUnknownIndicator>Y</CityUnknownIndicator>
		<CountryCodeUnknownIndicator>Y</CountryCodeUnknownIndicator>
		<RawStreetAddress1Text>987 Rocky Road</RawStreetAddress1Text>
		<StreetAddressUnknownIndicator>Y</StreetAddressUnknownIndicator>
		<ZIPCodeUnknownIndicator>Y</ZIPCodeUnknownIndicator>
	</Address>
	<PartyIdentification SeqNum="33">
		<PartyIdentificationNumberText>458789856</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="34">
		<PartyIdentificationNumberText>5589887789</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>10</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="35">
		<PartyIdentificationNumberText>4578958658</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>29</PartyIdentificationTypeCode>
	</PartyIdentification>
	<OrganizationClassificationTypeSubtype SeqNum="36">
		<OrganizationSubtypeID>533</OrganizationSubtypeID>
		<OrganizationTypeID>5</OrganizationTypeID>
	</OrganizationClassificationTypeSubtype>
	<OrganizationClassificationTypeSubtype SeqNum="37">
		<OrganizationSubtypeID>5999</OrganizationSubtypeID>
		<OrganizationTypeID>5</OrganizationTypeID>
		<OtherOrganizationSubTypeText>other S/F institution &#xA;description</OtherOrganizationSubTypeText>
	</OrganizationClassificationTypeSubtype>
	<PartyAssociation SeqNum="38">
		<Party SeqNum="39">
			<ActivityPartyTypeCode>46</ActivityPartyTypeCode>
			<Address SeqNum="40"></Address>
		</Party>
	</PartyAssociation>
</Party>`)

	parties[PartyBranch] = []byte(`<Party SeqNum="39">
	<ActivityPartyTypeCode>46</ActivityPartyTypeCode>
	<SellingLocationIndicator>Y</SellingLocationIndicator>
	<Address SeqNum="40">
		<RawCityText>Cityville</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<RawStreetAddress1Text>123 Address Way</RawStreetAddress1Text>
		<RawZIPCode>547898856</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="41">
		<PartyIdentificationNumberText>445564615654</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>14</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties[PartyFinancialInstitutionAccount] = []byte(`<Party SeqNum="75">
	<ActivityPartyTypeCode>41</ActivityPartyTypeCode>
	<NonUSFinancialInstitutionIndicator>Y</NonUSFinancialInstitutionIndicator>
	<PartyIdentification SeqNum="76">
		<PartyIdentificationNumberText>458789856</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>4</PartyIdentificationTypeCode>
	</PartyIdentification>
	<Account SeqNum="77">
		<AccountNumberText>1502417873</AccountNumberText>
		<PartyAccountAssociation SeqNum="78">
			<PartyAccountAssociationTypeCode>5</PartyAccountAssociationTypeCode>
			<AccountClosedIndicator>Y</AccountClosedIndicator>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="79">
		<AccountNumberText>5477887896</AccountNumberText>
		<PartyAccountAssociation SeqNum="80">
			<PartyAccountAssociationTypeCode>5</PartyAccountAssociationTypeCode>
			<AccountClosedIndicator>Y</AccountClosedIndicator>
		</PartyAccountAssociation>
	</Account>
</Party>`)

	parties[PartySubject] = []byte(`<Party SeqNum="61">
	<ActivityPartyTypeCode>33</ActivityPartyTypeCode>
	<AdmissionConfessionYesIndicator>Y</AdmissionConfessionYesIndicator>
	<BothPurchaserSenderPayeeReceiveIndicator>Y</BothPurchaserSenderPayeeReceiveIndicator>
	<FemaleGenderIndicator>Y</FemaleGenderIndicator>
	<IndividualBirthDateText>19801025</IndividualBirthDateText>
	<PartyName SeqNum="62">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Mann</RawEntityIndividualLastName>
		<RawIndividualFirstName>Janice</RawIndividualFirstName>
	</PartyName>
	<PartyName SeqNum="63">
		<PartyNameTypeCode>AKA</PartyNameTypeCode>
		<RawPartyFullName>Janda</RawPartyFullName>
	</PartyName>
	<PartyName SeqNum="64">
		<PartyNameTypeCode>AKA</PartyNameTypeCode>
		<RawPartyFullName>Jan</RawPartyFullName>
	</PartyName>
	<Address SeqNum="65">
		<CityUnknownIndicator>Y</CityUnknownIndicator>
		<CountryCodeUnknownIndicator>Y</CountryCodeUnknownIndicator>
		<StateCodeUnknownIndicator>Y</StateCodeUnknownIndicator>
		<StreetAddressUnknownIndicator>Y</StreetAddressUnknownIndicator>
		<ZIPCodeUnknownIndicator>Y</ZIPCodeUnknownIndicator>
	</Address>
	<PhoneNumber SeqNum="66">
		<PhoneNumberText>6194760276</PhoneNumberText>
		<PhoneNumberTypeCode>R</PhoneNumberTypeCode>
	</PhoneNumber>
	<PhoneNumber SeqNum="67">
		<PhoneNumberExtensionText>5584</PhoneNumberExtensionText>
		<PhoneNumberText>6589784589</PhoneNumberText>
		<PhoneNumberTypeCode>W</PhoneNumberTypeCode>
	</PhoneNumber>
	<PartyIdentification SeqNum="68">
		<IdentificationPresentUnknownIndicator>Y</IdentificationPresentUnknownIndicator>
		<OtherIssuerCountryText>US</OtherIssuerCountryText>
		<OtherIssuerStateText>CA</OtherIssuerStateText>
		<OtherPartyIdentificationTypeText>Student ID</OtherPartyIdentificationTypeText>
		<PartyIdentificationNumberText>660623559</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>999</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="69">
		<TINUnknownIndicator>Y</TINUnknownIndicator>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="70">
		<NAICSCode>7548</NAICSCode>
		<OccupationBusinessText>Mechanic</OccupationBusinessText>
	</PartyOccupationBusiness>
	<ElectronicAddress SeqNum="71">
		<ElectronicAddressText>jan@gmail.com</ElectronicAddressText>
		<ElectronicAddressTypeCode>E</ElectronicAddressTypeCode>
	</ElectronicAddress>
	<ElectronicAddress SeqNum="72">
		<ElectronicAddressText>janda.org</ElectronicAddressText>
		<ElectronicAddressTypeCode>U</ElectronicAddressTypeCode>
	</ElectronicAddress>
	<PartyAssociation SeqNum="73">
		<NoRelationshipToInstitutionIndicator>Y</NoRelationshipToInstitutionIndicator>
		<SubjectRelationshipFinancialInstitutionTINText>458789856</SubjectRelationshipFinancialInstitutionTINText>
	</PartyAssociation>
	<PartyAccountAssociation SeqNum="74">
		<PartyAccountAssociationTypeCode>7</PartyAccountAssociationTypeCode>
		<Party SeqNum="75">
			<ActivityPartyTypeCode>41</ActivityPartyTypeCode>
		</Party>
	</PartyAccountAssociation>
</Party>`)

	return parties
}

func TestParty(t *testing.T) {
	t.Run("Transmitter Party Record", func(t *testing.T) {
		sample := mocParties()[PartyTransmitter]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(4))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyTransmitter))
		require.Equal(t, len(party.PartyName), 1)
		require.Equal(t, len(party.Address), 1)
		require.Equal(t, len(party.PhoneNumber), 1)
		require.Equal(t, len(party.PartyIdentification), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(5))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateLegalName))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter legal name"))

		address := party.Address[0]
		require.Equal(t, address.SeqNum, fincen.SeqNumber(6))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Transmitter city"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("Transmitter street address"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("22113"))

		number := party.PhoneNumber[0]
		require.Equal(t, number.SeqNum, fincen.SeqNumber(7))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("7217894455"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(8))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("458985215"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("4"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Transmitter Contact Party Record", func(t *testing.T) {
		sample := mocParties()[PartyTransmitterContact]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(10))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyTransmitterContact))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(11))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateLegalName))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter contact legal name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Filing Institution Party Record", func(t *testing.T) {
		sample := mocParties()[PartyFilingInstitution]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(12))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyFilingInstitution))
		require.Equal(t, *party.PrimaryRegulatorTypeCode, ValidateFederalRegulatorCodeType("7"))
		require.Equal(t, len(party.PartyName), 2)
		require.Equal(t, len(party.Address), 1)
		require.Equal(t, len(party.PartyIdentification), 3)
		require.Equal(t, len(party.OrganizationClassificationTypeSubtype), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(13))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateLegalName))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Filer name"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(14))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateDoingBusiness))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Alternate name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Designated Contact Office Party Record", func(t *testing.T) {
		sample := mocParties()[PartyContactOffice]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(26))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyContactOffice))
		require.Equal(t, len(party.PartyName), 1)
		require.Equal(t, len(party.PhoneNumber), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(27))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateLegalName))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Designated contact office"))

		number := party.PhoneNumber[0]
		require.Equal(t, number.SeqNum, fincen.SeqNumber(28))
		require.Equal(t, *number.PhoneNumberExtensionText, fincen.RestrictString6("1234"))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("4157653838"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Law Enforcement Agency Party Record", func(t *testing.T) {
		sample := mocParties()[PartyLEContactAgency]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(21))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyLEContactAgency))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(22))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateLegalName))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("LE Contact Agency"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Law Enforcement Name Party Record", func(t *testing.T) {
		sample := mocParties()[PartyLEContactName]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(23))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyLEContactName))
		require.Equal(t, len(party.PartyName), 1)
		require.Equal(t, len(party.PhoneNumber), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(24))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateLegalName))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("LE Contact Name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Financial Institution Where Activity Occurred Party Record", func(t *testing.T) {
		sample := mocParties()[PartyFinancialInstitution]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(29))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyFinancialInstitution))
		require.Equal(t, *party.LossToFinancialAmountText, fincen.RestrictString15("15000"))
		require.Equal(t, *party.NoBranchActivityInvolvedIndicator, fincen.ValidateIndicatorNullType("Y"))
		require.Equal(t, *party.PayLocationIndicator, fincen.ValidateIndicatorNullType("Y"))
		require.Equal(t, *party.PrimaryRegulatorTypeCode, ValidateFederalRegulatorCodeType("4"))
		require.Equal(t, len(party.PartyName), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(30))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateLegalName))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Union Bank of California"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(31))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateDoingBusiness))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Doing Business As Name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Financial Institution Where Activity Occurred Party Record", func(t *testing.T) {
		sample := mocParties()["33"]
		party := PartyType{}

		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(61))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartySubject))
		require.Equal(t, *party.AdmissionConfessionYesIndicator, fincen.ValidateIndicatorNullType("Y"))
		require.Equal(t, *party.BothPurchaserSenderPayeeReceiveIndicator, fincen.ValidateIndicatorNullType("Y"))
		require.Equal(t, *party.FemaleGenderIndicator, fincen.ValidateIndicatorNullType("Y"))
		require.Equal(t, *party.IndividualBirthDateText, fincen.DateYYYYMMDDOrBlankTypeDOB("19801025"))
		require.Equal(t, len(party.PartyName), 3)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(62))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateLegalName))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Mann"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("Janice"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(63))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType(fincen.IndicateKnownAs))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Janda"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Branch Where Activity Occurred Party Record", func(t *testing.T) {
		sample := mocParties()[PartyBranch]
		party := AssociationParty{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(39))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyBranch))
		require.Equal(t, *party.SellingLocationIndicator, fincen.ValidateIndicatorNullType("Y"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(40))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Cityville"))

		identification := party.PartyIdentification
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(41))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("445564615654"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Financial Institution Where Account is Held Party Record", func(t *testing.T) {
		sample := mocParties()["41"]
		party := AccountAssociationParty{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(75))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType(PartyFinancialInstitutionAccount))
		require.Equal(t, *party.NonUSFinancialInstitutionIndicator, fincen.ValidateIndicatorNullType("Y"))
		require.Equal(t, len(party.Account), 2)

		identification := party.PartyIdentification
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(76))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("458789856"))

		acc := party.Account[0]
		require.Equal(t, acc.SeqNum, fincen.SeqNumber(77))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("1502417873"))

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
		activity.ActivityAssociation = &ActivityAssociationType{}
		activity.SuspiciousActivity = &SuspiciousActivityType{}

		require.Equal(t, "SARX", activity.FormTypeCode())
		require.Equal(t, "The Party has invalid min & max range", activity.Validate().Error())

		indicator := fincen.ValidateIndicatorNullType("Y")

		activity.EFilingPriorDocumentNumber = nil
		activity.ActivityAssociation.CorrectsAmendsPriorReportIndicator = &indicator
		require.Equal(t, "The EFilingPriorDocumentNumber is a required field", activity.Validate().Error())

		number := EFilingPriorDocumentNumberType("00000000000000")
		activity.EFilingPriorDocumentNumber = &number
		activity.ActivityAssociation.CorrectsAmendsPriorReportIndicator = nil
		require.Equal(t, "The EFilingPriorDocumentNumber has invalid value", activity.Validate().Error())

		activity.EFilingPriorDocumentNumber = nil
		activity.ActivityAssociation.CorrectsAmendsPriorReportIndicator = nil

		for i := 0; i < 6; i++ {
			activity.Party = append(activity.Party, &PartyType{})
		}
		require.Equal(t, "The Party type(transmitter) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: PartyTransmitter})
		require.Equal(t, "The Party type(transmitter contact) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: PartyTransmitterContact})
		require.Equal(t, "The Party type(filing institution) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: PartyFilingInstitution})
		require.Equal(t, "The Party type(contact office) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: PartyContactOffice})
		require.Equal(t, "The Party type(financial institution where activity occurred) is a required field", activity.Validate().Error())
	})

	t.Run("ActivityAssociationType", func(t *testing.T) {
		var sample ActivityAssociationType
		require.Equal(t, "The ActivityAssociation has invalid value", sample.Validate().Error())
		indicator := fincen.ValidateIndicatorNullType("Y")
		sample.InitialReportIndicator = &indicator
		sample.CorrectsAmendsPriorReportIndicator = &indicator
		require.Equal(t, "The ActivityAssociation has invalid value", sample.Validate().Error())
	})

	t.Run("PartyType", func(t *testing.T) {
		var sample PartyType
		require.Equal(t, "The ActivityPartyCode has invalid value", sample.Validate().Error())
		sample.ActivityPartyTypeCode = PartyTransmitter
		require.Equal(t, "The Party has invalid value", sample.Validate().Error())
		sample.ActivityPartyTypeCode = PartyTransmitterContact
		require.Equal(t, "The Party has invalid value", sample.Validate().Error())
		sample.ActivityPartyTypeCode = PartyFilingInstitution
		require.Equal(t, "The Party has invalid value", sample.Validate().Error())
		sample.ActivityPartyTypeCode = PartyContactOffice
		require.Equal(t, "The Party has invalid value", sample.Validate().Error())
		sample.ActivityPartyTypeCode = PartyLEContactAgency
		require.Equal(t, "The Party has invalid value", sample.Validate().Error())
		sample.ActivityPartyTypeCode = PartyFinancialInstitution
		require.Equal(t, "The Party has invalid value", sample.Validate().Error())
		sample.ActivityPartyTypeCode = PartySubject
		require.Equal(t, "The Party has invalid value", sample.Validate().Error())
	})

	t.Run("PartyNameType", func(t *testing.T) {
		var sample PartyNameType
		require.NoError(t, sample.Validate())
		require.NoError(t, sample.Validate("INVALID"))
		require.Equal(t, "The PartyName has invalid value", sample.Validate(PartyTransmitter).Error())
		require.Equal(t, "The PartyName has invalid value", sample.Validate(PartyFilingInstitution).Error())
		require.Equal(t, "The PartyName has invalid value", sample.Validate(PartyFinancialInstitution).Error())
		c := ValidatePartyNameCodeType("TEST")
		sample.PartyNameTypeCode = &c
		require.Equal(t, "The PartyName has invalid value", sample.Validate(PartyFinancialInstitution).Error())
	})

	t.Run("AddressType", func(t *testing.T) {
		var sample AddressType
		require.NoError(t, sample.Validate())
		require.NoError(t, sample.Validate("INVALID"))
		require.Equal(t, "The Address has invalid value", sample.Validate(PartyTransmitter).Error())
	})

	t.Run("PhoneNumberType", func(t *testing.T) {
		var sample PhoneNumberType
		require.NoError(t, sample.Validate())
		require.NoError(t, sample.Validate("INVALID"))
		require.Equal(t, "The PhoneNumber has invalid value", sample.Validate(PartyContactOffice).Error())
	})

	t.Run("PartyIdentificationType", func(t *testing.T) {
		var sample PartyIdentificationType
		require.NoError(t, sample.Validate())
		require.NoError(t, sample.Validate("INVALID"))
		require.Equal(t, "The PartyIdentification has invalid value", sample.Validate(PartyFilingInstitution).Error())
	})

	t.Run("PartyAssociationType", func(t *testing.T) {
		var sample PartyAssociationType
		require.NoError(t, sample.Validate())
		require.NoError(t, sample.Validate("INVALID"))
		for i := 0; i < 100; i++ {
			sample.Party = append(sample.Party, &AssociationParty{})
		}
		require.Equal(t, "The Party has invalid min & max range", sample.Validate(PartyFinancialInstitution).Error())
	})

	t.Run("PartyAccountAssociationType", func(t *testing.T) {
		var sample PartyAccountAssociationType
		require.Equal(t, "The PartyAccountAssociationCode has invalid value", sample.Validate().Error())
		require.Equal(t, "The PartyAccountAssociationCode has invalid value", sample.Validate("INVALID").Error())
		require.Equal(t, "The Party has invalid min & max range", sample.Validate(PartySubject).Error())
	})

	t.Run("AccountType", func(t *testing.T) {
		var sample AccountType
		require.Equal(t, "The PartyAccountAssociationCode has invalid value (PartyAccountAssociationType)", sample.Validate().Error())
		require.Equal(t, "The PartyAccountAssociationCode has invalid value (PartyAccountAssociationType)", sample.Validate("INVALID").Error())
		require.Equal(t, "The Account has invalid value", sample.Validate(PartyFinancialInstitutionAccount).Error())
	})

	t.Run("SuspiciousActivityType", func(t *testing.T) {
		var sample SuspiciousActivityType
		require.Equal(t, "The SuspiciousActivity has invalid min & max range", sample.Validate().Error())

		sample.SuspiciousActivityClassification = append(sample.SuspiciousActivityClassification, &SuspiciousActivityClassificationType{})
		require.Equal(t, "The DateYYYYMMDDType has invalid value", sample.Validate().Error())
	})
}
