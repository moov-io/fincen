// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"

	. "github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
	. "github.com/moov-io/fincen/pkg/cash_payments"
)

func main() {

	party1 := &PartyType{}
	party1.ActivityPartyTypeCode = ValidateActivityPartyCodeType("35")
	party1.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Transmitter Legal Name")),
		},
	}
	party1.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("McLean")),
		RawCountryCodeText:    Ptr(RestrictString2("US")),
		RawStateCodeText:      Ptr(RestrictString3("VA")),
		RawStreetAddress1Text: Ptr(RestrictString100("123 Street")),
		RawZIPCode:            Ptr(RestrictString9("55478")),
	}
	party1.PhoneNumber = &PhoneNumberType{
		PhoneNumberText: Ptr(RestrictString16("7031231234")),
	}
	party1.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("451898562")),
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("2"),
		},
	}
	fmt.Println("Validating party1:", party1.Validate())

	party2 := &PartyType{}
	party2.ActivityPartyTypeCode = ValidateActivityPartyCodeType("37")
	party2.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Transmitter contact legal name")),
		},
	}
	fmt.Println("Validating party2:", party2.Validate())

	party3 := &PartyType{}
	party3.ActivityPartyTypeCode = ValidateActivityPartyCodeType("16")
	party3.IndividualBirthDateText = Ptr(DateYYYYMMDDOrBlankType("19750101"))
	party3.PartyName = []*PartyNameType{
		{
			RawEntityIndividualLastName: Ptr(RestrictString150("Williams")),
			RawIndividualFirstName:      Ptr(RestrictString35("John")),
			RawIndividualMiddleName:     Ptr(RestrictString35("Holmes")),
		},
	}
	party3.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("Rockville")),
		RawCountryCodeText:    Ptr(RestrictString2("US")),
		RawStateCodeText:      Ptr(RestrictString3("MD")),
		RawStreetAddress1Text: Ptr(RestrictString100("456 Street")),
		RawZIPCode:            Ptr(RestrictString9("254789841")),
	}
	party3.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("125489563")),
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("1"),
		},
		{
			OtherIssuerStateText:          Ptr(RestrictString3("FL")),
			PartyIdentificationNumberText: Ptr(RestrictString25("G459851234")),
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("5"),
		},
	}
	party3.PartyOccupationBusiness = &PartyOccupationBusinessType{
		OccupationBusinessText: RestrictString30("Business description"),
	}
	fmt.Println("Validating party3:", party3.Validate())

	party4 := &PartyType{}
	party4.ActivityPartyTypeCode = ValidateActivityPartyCodeType("23")
	party4.PartyTypeCode = Ptr(ValidatePartyTypeCode("O"))
	party4.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode:           Ptr(ValidatePartyNameCodeType("L")),
			RawEntityIndividualLastName: Ptr(RestrictString150("Bob's Business")),
		},
		{
			PartyNameTypeCode:           Ptr(ValidatePartyNameCodeType("DBA")),
			RawEntityIndividualLastName: Ptr(RestrictString150("Bob's Business")),
		},
	}
	party4.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("Silver Spring")),
		RawCountryCodeText:    Ptr(RestrictString2("US")),
		RawStateCodeText:      Ptr(RestrictString3("MD")),
		RawStreetAddress1Text: Ptr(RestrictString100("789 Street")),
		RawZIPCode:            Ptr(RestrictString9("54789")),
	}
	party4.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("565898523")),
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("1"),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("554123985")),
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("2"),
		},
		{
			OtherIssuerStateText:          Ptr(RestrictString3("CS")),
			PartyIdentificationNumberText: Ptr(RestrictString25("5326547")),
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("999"),
		},
	}
	party4.PartyOccupationBusiness = &PartyOccupationBusinessType{
		OccupationBusinessText: RestrictString30("Business description"),
	}
	fmt.Println("Validating party4:", party4.Validate())

	party5 := &PartyType{}
	party5.ActivityPartyTypeCode = ValidateActivityPartyCodeType("3")
	party5.PartyName = []*PartyNameType{
		{
			RawIndividualTitleText: Ptr(RestrictString35("Authorizing person title")),
		},
	}
	fmt.Println("Validating party5:", party5.Validate())

	party6 := &PartyType{}
	party6.ActivityPartyTypeCode = ValidateActivityPartyCodeType("4")
	party6.PartyName = []*PartyNameType{
		{
			RawPartyFullName: Ptr(RestrictString150("Name of business that received cash")),
		},
	}
	party6.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("Richmond")),
		RawStateCodeText:      Ptr(RestrictString3("VA")),
		RawStreetAddress1Text: Ptr(RestrictString100("159 Street")),
		RawZIPCode:            Ptr(RestrictString9("66589")),
	}
	party6.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("565898523")),
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("2"),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("554123985")),
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("1"),
		},
	}
	party6.PartyOccupationBusiness = &PartyOccupationBusinessType{
		OccupationBusinessText: RestrictString30("Business description"),
	}
	fmt.Println("Validating party6:", party6.Validate())

	act := NewActivity()
	act.Party = []*PartyType{party1, party2, party3, party4, party5, party6}
	act.CurrencyTransactionActivity = &CurrencyTransactionActivityType{
		TransactionDateText: DateYYYYMMDDType("19801025"),
		CurrencyTransactionActivityDetail: []*CurrencyTransactionActivityDetailType{
			{},
			{},
		},
	}
	act.FilingDateText = "19801025"
	fmt.Println("Validating activity:", act.Validate())

	r := batch.NewReport(Form8300)
	r.AppendActivity(act)

	// generating attr
	r.GenerateAttrs()
	r.GenerateSeqNumbers()

	fmt.Println("Validating batch report:", r.Validate())

	buf, _ := xml.MarshalIndent(r, "", "  ")

	fmt.Println("\nXML Body:")
	fmt.Println(string(buf))
}
