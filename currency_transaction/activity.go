// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package currency_transaction

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

func checkInvolved(element string, elements ...string) bool {
	for _, elm := range elements {
		if element == elm {
			return true
		}
	}
	return false
}

type ActivityType struct {
	XMLName                     xml.Name                        `xml:"Activity"`
	SeqNum                      fincen.SeqNumber                `xml:"SeqNum,attr"`
	EFilingPriorDocumentNumber  int64                           `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText              fincen.DateYYYYMMDDType         `xml:"FilingDateText"`
	ActivityAssociation         ActivityAssociationType         `xml:"ActivityAssociation"`
	Party                       []PartyType                     `xml:"Party"`
	CurrencyTransactionActivity CurrencyTransactionActivityType `xml:"CurrencyTransactionActivity"`
}

func (r ActivityType) FormTypeCode() string {
	return "CTRX"
}

func (r ActivityType) fieldInclusion() error {
	if len(r.Party) < 6 || len(r.Party) > 2002 {
		return fincen.NewErrValueInvalid("Party")
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
		return fincen.NewErrValueInvalid("Party")
	}
	if _, ok := existed["37"]; !ok {
		return fincen.NewErrValueInvalid("Party")
	}
	if _, ok := existed["30"]; !ok {
		return fincen.NewErrValueInvalid("Party")
	}
	if _, ok := existed["8"]; !ok {
		return fincen.NewErrValueInvalid("Party")
	}
	if cnt, ok := existed["34"]; !ok || cnt > 999 {
		return fincen.NewErrValueInvalid("Party")
	}
	if cnt, ok := existed["50"]; !ok || cnt > 999 {
		return fincen.NewErrValueInvalid("Party")
	}

	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	XMLName                            xml.Name                      `xml:"ActivityAssociation"`
	SeqNum                             fincen.SeqNumber              `xml:"SeqNum,attr"`
	CorrectsAmendsPriorReportIndicator *fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator"`
	FinCENDirectBackFileIndicator      *fincen.ValidateIndicatorType `xml:"FinCENDirectBackFileIndicator"`
	InitialReportIndicator             *fincen.ValidateIndicatorType `xml:"InitialReportIndicator"`
}

func (r ActivityAssociationType) Validate(args ...string) error {

	indicatorCnt := 0

	if r.CorrectsAmendsPriorReportIndicator != nil {
		indicatorCnt++
	}
	if r.FinCENDirectBackFileIndicator != nil {
		indicatorCnt++
	}
	if r.InitialReportIndicator != nil {
		indicatorCnt++
	}

	if indicatorCnt != 1 {
		return fincen.NewErrValueInvalid("ActivityAssociationType")
	}

	return fincen.Validate(&r, args...)
}

