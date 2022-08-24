package suspicious_activity

import (
	"encoding/xml"
	"fmt"
	"github.com/moov-io/fincen"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// ActualAmountPaid creates a ActualAmountPaid
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
	<PartyIdentification SeqNum="9">
		<PartyIdentificationNumberText>PBSA1234</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>28</PartyIdentificationTypeCode>
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

	parties["18"] = []byte(`
<fc2:Party SeqNum="10">
<fc2:ActivityPartyTypeCode>37</fc2:ActivityPartyTypeCode>
<fc2:PartyName SeqNum="11">
<fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
<fc2:RawPartyFullName>Transmitter contact legal name</fc2:RawPartyFullName>
</fc2:PartyName>
</fc2:Party>
`)

	parties["34"] = []byte(`
<fc2:Party SeqNum="29">
<fc2:ActivityPartyTypeCode>34</fc2:ActivityPartyTypeCode>
<fc2:LossToFinancialAmountText>15000</fc2:LossToFinancialAmountText>
<fc2:NoBranchActivityInvolvedIndicator>Y</fc2:NoBranchActivityInvolvedIndicator>
<fc2:PayLocationIndicator>Y</fc2:PayLocationIndicator>
<fc2:PrimaryRegulatorTypeCode>4</fc2:PrimaryRegulatorTypeCode>
<fc2:PartyName SeqNum="30">
<fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
<fc2:RawPartyFullName>Union Bank of California</fc2:RawPartyFullName>
</fc2:PartyName>
<fc2:PartyName SeqNum="31">
<fc2:PartyNameTypeCode>DBA</fc2:PartyNameTypeCode>
<fc2:RawPartyFullName>Doing Business As Name</fc2:RawPartyFullName>
</fc2:PartyName>
<fc2:Address SeqNum="32">
<fc2:CityUnknownIndicator>Y</fc2:CityUnknownIndicator>
<fc2:CountryCodeUnknownIndicator>Y</fc2:CountryCodeUnknownIndicator>
<fc2:RawStreetAddress1Text>987 Rocky Road</fc2:RawStreetAddress1Text>
<fc2:StreetAddressUnknownIndicator>Y</fc2:StreetAddressUnknownIndicator>
<fc2:ZIPCodeUnknownIndicator>Y</fc2:ZIPCodeUnknownIndicator>
</fc2:Address>
<fc2:PartyIdentification SeqNum="33">
<fc2:PartyIdentificationNumberText>458789856</fc2:PartyIdentificationNumberText>
<fc2:PartyIdentificationTypeCode>2</fc2:PartyIdentificationTypeCode>
</fc2:PartyIdentification>
<fc2:PartyIdentification SeqNum="34">
<fc2:PartyIdentificationNumberText>5589887789</fc2:PartyIdentificationNumberText>
<fc2:PartyIdentificationTypeCode>10</fc2:PartyIdentificationTypeCode>
</fc2:PartyIdentification>
<fc2:PartyIdentification SeqNum="35">
<fc2:PartyIdentificationNumberText>4578958658</fc2:PartyIdentificationNumberText>
<fc2:PartyIdentificationTypeCode>29</fc2:PartyIdentificationTypeCode>
</fc2:PartyIdentification>
<fc2:OrganizationClassificationTypeSubtype SeqNum="36">
<fc2:OrganizationSubtypeID>533</fc2:OrganizationSubtypeID>
<fc2:OrganizationTypeID>5</fc2:OrganizationTypeID>
</fc2:OrganizationClassificationTypeSubtype>
<fc2:OrganizationClassificationTypeSubtype SeqNum="37">
<fc2:OrganizationSubtypeID>5999</fc2:OrganizationSubtypeID>
<fc2:OrganizationTypeID>5</fc2:OrganizationTypeID>
<fc2:OtherOrganizationSubTypeText>other S/F institution 
description</fc2:OtherOrganizationSubTypeText>
</fc2:OrganizationClassificationTypeSubtype>
<fc2:PartyAssociation SeqNum="38">
<fc2:Party SeqNum="39">
<fc2:ActivityPartyTypeCode>46</fc2:ActivityPartyTypeCode>
</fc2:Party>
</fc2:PartyAssociation>
</fc2:Party>
`)

	parties["46"] = []byte(`
<fc2:Party SeqNum="39">
<fc2:ActivityPartyTypeCode>46</fc2:ActivityPartyTypeCode>
<fc2:SellingLocationIndicator>Y</fc2:SellingLocationIndicator>
<fc2:Address SeqNum="40">
<fc2:RawCityText>Cityville</fc2:RawCityText>
<fc2:RawCountryCodeText>US</fc2:RawCountryCodeText>
<fc2:RawStateCodeText>VA</fc2:RawStateCodeText>
<fc2:RawStreetAddress1Text>123 Address Way</fc2:RawStreetAddress1Text>
<fc2:RawZIPCode>547898856</fc2:RawZIPCode>
</fc2:Address>
<fc2:PartyIdentification SeqNum="41">
<fc2:PartyIdentificationNumberText>445564615654</fc2:PartyIdentificationNumberText>
<fc2:PartyIdentificationTypeCode>14</fc2:PartyIdentificationTypeCode>
</fc2:PartyIdentification>
</fc2:Party>
`)

	parties["41"] = []byte(`
<fc2:Party SeqNum="75">
<fc2:ActivityPartyTypeCode>41</fc2:ActivityPartyTypeCode>
<fc2:NonUSFinancialInstitutionIndicator>Y</fc2:NonUSFinancialInstitutionIndicator>
<fc2:PartyIdentification SeqNum="76">
<fc2:PartyIdentificationNumberText>458789856</fc2:PartyIdentificationNumberText>
<fc2:PartyIdentificationTypeCode>4</fc2:PartyIdentificationTypeCode>
</fc2:PartyIdentification>
<fc2:Account SeqNum="77">
<fc2:AccountNumberText>1502417873</fc2:AccountNumberText>
<fc2:PartyAccountAssociation SeqNum="78">
<fc2:AccountClosedIndicator>Y</fc2:AccountClosedIndicator>
 <fc2:PartyAccountAssociationTypeCode>5</fc2:PartyAccountAssociationTypeCode>
</fc2:PartyAccountAssociation>
</fc2:Account>
<fc2:Account SeqNum="79">
<fc2:AccountNumberText>5477887896</fc2:AccountNumberText>
<fc2:PartyAccountAssociation SeqNum="80">
<fc2:AccountClosedIndicator>Y</fc2:AccountClosedIndicator>
 <fc2:PartyAccountAssociationTypeCode>5</fc2:PartyAccountAssociationTypeCode>
</fc2:PartyAccountAssociation>
</fc2:Account>
</fc2:Party>
`)

	parties["41"] = []byte(`
<fc2:Party SeqNum="61">
<fc2:ActivityPartyTypeCode>33</fc2:ActivityPartyTypeCode>
<fc2:AdmissionConfessionYesIndicator>Y</fc2:AdmissionConfessionYesIndicator>
<fc2:BothPurchaserSenderPayeeReceiveIndicator>Y</fc2:BothPurchaserSenderPayeeReceiveIndicator>
<fc2:FemaleGenderIndicator>Y</fc2:FemaleGenderIndicator>
<fc2:IndividualBirthDateText>19801025</fc2:IndividualBirthDateText>
<fc2:PartyName SeqNum="62">
<fc2:PartyNameTypeCode>L</fc2:PartyNameTypeCode>
<fc2:RawEntityIndividualLastName>Mann</fc2:RawEntityIndividualLastName>
<fc2:RawIndividualFirstName>Janice</fc2:RawIndividualFirstName>
</fc2:PartyName>
<fc2:PartyName SeqNum="63">
<fc2:PartyNameTypeCode>AKA</fc2:PartyNameTypeCode>
<fc2:RawPartyFullName>Janda</fc2:RawPartyFullName>
</fc2:PartyName>
<fc2:PartyName SeqNum="64">
<fc2:PartyNameTypeCode>AKA</fc2:PartyNameTypeCode>
<fc2:RawPartyFullName>Jan</fc2:RawPartyFullName>
</fc2:PartyName>
<fc2:Address SeqNum="65">
<fc2:CityUnknownIndicator>Y</fc2:CityUnknownIndicator>
<fc2:CountryCodeUnknownIndicator>Y</fc2:CountryCodeUnknownIndicator>
<fc2:StateCodeUnknownIndicator>Y</fc2:StateCodeUnknownIndicator>
<fc2:StreetAddressUnknownIndicator>Y</fc2:StreetAddressUnknownIndicator>
<fc2:ZIPCodeUnknownIndicator>Y</fc2:ZIPCodeUnknownIndicator>
</fc2:Address>
<fc2:PhoneNumber SeqNum="66">
<fc2:PhoneNumberText>6194760276</fc2:PhoneNumberText>
<fc2:PhoneNumberTypeCode>R</fc2:PhoneNumberTypeCode>
</fc2:PhoneNumber>
<fc2:PhoneNumber SeqNum="67">
<fc2:PhoneNumberExtensionText>5584</fc2:PhoneNumberExtensionText>
<fc2:PhoneNumberText>6589784589</fc2:PhoneNumberText>
<fc2:PhoneNumberTypeCode>W</fc2:PhoneNumberTypeCode>
</fc2:PhoneNumber>
<fc2:PartyIdentification SeqNum="68">
<fc2:IdentificationPresentUnknownIndicator>Y</fc2:IdentificationPresentUnknownIndicator>
<fc2:OtherIssuerCountryText>US</fc2:OtherIssuerCountryText>
<fc2:OtherIssuerStateText>CA</fc2:OtherIssuerStateText>
<fc2:OtherPartyIdentificationTypeText>Student ID</fc2:OtherPartyIdentificationTypeText>
<fc2:PartyIdentificationNumberText>660623559</fc2:PartyIdentificationNumberText>
<fc2:PartyIdentificationTypeCode>999</fc2:PartyIdentificationTypeCode>
</fc2:PartyIdentification>
<fc2:PartyIdentification SeqNum="69">
<fc2:TINUnknownIndicator>Y</fc2:TINUnknownIndicator>
</fc2:PartyIdentification>
<fc2:PartyOccupationBusiness SeqNum="70">
<fc2:NAICSCode>7548</fc2:NAICSCode>
<fc2:OccupationBusinessText>Mechanic</fc2:OccupationBusinessText>
</fc2:PartyOccupationBusiness>
<fc2:ElectronicAddress SeqNum="71">
<fc2:ElectronicAddressText>jan@gmail.com</fc2:ElectronicAddressText>
<fc2:ElectronicAddressTypeCode>E</fc2:ElectronicAddressTypeCode>
</fc2:ElectronicAddress>
<fc2:ElectronicAddress SeqNum="72">
<fc2:ElectronicAddressText>janda.org</fc2:ElectronicAddressText>
<fc2:ElectronicAddressTypeCode>U</fc2:ElectronicAddressTypeCode>
</fc2:ElectronicAddress>
<fc2:PartyAssociation SeqNum="73">
<fc2:NoRelationshipToInstitutionIndicator>Y</fc2:NoRelationshipToInstitutionIndicator>
<fc2:SubjectRelationshipFinancialInstitutionTINText>458789856</fc2:SubjectRelationshipFinancialInsti
tutionTINText>
</fc2:PartyAssociation>
<fc2:PartyAccountAssociation SeqNum="74">
<fc2:PartyAccountAssociationTypeCode>7</fc2:PartyAccountAssociationTypeCode>
<fc2:Party SeqNum="75">
<fc2:ActivityPartyTypeCode>41</fc2:ActivityPartyTypeCode>
</fc2:Party>
</fc2:PartyAccountAssociation>
</fc2:Party>
`)

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
		require.Equal(t, len(party.PartyIdentification), 2)

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

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, int64(9))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("PBSA1234"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("28"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		fmt.Println(err)
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
	})
}
