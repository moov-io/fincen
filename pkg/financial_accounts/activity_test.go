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

func TestElements(t *testing.T) {

	t.Run("ActivityType", func(t *testing.T) {
		activity := NewActivity()

		require.Equal(t, "FBARX", activity.FormTypeCode())
		require.Equal(t, "The Party has invalid min & max range", activity.Validate().Error())

		activity.ActivityAssociation = &ActivityAssociationType{}
		activity.ForeignAccountActivity = &ForeignAccountActivityType{}
		for i := 0; i < 3; i++ {
			activity.Party = append(activity.Party, &PartyType{})
		}
		require.Equal(t, "The Party(type 35) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "35"})
		require.Equal(t, "The Party(type 37) is a required field", activity.Validate().Error())

		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "37"})
		require.Equal(t, "The Party(type 15) is a required field", activity.Validate().Error())

		activity.Party = nil
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "35"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "37"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "15"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "37"})
		require.Equal(t, "The Party(type 37) has invalid min & max range", activity.Validate().Error())

		activity.Party = nil
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "35"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "37"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "15"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "15"})
		require.Equal(t, "The Party(type 15) has invalid min & max range", activity.Validate().Error())

		activity.Party = nil
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "35"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "37"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "15"})
		activity.Party = append(activity.Party, &PartyType{ActivityPartyTypeCode: "35"})
		require.Equal(t, "The Party(type 35) has invalid min & max range", activity.Validate().Error())
	})

	t.Run("PartyType", func(t *testing.T) {
		var sample PartyType
		indicatorYN := ValidateIndicatorYNType("Y")
		indicator := fincen.ValidateIndicatorNullType("Y")
		sample.PartyName = &PartyNameType{}

		require.Equal(t, "The ActivityPartyCode has invalid value", sample.Validate().Error())

		sample = PartyType{ActivityPartyTypeCode: "15"}

		require.Equal(t, "The FilerFinancialInterest25ForeignAccountIndicator is a required field", sample.Validate().Error())
		sample.FilerFinancialInterest25ForeignAccountIndicator = &indicatorYN
		require.Equal(t, "The Address is a required field", sample.Validate().Error())

		sample = PartyType{ActivityPartyTypeCode: "35"}

		require.Equal(t, "The Address is a required field", sample.Validate().Error())
		sample.Address = &AddressType{}
		require.Equal(t, "The PhoneNumber is a required field", sample.Validate().Error())

		sample = PartyType{ActivityPartyTypeCode: "INVALID"}

		sample.FilerTypeConsolidatedIndicator = &indicator
		require.Equal(t, "The FilerTypeConsolidatedIndicator should be omitted", sample.Validate().Error())
		sample.FilerTypeConsolidatedIndicator = nil
		sample.FilerTypeCorporationIndicator = &indicator
		require.Equal(t, "The FilerTypeCorporationIndicator should be omitted", sample.Validate().Error())
		sample.FilerTypeCorporationIndicator = nil
		sample.FilerTypeFiduciaryOtherIndicator = &indicator
		require.Equal(t, "The FilerTypeFiduciaryOtherIndicator should be omitted", sample.Validate().Error())
		sample.FilerTypeFiduciaryOtherIndicator = nil
		sample.FilerTypeIndividualIndicator = &indicator
		require.Equal(t, "The FilerTypeIndividualIndicator should be omitted", sample.Validate().Error())
		sample.FilerTypeIndividualIndicator = nil
		sample.FilerTypePartnershipIndicator = &indicator
		require.Equal(t, "The FilerTypePartnershipIndicator should be omitted", sample.Validate().Error())
		sample.FilerTypePartnershipIndicator = nil
		sample.SelfEmployedIndicator = &indicator
		require.Equal(t, "The SelfEmployedIndicator should be omitted", sample.Validate().Error())
		sample.SelfEmployedIndicator = nil
		sample.SignatureAuthoritiesIndicator = &indicatorYN
		require.Equal(t, "The SignatureAuthoritiesIndicator should be omitted", sample.Validate().Error())
		v50 := fincen.RestrictString50("")
		sample.FilerTypeOtherText = &v50
		sample.SignatureAuthoritiesIndicator = nil
		require.Equal(t, "The FilerTypeOtherText has invalid value", sample.Validate().Error())
		sample.FilerTypeOtherText = nil
		dd := fincen.DateYYYYMMDDOrBlankType("")
		sample.IndividualBirthDateText = &dd
		require.Equal(t, "The IndividualBirthDateText has invalid value", sample.Validate().Error())
		sample.IndividualBirthDateText = nil
		for i := 0; i < 3; i++ {
			sample.PartyIdentification = append(sample.PartyIdentification, &PartyIdentificationType{})
		}
		require.Equal(t, "The PartyIdentification has invalid min & max range", sample.Validate().Error())
	})

	t.Run("AccountType", func(t *testing.T) {
		var sample AccountType

		require.Equal(t, "The Party is a required field", sample.Validate().Error())
	})

	t.Run("AddressType", func(t *testing.T) {
		var sample AddressType

		require.NoError(t, sample.Validate())
		require.NoError(t, sample.Validate("INVALID"))

		require.Equal(t, "The RawCityText is a required field", sample.Validate("35").Error())
		v50 := fincen.RestrictString50("")
		sample.RawCityText = &v50
		require.Equal(t, "The RawCountryCodeText is a required field", sample.Validate("35").Error())
		v2 := fincen.RestrictString2("")
		sample.RawCountryCodeText = &v2
		require.Equal(t, "The RawStateCodeText is a required field", sample.Validate("35").Error())
		v3 := fincen.RestrictString3("")
		sample.RawStateCodeText = &v3
		require.Equal(t, "The RawStreetAddress1Text is a required field", sample.Validate("35").Error())
		v100 := fincen.RestrictString100("")
		sample.RawStreetAddress1Text = &v100
		require.Equal(t, "The RawZIPCode is a required field", sample.Validate("35").Error())
	})

	t.Run("AccountPartyType", func(t *testing.T) {
		var sample AccountPartyType
		indicator := fincen.ValidateIndicatorNullType("Y")
		sample.PartyName = &AccountPartyNameType{}
		sample.Address = &AccountAddressType{}

		require.Equal(t, "The AccountPartyCode has invalid value", sample.Validate().Error())

		sample.ActivityPartyTypeCode = "41"
		sample.PartyAsEntityOrganizationIndicator = &indicator
		require.Equal(t, "The PartyAsEntityOrganizationIndicator should be omitted", sample.Validate().Error())
		sample.PartyAsEntityOrganizationIndicator = nil
		sample.PartyIdentification = &AccountPartyIdentificationType{}
		require.Equal(t, "The PartyIdentification should be omitted", sample.Validate().Error())
	})

	t.Run("PartyIdentificationType", func(t *testing.T) {
		var sample PartyIdentificationType

		require.NoError(t, sample.Validate())
		require.NoError(t, sample.Validate("INVALID"))

		require.Equal(t, "The OtherIssuerCountryText is a required field", sample.Validate("15").Error())
		v2 := fincen.RestrictString2("")
		sample.OtherIssuerCountryText = &v2
		require.Equal(t, "The OtherPartyIdentificationTypeText is a required field", sample.Validate("15").Error())
		v50 := fincen.RestrictString50("")
		sample.OtherPartyIdentificationTypeText = &v50
		require.Equal(t, "The PartyIdentificationNumberText is a required field", sample.Validate("15").Error())
		v25 := fincen.RestrictString25("")
		sample.PartyIdentificationNumberText = &v25
		iCode := ValidateActivityPartyIdentificationCodeType("")
		sample.PartyIdentificationTypeCode = &iCode
		require.Equal(t, "The PartyIdentificationTypeCode has invalid value", sample.Validate("15").Error())
		require.Equal(t, "The PartyIdentificationTypeCode has invalid value", sample.Validate("35").Error())
		require.Equal(t, "The PartyIdentificationTypeCode has invalid value", sample.Validate("57").Error())
		require.Equal(t, "The PartyIdentificationTypeCode has invalid value", sample.Validate("56").Error())
	})

	t.Run("PartyNameType", func(t *testing.T) {
		var sample PartyNameType

		require.Equal(t, "The PartyNameCode has invalid value", sample.Validate().Error())
		require.Equal(t, "The PartyNameCode has invalid value", sample.Validate("INVALID").Error())

		v150 := fincen.RestrictString150("")
		sample.RawEntityIndividualLastName = &v150
		require.Equal(t, "The RawEntityIndividualLastName should be omitted", sample.Validate("INVALID").Error())
		sample.RawEntityIndividualLastName = nil
		v35 := fincen.RestrictString35("")
		sample.RawIndividualFirstName = &v35
		require.Equal(t, "The RawIndividualFirstName should be omitted", sample.Validate("INVALID").Error())
		sample.RawIndividualFirstName = nil
		sample.RawIndividualMiddleName = &v35
		require.Equal(t, "The RawIndividualMiddleName should be omitted", sample.Validate("INVALID").Error())
		sample.RawIndividualMiddleName = nil
		sample.RawIndividualNameSuffixText = &v35
		require.Equal(t, "The RawIndividualNameSuffixText should be omitted", sample.Validate("INVALID").Error())
		sample.RawIndividualNameSuffixText = nil
		v20 := fincen.RestrictString20("")
		sample.RawIndividualTitleText = &v20
		require.Equal(t, "The RawIndividualTitleText should be omitted", sample.Validate("INVALID").Error())
		sample.RawIndividualTitleText = nil
		require.Equal(t, "The RawPartyFullName is a required field", sample.Validate("35").Error())
	})

	t.Run("PhoneNumberType", func(t *testing.T) {
		var sample PhoneNumberType

		require.NoError(t, sample.Validate())
		require.NoError(t, sample.Validate("INVALID"))

		v6 := fincen.RestrictString6("")
		sample.PhoneNumberExtensionText = &v6
		require.Equal(t, "The PhoneNumberExtensionText should be omitted", sample.Validate("INVALID").Error())

		sample.PhoneNumberExtensionText = nil
		require.Equal(t, "The PhoneNumberText is a required field", sample.Validate("35").Error())
	})

	t.Run("AccountPartyNameType", func(t *testing.T) {
		var sample AccountPartyNameType

		require.Equal(t, "The PartyNameCode has invalid value", sample.Validate().Error())
		require.Equal(t, "The PartyNameCode has invalid value", sample.Validate("INVALID").Error())

		v150 := fincen.RestrictString150("")
		sample.RawEntityIndividualLastName = &v150
		require.Equal(t, "The RawEntityIndividualLastName should be omitted", sample.Validate("INVALID").Error())
		sample.RawEntityIndividualLastName = nil
		v35 := fincen.RestrictString35("")
		sample.RawIndividualFirstName = &v35
		require.Equal(t, "The RawIndividualFirstName should be omitted", sample.Validate("INVALID").Error())
		sample.RawIndividualFirstName = nil
		sample.RawIndividualMiddleName = &v35
		require.Equal(t, "The RawIndividualMiddleName should be omitted", sample.Validate("INVALID").Error())
		sample.RawIndividualMiddleName = nil
		sample.RawIndividualNameSuffixText = &v35
		require.Equal(t, "The RawIndividualNameSuffixText should be omitted", sample.Validate("INVALID").Error())
		sample.RawIndividualNameSuffixText = nil
		v20 := fincen.RestrictString20("")
		sample.RawIndividualTitleText = &v20
		require.Equal(t, "The RawIndividualTitleText should be omitted", sample.Validate("INVALID").Error())
		sample.RawIndividualTitleText = nil
		sample.RawPartyFullName = &v150
		require.Equal(t, "The RawPartyFullName should be omitted", sample.Validate("INVALID").Error())
	})

}