type PartyType struct {
	XMLName                                         xml.Name                                   `xml:"Party"`
	SeqNum                                          fincen.SeqNumber                           `xml:"SeqNum,attr"`
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType              `xml:"ActivityPartyTypeCode"`
	BirthDateUnknownIndicator                       *fincen.ValidateIndicatorType              `xml:"BirthDateUnknownIndicator,omitempty" json:",omitempty"`
	EFilingCoverageBeginningDateText                *fincen.DateYYYYMMDDType                   `xml:"EFilingCoverageBeginningDateText,omitempty" json:",omitempty"`
	EFilingCoverageEndDateText                      *fincen.DateYYYYMMDDType                   `xml:"EFilingCoverageEndDateText,omitempty" json:",omitempty"`
	FemaleGenderIndicator                           *fincen.ValidateIndicatorType              `xml:"FemaleGenderIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                         *fincen.DateYYYYMMDDOrBlankTypeDOB         `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	IndividualEntityCashInAmountText                *fincen.RestrictString15                   `xml:"IndividualEntityCashInAmountText,omitempty" json:",omitempty"`
	IndividualEntityCashOutAmountText               *fincen.RestrictString15                   `xml:"IndividualEntityCashOutAmountText,omitempty" json:",omitempty"`
	MaleGenderIndicator                             *fincen.ValidateIndicatorType              `xml:"MaleGenderIndicator,omitempty" json:",omitempty"`
	MultipleTransactionsPersonsIndividualsIndicator *fincen.ValidateIndicatorType              `xml:"MultipleTransactionsPersonsIndividualsIndicator,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator              *fincen.ValidateIndicatorType              `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode                        *ValidateFederalRegulatorCodeType          `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	UnknownGenderIndicator                          *fincen.ValidateIndicatorType              `xml:"UnknownGenderIndicator,omitempty" json:",omitempty"`
	PartyName                                       []PartyNameType                            `xml:"PartyName"`
	Address                                         *AddressType                               `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber                                     *PhoneNumberType                           `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification                             []PartyIdentificationType                  `xml:"PartyIdentification,omitempty" json:",omitempty"`
	OrganizationClassificationTypeSubtype           *OrganizationClassificationTypeSubtypeType `xml:"OrganizationClassificationTypeSubtype,omitempty" json:",omitempty"`
	PartyOccupationBusiness                         *PartyOccupationBusinessType               `xml:"PartyOccupationBusiness,omitempty" json:",omitempty"`
	ElectronicAddress                               *ElectronicAddressType                     `xml:"ElectronicAddress,omitempty" json:",omitempty"`
	Account                                         []AccountType                              `xml:"Account,omitempty" json:",omitempty"`
}

func (r PartyType) fieldInclusion() error {
	typeCode := string(r.ActivityPartyTypeCode)

	if r.BirthDateUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("BirthDateUnknownIndicator")
	}

	if r.EFilingCoverageBeginningDateText != nil && !checkInvolved(typeCode, "35", "37") {
		return fincen.NewErrFiledNotAssociated("EFilingCoverageBeginningDateText")
	}

	if r.EFilingCoverageEndDateText != nil && !checkInvolved(typeCode, "35", "37") {
		return fincen.NewErrFiledNotAssociated("EFilingCoverageEndDateText")
	}

	if r.FemaleGenderIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("FemaleGenderIndicator")
	}

	if r.IndividualBirthDateText != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("IndividualBirthDateText")
	}

	if r.IndividualEntityCashInAmountText != nil && !checkInvolved(typeCode, "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("IndividualEntityCashInAmountText")
	}

	if r.IndividualEntityCashOutAmountText != nil && !checkInvolved(typeCode, "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("IndividualEntityCashOutAmountText")
	}

	if r.MaleGenderIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("MaleGenderIndicator")
	}

	if r.MultipleTransactionsPersonsIndividualsIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("MultipleTransactionsPersonsIndividualsIndicator")
	}

	if r.PartyAsEntityOrganizationIndicator != nil && !checkInvolved(typeCode, "23", "58") {
		return fincen.NewErrFiledNotAssociated("PartyAsEntityOrganizationIndicator")
	}

	if r.PrimaryRegulatorTypeCode != nil && !checkInvolved(typeCode, "30", "34") {
		return fincen.NewErrFiledNotAssociated("PrimaryRegulatorTypeCode")
	}

	if r.UnknownGenderIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("UnknownGenderIndicator")
	}

	if len(r.PartyName) < 1 || len(r.PartyName) > 2 {
		return fincen.NewErrMinMaxRange("PartyName")
	}

	if r.Address != nil && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("UnknownGenderIndicator")
	}

	if r.PhoneNumber != nil && !checkInvolved(typeCode, "35", "8", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("PhoneNumber")
	}

	if len(r.PartyIdentification) > 2 {
		return fincen.NewErrMinMaxRange("PartyIdentification")
	}

	if len(r.PartyIdentification) > 0 && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("PartyIdentification")
	}

	if r.OrganizationClassificationTypeSubtype != nil && !checkInvolved(typeCode, "30", "34") {
		return fincen.NewErrFiledNotAssociated("OrganizationClassificationTypeSubtype")
	}

	if r.PartyOccupationBusiness != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("PartyOccupationBusiness")
	}

	if r.ElectronicAddress != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("ElectronicAddress")
	}

	if len(r.Account) > 198 {
		return fincen.NewErrMinMaxRange("Account")
	}

	if len(r.Account) > 0 && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("Account")
	}

	return nil
}

func (r PartyType) Validate(args ...string) error {
	if err := r.fieldInclusion(); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type PartyNameType struct {
	XMLName                        xml.Name                      `xml:"PartyName"`
	SeqNum                         fincen.SeqNumber              `xml:"SeqNum,attr"`
	EntityLastNameUnknownIndicator *fincen.ValidateIndicatorType `xml:"EntityLastNameUnknownIndicator,omitempty" json:",omitempty"`
	FirstNameUnknownIndicator      *fincen.ValidateIndicatorType `xml:"FirstNameUnknownIndicator,omitempty" json:",omitempty"`
	PartyNameTypeCode              *ValidatePartyNameCodeType    `xml:"PartyNameTypeCode,omitempty" json:",omitempty"`
	RawEntityIndividualLastName    *fincen.RestrictString150     `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName         *fincen.RestrictString35      `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName        *fincen.RestrictString35      `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualNameSuffixText    *fincen.RestrictString35      `xml:"RawIndividualNameSuffixText,omitempty" json:",omitempty"`
	RawPartyFullName               *fincen.RestrictString150     `xml:"RawPartyFullName,omitempty" json:",omitempty"`
}

