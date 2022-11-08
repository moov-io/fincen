// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"

	. "github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
	. "github.com/moov-io/fincen/pkg/suspicious_activity"
)

func main() {
	party1 := &PartyType{}
	party1.ActivityPartyTypeCode = ValidateActivityPartyCodeType("35")
	party1.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Transmitter legal name")),
		},
	}
	party1.Address = []*AddressType{
		{
			RawCityText:           Ptr(RestrictString50("Transmitter city")),
			RawCountryCodeText:    Ptr(RestrictString2("US")),
			RawStateCodeText:      Ptr(RestrictString3("VA")),
			RawStreetAddress1Text: Ptr(RestrictString100("Transmitter street address")),
			RawZIPCode:            Ptr(RestrictString9("22113")),
		},
	}
	party1.PhoneNumber = []*PhoneNumberType{
		{
			PhoneNumberText: Ptr(RestrictString16("7217894455")),
		},
	}
	party1.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("458985215")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("4")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("458985215")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("28")),
		},
	}
	fmt.Println("Validating party1:", party1.Validate())

	party2 := &PartyType{}
	party2.ActivityPartyTypeCode = ValidateActivityPartyCodeType("37")
	party2.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Transmitter legal name")),
		},
	}
	fmt.Println("Validating party2:", party2.Validate())

	party3 := &PartyType{}
	party3.ActivityPartyTypeCode = ValidateActivityPartyCodeType("30")
	party3.PrimaryRegulatorTypeCode = Ptr(ValidateFederalRegulatorCodeType("7"))
	party3.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Filer name")),
		},
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("DBA")),
			RawPartyFullName:  Ptr(RestrictString150("Alternate name")),
		},
	}
	party3.Address = []*AddressType{
		{
			RawCityText:           Ptr(RestrictString50("Rockville")),
			RawCountryCodeText:    Ptr(RestrictString2("US")),
			RawStateCodeText:      Ptr(RestrictString3("MD")),
			RawStreetAddress1Text: Ptr(RestrictString100("123 Viers Mill Road")),
			RawZIPCode:            Ptr(RestrictString9("20905")),
		},
	}
	party3.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("125478965")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("4")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("5558789654")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("10")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("451256558")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("29")),
		},
	}
	party3.OrganizationClassificationTypeSubtype = []*OrganizationClassificationTypeSubtypeType{
		{
			OrganizationSubtypeID: Ptr(ValidateOrganizationSubtypeCodeSarType(535)),
			OrganizationTypeID:    Ptr(ValidateOrganizationCodeType(5)),
		},
		{
			OrganizationSubtypeID: Ptr(ValidateOrganizationSubtypeCodeSarType(529)),
			OrganizationTypeID:    Ptr(ValidateOrganizationCodeType(5)),
		},
	}
	fmt.Println("Validating party3:", party3.Validate())

	party4 := &PartyType{}
	party4.ActivityPartyTypeCode = ValidateActivityPartyCodeType("8")
	party4.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Designated contact office")),
		},
	}
	party4.PhoneNumber = []*PhoneNumberType{
		{
			PhoneNumberExtensionText: Ptr(RestrictString6("1234")),
			PhoneNumberText:          Ptr(RestrictString16("4157653838")),
		},
	}
	fmt.Println("Validating party4:", party4.Validate())

	party5 := &PartyType{}
	party5.ActivityPartyTypeCode = ValidateActivityPartyCodeType("34")
	party5.LossToFinancialAmountText = Ptr(RestrictString15("15000"))
	party5.NoBranchActivityInvolvedIndicator = Ptr(ValidateIndicatorType("Y"))
	party5.PayLocationIndicator = Ptr(ValidateIndicatorType("Y"))
	party5.PrimaryRegulatorTypeCode = Ptr(ValidateFederalRegulatorCodeType("4"))
	party5.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Union Bank of California")),
		},
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("DBA")),
			RawPartyFullName:  Ptr(RestrictString150("Doing Business As Name")),
		},
	}
	party5.Address = []*AddressType{
		{
			CityUnknownIndicator:          Ptr(ValidateIndicatorType("Y")),
			CountryCodeUnknownIndicator:   Ptr(ValidateIndicatorType("Y")),
			RawStreetAddress1Text:         Ptr(RestrictString100("987 Rocky Road")),
			StreetAddressUnknownIndicator: Ptr(ValidateIndicatorType("Y")),
			ZIPCodeUnknownIndicator:       Ptr(ValidateIndicatorType("Y")),
		},
	}
	party5.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("458789856")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("4")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("5589887789")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("10")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("4578958658")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("29")),
		},
	}
	party5.OrganizationClassificationTypeSubtype = []*OrganizationClassificationTypeSubtypeType{
		{
			OrganizationSubtypeID: Ptr(ValidateOrganizationSubtypeCodeSarType(533)),
			OrganizationTypeID:    Ptr(ValidateOrganizationCodeType(5)),
		},
		{
			OrganizationSubtypeID: Ptr(ValidateOrganizationSubtypeCodeSarType(5999)),
			OrganizationTypeID:    Ptr(ValidateOrganizationCodeType(5)),
		},
	}
	party5.PartyAssociation = []*PartyAssociationType{
		{
			Party: []*AssociationParty{
				{
					ActivityPartyTypeCode: ValidateActivityPartyCodeType("46"),
					Address:               &AddressType{},
				},
			},
		},
	}
	fmt.Println("Validating party5:", party5.Validate())

	party6 := &PartyType{}
	party6.ActivityPartyTypeCode = ValidateActivityPartyCodeType("33")
	party6.AdmissionConfessionYesIndicator = Ptr(ValidateIndicatorType("Y"))
	party6.BothPurchaserSenderPayeeReceiveIndicator = Ptr(ValidateIndicatorType("Y"))
	party6.FemaleGenderIndicator = Ptr(ValidateIndicatorType("Y"))
	party6.IndividualBirthDateText = Ptr(DateYYYYMMDDOrBlankTypeDOB("19801025"))
	party6.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode:           Ptr(ValidatePartyNameCodeType("L")),
			RawEntityIndividualLastName: Ptr(RestrictString150("Mann")),
			RawIndividualFirstName:      Ptr(RestrictString35("Janice")),
		},
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("DBA")),
			RawPartyFullName:  Ptr(RestrictString150("Janda")),
		},
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("AKA")),
			RawPartyFullName:  Ptr(RestrictString150("Jan")),
		},
	}
	party6.Address = []*AddressType{
		{
			CityUnknownIndicator:          Ptr(ValidateIndicatorType("Y")),
			CountryCodeUnknownIndicator:   Ptr(ValidateIndicatorType("Y")),
			StateCodeUnknownIndicator:     Ptr(ValidateIndicatorType("Y")),
			StreetAddressUnknownIndicator: Ptr(ValidateIndicatorType("Y")),
			ZIPCodeUnknownIndicator:       Ptr(ValidateIndicatorType("Y")),
		},
	}
	party6.PhoneNumber = []*PhoneNumberType{
		{
			PhoneNumberText:     Ptr(RestrictString16("7217894455")),
			PhoneNumberTypeCode: Ptr(ValidatePhoneNumberCodeType("R")),
		},
		{
			PhoneNumberText:          Ptr(RestrictString16("6589784589")),
			PhoneNumberTypeCode:      Ptr(ValidatePhoneNumberCodeType("W")),
			PhoneNumberExtensionText: Ptr(RestrictString6("5584")),
		},
	}
	party6.PartyIdentification = []*PartyIdentificationType{
		{
			IdentificationPresentUnknownIndicator: Ptr(ValidateIndicatorType("Y")),
			OtherIssuerCountryText:                Ptr(RestrictString2("US")),
			OtherIssuerStateText:                  Ptr(RestrictString3("CA")),
			OtherPartyIdentificationTypeText:      Ptr(RestrictString50("Student ID")),
			PartyIdentificationNumberText:         Ptr(RestrictString25("660623559")),
			PartyIdentificationTypeCode:           Ptr(ValidatePartyIdentificationCodeType("999")),
		},
		{
			TINUnknownIndicator: Ptr(ValidateIndicatorType("Y")),
		},
	}
	party6.PartyOccupationBusiness = &PartyOccupationBusinessType{
		NAICSCode:              Ptr(RestrictString6("7548")),
		OccupationBusinessText: Ptr(RestrictString50("7548")),
	}
	party6.ElectronicAddress = []*ElectronicAddressType{
		{
			ElectronicAddressText:     RestrictString517("jan@gmail.com"),
			ElectronicAddressTypeCode: ValidateElectronicAddressTypeCode("E"),
		},
		{
			ElectronicAddressText:     RestrictString517("janda.org"),
			ElectronicAddressTypeCode: ValidateElectronicAddressTypeCode("U"),
		},
	}
	party6.PartyAssociation = []*PartyAssociationType{
		{
			NoRelationshipToInstitutionIndicator:           Ptr(ValidateIndicatorType("Y")),
			SubjectRelationshipFinancialInstitutionTINText: Ptr(RestrictString25("458789856")),
		},
	}
	party6.PartyAccountAssociation = &PartyAccountAssociationType{
		Party: []*AccountAssociationParty{
			{
				ActivityPartyTypeCode: ValidateActivityPartyCodeType("41"),
			},
		},
		PartyAccountAssociationTypeCode: ValidatePartyAccountAssociationCodeType("7"),
	}
	fmt.Println("Validating party6:", party6.Validate())

	act := NewActivity()
	act.Party = []*PartyType{party1, party2, party3, party4, party5, party6}
	act.SuspiciousActivity = &SuspiciousActivityType{
		SuspiciousActivityFromDateText: DateYYYYMMDDType("19801025"),
		SuspiciousActivityClassification: []*SuspiciousActivityClassificationType{
			{
				SuspiciousActivitySubtypeID: ValidateSuspiciousActivitySubtypeID(106),
				SuspiciousActivityTypeID:    ValidateSuspiciousActivityTypeID(1),
			},
		},
	}
	act.EFilingPriorDocumentNumber = Ptr(EFilingPriorDocumentNumberType("00000000000001"))
	act.FilingDateText = "19801025"
	act.ActivityAssociation = &ActivityAssociationType{
		CorrectsAmendsPriorReportIndicator: Ptr(ValidateIndicatorType("Y")),
	}
	fmt.Println("Validating activity:", act.Validate())

	r := batch.NewReport(Report111)
	r.AppendActivity(act)

	// generating attr
	r.GenerateAttrs()
	r.GenerateSeqNumbers()

	fmt.Println("Validating batch report:", r.Validate())

	buf, _ := xml.MarshalIndent(r, "", "  ")

	fmt.Println("\nXML Body:")
	fmt.Println(string(buf))
}
