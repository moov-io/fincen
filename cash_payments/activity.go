// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Report of Cash Payments Over $10,000 Received in a Trade or Business (FinCEN Form 8300)

package cash_payments

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

type ActivityType struct {
	XMLName                                     xml.Name                          `xml:"Activity"`
	SeqNum                                      fincen.SeqNumber                  `xml:"SeqNum,attr"`
	EFilingPriorDocumentNumber                  int64                             `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                              fincen.DateYYYYMMDDType           `xml:"FilingDateText"`
	MultipleSubjectsIndicator                   *fincen.ValidateIndicatorType     `xml:"MultipleSubjectsIndicator,omitempty" json:",omitempty"`
	SuspiciousTransactionIndicator              *fincen.ValidateIndicatorType     `xml:"SuspiciousTransactionIndicator,omitempty" json:",omitempty"`
	TransactionOnBehalfMultiplePersonsIndicator *fincen.ValidateIndicatorType     `xml:"TransactionOnBehalfMultiplePersonsIndicator,omitempty" json:",omitempty"`
	ActivityAssociation                         *ActivityAssociationType          `xml:"ActivityAssociation"`
	Party                                       []PartyType                       `xml:"Party"`
	CurrencyTransactionActivity                 CurrencyTransactionActivityType   `xml:"CurrencyTransactionActivity"`
	ActivityNarrativeInformation                *ActivityNarrativeInformationType `xml:"ActivityNarrativeInformation,omitempty" json:",omitempty"`
}

func (r ActivityType) FormTypeCode() string {
	return "8300X"
}

func (r ActivityType) fieldInclusion() error {
	if len(r.Party) < 4 || len(r.Party) > 203 {
		return fincen.NewErrMinMaxRange("Party")
	}

	return nil
}

func (r ActivityType) Validate(args ...string) error {

	if err := r.fieldInclusion(); err != nil {
		return err
	}

	existed := make(map[string]int)
	for _, p := range r.Party {
		typeCode := string(p.ActivityPartyTypeCode)
		if cnt, ok := existed[typeCode]; !ok {
			existed[typeCode] = cnt + 1
		} else {
			existed[typeCode] = 1
		}
	}

	if _, ok := existed["35"]; !ok {
		return fincen.NewErrFiledNotAssociated("Party(type 35)")
	}
	if _, ok := existed["37"]; !ok {
		return fincen.NewErrFiledNotAssociated("Party(type 37)")
	}
	if _, ok := existed["3"]; !ok {
		return fincen.NewErrFiledNotAssociated("Party(type 3)")
	}
	if _, ok := existed["4"]; !ok {
		return fincen.NewErrFiledNotAssociated("Party(type 4)")
	}
	if cnt, ok := existed["16"]; !ok || cnt > 99 {
		return fincen.NewErrMinMaxRange("Party(type 16)")
	}
	if cnt, ok := existed["23"]; !ok || cnt > 99 {
		return fincen.NewErrMinMaxRange("Party(type 23)")
	}
	if cnt, ok := existed["3"]; !ok || cnt > 1 {
		return fincen.NewErrMinMaxRange("Party(type 3)")
	}
	if cnt, ok := existed["4"]; !ok || cnt > 1 {
		return fincen.NewErrMinMaxRange("Party(type 4)")
	}

	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	XMLName                            xml.Name                      `xml:"ActivityAssociation"`
	SeqNum                             fincen.SeqNumber              `xml:"SeqNum,attr"`
	CorrectsAmendsPriorReportIndicator *fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator,omitempty" json:",omitempty"`
	InitialReportIndicator             *fincen.ValidateIndicatorType `xml:"InitialReportIndicator,omitempty" json:",omitempty"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyType struct {
	XMLName                 xml.Name                        `xml:"Party"`
	SeqNum                  fincen.SeqNumber                `xml:"SeqNum,attr"`
	ActivityPartyTypeCode   ValidateActivityPartyCodeType   `xml:"ActivityPartyTypeCode"`
	IndividualBirthDateText *fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	PartyTypeCode           *fincen.ValidatePartyTypeCode   `xml:"PartyTypeCode,omitempty" json:",omitempty"`
	PartyName               []PartyNameType                 `xml:"PartyName,omitempty" json:",omitempty"`
	Address                 *AddressType                    `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber             *PhoneNumberType                `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification     []PartyIdentificationType       `xml:"PartyIdentification,omitempty" json:",omitempty"`
	PartyOccupationBusiness *PartyOccupationBusinessType    `xml:"PartyOccupationBusiness,omitempty" json:",omitempty"`
}

func (r PartyType) fieldInclusion() error {
	typeCode := string(r.ActivityPartyTypeCode)

	if typeCode == "35" {
		if r.PartyName == nil {
			return fincen.NewErrFiledNotAssociated("PartyName")
		}
		if r.Address == nil {
			return fincen.NewErrFiledNotAssociated("Address")
		}
		if r.PhoneNumber == nil {
			return fincen.NewErrFiledNotAssociated("PhoneNumber")
		}
		if len(r.PartyIdentification) < 1 {
			return fincen.NewErrFiledNotAssociated("PartyIdentification")
		}
		if len(r.PartyIdentification) > 1 {
			return fincen.NewErrMinMaxRange("PartyIdentification")
		}
	}

	if typeCode == "37" || typeCode == "3" {
		if r.PartyName == nil {
			return fincen.NewErrFiledNotAssociated("PartyName")
		}
	}

	if typeCode == "16" {
		if r.PartyName == nil {
			return fincen.NewErrFiledNotAssociated("PartyName")
		}
		if r.Address == nil {
			return fincen.NewErrFiledNotAssociated("Address")
		}
		if len(r.PartyIdentification) < 1 {
			return fincen.NewErrFiledNotAssociated("PartyIdentification")
		}
		if len(r.PartyIdentification) > 2 {
			return fincen.NewErrMinMaxRange("PartyIdentification")
		}
	}

	if typeCode == "23" {
		if r.PartyTypeCode == nil {
			return fincen.NewErrFiledNotAssociated("PartyTypeCode")
		}
		if len(r.PartyName) > 2 {
			return fincen.NewErrMinMaxRange("PartyName")
		}
		if len(r.PartyIdentification) > 3 {
			return fincen.NewErrMinMaxRange("PartyIdentification")
		}
	}

	if typeCode == "4" {
		if r.PartyName == nil {
			return fincen.NewErrFiledNotAssociated("PartyName")
		}
		if len(r.PartyIdentification) < 1 {
			return fincen.NewErrFiledNotAssociated("PartyIdentification")
		}
		if len(r.PartyIdentification) > 2 {
			return fincen.NewErrMinMaxRange("PartyIdentification")
		}
	}

	return nil
}

func (r PartyType) Validate(args ...string) error {
	if err := r.fieldInclusion(); err != nil {
		return err
	}

	return fincen.Validate(&r, string(r.ActivityPartyTypeCode))
}

type PartyNameType struct {
	XMLName                     xml.Name                   `xml:"PartyName"`
	SeqNum                      fincen.SeqNumber           `xml:"SeqNum,attr"`
	PartyNameTypeCode           *ValidatePartyNameCodeType `xml:"PartyNameTypeCode,omitempty" json:",omitempty"`
	RawEntityIndividualLastName *fincen.RestrictString150  `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName      *fincen.RestrictString35   `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName     *fincen.RestrictString35   `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualTitleText      *fincen.RestrictString35   `xml:"RawIndividualTitleText,omitempty" json:",omitempty"`
	RawPartyFullName            *fincen.RestrictString150  `xml:"RawPartyFullName,omitempty" json:",omitempty"`
}

