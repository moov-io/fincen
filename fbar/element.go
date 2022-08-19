// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fbar

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

type EFilingBatchXML struct {
	XMLName                xml.Name `xml:"EFilingBatchXML"`
	FormTypeCode           string   `xml:"FormTypeCode"`
	Activity               []string `xml:"Activity"`
	ActivityCount          int64    `xml:"ActivityCount,attr"`
	PartyCount             int64    `xml:"PartyCount,attr"`
	AccountCount           int64    `xml:"AccountCount,attr"`
	JointlyOwnedOwnerCount int64    `xml:"JointlyOwnedOwnerCount,attr"`
	NoFIOwnerCount         int64    `xml:"NoFIOwnerCount,attr"`
	ConsolidatedOwnerCount int64    `xml:"ConsolidatedOwnerCount,attr"`
}

func (r EFilingBatchXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Account struct {
	AccountMaximumValueAmountText fincen.RestrictString15                   `xml:"AccountMaximumValueAmountText,omitempty"`
	AccountNumberText             fincen.RestrictString40                   `xml:"AccountNumberText,omitempty"`
	AccountTypeCode               fincen.ValidateAccountTypeCodeType        `xml:"AccountTypeCode,omitempty"`
	EFilingAccountTypeCode        fincen.ValidateEFilingAccountTypeCodeType `xml:"EFilingAccountTypeCode"`
	JointOwnerQuantityText        fincen.RestrictString3                    `xml:"JointOwnerQuantityText,omitempty"`
	OtherAccountTypeText          fincen.RestrictString50                   `xml:"OtherAccountTypeText,omitempty"`
	UnknownMaximumValueIndicator  fincen.ValidateIndicatorType              `xml:"UnknownMaximumValueIndicator,omitempty"`
	Party                         []Anon1                                   `xml:"Party"`
	SeqNum                        int64                                     `xml:"SeqNum,attr"`
}

func (r Account) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AccountType struct {
	AccountMaximumValueAmountText fincen.RestrictString15                   `xml:"AccountMaximumValueAmountText,omitempty"`
	AccountNumberText             fincen.RestrictString40                   `xml:"AccountNumberText,omitempty"`
	AccountTypeCode               fincen.ValidateAccountTypeCodeType        `xml:"AccountTypeCode,omitempty"`
	EFilingAccountTypeCode        fincen.ValidateEFilingAccountTypeCodeType `xml:"EFilingAccountTypeCode"`
	JointOwnerQuantityText        fincen.RestrictString3                    `xml:"JointOwnerQuantityText,omitempty"`
	OtherAccountTypeText          fincen.RestrictString50                   `xml:"OtherAccountTypeText,omitempty"`
	UnknownMaximumValueIndicator  fincen.ValidateIndicatorType              `xml:"UnknownMaximumValueIndicator,omitempty"`
	SeqNum                        int64                                     `xml:"SeqNum,attr"`
}

func (r AccountType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Activity struct {
	ApprovalOfficialSignatureDateText fincen.DateYYYYMMDDType      `xml:"ApprovalOfficialSignatureDateText"`
	EFilingPriorDocumentNumber        int64                        `xml:"EFilingPriorDocumentNumber,omitempty"`
	PreparerFilingSignatureIndicator  fincen.ValidateIndicatorType `xml:"PreparerFilingSignatureIndicator,omitempty"`
	ThirdPartyPreparerIndicator       fincen.ValidateIndicatorType `xml:"ThirdPartyPreparerIndicator,omitempty"`
	ActivityAssociation               ActivityAssociationType      `xml:"ActivityAssociation"`
	Party                             []string                     `xml:"Party"`
	Account                           []string                     `xml:"Account,omitempty"`
	ForeignAccountActivity            string                       `xml:"ForeignAccountActivity"`
	ActivityNarrativeInformation      string                       `xml:"ActivityNarrativeInformation,omitempty"`
	SeqNum                            int64                        `xml:"SeqNum,attr"`
}

func (r Activity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	CorrectsAmendsPriorReportIndicator fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator"`
	SeqNum                             int64                        `xml:"SeqNum,attr"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityNarrativeInformation struct {
	ActivityNarrativeSequenceNumber int                       `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString4000 `xml:"ActivityNarrativeText"`
	SeqNum                          int64                     `xml:"SeqNum,attr"`
}

func (r ActivityNarrativeInformation) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityNarrativeInformationType struct {
	ActivityNarrativeSequenceNumber int                       `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString4000 `xml:"ActivityNarrativeText"`
	SeqNum                          int64                     `xml:"SeqNum,attr"`
}

func (r ActivityNarrativeInformationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityType struct {
	ApprovalOfficialSignatureDateText fincen.DateYYYYMMDDType      `xml:"ApprovalOfficialSignatureDateText"`
	EFilingPriorDocumentNumber        int64                        `xml:"EFilingPriorDocumentNumber,omitempty"`
	PreparerFilingSignatureIndicator  fincen.ValidateIndicatorType `xml:"PreparerFilingSignatureIndicator,omitempty"`
	ThirdPartyPreparerIndicator       fincen.ValidateIndicatorType `xml:"ThirdPartyPreparerIndicator,omitempty"`
	SeqNum                            int64                        `xml:"SeqNum,attr"`
}

func (r ActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AddressType struct {
	RawCityText           fincen.RestrictString50  `xml:"RawCityText,omitempty"`
	RawCountryCodeText    fincen.RestrictString2   `xml:"RawCountryCodeText,omitempty"`
	RawStateCodeText      fincen.RestrictString3   `xml:"RawStateCodeText,omitempty"`
	RawStreetAddress1Text fincen.RestrictString100 `xml:"RawStreetAddress1Text,omitempty"`
	RawZIPCode            fincen.RestrictString9   `xml:"RawZIPCode,omitempty"`
	SeqNum                int64                    `xml:"SeqNum,attr"`
}

func (r AddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Anon1 struct {
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType  `xml:"ActivityPartyTypeCode"`
	FilerFinancialInterest25ForeignAccountIndicator fincen.ValidateIndicatorYNType `xml:"FilerFinancialInterest25ForeignAccountIndicator,omitempty"`
	FilerTypeConsolidatedIndicator                  fincen.ValidateIndicatorType   `xml:"FilerTypeConsolidatedIndicator,omitempty"`
	FilerTypeCorporationIndicator                   fincen.ValidateIndicatorType   `xml:"FilerTypeCorporationIndicator,omitempty"`
	FilerTypeFiduciaryOtherIndicator                fincen.ValidateIndicatorType   `xml:"FilerTypeFiduciaryOtherIndicator,omitempty"`
	FilerTypeIndividualIndicator                    fincen.ValidateIndicatorType   `xml:"FilerTypeIndividualIndicator,omitempty"`
	FilerTypeOtherText                              fincen.RestrictString50        `xml:"FilerTypeOtherText,omitempty"`
	FilerTypePartnershipIndicator                   fincen.ValidateIndicatorType   `xml:"FilerTypePartnershipIndicator,omitempty"`
	IndividualBirthDateText                         fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty"`
	PartyAsEntityOrganizationIndicator              fincen.ValidateIndicatorType   `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	SelfEmployedIndicator                           fincen.ValidateIndicatorType   `xml:"SelfEmployedIndicator,omitempty"`
	SignatureAuthoritiesIndicator                   fincen.ValidateIndicatorYNType `xml:"SignatureAuthoritiesIndicator,omitempty"`
	PartyName                                       PartyNameType                  `xml:"PartyName"`
	Address                                         AddressType                    `xml:"Address"`
	PartyIdentification                             PartyIdentificationType        `xml:"PartyIdentification,omitempty"`
	SeqNum                                          int64                          `xml:"SeqNum,attr"`
}

func (r Anon1) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ForeignAccountActivity struct {
	ForeignAccountHeldQuantityText   fincen.RestrictString4                  `xml:"ForeignAccountHeldQuantityText,omitempty"`
	LateFilingReasonCode             fincen.ValidateLateFilingReasonCodeType `xml:"LateFilingReasonCode,omitempty"`
	ReportCalendarYearText           fincen.DateYYYYType                     `xml:"ReportCalendarYearText"`
	SignatureAuthoritiesQuantityText fincen.RestrictString4                  `xml:"SignatureAuthoritiesQuantityText,omitempty"`
	SeqNum                           int64                                   `xml:"SeqNum,attr"`
}

func (r ForeignAccountActivity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ForeignAccountActivityType struct {
	ForeignAccountHeldQuantityText   fincen.RestrictString4                  `xml:"ForeignAccountHeldQuantityText,omitempty"`
	LateFilingReasonCode             fincen.ValidateLateFilingReasonCodeType `xml:"LateFilingReasonCode,omitempty"`
	ReportCalendarYearText           fincen.DateYYYYType                     `xml:"ReportCalendarYearText"`
	SignatureAuthoritiesQuantityText fincen.RestrictString4                  `xml:"SignatureAuthoritiesQuantityText,omitempty"`
	SeqNum                           int64                                   `xml:"SeqNum,attr"`
}

func (r ForeignAccountActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Party struct {
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType  `xml:"ActivityPartyTypeCode"`
	FilerFinancialInterest25ForeignAccountIndicator fincen.ValidateIndicatorYNType `xml:"FilerFinancialInterest25ForeignAccountIndicator,omitempty"`
	FilerTypeConsolidatedIndicator                  fincen.ValidateIndicatorType   `xml:"FilerTypeConsolidatedIndicator,omitempty"`
	FilerTypeCorporationIndicator                   fincen.ValidateIndicatorType   `xml:"FilerTypeCorporationIndicator,omitempty"`
	FilerTypeFiduciaryOtherIndicator                fincen.ValidateIndicatorType   `xml:"FilerTypeFiduciaryOtherIndicator,omitempty"`
	FilerTypeIndividualIndicator                    fincen.ValidateIndicatorType   `xml:"FilerTypeIndividualIndicator,omitempty"`
	FilerTypeOtherText                              fincen.RestrictString50        `xml:"FilerTypeOtherText,omitempty"`
	FilerTypePartnershipIndicator                   fincen.ValidateIndicatorType   `xml:"FilerTypePartnershipIndicator,omitempty"`
	IndividualBirthDateText                         fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty"`
	PartyAsEntityOrganizationIndicator              fincen.ValidateIndicatorType   `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	SelfEmployedIndicator                           fincen.ValidateIndicatorType   `xml:"SelfEmployedIndicator,omitempty"`
	SignatureAuthoritiesIndicator                   fincen.ValidateIndicatorYNType `xml:"SignatureAuthoritiesIndicator,omitempty"`
	PartyName                                       PartyNameType                  `xml:"PartyName"`
	Address                                         AddressType                    `xml:"Address,omitempty"`
	PhoneNumber                                     PhoneNumberType                `xml:"PhoneNumber,omitempty"`
	PartyIdentification                             []PartyIdentificationType      `xml:"PartyIdentification,omitempty"`
	SeqNum                                          int64                          `xml:"SeqNum,attr"`
}

func (r Party) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyIdentificationType struct {
	OtherIssuerCountryText           fincen.RestrictString2              `xml:"OtherIssuerCountryText,omitempty"`
	OtherPartyIdentificationTypeText fincen.RestrictString50             `xml:"OtherPartyIdentificationTypeText,omitempty"`
	PartyIdentificationNumberText    fincen.RestrictString25             `xml:"PartyIdentificationNumberText"`
	PartyIdentificationTypeCode      ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode"`
	SeqNum                           int64                               `xml:"SeqNum,attr"`
}

func (r PartyIdentificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyNameType struct {
	PartyNameTypeCode           ValidateActivityPartyCodeType `xml:"PartyNameTypeCode"`
	RawEntityIndividualLastName fincen.RestrictString150      `xml:"RawEntityIndividualLastName,omitempty"`
	RawIndividualFirstName      fincen.RestrictString35       `xml:"RawIndividualFirstName,omitempty"`
	RawIndividualMiddleName     fincen.RestrictString35       `xml:"RawIndividualMiddleName,omitempty"`
	RawIndividualNameSuffixText fincen.RestrictString35       `xml:"RawIndividualNameSuffixText,omitempty"`
	RawIndividualTitleText      fincen.RestrictString20       `xml:"RawIndividualTitleText,omitempty"`
	RawPartyFullName            fincen.RestrictString150      `xml:"RawPartyFullName,omitempty"`
	SeqNum                      int64                         `xml:"SeqNum,attr"`
}

func (r PartyNameType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyType struct {
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType  `xml:"ActivityPartyTypeCode"`
	FilerFinancialInterest25ForeignAccountIndicator fincen.ValidateIndicatorYNType `xml:"FilerFinancialInterest25ForeignAccountIndicator,omitempty"`
	FilerTypeConsolidatedIndicator                  fincen.ValidateIndicatorType   `xml:"FilerTypeConsolidatedIndicator,omitempty"`
	FilerTypeCorporationIndicator                   fincen.ValidateIndicatorType   `xml:"FilerTypeCorporationIndicator,omitempty"`
	FilerTypeFiduciaryOtherIndicator                fincen.ValidateIndicatorType   `xml:"FilerTypeFiduciaryOtherIndicator,omitempty"`
	FilerTypeIndividualIndicator                    fincen.ValidateIndicatorType   `xml:"FilerTypeIndividualIndicator,omitempty"`
	FilerTypeOtherText                              fincen.RestrictString50        `xml:"FilerTypeOtherText,omitempty"`
	FilerTypePartnershipIndicator                   fincen.ValidateIndicatorType   `xml:"FilerTypePartnershipIndicator,omitempty"`
	IndividualBirthDateText                         fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty"`
	PartyAsEntityOrganizationIndicator              fincen.ValidateIndicatorType   `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	SelfEmployedIndicator                           fincen.ValidateIndicatorType   `xml:"SelfEmployedIndicator,omitempty"`
	SignatureAuthoritiesIndicator                   fincen.ValidateIndicatorYNType `xml:"SignatureAuthoritiesIndicator,omitempty"`
	SeqNum                                          int64                          `xml:"SeqNum,attr"`
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
