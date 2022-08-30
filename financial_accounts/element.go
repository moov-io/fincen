// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package financial_accounts

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

type ActivityType struct {
	XMLName                           xml.Name                      `xml:"Activity"`
	SeqNum                            fincen.SeqNumber              `xml:"SeqNum,attr"`
	ApprovalOfficialSignatureDateText fincen.DateYYYYMMDDType       `xml:"ApprovalOfficialSignatureDateText"`
	EFilingPriorDocumentNumber        int64                         `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	PreparerFilingSignatureIndicator  *fincen.ValidateIndicatorType `xml:"PreparerFilingSignatureIndicator,omitempty" json:",omitempty"`
	ThirdPartyPreparerIndicator       *fincen.ValidateIndicatorType `xml:"ThirdPartyPreparerIndicator,omitempty" json:",omitempty"`
	ActivityAssociation               ActivityAssociationType       `xml:"ActivityAssociation"`
	Party                             []string                      `xml:"Party"`
	Account                           []string                      `xml:"Account,omitempty" json:",omitempty"`
	ForeignAccountActivity            string                        `xml:"ForeignAccountActivity"`
	ActivityNarrativeInformation      string                        `xml:"ActivityNarrativeInformation,omitempty" json:",omitempty"`
}

func (r ActivityType) FormTypeCode() string {
	return "FBARX"
}

func (r ActivityType) Validate(args ...string) error {
	if len(r.Party) < 4 || len(r.Party) > 104 {
		return fincen.NewErrValueInvalid("Party")
	}

	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	XMLName                            xml.Name                     `xml:"ActivityAssociation"`
	SeqNum                             fincen.SeqNumber             `xml:"SeqNum,attr"`
	CorrectsAmendsPriorReportIndicator fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Account struct {
	AccountMaximumValueAmountText *fincen.RestrictString15                  `xml:"AccountMaximumValueAmountText,omitempty" json:",omitempty"`
	AccountNumberText             *fincen.RestrictString40                  `xml:"AccountNumberText,omitempty" json:",omitempty"`
	AccountTypeCode               *fincen.ValidateAccountTypeCodeType       `xml:"AccountTypeCode,omitempty" json:",omitempty"`
	EFilingAccountTypeCode        fincen.ValidateEFilingAccountTypeCodeType `xml:"EFilingAccountTypeCode"`
	JointOwnerQuantityText        *fincen.RestrictString3                   `xml:"JointOwnerQuantityText,omitempty" json:",omitempty"`
	OtherAccountTypeText          *fincen.RestrictString50                  `xml:"OtherAccountTypeText,omitempty" json:",omitempty"`
	UnknownMaximumValueIndicator  *fincen.ValidateIndicatorType             `xml:"UnknownMaximumValueIndicator,omitempty" json:",omitempty"`
	Party                         []AccountParty                            `xml:"Party"`
	SeqNum                        int64                                     `xml:"SeqNum,attr"`
}

func (r Account) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AccountType struct {
	AccountMaximumValueAmountText *fincen.RestrictString15                  `xml:"AccountMaximumValueAmountText,omitempty" json:",omitempty"`
	AccountNumberText             *fincen.RestrictString40                  `xml:"AccountNumberText,omitempty" json:",omitempty"`
	AccountTypeCode               *fincen.ValidateAccountTypeCodeType       `xml:"AccountTypeCode,omitempty" json:",omitempty"`
	EFilingAccountTypeCode        fincen.ValidateEFilingAccountTypeCodeType `xml:"EFilingAccountTypeCode"`
	JointOwnerQuantityText        *fincen.RestrictString3                   `xml:"JointOwnerQuantityText,omitempty" json:",omitempty"`
	OtherAccountTypeText          *fincen.RestrictString50                  `xml:"OtherAccountTypeText,omitempty" json:",omitempty"`
	UnknownMaximumValueIndicator  *fincen.ValidateIndicatorType             `xml:"UnknownMaximumValueIndicator,omitempty" json:",omitempty"`
	SeqNum                        int64                                     `xml:"SeqNum,attr"`
}

func (r AccountType) Validate(args ...string) error {
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

type AddressType struct {
	RawCityText           *fincen.RestrictString50  `xml:"RawCityText,omitempty" json:",omitempty"`
	RawCountryCodeText    *fincen.RestrictString2   `xml:"RawCountryCodeText,omitempty" json:",omitempty"`
	RawStateCodeText      *fincen.RestrictString3   `xml:"RawStateCodeText,omitempty" json:",omitempty"`
	RawStreetAddress1Text *fincen.RestrictString100 `xml:"RawStreetAddress1Text,omitempty" json:",omitempty"`
	RawZIPCode            *fincen.RestrictString9   `xml:"RawZIPCode,omitempty" json:",omitempty"`
	SeqNum                int64                     `xml:"SeqNum,attr"`
}

func (r AddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AccountParty struct {
	XMLName                                         xml.Name                        `xml:"Party"`
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType   `xml:"ActivityPartyTypeCode"`
	FilerFinancialInterest25ForeignAccountIndicator *fincen.ValidateIndicatorYNType `xml:"FilerFinancialInterest25ForeignAccountIndicator,omitempty" json:",omitempty"`
	FilerTypeConsolidatedIndicator                  *fincen.ValidateIndicatorType   `xml:"FilerTypeConsolidatedIndicator,omitempty" json:",omitempty"`
	FilerTypeCorporationIndicator                   *fincen.ValidateIndicatorType   `xml:"FilerTypeCorporationIndicator,omitempty" json:",omitempty"`
	FilerTypeFiduciaryOtherIndicator                *fincen.ValidateIndicatorType   `xml:"FilerTypeFiduciaryOtherIndicator,omitempty" json:",omitempty"`
	FilerTypeIndividualIndicator                    *fincen.ValidateIndicatorType   `xml:"FilerTypeIndividualIndicator,omitempty" json:",omitempty"`
	FilerTypeOtherText                              *fincen.RestrictString50        `xml:"FilerTypeOtherText,omitempty" json:",omitempty"`
	FilerTypePartnershipIndicator                   *fincen.ValidateIndicatorType   `xml:"FilerTypePartnershipIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                         *fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator              *fincen.ValidateIndicatorType   `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	SelfEmployedIndicator                           *fincen.ValidateIndicatorType   `xml:"SelfEmployedIndicator,omitempty" json:",omitempty"`
	SignatureAuthoritiesIndicator                   *fincen.ValidateIndicatorYNType `xml:"SignatureAuthoritiesIndicator,omitempty" json:",omitempty"`
	PartyName                                       PartyNameType                   `xml:"PartyName"`
	Address                                         AddressType                     `xml:"Address"`
	PartyIdentification                             *PartyIdentificationType        `xml:"PartyIdentification,omitempty" json:",omitempty"`
	SeqNum                                          int64                           `xml:"SeqNum,attr"`
}

