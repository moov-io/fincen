// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"

	. "github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
	. "github.com/moov-io/fincen/pkg/exempt_designation"
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
		RawCityText:           "Transmitter city",
		RawCountryCodeText:    Ptr(RestrictString2("US")),
		RawStateCodeText:      "VA",
		RawStreetAddress1Text: "Transmitter street address",
		RawZIPCode:            "22113",
	}
	party1.PhoneNumber = &PhoneNumberType{
		PhoneNumberText: Ptr(RestrictString16("7217894455")),
	}
	party1.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: "458985215",
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

	party3 := &PartyType{
		ActivityPartyTypeCode:              ValidateActivityPartyCodeType("11"),
		PartyAsEntityOrganizationIndicator: Ptr(ValidateIndicatorNullType("Y")),
	}
	party3.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode:           Ptr(ValidatePartyNameCodeType("L")),
			RawEntityIndividualLastName: Ptr(RestrictString150("Johnson Wedding Company")),
		},
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("DBA")),
			RawPartyFullName:  Ptr(RestrictString150("JWC")),
		},
	}
	party3.Address = &AddressType{
		RawCityText:           "Rockcreek",
		RawStateCodeText:      "NY",
		RawStreetAddress1Text: "987 Flint Street",
		RawZIPCode:            "21121",
	}
	party3.PhoneNumber = &PhoneNumberType{
		PhoneNumberExtensionText: Ptr(RestrictString6("4412")),
		PhoneNumberText:          Ptr(RestrictString16("3659855214")),
	}
	party3.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: "212121211",
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("2"),
		},
	}
	party3.PartyOccupationBusiness = &PartyOccupationBusinessType{
		NAICSCode:              Ptr(RestrictString6("554")),
		OccupationBusinessText: Ptr(RestrictString30("Wedding Planner")),
	}
	party3.ElectronicAddress = &ElectronicAddressType{
		ElectronicAddressText: "johnson@wedding.org",
	}
	fmt.Println("Validating party3:", party3.Validate())

	party4 := &PartyType{
		ActivityPartyTypeCode:    ValidateActivityPartyCodeType("45"),
		PrimaryRegulatorTypeCode: Ptr(ValidateFederalRegulatorCodeType("1")),
	}
	party4.PartyName = []*PartyNameType{
		{
			RawPartyFullName: Ptr(RestrictString150("Bank of the West")),
		},
	}
	party4.Address = &AddressType{
		RawCityText:           "RawPartyFullName",
		RawStateCodeText:      "AZ",
		RawStreetAddress1Text: "555 Rolling Creek Way",
		RawZIPCode:            "554478985",
	}
	party4.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: "318181818",
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("2"),
		},
		{
			PartyIdentificationNumberText: "445581258",
			PartyIdentificationTypeCode:   ValidatePartyIdentificationCodeType("14"),
		},
	}
	fmt.Println("Validating party4:", party4.Validate())

	act := NewActivity()
	act.Party = []*PartyType{party1, party2, party3, party4}
	act.FilingDateText = "19801025"
	act.ActivityAssociation = &ActivityAssociationType{
		InitialDesignationIndicator: Ptr(ValidateIndicatorNullType("Y")),
	}
	act.DesignationExemptActivity = &DesignationExemptActivityType{
		ExemptBasisTypeCode:          "C",
		ExemptEffectiveBeginDateText: "20190208",
	}
	fmt.Println("Validating activity:", act.Validate())

	r := batch.NewReport(Report110)
	r.AppendActivity(act)

	// generating attr
	r.GenerateAttrs()
	r.GenerateSeqNumbers()

	fmt.Println("Validating batch report:", r.Validate())

	buf, _ := xml.MarshalIndent(r, "", "  ")

	fmt.Println("\nXML Body:")
	fmt.Println(string(buf))
}
