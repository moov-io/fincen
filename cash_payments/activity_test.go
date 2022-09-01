package cash_payments

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
		<RawCityText>McLean</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>VA</RawStateCodeText>
		<RawStreetAddress1Text>123 Street</RawStreetAddress1Text>
		<RawZIPCode>55478</RawZIPCode>
	</Address>
	<PhoneNumber SeqNum="6">
		<PhoneNumberText>7031231234</PhoneNumberText>
	</PhoneNumber>
	<PartyIdentification SeqNum="7">
		<PartyIdentificationNumberText>451898562</PartyIdentificationNumberText>
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

	parties["16"] = []byte(`<Party SeqNum="95">
	<ActivityPartyTypeCode>16</ActivityPartyTypeCode>
	<IndividualBirthDateText>19750101</IndividualBirthDateText>
	<PartyName SeqNum="96">
		<RawEntityIndividualLastName>Williams</RawEntityIndividualLastName>
		<RawIndividualFirstName>John</RawIndividualFirstName>
		<RawIndividualMiddleName>Holmes</RawIndividualMiddleName>
	</PartyName>
	<Address SeqNum="97">
		<RawCityText>Rockville</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>MD</RawStateCodeText>
		<RawStreetAddress1Text>456 Street</RawStreetAddress1Text>
		<RawZIPCode>254789841</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="98">
		<PartyIdentificationNumberText>125489563</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>1</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="99">
		<OtherIssuerStateText>FL</OtherIssuerStateText>
		<PartyIdentificationNumberText>G459851234</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>5</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="100">
		<OccupationBusinessText>Business description</OccupationBusinessText>
	</PartyOccupationBusiness>
</Party>`)

	parties["23"] = []byte(`<Party SeqNum="253">
	<ActivityPartyTypeCode>23</ActivityPartyTypeCode>
	<PartyTypeCode>O</PartyTypeCode>
	<PartyName SeqNum="254">
		<PartyNameTypeCode>L</PartyNameTypeCode>
		<RawEntityIndividualLastName>Bob&#39;s Business</RawEntityIndividualLastName>
	</PartyName>
	<PartyName SeqNum="255">
		<PartyNameTypeCode>DBA</PartyNameTypeCode>
		<RawPartyFullName>Bobby&#39;s Business</RawPartyFullName>
	</PartyName>
	<Address SeqNum="256">
		<RawCityText>Silver Spring</RawCityText>
		<RawCountryCodeText>US</RawCountryCodeText>
		<RawStateCodeText>MD</RawStateCodeText>
		<RawStreetAddress1Text>789 Street</RawStreetAddress1Text>
		<RawZIPCode>54789</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="257">
		<PartyIdentificationNumberText>565898523</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>1</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="258">
		<PartyIdentificationNumberText>554123985</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="259">
		<OtherIssuerStateText>CA</OtherIssuerStateText>
		<PartyIdentificationNumberText>5326547</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>999</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="260">
		<OccupationBusinessText>Business description</OccupationBusinessText>
	</PartyOccupationBusiness>
</Party>`)

	parties["4"] = []byte(`<Party SeqNum="36">
	<ActivityPartyTypeCode>4</ActivityPartyTypeCode>
	<PartyName SeqNum="37">
		<RawPartyFullName>Name of business that received cash</RawPartyFullName>
	</PartyName>
	<Address SeqNum="38">
		<RawCityText>Richmond</RawCityText>
		<RawStateCodeText>VA</RawStateCodeText>
		<RawStreetAddress1Text>159 Street</RawStreetAddress1Text>
		<RawZIPCode>66589</RawZIPCode>
	</Address>
	<PartyIdentification SeqNum="39">
		<PartyIdentificationNumberText>565898523</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>2</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyIdentification SeqNum="40">
		<PartyIdentificationNumberText>554123985</PartyIdentificationNumberText>
		<PartyIdentificationTypeCode>1</PartyIdentificationTypeCode>
	</PartyIdentification>
	<PartyOccupationBusiness SeqNum="41">
		<OccupationBusinessText>Business description</OccupationBusinessText>
	</PartyOccupationBusiness>
</Party>`)

	parties["3"] = []byte(`<Party SeqNum="88">
	<ActivityPartyTypeCode>3</ActivityPartyTypeCode>
	<PartyName SeqNum="89">
		<RawIndividualTitleText>Authorizing person title</RawIndividualTitleText>
	</PartyName>
</Party>`)

	parties["8"] = []byte(`<Party SeqNum="174">
	<ActivityPartyTypeCode>8</ActivityPartyTypeCode>
	<PartyName SeqNum="175">
		<RawPartyFullName>Contact name</RawPartyFullName>
	</PartyName>
	<PhoneNumber SeqNum="176">
		<PhoneNumberText>4589774512</PhoneNumberText>
	</PhoneNumber>
</Party>`)

	return parties
}