func (r AccountParty) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ForeignAccountActivity struct {
	ForeignAccountHeldQuantityText   *fincen.RestrictString4                  `xml:"ForeignAccountHeldQuantityText,omitempty" json:",omitempty"`
	LateFilingReasonCode             *fincen.ValidateLateFilingReasonCodeType `xml:"LateFilingReasonCode,omitempty" json:",omitempty"`
	ReportCalendarYearText           fincen.DateYYYYType                      `xml:"ReportCalendarYearText"`
	SignatureAuthoritiesQuantityText *fincen.RestrictString4                  `xml:"SignatureAuthoritiesQuantityText,omitempty" json:",omitempty"`
	SeqNum                           int64                                    `xml:"SeqNum,attr"`
}

func (r ForeignAccountActivity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ForeignAccountActivityType struct {
	ForeignAccountHeldQuantityText   *fincen.RestrictString4                  `xml:"ForeignAccountHeldQuantityText,omitempty" json:",omitempty"`
	LateFilingReasonCode             *fincen.ValidateLateFilingReasonCodeType `xml:"LateFilingReasonCode,omitempty" json:",omitempty"`
	ReportCalendarYearText           fincen.DateYYYYType                      `xml:"ReportCalendarYearText"`
	SignatureAuthoritiesQuantityText *fincen.RestrictString4                  `xml:"SignatureAuthoritiesQuantityText,omitempty" json:",omitempty"`
	SeqNum                           int64                                    `xml:"SeqNum,attr"`
}

func (r ForeignAccountActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Party struct {
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType   `xml:"ActivityPartyTypeCode"`
	FilerFinancialInterest25ForeignAccountIndicator *fincen.ValidateIndicatorYNType `xml:"FilerFinancialInterest25ForeignAccountIndicator,omitempty" json:",omitempty"`
	FilerTypeConsolidatedIndicator                  *fincen.ValidateIndicatorType   `xml:"FilerTypeConsolidatedIndicator,omitempty" json:",omitempty"`
	FilerTypeCorporationIndicator                   *fincen.ValidateIndicatorType   `xml:"FilerTypeCorporationIndicator,omitempty" json:",omitempty"`
	FilerTypeFiduciaryOtherIndicator                *fincen.ValidateIndicatorType   `xml:"FilerTypeFiduciaryOtherIndicator,omitempty" json:",omitempty"`
	FilerTypeIndividualIndicator                    *fincen.ValidateIndicatorType   `xml:"FilerTypeIndividualIndicator,omitempty" json:",omitempty"`
	FilerTypeOtherText                              *fincen.RestrictString50        `xml:"FilerTypeOtherText,omitempty" json:",omitempty"`
	FilerTypePartnershipIndicator                   *fincen.ValidateIndicatorType   `xml:"FilerTypePartnershipIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                         *fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator              *fincen.ValidateIndicatorType   `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	SelfEmployedIndicator                           *fincen.ValidateIndicatorType   `xml:"SelfEmployedIndicator,omitempty" json:",omitempty"`
	SignatureAuthoritiesIndicator                   *fincen.ValidateIndicatorYNType `xml:"SignatureAuthoritiesIndicator,omitempty" json:",omitempty"`
	PartyName                                       PartyNameType                   `xml:"PartyName"`
	Address                                         *AddressType                    `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber                                     *PhoneNumberType                `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification                             []PartyIdentificationType       `xml:"PartyIdentification,omitempty" json:",omitempty"`
	SeqNum                                          int64                           `xml:"SeqNum,attr"`
}

func (r Party) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyIdentificationType struct {
	OtherIssuerCountryText           *fincen.RestrictString2             `xml:"OtherIssuerCountryText,omitempty" json:",omitempty"`
	OtherPartyIdentificationTypeText *fincen.RestrictString50            `xml:"OtherPartyIdentificationTypeText,omitempty" json:",omitempty"`
	PartyIdentificationNumberText    fincen.RestrictString25             `xml:"PartyIdentificationNumberText"`
	PartyIdentificationTypeCode      ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode"`
	SeqNum                           int64                               `xml:"SeqNum,attr"`
}

func (r PartyIdentificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyNameType struct {
	PartyNameTypeCode           ValidateActivityPartyCodeType `xml:"PartyNameTypeCode"`
	RawEntityIndividualLastName *fincen.RestrictString150     `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName      *fincen.RestrictString35      `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName     *fincen.RestrictString35      `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualNameSuffixText *fincen.RestrictString35      `xml:"RawIndividualNameSuffixText,omitempty" json:",omitempty"`
	RawIndividualTitleText      *fincen.RestrictString20      `xml:"RawIndividualTitleText,omitempty" json:",omitempty"`
	RawPartyFullName            *fincen.RestrictString150     `xml:"RawPartyFullName,omitempty" json:",omitempty"`
	SeqNum                      int64                         `xml:"SeqNum,attr"`
}

func (r PartyNameType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyType struct {
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType   `xml:"ActivityPartyTypeCode"`
	FilerFinancialInterest25ForeignAccountIndicator *fincen.ValidateIndicatorYNType `xml:"FilerFinancialInterest25ForeignAccountIndicator,omitempty" json:",omitempty"`
	FilerTypeConsolidatedIndicator                  *fincen.ValidateIndicatorType   `xml:"FilerTypeConsolidatedIndicator,omitempty" json:",omitempty"`
	FilerTypeCorporationIndicator                   *fincen.ValidateIndicatorType   `xml:"FilerTypeCorporationIndicator,omitempty" json:",omitempty"`
	FilerTypeFiduciaryOtherIndicator                *fincen.ValidateIndicatorType   `xml:"FilerTypeFiduciaryOtherIndicator,omitempty" json:",omitempty"`
	FilerTypeIndividualIndicator                    *fincen.ValidateIndicatorType   `xml:"FilerTypeIndividualIndicator,omitempty" json:",omitempty"`
	FilerTypeOtherText                              *fincen.RestrictString50        `xml:"FilerTypeOtherText,omitempty" json:",omitempty"`
	FilerTypePartnershipIndicator                   *fincen.ValidateIndicatorType   `xml:"FilerTypePartnershipIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                         *fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator              *fincen.ValidateIndicatorType   `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	SelfEmployedIndicator                           *fincen.ValidateIndicatorType   `xml:"SelfEmployedIndicator,omitempty" json:",omitempty"`
	SignatureAuthoritiesIndicator                   *fincen.ValidateIndicatorYNType `xml:"SignatureAuthoritiesIndicator,omitempty" json:",omitempty"`
	SeqNum                                          int64                           `xml:"SeqNum,attr"`
}

func (r PartyType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PhoneNumberType struct {
	PhoneNumberExtensionText *fincen.RestrictString6  `xml:"PhoneNumberExtensionText,omitempty" json:",omitempty"`
	PhoneNumberText          *fincen.RestrictString16 `xml:"PhoneNumberText,omitempty" json:",omitempty"`
	SeqNum                   int64                    `xml:"SeqNum,attr"`
}

func (r PhoneNumberType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