func (r PartyNameType) fieldInclusion(typeCode string) error {
	if r.EntityLastNameUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("EntityLastNameUnknownIndicator")
	}

	if r.FirstNameUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("FirstNameUnknownIndicator")
	}

	if r.PartyNameTypeCode != nil && string(*r.PartyNameTypeCode) != "L" &&
		!checkInvolved(typeCode, "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrValueInvalid("PartyNameTypeCode")
	}

	if r.RawEntityIndividualLastName != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawEntityIndividualLastName")
	}

	if r.RawIndividualFirstName != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawIndividualFirstName")
	}

	if r.RawIndividualMiddleName != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawIndividualMiddleName")
	}

	if r.RawIndividualNameSuffixText != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawIndividualNameSuffixText")
	}

	if r.RawPartyFullName != nil && checkInvolved(typeCode, "50", "17", "23", "58") {
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
	XMLName                       xml.Name                      `xml:"Address"`
	SeqNum                        fincen.SeqNumber              `xml:"SeqNum,attr"`
	CityUnknownIndicator          *fincen.ValidateIndicatorType `xml:"CityUnknownIndicator,omitempty" json:",omitempty"`
	CountryCodeUnknownIndicator   *fincen.ValidateIndicatorType `xml:"CountryCodeUnknownIndicator,omitempty" json:",omitempty"`
	RawCityText                   *fincen.RestrictString50      `xml:"RawCityText,omitempty" json:",omitempty"`
	RawCountryCodeText            *fincen.RestrictString2       `xml:"RawCountryCodeText,omitempty" json:",omitempty"`
	RawStateCodeText              *fincen.RestrictString3       `xml:"RawStateCodeText,omitempty" json:",omitempty"`
	RawStreetAddress1Text         *fincen.RestrictString100     `xml:"RawStreetAddress1Text,omitempty" json:",omitempty"`
	RawZIPCode                    *fincen.RestrictString9       `xml:"RawZIPCode,omitempty" json:",omitempty"`
	StateCodeUnknownIndicator     *fincen.ValidateIndicatorType `xml:"StateCodeUnknownIndicator,omitempty" json:",omitempty"`
	StreetAddressUnknownIndicator *fincen.ValidateIndicatorType `xml:"StreetAddressUnknownIndicator,omitempty" json:",omitempty"`
	ZIPCodeUnknownIndicator       *fincen.ValidateIndicatorType `xml:"ZIPCodeUnknownIndicator,omitempty" json:",omitempty"`
}

func (r AddressType) fieldInclusion(typeCode string) error {
	if r.CityUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("CityUnknownIndicator")
	}

	if r.CountryCodeUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("CountryCodeUnknownIndicator")
	}

	if r.RawCityText != nil && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawCityText")
	}

	if r.RawCountryCodeText != nil && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawCountryCodeText")
	}

	if r.RawStateCodeText != nil && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawStateCodeText")
	}

	if r.RawStreetAddress1Text != nil && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawStreetAddress1Text")
	}

	if r.RawZIPCode != nil && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("RawZIPCode")
	}

	if r.StateCodeUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("StateCodeUnknownIndicator")
	}

	if r.StreetAddressUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("StreetAddressUnknownIndicator")
	}

	if r.ZIPCodeUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("ZIPCodeUnknownIndicator")
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
	XMLName                  xml.Name                 `xml:"PhoneNumber"`
	SeqNum                   fincen.SeqNumber         `xml:"SeqNum,attr"`
	PhoneNumberExtensionText *fincen.RestrictString6  `xml:"PhoneNumberExtensionText,omitempty" json:",omitempty"`
	PhoneNumberText          *fincen.RestrictString16 `xml:"PhoneNumberText,omitempty" json:",omitempty"`
}

