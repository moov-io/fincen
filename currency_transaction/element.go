// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package currency_transaction

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

type EFilingBatchXML struct {
	XMLName       xml.Name `xml:"EFilingBatchXML"`
	FormTypeCode  string   `xml:"FormTypeCode"`
	Activity      []string `xml:"Activity"`
	TotalAmount   float64  `xml:"TotalAmount,attr"`
	PartyCount    int64    `xml:"PartyCount,attr"`
	ActivityCount int64    `xml:"ActivityCount,attr"`
}

func (r EFilingBatchXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Account struct {
	AccountNumberText       fincen.RestrictString40     `xml:"AccountNumberText,omitempty"`
	PartyAccountAssociation PartyAccountAssociationType `xml:"PartyAccountAssociation,omitempty"`
	SeqNum                  int64                       `xml:"SeqNum,attr"`
}

func (r Account) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AccountType struct {
	AccountNumberText fincen.RestrictString40 `xml:"AccountNumberText,omitempty"`
	SeqNum            int64                   `xml:"SeqNum,attr"`
}

func (r AccountType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Activity struct {
	EFilingPriorDocumentNumber  int64                   `xml:"EFilingPriorDocumentNumber,omitempty"`
	FilingDateText              fincen.DateYYYYMMDDType `xml:"FilingDateText"`
	ActivityAssociation         ActivityAssociationType `xml:"ActivityAssociation"`
	Party                       []string                `xml:"Party"`
	CurrencyTransactionActivity string                  `xml:"CurrencyTransactionActivity"`
	SeqNum                      int64                   `xml:"SeqNum,attr"`
}

func (r Activity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	CorrectsAmendsPriorReportIndicator fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator"`
	FinCENDirectBackFileIndicator      fincen.ValidateIndicatorType `xml:"FinCENDirectBackFileIndicator"`
	InitialReportIndicator             fincen.ValidateIndicatorType `xml:"InitialReportIndicator"`
	SeqNum                             int64                        `xml:"SeqNum,attr"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityType struct {
	EFilingPriorDocumentNumber int64                   `xml:"EFilingPriorDocumentNumber,omitempty"`
	FilingDateText             fincen.DateYYYYMMDDType `xml:"FilingDateText"`
	SeqNum                     int64                   `xml:"SeqNum,attr"`
}

func (r ActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AddressType struct {
	CityUnknownIndicator          fincen.ValidateIndicatorType `xml:"CityUnknownIndicator,omitempty"`
	CountryCodeUnknownIndicator   fincen.ValidateIndicatorType `xml:"CountryCodeUnknownIndicator,omitempty"`
	RawCityText                   fincen.RestrictString50      `xml:"RawCityText,omitempty"`
	RawCountryCodeText            fincen.RestrictString2       `xml:"RawCountryCodeText,omitempty"`
	RawStateCodeText              fincen.RestrictString3       `xml:"RawStateCodeText,omitempty"`
	RawStreetAddress1Text         fincen.RestrictString100     `xml:"RawStreetAddress1Text,omitempty"`
	RawZIPCode                    fincen.RestrictString9       `xml:"RawZIPCode,omitempty"`
	StateCodeUnknownIndicator     fincen.ValidateIndicatorType `xml:"StateCodeUnknownIndicator,omitempty"`
	StreetAddressUnknownIndicator fincen.ValidateIndicatorType `xml:"StreetAddressUnknownIndicator,omitempty"`
	ZIPCodeUnknownIndicator       fincen.ValidateIndicatorType `xml:"ZIPCodeUnknownIndicator,omitempty"`
	SeqNum                        int64                        `xml:"SeqNum,attr"`
}

func (r AddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivity struct {
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
	SeqNum                            int64                                   `xml:"SeqNum,attr"`
}

func (r CurrencyTransactionActivity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityDetailType struct {
	CurrencyTransactionActivityDetailTypeCode  fincen.ValidateCurrencyTransactionActivityDetailCodeType `xml:"CurrencyTransactionActivityDetailTypeCode"`
	DetailTransactionAmountText                fincen.RestrictString15                                  `xml:"DetailTransactionAmountText"`
	OtherCurrencyTransactionActivityDetailText fincen.RestrictString50                                  `xml:"OtherCurrencyTransactionActivityDetailText"`
	OtherForeignCurrencyCountryText            fincen.RestrictString2                                   `xml:"OtherForeignCurrencyCountryText"`
	SeqNum                                     int64                                                    `xml:"SeqNum,attr"`
}

func (r CurrencyTransactionActivityDetailType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityType struct {
	AggregateTransactionIndicator fincen.ValidateIndicatorType `xml:"AggregateTransactionIndicator"`
	ArmoredCarServiceIndicator    fincen.ValidateIndicatorType `xml:"ArmoredCarServiceIndicator"`
	ATMIndicator                  fincen.ValidateIndicatorType `xml:"ATMIndicator"`
	MailDepositShipmentIndicator  fincen.ValidateIndicatorType `xml:"MailDepositShipmentIndicator"`
	NightDepositIndicator         fincen.ValidateIndicatorType `xml:"NightDepositIndicator"`
	SharedBranchingIndicator      fincen.ValidateIndicatorType `xml:"SharedBranchingIndicator"`
	TotalCashInReceiveAmountText  fincen.RestrictString15      `xml:"TotalCashInReceiveAmountText"`
	TotalCashOutAmountText        fincen.RestrictString15      `xml:"TotalCashOutAmountText"`
	TransactionDateText           fincen.DateYYYYMMDDType      `xml:"TransactionDateText"`
	SeqNum                        int64                        `xml:"SeqNum,attr"`
}

func (r CurrencyTransactionActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ElectronicAddressType struct {
	ElectronicAddressText fincen.RestrictString517 `xml:"ElectronicAddressText,omitempty"`
	SeqNum                int64                    `xml:"SeqNum,attr"`
}

func (r ElectronicAddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type OrganizationClassificationTypeSubtypeType struct {
	OrganizationSubtypeID        fincen.ValidateOrganizationSubtypeCodeCtrType `xml:"OrganizationSubtypeID,omitempty"`
	OrganizationTypeID           fincen.ValidateOrganizationCodeType           `xml:"OrganizationTypeID"`
	OtherOrganizationSubTypeText fincen.RestrictString50                       `xml:"OtherOrganizationSubTypeText,omitempty"`
	OtherOrganizationTypeText    fincen.RestrictString50                       `xml:"OtherOrganizationTypeText,omitempty"`
	SeqNum                       int64                                         `xml:"SeqNum,attr"`
}

func (r OrganizationClassificationTypeSubtypeType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Party struct {
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType             `xml:"ActivityPartyTypeCode"`
	BirthDateUnknownIndicator                       fincen.ValidateIndicatorType              `xml:"BirthDateUnknownIndicator,omitempty"`
	EFilingCoverageBeginningDateText                fincen.DateYYYYMMDDType                   `xml:"EFilingCoverageBeginningDateText,omitempty"`
	EFilingCoverageEndDateText                      fincen.DateYYYYMMDDType                   `xml:"EFilingCoverageEndDateText,omitempty"`
	FemaleGenderIndicator                           fincen.ValidateIndicatorType              `xml:"FemaleGenderIndicator,omitempty"`
	IndividualBirthDateText                         fincen.DateYYYYMMDDOrBlankTypeDOB         `xml:"IndividualBirthDateText,omitempty"`
	IndividualEntityCashInAmountText                fincen.RestrictString15                   `xml:"IndividualEntityCashInAmountText,omitempty"`
	IndividualEntityCashOutAmountText               fincen.RestrictString15                   `xml:"IndividualEntityCashOutAmountText,omitempty"`
	MaleGenderIndicator                             fincen.ValidateIndicatorType              `xml:"MaleGenderIndicator,omitempty"`
	MultipleTransactionsPersonsIndividualsIndicator fincen.ValidateIndicatorType              `xml:"MultipleTransactionsPersonsIndividualsIndicator,omitempty"`
	PartyAsEntityOrganizationIndicator              fincen.ValidateIndicatorType              `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	PrimaryRegulatorTypeCode                        ValidateFederalRegulatorCodeType          `xml:"PrimaryRegulatorTypeCode,omitempty"`
	UnknownGenderIndicator                          fincen.ValidateIndicatorType              `xml:"UnknownGenderIndicator,omitempty"`
	PartyName                                       []PartyNameType                           `xml:"PartyName"`
	Address                                         AddressType                               `xml:"Address,omitempty"`
	PhoneNumber                                     PhoneNumberType                           `xml:"PhoneNumber,omitempty"`
	PartyIdentification                             []PartyIdentificationType                 `xml:"PartyIdentification,omitempty"`
	OrganizationClassificationTypeSubtype           OrganizationClassificationTypeSubtypeType `xml:"OrganizationClassificationTypeSubtype,omitempty"`
	PartyOccupationBusiness                         PartyOccupationBusinessType               `xml:"PartyOccupationBusiness,omitempty"`
	ElectronicAddress                               ElectronicAddressType                     `xml:"ElectronicAddress,omitempty"`
	Account                                         []Account                                 `xml:"Account,omitempty"`
	SeqNum                                          int64                                     `xml:"SeqNum,attr"`
}

func (r Party) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAccountAssociationType struct {
	PartyAccountAssociationTypeCode fincen.ValidatePartyAccountAssociationCodeType `xml:"PartyAccountAssociationTypeCode,omitempty"`
	SeqNum                          int64                                          `xml:"SeqNum,attr"`
}

func (r PartyAccountAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyIdentificationType struct {
	IdentificationPresentUnknownIndicator fincen.ValidateIndicatorType        `xml:"IdentificationPresentUnknownIndicator,omitempty"`
	OtherIssuerCountryText                fincen.RestrictString2              `xml:"OtherIssuerCountryText,omitempty"`
	OtherIssuerStateText                  fincen.RestrictString3              `xml:"OtherIssuerStateText,omitempty"`
	OtherPartyIdentificationTypeText      fincen.RestrictString50             `xml:"OtherPartyIdentificationTypeText,omitempty"`
	PartyIdentificationNumberText         fincen.RestrictString25             `xml:"PartyIdentificationNumberText,omitempty"`
	PartyIdentificationTypeCode           ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode,omitempty"`
	TINUnknownIndicator                   fincen.ValidateIndicatorType        `xml:"TINUnknownIndicator,omitempty"`
	SeqNum                                int64                               `xml:"SeqNum,attr"`
}

func (r PartyIdentificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyNameType struct {
	EntityLastNameUnknownIndicator fincen.ValidateIndicatorType `xml:"EntityLastNameUnknownIndicator,omitempty"`
	FirstNameUnknownIndicator      fincen.ValidateIndicatorType `xml:"FirstNameUnknownIndicator,omitempty"`
	PartyNameTypeCode              ValidatePartyNameCodeType    `xml:"PartyNameTypeCode,omitempty"`
	RawEntityIndividualLastName    fincen.RestrictString150     `xml:"RawEntityIndividualLastName,omitempty"`
	RawIndividualFirstName         fincen.RestrictString35      `xml:"RawIndividualFirstName,omitempty"`
	RawIndividualMiddleName        fincen.RestrictString35      `xml:"RawIndividualMiddleName,omitempty"`
	RawIndividualNameSuffixText    fincen.RestrictString35      `xml:"RawIndividualNameSuffixText,omitempty"`
	RawPartyFullName               fincen.RestrictString150     `xml:"RawPartyFullName,omitempty"`
	SeqNum                         int64                        `xml:"SeqNum,attr"`
}

func (r PartyNameType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyOccupationBusinessType struct {
	NAICSCode              fincen.RestrictString6  `xml:"NAICSCode,omitempty"`
	OccupationBusinessText fincen.RestrictString50 `xml:"OccupationBusinessText,omitempty"`
	SeqNum                 int64                   `xml:"SeqNum,attr"`
}

func (r PartyOccupationBusinessType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyType struct {
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	BirthDateUnknownIndicator                       fincen.ValidateIndicatorType      `xml:"BirthDateUnknownIndicator,omitempty"`
	EFilingCoverageBeginningDateText                fincen.DateYYYYMMDDType           `xml:"EFilingCoverageBeginningDateText,omitempty"`
	EFilingCoverageEndDateText                      fincen.DateYYYYMMDDType           `xml:"EFilingCoverageEndDateText,omitempty"`
	FemaleGenderIndicator                           fincen.ValidateIndicatorType      `xml:"FemaleGenderIndicator,omitempty"`
	IndividualBirthDateText                         fincen.DateYYYYMMDDOrBlankTypeDOB `xml:"IndividualBirthDateText,omitempty"`
	IndividualEntityCashInAmountText                fincen.RestrictString15           `xml:"IndividualEntityCashInAmountText,omitempty"`
	IndividualEntityCashOutAmountText               fincen.RestrictString15           `xml:"IndividualEntityCashOutAmountText,omitempty"`
	MaleGenderIndicator                             fincen.ValidateIndicatorType      `xml:"MaleGenderIndicator,omitempty"`
	MultipleTransactionsPersonsIndividualsIndicator fincen.ValidateIndicatorType      `xml:"MultipleTransactionsPersonsIndividualsIndicator,omitempty"`
	PartyAsEntityOrganizationIndicator              fincen.ValidateIndicatorType      `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	PrimaryRegulatorTypeCode                        ValidateFederalRegulatorCodeType  `xml:"PrimaryRegulatorTypeCode,omitempty"`
	UnknownGenderIndicator                          fincen.ValidateIndicatorType      `xml:"UnknownGenderIndicator,omitempty"`
	SeqNum                                          int64                             `xml:"SeqNum,attr"`
}

func (r PartyType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PhoneNumberType struct {
	PhoneNumberExtensionText fincen.RestrictString6  `xml:"PhoneNumberExtensionText,omitempty"`
	PhoneNumberText          fincen.RestrictString16 `xml:"PhoneNumberText,omitempty"`
	SeqNum                   int64                   `xml:"SeqNum,attr"`
}

func (r PhoneNumberType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