func (r PartyNameType) fieldInclusion(typeCode string) error {
	if fincen.CheckInvolved(typeCode, "35", "37") {
		if r.PartyNameTypeCode == nil || *r.PartyNameTypeCode != "L" {
			return fincen.NewErrValueInvalid("PartyNameTypeCode")
		}
		if r.RawPartyFullName == nil {
			return fincen.NewErrFiledNotAssociated("RawPartyFullName")
		}
	}

	if fincen.CheckInvolved(typeCode, "23") {
		if r.PartyNameTypeCode != nil && *r.PartyNameTypeCode != "DBA" {
			if r.RawPartyFullName == nil {
				return fincen.NewErrFiledNotAssociated("RawPartyFullName")
			}
		}
	}

	if fincen.CheckInvolved(typeCode, "4") && r.RawPartyFullName == nil {
		return fincen.NewErrFiledNotAssociated("RawPartyFullName")
	}

	if fincen.CheckInvolved(typeCode, "3") && r.RawIndividualTitleText == nil {
		return fincen.NewErrFiledNotAssociated("RawIndividualTitleText")
	}

	if fincen.CheckInvolved(typeCode, "8") && r.RawPartyFullName == nil {
		return fincen.NewErrFiledNotAssociated("RawPartyFullName")
	}

	return nil
}