func (r PhoneNumberType) fieldInclusion(typeCode string) error {
	if r.PhoneNumberExtensionText != nil && !checkInvolved(typeCode, "8", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("PhoneNumberExtensionText")
	}

	if r.PhoneNumberText != nil && !checkInvolved(typeCode, "35", "8", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("PhoneNumberText")
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
	XMLName                               xml.Name                             `xml:"PartyIdentification"`
	SeqNum                                fincen.SeqNumber                     `xml:"SeqNum,attr"`
	IdentificationPresentUnknownIndicator *fincen.ValidateIndicatorType        `xml:"IdentificationPresentUnknownIndicator,omitempty" json:",omitempty"`
	OtherIssuerCountryText                *fincen.RestrictString2              `xml:"OtherIssuerCountryText,omitempty" json:",omitempty"`
	OtherIssuerStateText                  *fincen.RestrictString3              `xml:"OtherIssuerStateText,omitempty" json:",omitempty"`
	OtherPartyIdentificationTypeText      *fincen.RestrictString50             `xml:"OtherPartyIdentificationTypeText,omitempty" json:",omitempty"`
	PartyIdentificationNumberText         *fincen.RestrictString25             `xml:"PartyIdentificationNumberText,omitempty" json:",omitempty"`
	PartyIdentificationTypeCode           *ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode,omitempty" json:",omitempty"`
	TINUnknownIndicator                   *fincen.ValidateIndicatorType        `xml:"TINUnknownIndicator,omitempty" json:",omitempty"`
}

func (r PartyIdentificationType) fieldInclusion(typeCode string) error {
	if r.IdentificationPresentUnknownIndicator != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("IdentificationPresentUnknownIndicator")
	}

	if r.OtherIssuerCountryText != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("OtherIssuerCountryText")
	}

	if r.OtherIssuerStateText != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("OtherIssuerStateText")
	}

	if r.OtherPartyIdentificationTypeText != nil && !checkInvolved(typeCode, "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("OtherPartyIdentificationTypeText")
	}

	if r.PartyIdentificationNumberText != nil && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("PartyIdentificationNumberText")
	}

	if r.PartyIdentificationTypeCode != nil && !checkInvolved(typeCode, "35", "30", "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("PartyIdentificationTypeCode")
	}

	if r.TINUnknownIndicator != nil && !checkInvolved(typeCode, "34", "50", "17", "23", "58") {
		return fincen.NewErrFiledNotAssociated("TINUnknownIndicator")
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

	if r.PartyIdentificationTypeCode != nil {
		code := string(*r.PartyIdentificationTypeCode)
		switch code {
		case "4", "28":
			if !checkInvolved(typeCode, "35") {
				return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
			}
		case "2":
			if !checkInvolved(typeCode, "30", "34", "50", "17", "23", "58") {
				return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
			}
		case "1", "9", "5", "6", "7", "999":
			if !checkInvolved(typeCode, "50", "17", "23", "58") {
				return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
			}
		case "10", "11", "12", "13", "14":
			if !checkInvolved(typeCode, "30", "34") {
				return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
			}
		}
	}

	return fincen.Validate(&r, args...)
}

type OrganizationClassificationTypeSubtypeType struct {
	XMLName                      xml.Name                                       `xml:"OrganizationClassificationTypeSubtype"`
	SeqNum                       fincen.SeqNumber                               `xml:"SeqNum,attr"`
	OrganizationSubtypeID        *fincen.ValidateOrganizationSubtypeCodeCtrType `xml:"OrganizationSubtypeID,omitempty" json:",omitempty"`
	OrganizationTypeID           fincen.ValidateOrganizationCodeType            `xml:"OrganizationTypeID"`
	OtherOrganizationSubTypeText *fincen.RestrictString50                       `xml:"OtherOrganizationSubTypeText,omitempty" json:",omitempty"`
	OtherOrganizationTypeText    *fincen.RestrictString50                       `xml:"OtherOrganizationTypeText,omitempty" json:",omitempty"`
}

