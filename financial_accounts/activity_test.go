// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package financial_accounts

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/moov-io/fincen"

	"github.com/stretchr/testify/require"
)

func mocParties() map[string][]byte {
	parties := make(map[string][]byte)

	parties["35"] = []byte(`<Party SeqNum="3">
	<ActivityPartyTypeCode>35</ActivityPartyTypeCode>
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
	<PartyName SeqNum="10">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Transmitter Contact Name</RawPartyFullName>
	</PartyName>
</Party>`)

	parties["15"] = []byte(`<Party SeqNum="19">
	<ActivityPartyTypeCode>15</ActivityPartyTypeCode>
	<FilerFinancialInterest25ForeignAccountIndicator>N</FilerFinancialInterest25ForeignAccountIndicator>
	<FilerTypeIndividualIndicator>Y</FilerTypeIndividualIndicator>
	<IndividualBirthDateText>19700515</IndividualBirthDateText>
	<SignatureAuthoritiesIndicator>N</SignatureAuthoritiesIndicator>
	<PartyName SeqNum="20">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Johnson</RawEntityIndividualLastName>
		<RawIndividualFirstName>John</RawIndividualFirstName>
		<RawIndividualMiddleName>Martin</RawIndividualMiddleName>
		<RawIndividualNameSuffixText>Sr.</RawIndividualNameSuffixText>
	</PartyName>
	<Address SeqNum="21">
		<RawCityText>Detroit</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>MI</RawStateCodeText>
		<RawStreetAddress1Text>555 Rock Avenue</RawStreetAddress1Text>
		<RawZIPCode>48127</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="22">
		<PartyIdentificationNumberText>115478895</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>1</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="23">
		<OtherIssuerCountryText>MX</OtherIssuerCountryText>
		<PartyIdentificationNumberText>55881266698547</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>9</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties["57"] = []byte(`<Party SeqNum="27">
	<ActivityPartyTypeCode>57</ActivityPartyTypeCode>
	<SelfEmployedIndicator>Y</SelfEmployedIndicator>
	<PartyName SeqNum="28">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Jordan</RawEntityIndividualLastName>
		<RawIndividualFirstName>Kate</RawIndividualFirstName>
		<RawIndividualMiddleName>P</RawIndividualMiddleName>
	</PartyName>
	<Address SeqNum="29">
		<RawCityText>Rockville</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>MD</RawStateCodeText>
		<RawStreetAddress1Text>897 Falling Tree Way</RawStreetAddress1Text>
		<RawZIPCode>20899</RawZIPCode>
	</Address>
	<PhoneNumber SeqNum="30">
		<PhoneNumberExtensionText>54</PhoneNumberExtensionText>
		<PhoneNumberText>4875896521</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="31">
		<PartyIdentificationNumberText>558898745</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>31</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties["56"] = []byte(`<Party SeqNum="98">
	<ActivityPartyTypeCode>56</ActivityPartyTypeCode>
	<PartyName SeqNum="99">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Joe’s Preparing Firm</RawPartyFullName>
	</PartyName>
	<PartyIdentification SeqNum="100">
		<PartyIdentificationNumberText>123658984</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	parties["41"] = []byte(`<Party SeqNum="87">
	<ActivityPartyTypeCode>41</ActivityPartyTypeCode>
	<PartyName SeqNum="88">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawPartyFullName>Chester Bank and Trust</RawPartyFullName>
	</PartyName>
	<Address SeqNum="89">
		<RawCityText>Chester</RawCityText>
		<RawCountryCodeText>CA</RawCountryCodeText>
		<RawStateCodeText>NS</RawStateCodeText>
		<RawStreetAddress1Text>14 Middlespring Road</RawStreetAddress1Text>
		<RawZIPCode>B0J1J0</RawZIPCode>
	</Address>
</Party>`)

	parties["42"] = []byte(`<Party SeqNum="90">
	<ActivityPartyTypeCode>42</ActivityPartyTypeCode>
	<PartyName SeqNum="91">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Entity Name</RawEntityIndividualLastName>
	</PartyName>
	<Address SeqNum="92">
		<RawCityText>Vienna</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<RawStreetAddress1Text>987 Speed Way</RawStreetAddress1Text>
		<RawZIPCode>66589</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="93">
		<PartyIdentificationNumberText>887458998</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
</Party>`)

	return parties
}

func TestParty(t *testing.T) {
	t.Run("Transmitter Party Requirements", func(t *testing.T) {
		sample := mocParties()["35"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(3))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("35"))
		require.Equal(t, len(party.PartyIdentification), 2)

		name := party.PartyName
		require.Equal(t, name.SeqNum, fincen.SeqNumber(4))
		require.Equal(t, name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
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
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidateActivityPartyIdentificationCodeType("4"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(8))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("PTCC1234"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidateActivityPartyIdentificationCodeType("28"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Transmitter Contact Party Requirements", func(t *testing.T) {
		sample := mocParties()["37"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(9))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("37"))

		name := party.PartyName
		require.Equal(t, name.SeqNum, fincen.SeqNumber(10))
		require.Equal(t, name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter Contact Name"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Foreign Account Filer Party Requirements", func(t *testing.T) {
		sample := mocParties()["15"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(19))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("15"))

		name := party.PartyName
		require.Equal(t, name.SeqNum, fincen.SeqNumber(20))
		require.Equal(t, name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Johnson"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("John"))
		require.Equal(t, *name.RawIndividualMiddleName, fincen.RestrictString35("Martin"))
		require.Equal(t, *name.RawIndividualNameSuffixText, fincen.RestrictString35("Sr."))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(21))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Detroit"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("MI"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("555 Rock Avenue"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("48127"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(22))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("115478895"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidateActivityPartyIdentificationCodeType("1"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(23))
		require.Equal(t, *identification.OtherIssuerCountryText, fincen.RestrictString2("MX"))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("55881266698547"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidateActivityPartyIdentificationCodeType("9"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Third party preparer", func(t *testing.T) {
		sample := mocParties()["57"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(27))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("57"))

		name := party.PartyName
		require.Equal(t, name.SeqNum, fincen.SeqNumber(28))
		require.Equal(t, name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Jordan"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("Kate"))
		require.Equal(t, *name.RawIndividualMiddleName, fincen.RestrictString35("P"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(29))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Rockville"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("MD"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("897 Falling Tree Way"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("20899"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(30))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("4875896521"))
		require.Equal(t, *number.PhoneNumberExtensionText, fincen.RestrictString6("54"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(31))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("558898745"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidateActivityPartyIdentificationCodeType("31"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Third Party Preparer Firm Party Requirements", func(t *testing.T) {
		sample := mocParties()["56"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(98))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("56"))

		name := party.PartyName
		require.Equal(t, name.SeqNum, fincen.SeqNumber(99))
		require.Equal(t, name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Joe’s Preparing Firm"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(100))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("123658984"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidateActivityPartyIdentificationCodeType("2"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Account Party (41)", func(t *testing.T) {
		sample := mocParties()["41"]
		party := AccountPartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(87))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateAccountPartyCodeType("41"))

		name := party.PartyName
		require.Equal(t, name.SeqNum, fincen.SeqNumber(88))
		require.Equal(t, name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Chester Bank and Trust"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(89))
		require.Equal(t, address.RawCityText, fincen.RestrictString50("Chester"))
		require.Equal(t, address.RawCountryCodeText, fincen.RestrictString2("CA"))
		require.Equal(t, address.RawStateCodeText, fincen.RestrictString3("NS"))
		require.Equal(t, address.RawStreetAddress1Text, fincen.RestrictString100("14 Middlespring Road"))
		require.Equal(t, address.RawZIPCode, fincen.RestrictString9("B0J1J0"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Account Party (42)", func(t *testing.T) {
		sample := mocParties()["42"]
		party := AccountPartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(90))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateAccountPartyCodeType("42"))

		name := party.PartyName
		require.Equal(t, name.SeqNum, fincen.SeqNumber(91))
		require.Equal(t, name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Entity Name"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(92))
		require.Equal(t, address.RawCityText, fincen.RestrictString50("Vienna"))
		require.Equal(t, address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, address.RawStreetAddress1Text, fincen.RestrictString100("987 Speed Way"))
		require.Equal(t, address.RawZIPCode, fincen.RestrictString9("66589"))

		identification := party.PartyIdentification
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(93))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("887458998"))
		require.Equal(t, *identification.PartyIdentificationTypeCode, ValidateAccountPartyIdentificationCodeType("2"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})
}