func (r PartyNameType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	if err := r.fieldInclusion(typeCode); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type AddressType struct {
	XMLName               xml.Name                  `xml:"Address"`
	SeqNum                fincen.SeqNumber          `xml:"SeqNum,attr"`
	RawCityText           *fincen.RestrictString50  `xml:"RawCityText,omitempty" json:",omitempty"`
	RawCountryCodeText    *fincen.RestrictString2   `xml:"RawCountryCodeText,omitempty" json:",omitempty"`
	RawStateCodeText      *fincen.RestrictString3   `xml:"RawStateCodeText,omitempty" json:",omitempty"`
	RawStreetAddress1Text *fincen.RestrictString100 `xml:"RawStreetAddress1Text,omitempty" json:",omitempty"`
	RawZIPCode            *fincen.RestrictString9   `xml:"RawZIPCode,omitempty" json:",omitempty"`
}

func (r AddressType) fieldInclusion(typeCode string) error {
	if fincen.CheckInvolved(typeCode, "35") {
		if r.RawCityText == nil {
			return fincen.NewErrFiledNotAssociated("RawCityText")
		}
		if r.RawCountryCodeText == nil {
			return fincen.NewErrFiledNotAssociated("RawCountryCodeText")
		}
		if r.RawStateCodeText == nil {
			return fincen.NewErrFiledNotAssociated("RawStateCodeText")
		}
		if r.RawStreetAddress1Text == nil {
			return fincen.NewErrFiledNotAssociated("RawStreetAddress1Text")
		}
		if r.RawZIPCode == nil {
			return fincen.NewErrFiledNotAssociated("RawZIPCode")
		}
	}

	return nil
}

func (r AddressType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	if err := r.fieldInclusion(typeCode); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type PhoneNumberType struct {
	XMLName         xml.Name                 `xml:"PhoneNumber"`
	SeqNum          fincen.SeqNumber         `xml:"SeqNum,attr"`
	PhoneNumberText *fincen.RestrictString16 `xml:"PhoneNumberText,omitempty" json:",omitempty"`
}

func (r PhoneNumberType) fieldInclusion(typeCode string) error {
	if fincen.CheckInvolved(typeCode, "35", "8") {
		if r.PhoneNumberText == nil {
			return fincen.NewErrFiledNotAssociated("PhoneNumberText")
		}
	}

	return nil
}

func (r PhoneNumberType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	if err := r.fieldInclusion(typeCode); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type PartyIdentificationType struct {
	XMLName                       xml.Name                            `xml:"PartyIdentification"`
	SeqNum                        fincen.SeqNumber                    `xml:"SeqNum,attr"`
	OtherIssuerStateText          *fincen.RestrictString3             `xml:"OtherIssuerStateText,omitempty" json:",omitempty"`
	PartyIdentificationNumberText *fincen.RestrictString25            `xml:"PartyIdentificationNumberText,omitempty" json:",omitempty"`
	PartyIdentificationTypeCode   ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode"`
}

func (r PartyIdentificationType) fieldInclusion(typeCode string) error {
	if fincen.CheckInvolved(typeCode, "35") {
		if r.PartyIdentificationNumberText == nil {
			return fincen.NewErrFiledNotAssociated("PartyIdentificationNumberText")
		}
		if string(r.PartyIdentificationTypeCode) != "28" && string(r.PartyIdentificationTypeCode) != "2" {
			return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
		}
	}

	if fincen.CheckInvolved(typeCode, "4") {
		if r.PartyIdentificationNumberText == nil {
			return fincen.NewErrFiledNotAssociated("PartyIdentificationNumberText")
		}
	}

	return nil
}

func (r PartyIdentificationType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	if err := r.fieldInclusion(typeCode); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type PartyOccupationBusinessType struct {
	XMLName                xml.Name                `xml:"PartyOccupationBusiness"`
	SeqNum                 fincen.SeqNumber        `xml:"SeqNum,attr"`
	OccupationBusinessText fincen.RestrictString30 `xml:"OccupationBusinessText"`
}

func (r PartyOccupationBusinessType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityType struct {
	XMLName                           xml.Name                                `xml:"CurrencyTransactionActivity"`
	SeqNum                            fincen.SeqNumber                        `xml:"SeqNum,attr"`
	InstallmentPaymentOtherIndicator  *fincen.ValidateIndicatorType           `xml:"InstallmentPaymentOtherIndicator,omitempty" json:",omitempty"`
	Total100DollarBillInAmountText    string                                  `xml:"Total100DollarBillInAmountText,omitempty" json:",omitempty"`
	TotalCashInReceiveAmountText      string                                  `xml:"TotalCashInReceiveAmountText"`
	TotalPriceAmountText              string                                  `xml:"TotalPriceAmountText,omitempty" json:",omitempty"`
	TransactionDateText               fincen.DateYYYYMMDDType                 `xml:"TransactionDateText"`
	CurrencyTransactionActivityDetail []CurrencyTransactionActivityDetailType `xml:"CurrencyTransactionActivityDetail"`
}

func (r CurrencyTransactionActivityType) fieldInclusion() error {
	if len(r.CurrencyTransactionActivityDetail) < 2 || len(r.CurrencyTransactionActivityDetail) > 9 {
		return fincen.NewErrMinMaxRange("CurrencyTransactionActivityDetail")
	}

	return nil
}

func (r CurrencyTransactionActivityType) Validate(args ...string) error {
	if err := r.fieldInclusion(); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type ActivityNarrativeInformationType struct {
	XMLName                         xml.Name                                `xml:"ActivityNarrativeInformation"`
	SeqNum                          fincen.SeqNumber                        `xml:"SeqNum,attr"`
	ActivityNarrativeSequenceNumber ValidateActivityNarrativeSequenceNumber `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString750                `xml:"ActivityNarrativeText"`
}

func (r ActivityNarrativeInformationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityDetailType struct {
	XMLName                                   xml.Name                                                  `xml:"CurrencyTransactionActivityDetail"`
	SeqNum                                    fincen.SeqNumber                                          `xml:"SeqNum,attr"`
	CurrencyTransactionActivityDetailTypeCode *fincen.ValidateCurrencyTransactionActivityDetailCodeType `xml:"CurrencyTransactionActivityDetailTypeCode,omitempty" json:",omitempty"`
	DetailTransactionAmountText               *fincen.RestrictString15                                  `xml:"DetailTransactionAmountText,omitempty" json:",omitempty"`
	DetailTransactionDescription              string                                                    `xml:"DetailTransactionDescription,omitempty" json:",omitempty"`
	InstrumentProductServiceTypeCode          *fincen.ValidateInstrumentProductServiceTypeCode          `xml:"InstrumentProductServiceTypeCode,omitempty" json:",omitempty"`
	IssuerNameText                            string                                                    `xml:"IssuerNameText,omitempty" json:",omitempty"`
	OtherForeignCurrencyCountryText           *fincen.RestrictString2                                   `xml:"OtherForeignCurrencyCountryText,omitempty" json:",omitempty"`
}

func (r CurrencyTransactionActivityDetailType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
