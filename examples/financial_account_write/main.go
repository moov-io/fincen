// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"

	. "github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
	. "github.com/moov-io/fincen/pkg/financial_accounts"
)

func main() {
	party1 := &PartyType{}
	party1.ActivityPartyTypeCode = ValidateActivityPartyCodeType("35")
	party1.PartyName = &PartyNameType{
		PartyNameTypeCode: ValidatePartyNameCodeType("L"),
		RawPartyFullName:  Ptr(RestrictString150("Transmitter legal name")),
	}
	party1.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("Transmitter city")),
		RawCountryCodeText:    Ptr(RestrictString2("US")),
		RawStateCodeText:      Ptr(RestrictString3("VA")),
		RawStreetAddress1Text: Ptr(RestrictString100("123 Address Road")),
		RawZIPCode:            Ptr(RestrictString9("22102")),
	}
	party1.PhoneNumber = &PhoneNumberType{
		PhoneNumberText: Ptr(RestrictString16("7039991234")),
	}
	party1.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("547898569")),
			PartyIdentificationTypeCode:   Ptr(ValidateActivityPartyIdentificationCodeType("4")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("PTCC1234")),
			PartyIdentificationTypeCode:   Ptr(ValidateActivityPartyIdentificationCodeType("28")),
		},
	}
	fmt.Println("Validating party1:", party1.Validate())

	party2 := &PartyType{}
	party2.ActivityPartyTypeCode = ValidateActivityPartyCodeType("37")
	party2.PartyName = &PartyNameType{
		PartyNameTypeCode: ValidatePartyNameCodeType("L"),
		RawPartyFullName:  Ptr(RestrictString150("Transmitter Contact name")),
	}
	fmt.Println("Validating party2:", party2.Validate())

	party3 := &PartyType{
		ActivityPartyTypeCode:                           ValidateActivityPartyCodeType("15"),
		FilerFinancialInterest25ForeignAccountIndicator: Ptr(ValidateIndicatorYNType("N")),
		FilerTypeIndividualIndicator:                    Ptr(ValidateIndicatorNullType("Y")),
		IndividualBirthDateText:                         Ptr(DateYYYYMMDDOrBlankType("19700515")),
		SignatureAuthoritiesIndicator:                   Ptr(ValidateIndicatorYNType("N")),
	}
	party3.PartyName = &PartyNameType{
		PartyNameTypeCode:           ValidatePartyNameCodeType("L"),
		RawEntityIndividualLastName: Ptr(RestrictString150("Johnson")),
		RawIndividualFirstName:      Ptr(RestrictString35("John")),
		RawIndividualMiddleName:     Ptr(RestrictString35("Martin")),
		RawIndividualNameSuffixText: Ptr(RestrictString35("Sr.")),
	}
	party3.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("Detroit")),
		RawCountryCodeText:    Ptr(RestrictString2("US")),
		RawStateCodeText:      Ptr(RestrictString3("MI")),
		RawStreetAddress1Text: Ptr(RestrictString100("555 Rock Avenue")),
		RawZIPCode:            Ptr(RestrictString9("48127")),
	}
	party3.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("115478895")),
			PartyIdentificationTypeCode:   Ptr(ValidateActivityPartyIdentificationCodeType("1")),
		},
		{
			OtherIssuerCountryText:        Ptr(RestrictString2("MX")),
			PartyIdentificationNumberText: Ptr(RestrictString25("55881266698547")),
			PartyIdentificationTypeCode:   Ptr(ValidateActivityPartyIdentificationCodeType("9")),
		},
	}
	fmt.Println("Validating party3:", party3.Validate())

	act := NewActivity()
	act.Party = []*PartyType{party1, party2, party3}
	act.ForeignAccountActivity = &ForeignAccountActivityType{
		ForeignAccountHeldQuantityText:   Ptr(RestrictString4("31")),
		LateFilingReasonCode:             Ptr(ValidateLateFilingReasonCodeType("999")),
		ReportCalendarYearText:           "2016",
		SignatureAuthoritiesQuantityText: Ptr(RestrictString4("52")),
	}
	act.ApprovalOfficialSignatureDateText = "19801025"
	act.ActivityAssociation = &ActivityAssociationType{
		CorrectsAmendsPriorReportIndicator: "Y",
	}
	fmt.Println("Validating activity:", act.Validate())

	r := batch.NewReport(Report114)
	r.AppendActivity(act)

	// generating attr
	r.GenerateAttrs()
	r.GenerateSeqNumbers()

	fmt.Println("Validating batch report:", r.Validate())

	buf, _ := xml.MarshalIndent(r, "", "  ")

	fmt.Println("\nXML Body:")
	fmt.Println(string(buf))
}