func (r OrganizationClassificationTypeSubtypeType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyOccupationBusinessType struct {
	XMLName                xml.Name                 `xml:"PartyOccupationBusiness"`
	SeqNum                 fincen.SeqNumber         `xml:"SeqNum,attr"`
	NAICSCode              *fincen.RestrictString6  `xml:"NAICSCode,omitempty" json:",omitempty"`
	OccupationBusinessText *fincen.RestrictString50 `xml:"OccupationBusinessText,omitempty" json:",omitempty"`
}

func (r PartyOccupationBusinessType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ElectronicAddressType struct {
	XMLName               xml.Name                  `xml:"ElectronicAddress"`
	SeqNum                fincen.SeqNumber          `xml:"SeqNum,attr"`
	ElectronicAddressText *fincen.RestrictString517 `xml:"ElectronicAddressText,omitempty" json:",omitempty"`
}

func (r ElectronicAddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AccountType struct {
	XMLName                 xml.Name                    `xml:"Account"`
	SeqNum                  fincen.SeqNumber            `xml:"SeqNum,attr"`
	AccountNumberText       *fincen.RestrictString40    `xml:"AccountNumberText,omitempty" json:",omitempty"`
	PartyAccountAssociation PartyAccountAssociationType `xml:"PartyAccountAssociation"`
}

func (r AccountType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAccountAssociationType struct {
	XMLName                         xml.Name                                `xml:"PartyAccountAssociation"`
	SeqNum                          fincen.SeqNumber                        `xml:"SeqNum,attr"`
	PartyAccountAssociationTypeCode ValidatePartyAccountAssociationCodeType `xml:"PartyAccountAssociationTypeCode"`
}

func (r PartyAccountAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityType struct {
	XMLName                           xml.Name                                `xml:"CurrencyTransactionActivity"`
	SeqNum                            fincen.SeqNumber                        `xml:"SeqNum,attr"`
	AggregateTransactionIndicator     fincen.ValidateIndicatorType            `xml:"AggregateTransactionIndicator"`
	ArmoredCarServiceIndicator        fincen.ValidateIndicatorType            `xml:"ArmoredCarServiceIndicator"`
	ATMIndicator                      fincen.ValidateIndicatorType            `xml:"ATMIndicator"`
	MailDepositShipmentIndicator      fincen.ValidateIndicatorType            `xml:"MailDepositShipmentIndicator"`
	NightDepositIndicator             fincen.ValidateIndicatorType            `xml:"NightDepositIndicator"`
	SharedBranchingIndicator          fincen.ValidateIndicatorType            `xml:"SharedBranchingIndicator"`
	TotalCashInReceiveAmountText      fincen.RestrictString15                 `xml:"TotalCashInReceiveAmountText"`
	TotalCashOutAmountText            fincen.RestrictString15                 `xml:"TotalCashOutAmountText"`
	TransactionDateText               fincen.DateYYYYMMDDType                 `xml:"TransactionDateText"`
	CurrencyTransactionActivityDetail []CurrencyTransactionActivityDetailType `xml:"CurrencyTransactionActivityDetail"`
}

func (r CurrencyTransactionActivityType) fieldInclusion() error {
	if len(r.CurrencyTransactionActivityDetail) < 1 || len(r.CurrencyTransactionActivityDetail) > 219 {
		return fincen.NewErrMinMaxRange("CurrencyTransactionActivity")
	}
	return nil
}

func (r CurrencyTransactionActivityType) Validate(args ...string) error {
	if err := r.fieldInclusion(); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityDetailType struct {
	XMLName                                    xml.Name                                                 `xml:"CurrencyTransactionActivityDetail"`
	SeqNum                                     fincen.SeqNumber                                         `xml:"SeqNum,attr"`
	CurrencyTransactionActivityDetailTypeCode  fincen.ValidateCurrencyTransactionActivityDetailCodeType `xml:"CurrencyTransactionActivityDetailTypeCode"`
	DetailTransactionAmountText                fincen.RestrictString15                                  `xml:"DetailTransactionAmountText"`
	OtherCurrencyTransactionActivityDetailText fincen.RestrictString50                                  `xml:"OtherCurrencyTransactionActivityDetailText"`
	OtherForeignCurrencyCountryText            fincen.RestrictString2                                   `xml:"OtherForeignCurrencyCountryText"`
}

func (r CurrencyTransactionActivityDetailType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
