// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package exempt_designation

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/moov-io/fincen"

	"github.com/stretchr/testify/require"
)

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
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties["37"] = []byte(`<Party SeqNum="10">
	<ActivityPartyTypeCode>37</ActivityPartyTypeCode>
	<PartyName SeqNum="11">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Transmitter contact legal name</RawPartyFullName>
	</PartyName>
</Party>`)

	parties["11"] = []byte(`<Party SeqNum="10">
	<ActivityPartyTypeCode>11</ActivityPartyTypeCode>
	<PartyAsEntityOrganizationIndicator>Y</PartyAsEntityOrganizationIndicator>
	<PartyName SeqNum="24">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Johnson Wedding Company</RawEntityIndividualLastName>
	</PartyName>
	<PartyName SeqNum="25">
		<PartyNameTypeCode>DBA</PartyNameTypeCode>
		<RawPartyFullName>JWC</RawPartyFullName>
	</PartyName>
	<Address SeqNum="26">
		<RawCityText>Rockcreek</RawCityText>
		<RawStateCodeText>NY</RawStateCodeText>
		<RawStreetAddress1Text>987 Flint Street</RawStreetAddress1Text>
		<RawZIPCode>21121</RawZIPCode>
	</Address>
	<PhoneNumber SeqNum="27">
		<PhoneNumberExtensionText>4412</PhoneNumberExtensionText>
		<PhoneNumberText>3659855214</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="28">
		<PartyIdentificationNumberText>212121211</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="29">
		<NAICSCode>554</NAICSCode>
		<OccupationBusinessText>Wedding Planner</OccupationBusinessText>
	</PartyOccupationBusiness>
	<ElectronicAddress SeqNum="30">
		<ElectronicAddressText>johnson@wedding.org</ElectronicAddressText>
	</ElectronicAddress>
</Party>`)

	parties["45"] = []byte(`<Party SeqNum="11">
	<ActivityPartyTypeCode>45</ActivityPartyTypeCode>
	<PrimaryRegulatorTypeCode>1</PrimaryRegulatorTypeCode>
	<PartyName SeqNum="31">
		<RawPartyFullName>Bank of the West</RawPartyFullName>
	</PartyName>
	<Address SeqNum="32">
		<RawCityText>Cityville</RawCityText>
		<RawStateCodeText>AZ</RawStateCodeText>
		<RawStreetAddress1Text>555 Rolling Creek Way</RawStreetAddress1Text>
		<RawZIPCode>554478985</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="33">
		<PartyIdentificationNumberText>318181818</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="34">
		<PartyIdentificationNumberText>445581258</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>14</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties["12"] = []byte(`<Party SeqNum="11">
	<ActivityPartyTypeCode>12</ActivityPartyTypeCode>
	<PrimaryRegulatorTypeCode>1</PrimaryRegulatorTypeCode>
	<PartyName SeqNum="31">
		<RawPartyFullName>West Down Station Bank</RawPartyFullName>
	</PartyName>
	<Address SeqNum="32">
		<RawCityText>Cityville</RawCityText>
		<RawStateCodeText>AZ</RawStateCodeText>
		<RawStreetAddress1Text>19 Rolling Meadows Avenue</RawStreetAddress1Text>
		<RawZIPCode>44578</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="33">
		<PartyIdentificationNumberText>665523259</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="34">
		<PartyIdentificationNumberText>332211454</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>14</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties["3"] = []byte(`<Party SeqNum="88">
	<ActivityPartyTypeCode>3</ActivityPartyTypeCode>
	<PartyName SeqNum="89">
		<RawIndividualTitleText>Authorizing person title</RawIndividualTitleText>
		<RawPartyFullName>Authorizing person name</RawPartyFullName>
	</PartyName>
</Party>`)

	return parties
}

