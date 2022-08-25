package suspicious_activity

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/moov-io/fincen"

	"github.com/stretchr/testify/require"
)

func mockEFilingBatchXML() *EFilingBatchXML {
	batch := EFilingBatchXML{}
	return &batch
}

func mocParties() map[string][]byte {
	parties := make(map[string][]byte)

	parties["35"] = []byte(`<Party SeqNum="4">
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
</Party>`)

	parties["37"] = []byte(`<Party SeqNum="10">
	<ActivityPartyTypeCode>37</ActivityPartyTypeCode>
	<PartyName SeqNum="11">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Transmitter contact legal name</RawPartyFullName>
	</PartyName>
</Party>`)

	parties["30"] = []byte(`<Party SeqNum="12">
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
		<PartyIdentificationTypeCode>29</PartyIdentificationTypeCode>
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

	parties["8"] = []byte(`<Party SeqNum="26">
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

	parties["18"] = []byte(`<Party SeqNum="21">
	<ActivityPartyTypeCode>18</ActivityPartyTypeCode>
	<PartyName SeqNum="22">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>LE Contact Agency</RawPartyFullName>
	</PartyName>
</Party>`)

	parties["19"] = []byte(`<Party SeqNum="23">
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

	parties["34"] = []byte(`<Party SeqNum="29">
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
			<Address SeqNum="0"></Address>
		</Party>
	</PartyAssociation>
</Party>`)

	parties["46"] = []byte(`<Party SeqNum="39">
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

	parties["41"] = []byte(`<Party SeqNum="75">
	<ActivityPartyTypeCode>41</ActivityPartyTypeCode>
	<NonUSFinancialInstitutionIndicator>Y</NonUSFinancialInstitutionIndicator>
	<PartyIdentification SeqNum="76">
		<PartyIdentificationNumberText>458789856</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>4</PartyIdentificationTypeCode>
	</PartyIdentification>
	<Account SeqNum="77">
		<AccountNumberText>1502417873</AccountNumberText>
		<PartyAccountAssociation SeqNum="78">
			<AccountClosedIndicator>Y</AccountClosedIndicator>
			<PartyAccountAssociationTypeCode>5</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
	<Account SeqNum="79">
		<AccountNumberText>5477887896</AccountNumberText>
		<PartyAccountAssociation SeqNum="80">
			<AccountClosedIndicator>Y</AccountClosedIndicator>
			<PartyAccountAssociationTypeCode>5</PartyAccountAssociationTypeCode>
		</PartyAccountAssociation>
	</Account>
</Party>`)

	parties["33"] = []byte(`<Party SeqNum="61">
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
		<Party SeqNum="75">
			<ActivityPartyTypeCode>41</ActivityPartyTypeCode>
		</Party>
		<PartyAccountAssociationTypeCode>7</PartyAccountAssociationTypeCode>
	</PartyAccountAssociation>
</Party>`)

	return parties
}

func TestPartyByTypeCode(t *testing.T) {
	t.Run("Transmitter Party Record", func(t *testing.T) {
		sample := mocParties()["35"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, int64(4))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("35"))
		require.Equal(t, len(party.PartyName), 1)
		require.Equal(t, len(party.Address), 1)
		require.Equal(t, len(party.PhoneNumber), 1)
		require.Equal(t, len(party.PartyIdentification), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, int64(5))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter legal name"))

		address := party.Address[0]
		require.Equal(t, address.SeqNum, int64(6))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Transmitter city"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("Transmitter street address"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("22113"))

		number := party.PhoneNumber[0]
		require.Equal(t, number.SeqNum, int64(7))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("7217894455"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, int64(8))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("458985215"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("4"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Transmitter Contact Party Record", func(t *testing.T) {
		sample := mocParties()["37"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, int64(10))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("37"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, int64(11))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter contact legal name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Filing Institution Party Record", func(t *testing.T) {
		sample := mocParties()["30"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, int64(12))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("30"))
		require.Equal(t, *party.PrimaryRegulatorTypeCode, ValidateFederalRegulatorCodeType("7"))
		require.Equal(t, len(party.PartyName), 2)
		require.Equal(t, len(party.Address), 1)
		require.Equal(t, len(party.PartyIdentification), 3)
		require.Equal(t, len(party.OrganizationClassificationTypeSubtype), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, int64(13))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Filer name"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, int64(14))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("DBA"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Alternate name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Designated Contact Office Party Record", func(t *testing.T) {
		sample := mocParties()["8"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, int64(26))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("8"))
		require.Equal(t, len(party.PartyName), 1)
		require.Equal(t, len(party.PhoneNumber), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, int64(27))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Designated contact office"))

		number := party.PhoneNumber[0]
		require.Equal(t, number.SeqNum, int64(28))
		require.Equal(t, *number.PhoneNumberExtensionText, fincen.RestrictString6("1234"))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("4157653838"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Law Enforcement Agency Party Record", func(t *testing.T) {
		sample := mocParties()["18"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, int64(21))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("18"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, int64(22))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("LE Contact Agency"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Law Enforcement Name Party Record", func(t *testing.T) {
		sample := mocParties()["19"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, int64(23))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("19"))
		require.Equal(t, len(party.PartyName), 1)
		require.Equal(t, len(party.PhoneNumber), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, int64(24))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("LE Contact Name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Financial Institution Where Activity Occurred Party Record", func(t *testing.T) {
		sample := mocParties()["34"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, int64(29))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("34"))
		require.Equal(t, *party.LossToFinancialAmountText, fincen.RestrictString15("15000"))
		require.Equal(t, *party.NoBranchActivityInvolvedIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.PayLocationIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.PrimaryRegulatorTypeCode, ValidateFederalRegulatorCodeType("4"))
		require.Equal(t, len(party.PartyName), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, int64(30))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Union Bank of California"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, int64(31))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("DBA"))
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

		require.Equal(t, party.SeqNum, int64(61))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("33"))
		require.Equal(t, *party.AdmissionConfessionYesIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.BothPurchaserSenderPayeeReceiveIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.FemaleGenderIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, *party.IndividualBirthDateText, fincen.DateYYYYMMDDOrBlankTypeDOB("19801025"))
		require.Equal(t, len(party.PartyName), 3)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, int64(62))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Mann"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("Janice"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, int64(63))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("AKA"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Janda"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Branch Where Activity Occurred Party Record", func(t *testing.T) {
		sample := mocParties()["46"]
		party := AssociationParty{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, int64(39))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("46"))
		require.Equal(t, *party.SellingLocationIndicator, fincen.ValidateIndicatorType("Y"))

		address := party.Address
		require.Equal(t, address.SeqNum, int64(40))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Cityville"))

		identification := party.PartyIdentification
		require.Equal(t, identification.SeqNum, int64(41))
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

		require.Equal(t, party.SeqNum, int64(75))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("41"))
		require.Equal(t, *party.NonUSFinancialInstitutionIndicator, fincen.ValidateIndicatorType("Y"))
		require.Equal(t, len(party.Account), 2)

		identification := party.PartyIdentification
		require.Equal(t, identification.SeqNum, int64(76))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("458789856"))

		acc := party.Account[0]
		require.Equal(t, acc.SeqNum, int64(77))
		require.Equal(t, *acc.AccountNumberText, fincen.RestrictString40("1502417873"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})
}