func TestParty(t *testing.T) {
	t.Run("Transmitter", func(t *testing.T) {
		sample := mocParties()["35"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(3))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("35"))
		require.Equal(t, len(party.PartyName), 1)
		require.Equal(t, len(party.PartyIdentification), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(4))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Transmitter Legal Name"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(5))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("McLean"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("123 Street"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("55478"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(6))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("7031231234"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(7))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("451898562"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Transmitter contact", func(t *testing.T) {
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

	t.Run("Individual from whom the cash was received", func(t *testing.T) {
		sample := mocParties()["16"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(95))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("16"))
		require.Equal(t, *party.IndividualBirthDateText, fincen.DateYYYYMMDDOrBlankType("19750101"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(96))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Williams"))
		require.Equal(t, *name.RawIndividualFirstName, fincen.RestrictString35("John"))
		require.Equal(t, *name.RawIndividualMiddleName, fincen.RestrictString35("Holmes"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(97))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Rockville"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("MD"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("456 Street"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("254789841"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(98))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("125489563"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("1"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(99))
		require.Equal(t, *identification.OtherIssuerStateText, fincen.RestrictString3("FL"))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("G459851234"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("5"))

		business := party.PartyOccupationBusiness
		require.Equal(t, business.SeqNum, fincen.SeqNumber(100))
		require.Equal(t, business.OccupationBusinessText, fincen.RestrictString30("Business description"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Person on whose behalf transaction conducted", func(t *testing.T) {
		sample := mocParties()["23"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(253))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("23"))
		require.Equal(t, *party.PartyTypeCode, fincen.ValidatePartyTypeCode("O"))
		require.Equal(t, len(party.PartyName), 2)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(254))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("L"))
		require.Equal(t, *name.RawEntityIndividualLastName, fincen.RestrictString150("Bob's Business"))

		name = party.PartyName[1]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(255))
		require.Equal(t, *name.PartyNameTypeCode, ValidatePartyNameCodeType("DBA"))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Bobby's Business"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(256))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Silver Spring"))
		require.Equal(t, *address.RawCountryCodeText, fincen.RestrictString2("US"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("MD"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("789 Street"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("54789"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(257))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("565898523"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("1"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(258))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("554123985"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		identification = party.PartyIdentification[2]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(259))
		require.Equal(t, *identification.OtherIssuerStateText, fincen.RestrictString3("CA"))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("5326547"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("999"))

		business := party.PartyOccupationBusiness
		require.Equal(t, business.SeqNum, fincen.SeqNumber(260))
		require.Equal(t, business.OccupationBusinessText, fincen.RestrictString30("Business description"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Business that received cash", func(t *testing.T) {
		sample := mocParties()["4"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(36))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("4"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(37))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Name of business that received cash"))

		address := party.Address
		require.Equal(t, address.SeqNum, fincen.SeqNumber(38))
		require.Equal(t, *address.RawCityText, fincen.RestrictString50("Richmond"))
		require.Equal(t, *address.RawStateCodeText, fincen.RestrictString3("VA"))
		require.Equal(t, *address.RawStreetAddress1Text, fincen.RestrictString100("159 Street"))
		require.Equal(t, *address.RawZIPCode, fincen.RestrictString9("66589"))

		identification := party.PartyIdentification[0]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(39))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("565898523"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("2"))

		identification = party.PartyIdentification[1]
		require.Equal(t, identification.SeqNum, fincen.SeqNumber(40))
		require.Equal(t, *identification.PartyIdentificationNumberText, fincen.RestrictString25("554123985"))
		require.Equal(t, identification.PartyIdentificationTypeCode, ValidatePartyIdentificationCodeType("1"))

		business := party.PartyOccupationBusiness
		require.Equal(t, business.SeqNum, fincen.SeqNumber(41))
		require.Equal(t, business.OccupationBusinessText, fincen.RestrictString30("Business description"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Authorized official", func(t *testing.T) {
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

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})

	t.Run("Contact for assistance", func(t *testing.T) {
		sample := mocParties()["8"]
		party := PartyType{}
		err := xml.Unmarshal(sample, &party)
		require.NoError(t, err)

		require.Equal(t, party.SeqNum, fincen.SeqNumber(174))
		require.Equal(t, party.ActivityPartyTypeCode, ValidateActivityPartyCodeType("8"))
		require.Equal(t, len(party.PartyName), 1)

		name := party.PartyName[0]
		require.Equal(t, name.SeqNum, fincen.SeqNumber(175))
		require.Equal(t, *name.RawPartyFullName, fincen.RestrictString150("Contact name"))

		number := party.PhoneNumber
		require.Equal(t, number.SeqNum, fincen.SeqNumber(176))
		require.Equal(t, *number.PhoneNumberText, fincen.RestrictString16("4589774512"))

		buf, err := xml.MarshalIndent(&party, "", "\t")
		require.NoError(t, err)
		require.Equal(t, reflect.DeepEqual(sample, buf), true)

		err = party.Validate()
		require.NoError(t, err)
	})
}
