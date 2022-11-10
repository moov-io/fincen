// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"

	. "github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/batch"
	. "github.com/moov-io/fincen/pkg/currency_transaction"
)

func main() {
	party1 := &PartyType{
		ActivityPartyTypeCode:            ValidateActivityPartyCodeType("35"),
		EFilingCoverageBeginningDateText: Ptr(DateYYYYMMDDType("20170131")),
		EFilingCoverageEndDateText:       Ptr(DateYYYYMMDDType("20170131")),
	}
	party1.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Transmitter Legal Name")),
		},
	}
	party1.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("Vienna")),
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
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("4")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("PTCC1234")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("28")),
		},
	}
	fmt.Println("Validating party1:", party1.Validate())

	party2 := &PartyType{
		ActivityPartyTypeCode:            ValidateActivityPartyCodeType("37"),
		EFilingCoverageBeginningDateText: Ptr(DateYYYYMMDDType("20170101")),
		EFilingCoverageEndDateText:       Ptr(DateYYYYMMDDType("20170131")),
	}
	party2.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Transmitter Contact Name")),
		},
	}
	fmt.Println("Validating party2:", party2.Validate())

	party3 := &PartyType{
		ActivityPartyTypeCode:    ValidateActivityPartyCodeType("30"),
		PrimaryRegulatorTypeCode: Ptr(ValidateFederalRegulatorCodeType("9")),
	}
	party3.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Filing Institution Legal Name")),
		},
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("DBA")),
			RawPartyFullName:  Ptr(RestrictString150("Filing Institution Alternate DBA Name")),
		},
	}
	party3.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("Vienna")),
		RawCountryCodeText:    Ptr(RestrictString2("US")),
		RawStateCodeText:      Ptr(RestrictString3("VA")),
		RawStreetAddress1Text: Ptr(RestrictString100("456 Address Way")),
		RawZIPCode:            Ptr(RestrictString9("554789985")),
	}
	party3.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("554785215")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("2")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("15478564")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("10")),
		},
	}
	party3.OrganizationClassificationTypeSubtype = &OrganizationClassificationTypeSubtypeType{
		OrganizationTypeID: ValidateOrganizationCodeType(2),
	}
	fmt.Println("Validating party3:", party3.Validate())

	party4 := &PartyType{
		ActivityPartyTypeCode: ValidateActivityPartyCodeType("8"),
	}
	party4.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Contact Office Legal Name")),
		},
	}
	party4.PhoneNumber = &PhoneNumberType{
		PhoneNumberExtensionText: Ptr(RestrictString6("2210")),
		PhoneNumberText:          Ptr(RestrictString16("7039874589")),
	}
	fmt.Println("Validating party4:", party4.Validate())

	party5 := &PartyType{
		ActivityPartyTypeCode:            ValidateActivityPartyCodeType("34"),
		IndividualEntityCashInAmountText: Ptr(RestrictString15("15000")),
		PrimaryRegulatorTypeCode:         Ptr(ValidateFederalRegulatorCodeType("9")),
	}
	party5.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("L")),
			RawPartyFullName:  Ptr(RestrictString150("Filing Institution Legal Name")),
		},
		{
			PartyNameTypeCode: Ptr(ValidatePartyNameCodeType("DBA")),
			RawPartyFullName:  Ptr(RestrictString150("Filing Institution Alternate DBA Name")),
		},
	}
	party5.Address = &AddressType{
		RawCityText:           Ptr(RestrictString50("Vienna")),
		RawCountryCodeText:    Ptr(RestrictString2("US")),
		RawStateCodeText:      Ptr(RestrictString3("VA")),
		RawStreetAddress1Text: Ptr(RestrictString100("789 Address Court")),
		RawZIPCode:            Ptr(RestrictString9("55478")),
	}
	party5.PartyIdentification = []*PartyIdentificationType{
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("145878965")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("2")),
		},
		{
			PartyIdentificationNumberText: Ptr(RestrictString25("45899856")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("11")),
		},
	}
	party5.OrganizationClassificationTypeSubtype = &OrganizationClassificationTypeSubtypeType{
		OrganizationTypeID:           ValidateOrganizationCodeType(1),
		OrganizationSubtypeID:        Ptr(ValidateOrganizationSubtypeCodeCtrType(1999)),
		OtherOrganizationSubTypeText: Ptr(RestrictString50("Other casino")),
	}
	fmt.Println("Validating party5:", party5.Validate())

	party6 := &PartyType{
		ActivityPartyTypeCode:                           ValidateActivityPartyCodeType("50"),
		FemaleGenderIndicator:                           Ptr(ValidateIndicatorNullType("Y")),
		IndividualBirthDateText:                         Ptr(DateYYYYMMDDOrBlankTypeDOB("19750120")),
		IndividualEntityCashInAmountText:                Ptr(RestrictString15("20000")),
		MultipleTransactionsPersonsIndividualsIndicator: Ptr(ValidateIndicatorNullType("Y")),
	}
	party6.PartyName = []*PartyNameType{
		{
			PartyNameTypeCode:           Ptr(ValidatePartyNameCodeType("L")),
			RawEntityIndividualLastName: Ptr(RestrictString150("Doe")),
			RawIndividualFirstName:      Ptr(RestrictString35("John")),
			RawIndividualMiddleName:     Ptr(RestrictString35("Johnson")),
			RawIndividualNameSuffixText: Ptr(RestrictString35("Jr.")),
		},
		{
			PartyNameTypeCode:           Ptr(ValidatePartyNameCodeType("AKA")),
			RawEntityIndividualLastName: Ptr(RestrictString150("JJ")),
		},
	}
	party6.Address = &AddressType{
		RawCityText:                   Ptr(RestrictString50("Vienna")),
		RawCountryCodeText:            Ptr(RestrictString2("US")),
		RawStateCodeText:              Ptr(RestrictString3("VA")),
		StreetAddressUnknownIndicator: Ptr(ValidateIndicatorNullType("Y")),
		ZIPCodeUnknownIndicator:       Ptr(ValidateIndicatorNullType("Y")),
	}
	party6.PhoneNumber = &PhoneNumberType{
		PhoneNumberText: Ptr(RestrictString16("1458987456")),
	}
	party6.PartyIdentification = []*PartyIdentificationType{
		{
			TINUnknownIndicator: Ptr(ValidateIndicatorNullType("Y")),
		},
		{
			OtherIssuerCountryText:        Ptr(RestrictString2("US")),
			OtherIssuerStateText:          Ptr(RestrictString3("TX")),
			PartyIdentificationNumberText: Ptr(RestrictString25("44589774512")),
			PartyIdentificationTypeCode:   Ptr(ValidatePartyIdentificationCodeType("5")),
		},
	}
	party6.PartyOccupationBusiness = &PartyOccupationBusinessType{
		NAICSCode:              Ptr(RestrictString6("6214")),
		OccupationBusinessText: Ptr(RestrictString50("Outpatient Care Centers")),
	}
	party6.ElectronicAddress = &ElectronicAddressType{
		ElectronicAddressText: Ptr(RestrictString517("email123@fincen.gov")),
	}
	party6.Account = []*AccountType{
		{
			AccountNumberText: Ptr(RestrictString40("1115478569")),
			PartyAccountAssociation: &PartyAccountAssociationType{
				PartyAccountAssociationTypeCode: ValidatePartyAccountAssociationCodeType("8"),
			},
		},
		{
			AccountNumberText: Ptr(RestrictString40("3365998541")),
			PartyAccountAssociation: &PartyAccountAssociationType{
				PartyAccountAssociationTypeCode: ValidatePartyAccountAssociationCodeType("8"),
			},
		},
		{
			AccountNumberText: Ptr(RestrictString40("4857985691")),
			PartyAccountAssociation: &PartyAccountAssociationType{
				PartyAccountAssociationTypeCode: ValidatePartyAccountAssociationCodeType("8"),
			},
		},
	}
	fmt.Println("Validating party6:", party6.Validate())

	act := NewActivity()
	act.Party = []*PartyType{party1, party2, party3, party4, party5, party6}
	act.FilingDateText = "19801025"
	act.ActivityAssociation = &ActivityAssociationType{
		CorrectsAmendsPriorReportIndicator: "Y",
	}
	act.CurrencyTransactionActivity = &CurrencyTransactionActivityType{
		AggregateTransactionIndicator: "Y",
		ArmoredCarServiceIndicator:    "Y",
		ATMIndicator:                  "Y",
		MailDepositShipmentIndicator:  "Y",
		NightDepositIndicator:         "Y",
		SharedBranchingIndicator:      "Y",
		TotalCashInReceiveAmountText:  "4700",
		TransactionDateText:           "20170115",
		CurrencyTransactionActivityDetail: []*CurrencyTransactionActivityDetailType{
			{
				CurrencyTransactionActivityDetailTypeCode: "55",
				DetailTransactionAmountText:               "1500",
			},
			{
				CurrencyTransactionActivityDetailTypeCode: "46",
				DetailTransactionAmountText:               "1500",
			},
			{
				CurrencyTransactionActivityDetailTypeCode: "23",
				DetailTransactionAmountText:               "1500",
			},
		},
	}
	fmt.Println("Validating activity:", act.Validate())

	r := batch.NewReport(Report112)
	r.AppendActivity(act)

	// generating attr
	r.GenerateAttrs()
	r.GenerateSeqNumbers()

	fmt.Println("Validating batch report:", r.Validate())

	buf, _ := xml.MarshalIndent(r, "", "  ")

	fmt.Println("\nXML Body:")
	fmt.Println(string(buf))
}