func TestParty(t *testing.T) {
	t.Run("Transmitter", func(t *testing.T) {
		sample := mocParties()["35"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(4))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("35"))
		require.Equal(t, len(party.PartyName), 1)
		require.NotNil(t, party.Address)
		require.NotNil(t, party.PhoneNumber)
		require.Equal(t, len(party.PartyIdentification), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(5))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter legal name"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(6))
		require.Equal(t, address.RawCityText, fincen.RestrictString50("Transmitter city"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, address.RawStreetAddress1Text, fincen.RestrictString100("Transmitter street address"))
		require.Equal(t, address.RawZIPCode, fincen.RestrictString9("22113"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(7))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("7217894455"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(8))
		require.Equal(t, identification.PartyIdentificationNumberText, fincen.RestrictString25("458985215"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Transmitter Contact", func(t *testing.T) {
		sample := mocParties()["37"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(10))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("37"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(11))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter contact legal name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Exempt Party", func(t *testing.T) {
		sample := mocParties()["11"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(10))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("11"))
		require.Equal(t, len(party.PartyName), 2)
		require.NotNil(t, party.Address)
		require.NotNil(t, party.PhoneNumber)
		require.NotNil(t, party.ElectronicAddress)
		require.Equal(t, len(party.PartyIdentification), 1)
		require.Equal(t, *party.PartyAsEntityOrganizationIndicator, fincen.ValidateIndicatorType("Y"))

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(24))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Johnson Wedding Company"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(25))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("DBA"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("JWC"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(26))
		require.Equal(t, address.RawCityText, fincen.RestrictString50("Rockcreek"))
		require.Equal(t, address.RawStateCodeText, fincen.RestrictString3("NY"))
		require.Equal(t, address.RawStreetAddress1Text, fincen.RestrictString100("987 Flint Street"))
		require.Equal(t, address.RawZIPCode, fincen.RestrictString9("21121"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(27))
		require.Equal(t, *number.PhoneNumberExtensionText, fincen.RestrictString6("4412"))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("3659855214"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(28))
		require.Equal(t, identification.PartyIdentificationNumberText, fincen.RestrictString25("212121211"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		occ := party.PartyOccupationBusiness
		require.Equal(t, occ.SeqNum, fincen.SeqNumber(29))
		require.Equal(t, *occ.NAICSCode, fincen.RestrictString6("554"))
		require.Equal(t, *occ.OccupationBusinessText, fincen.RestrictString30("Wedding Planner"))

		eAddress := party.ElectronicAddress
		require.Equal(t, eAddress.SeqNum, fincen.SeqNumber(30))
		require.Equal(t, eAddress.ElectronicAddressText, fincen.RestrictString100("johnson@wedding.org"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Exempt Filer Bank", func(t *testing.T) {
		sample := mocParties()["45"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(11))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("45"))
		require.Equal(t, *party.PrimaryRegulatorTypeCode, ValidateFederalRegulatorCodeType("1"))
		require.Equal(t, len(party.PartyName), 1)
		require.NotNil(t, party.Address)
		require.Equal(t, len(party.PartyIdentification), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(31))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Bank of the West"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(32))
		require.Equal(t, address.RawCityText, fincen.RestrictString50("Cityville"))
		require.Equal(t, address.RawStateCodeText, fincen.RestrictString3("AZ"))
		require.Equal(t, address.RawStreetAddress1Text, fincen.RestrictString100("555 Rolling Creek Way"))
		require.Equal(t, address.RawZIPCode, fincen.RestrictString9("554478985"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(33))
		require.Equal(t, identification.PartyIdentificationNumberText, fincen.RestrictString25("318181818"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(34))
		require.Equal(t, identification.PartyIdentificationNumberText, fincen.RestrictString25("445581258"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("14"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Exempt Affiliated Bank", func(t *testing.T) {
		sample := mocParties()["12"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(11))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("12"))
		require.Equal(t, *party.PrimaryRegulatorTypeCode, ValidateFederalRegulatorCodeType("1"))
		require.Equal(t, len(party.PartyName), 1)
		require.NotNil(t, party.Address)
		require.Equal(t, len(party.PartyIdentification), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(31))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("West Down Station Bank"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(32))
		require.Equal(t, address.RawCityText, fincen.RestrictString50("Cityville"))
		require.Equal(t, address.RawStateCodeText, fincen.RestrictString3("AZ"))
		require.Equal(t, address.RawStreetAddress1Text, fincen.RestrictString100("19 Rolling Meadows Avenue"))
		require.Equal(t, address.RawZIPCode, fincen.RestrictString9("44578"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(33))
		require.Equal(t, identification.PartyIdentificationNumberText, fincen.RestrictString25("665523259"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(34))
		require.Equal(t, identification.PartyIdentificationNumberText, fincen.RestrictString25("332211454"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("14"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Authorized Official", func(t *testing.T) {
		sample := mocParties()["3"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)

		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(88))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("3"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(89))
		require.Equal(t, *name.RawIndividualTitleText, fincen.RestrictString35("Authorizing person title"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Authorizing person name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})
}
