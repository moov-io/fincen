// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// FinCEN Report of Foreign Bank and Financial Accounts (FBAR) (FinCEN Report 114)

package financial_accounts

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

func NewActivity() *ActivityType {
	return &ActivityType{}
}

type ActivityType struct {
	XMLName                           xml.Name                          `xml:"Activity"`
	SeqNum                            fincen.SeqNumber                  `xml:"SeqNum,attr"`
	ApprovalOfficialSignatureDateText fincen.DateYYYYMMDDType           `xml:"ApprovalOfficialSignatureDateText"`
	EFilingPriorDocumentNumber        int64                             `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	PreparerFilingSignatureIndicator  *fincen.ValidateIndicatorType     `xml:"PreparerFilingSignatureIndicator,omitempty" json:",omitempty"`
	ThirdPartyPreparerIndicator       *fincen.ValidateIndicatorType     `xml:"ThirdPartyPreparerIndicator,omitempty" json:",omitempty"`
	ActivityAssociation               *ActivityAssociationType          `xml:"ActivityAssociation"`
	Party                             []*PartyType                      `xml:"Party"`
	Account                           []*AccountType                    `xml:"Account,omitempty" json:",omitempty"`
	ForeignAccountActivity            *ForeignAccountActivityType       `xml:"ForeignAccountActivity"`
	ActivityNarrativeInformation      *ActivityNarrativeInformationType `xml:"ActivityNarrativeInformation,omitempty" json:",omitempty"`
}

func (r ActivityType) FormTypeCode() string {
	return fincen.Report114
}

func (r ActivityType) TotalAmount() float64 {
	// Dummy
	return 0
}

func (r ActivityType) PartyCount(args ...string) int64 {
	var count int64
	for _, party := range r.Party {
		typeCode := string(party.ActivityPartyTypeCode)
		if fincen.CheckInvolved(typeCode, args...) {
			count++
		}
	}

	return count
}

func (r ActivityType) fieldInclusion() error {

	if len(r.Party) < 3 || len(r.Party) > 5 {
		return fincen.NewErrMinMaxRange("Party")
	}

	if len(r.Account) > 9999 {
		return fincen.NewErrMinMaxRange("Account")
	}

	if r.ActivityAssociation == nil {
		return fincen.NewErrFieldRequired("ActivityAssociation")
	}

	if r.ForeignAccountActivity == nil {
		return fincen.NewErrFieldRequired("ForeignAccountActivity")
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
		if cnt, ok := existed[typeCode]; ok {
			existed[typeCode] = cnt + 1
		} else {
			existed[typeCode] = 1
		}
	}

	if _, ok := existed["35"]; !ok {
		return fincen.NewErrFieldRequired("Party(type 35)")
	}
	if _, ok := existed["37"]; !ok {
		return fincen.NewErrFieldRequired("Party(type 37)")
	}
	if _, ok := existed["15"]; !ok {
		return fincen.NewErrFieldRequired("Party(type 15)")
	}
	if cnt, ok := existed["35"]; ok && cnt > 1 {
		return fincen.NewErrMinMaxRange("Party(type 35)")
	}
	if cnt, ok := existed["37"]; ok && cnt > 1 {
		return fincen.NewErrMinMaxRange("Party(type 37)")
	}
	if cnt, ok := existed["15"]; ok && cnt > 1 {
		return fincen.NewErrMinMaxRange("Party(type 15)")
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

type PartyType struct {
	XMLName                                         xml.Name                        `xml:"Party"`
	SeqNum                                          fincen.SeqNumber                `xml:"SeqNum,attr"`
	ActivityPartyTypeCode                           ValidateActivityPartyCodeType   `xml:"ActivityPartyTypeCode"`
	FilerFinancialInterest25ForeignAccountIndicator *fincen.ValidateIndicatorYNType `xml:"FilerFinancialInterest25ForeignAccountIndicator,omitempty" json:",omitempty"`
	FilerTypeConsolidatedIndicator                  *fincen.ValidateIndicatorType   `xml:"FilerTypeConsolidatedIndicator,omitempty" json:",omitempty"`
	FilerTypeCorporationIndicator                   *fincen.ValidateIndicatorType   `xml:"FilerTypeCorporationIndicator,omitempty" json:",omitempty"`
	FilerTypeFiduciaryOtherIndicator                *fincen.ValidateIndicatorType   `xml:"FilerTypeFiduciaryOtherIndicator,omitempty" json:",omitempty"`
	FilerTypeIndividualIndicator                    *fincen.ValidateIndicatorType   `xml:"FilerTypeIndividualIndicator,omitempty" json:",omitempty"`
	FilerTypeOtherText                              *fincen.RestrictString50        `xml:"FilerTypeOtherText,omitempty" json:",omitempty"`
	FilerTypePartnershipIndicator                   *fincen.ValidateIndicatorType   `xml:"FilerTypePartnershipIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                         *fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	SelfEmployedIndicator                           *fincen.ValidateIndicatorType   `xml:"SelfEmployedIndicator,omitempty" json:",omitempty"`
	SignatureAuthoritiesIndicator                   *fincen.ValidateIndicatorYNType `xml:"SignatureAuthoritiesIndicator,omitempty" json:",omitempty"`
	PartyName                                       *PartyNameType                  `xml:"PartyName"`
	Address                                         *AddressType                    `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber                                     *PhoneNumberType                `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification                             []*PartyIdentificationType      `xml:"PartyIdentification,omitempty" json:",omitempty"`
}

func (r PartyType) fieldInclusion() error {
	typeCode := string(r.ActivityPartyTypeCode)

	if typeCode == "15" {
		if r.FilerFinancialInterest25ForeignAccountIndicator == nil {
			return fincen.NewErrFieldRequired("FilerFinancialInterest25ForeignAccountIndicator")
		}
	}

	if typeCode != "15" && r.FilerTypeConsolidatedIndicator != nil {
		return fincen.NewErrFiledOmitted("FilerTypeConsolidatedIndicator")
	}

	if typeCode != "15" && r.FilerTypeCorporationIndicator != nil {
		return fincen.NewErrFiledOmitted("FilerTypeCorporationIndicator")
	}

	if typeCode != "15" && r.FilerTypeFiduciaryOtherIndicator != nil {
		return fincen.NewErrFiledOmitted("FilerTypeFiduciaryOtherIndicator")
	}

	if typeCode != "15" && r.FilerTypeIndividualIndicator != nil {
		return fincen.NewErrFiledOmitted("FilerTypeIndividualIndicator")
	}

	if typeCode != "15" && r.FilerTypePartnershipIndicator != nil {
		return fincen.NewErrFiledOmitted("FilerTypePartnershipIndicator")
	}

	if typeCode != "57" && r.SelfEmployedIndicator != nil {
		return fincen.NewErrFiledOmitted("SelfEmployedIndicator")
	}

	if typeCode != "15" && r.SignatureAuthoritiesIndicator != nil {
		return fincen.NewErrFiledOmitted("SignatureAuthoritiesIndicator")
	}

	if r.FilerTypeOtherText != nil && !(typeCode == "15" && r.FilerTypeFiduciaryOtherIndicator != nil) {
		return fincen.NewErrValueInvalid("FilerTypeOtherText")
	}

	if r.IndividualBirthDateText != nil && !(typeCode == "15" && r.FilerTypeIndividualIndicator != nil) {
		return fincen.NewErrValueInvalid("IndividualBirthDateText")
	}

	if r.Address == nil && fincen.CheckInvolved(typeCode, "35", "15", "57") {
		return fincen.NewErrFieldRequired("Address")
	}

	if r.PhoneNumber == nil && fincen.CheckInvolved(typeCode, "35", "57") {
		return fincen.NewErrFieldRequired("PhoneNumber")
	}

	if len(r.PartyIdentification) > 2 {
		return fincen.NewErrMinMaxRange("PartyIdentification")
	}

	if r.PartyName == nil {
		return fincen.NewErrFieldRequired("PartyName")
	}

	return nil
}

func (r PartyType) Validate(args ...string) error {
	if err := r.fieldInclusion(); err != nil {
		return err
	}

	if r.FilerTypeIndividualIndicator != nil {
		return fincen.Validate(&r, string(r.ActivityPartyTypeCode), "true")
	}

	return fincen.Validate(&r, string(r.ActivityPartyTypeCode))
}

type AccountType struct {
	XMLName                       xml.Name                                  `xml:"Account"`
	SeqNum                        fincen.SeqNumber                          `xml:"SeqNum,attr"`
	AccountMaximumValueAmountText *fincen.RestrictString15                  `xml:"AccountMaximumValueAmountText,omitempty" json:",omitempty"`
	AccountNumberText             *fincen.RestrictString40                  `xml:"AccountNumberText,omitempty" json:",omitempty"`
	AccountTypeCode               *fincen.ValidateAccountTypeCodeType       `xml:"AccountTypeCode,omitempty" json:",omitempty"`
	EFilingAccountTypeCode        fincen.ValidateEFilingAccountTypeCodeType `xml:"EFilingAccountTypeCode"`
	JointOwnerQuantityText        *fincen.RestrictString3                   `xml:"JointOwnerQuantityText,omitempty" json:",omitempty"`
	OtherAccountTypeText          *fincen.RestrictString50                  `xml:"OtherAccountTypeText,omitempty" json:",omitempty"`
	UnknownMaximumValueIndicator  *fincen.ValidateIndicatorType             `xml:"UnknownMaximumValueIndicator,omitempty" json:",omitempty"`
	Party                         []*AccountPartyType                       `xml:"Party"`
}

func (r AccountType) Validate(args ...string) error {
	if len(r.Party) < 1 {
		return fincen.NewErrFieldRequired("Party")
	}

	return fincen.Validate(&r, args...)
}

type ActivityNarrativeInformationType struct {
	XMLName                         xml.Name                  `xml:"ActivityNarrativeInformation"`
	SeqNum                          fincen.SeqNumber          `xml:"SeqNum,attr"`
	ActivityNarrativeSequenceNumber int                       `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString4000 `xml:"ActivityNarrativeText"`
}

func (r ActivityNarrativeInformationType) Validate(args ...string) error {
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
	if r.RawCityText == nil && fincen.CheckInvolved(typeCode, "35", "15", "57") {
		return fincen.NewErrFieldRequired("RawCityText")
	}

	if r.RawCountryCodeText == nil && fincen.CheckInvolved(typeCode, "35", "15", "57") {
		return fincen.NewErrFieldRequired("RawCountryCodeText")
	}

	if r.RawStateCodeText == nil && fincen.CheckInvolved(typeCode, "35", "15", "57") {
		return fincen.NewErrFieldRequired("RawStateCodeText")
	}

	if r.RawStreetAddress1Text == nil && fincen.CheckInvolved(typeCode, "35", "15", "57") {
		return fincen.NewErrFieldRequired("RawStreetAddress1Text")
	}

	if r.RawZIPCode == nil && fincen.CheckInvolved(typeCode, "35", "15", "57") {
		return fincen.NewErrFieldRequired("RawZIPCode")
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

type AccountPartyType struct {
	XMLName                            xml.Name                        `xml:"Party"`
	SeqNum                             fincen.SeqNumber                `xml:"SeqNum,attr"`
	ActivityPartyTypeCode              ValidateAccountPartyCodeType    `xml:"ActivityPartyTypeCode"`
	PartyAsEntityOrganizationIndicator *fincen.ValidateIndicatorType   `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PartyName                          *AccountPartyNameType           `xml:"PartyName"`
	Address                            *AccountAddressType             `xml:"Address"`
	PartyIdentification                *AccountPartyIdentificationType `xml:"PartyIdentification,omitempty" json:",omitempty"`
}

func (r AccountPartyType) fieldInclusion() error {
	typeCode := string(r.ActivityPartyTypeCode)

	if !fincen.CheckInvolved(typeCode, "42", "43") && r.PartyAsEntityOrganizationIndicator != nil {
		return fincen.NewErrFiledOmitted("PartyAsEntityOrganizationIndicator")
	}

	if typeCode == "41" && r.PartyIdentification != nil {
		return fincen.NewErrFiledOmitted("PartyIdentification")
	}

	if r.PartyName == nil {
		return fincen.NewErrFieldRequired("PartyName")
	}

	if r.Address == nil {
		return fincen.NewErrFieldRequired("Address")
	}

	return nil
}

func (r AccountPartyType) Validate(args ...string) error {
	if err := r.fieldInclusion(); err != nil {
		return err
	}

	return fincen.Validate(&r, string(r.ActivityPartyTypeCode))
}

type ForeignAccountActivityType struct {
	XMLName                          xml.Name                                 `xml:"ForeignAccountActivity"`
	SeqNum                           fincen.SeqNumber                         `xml:"SeqNum,attr"`
	ForeignAccountHeldQuantityText   *fincen.RestrictString4                  `xml:"ForeignAccountHeldQuantityText,omitempty" json:",omitempty"`
	LateFilingReasonCode             *fincen.ValidateLateFilingReasonCodeType `xml:"LateFilingReasonCode,omitempty" json:",omitempty"`
	ReportCalendarYearText           fincen.DateYYYYType                      `xml:"ReportCalendarYearText"`
	SignatureAuthoritiesQuantityText *fincen.RestrictString4                  `xml:"SignatureAuthoritiesQuantityText,omitempty" json:",omitempty"`
}

func (r ForeignAccountActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyIdentificationType struct {
	XMLName                          xml.Name                                     `xml:"PartyIdentification"`
	SeqNum                           fincen.SeqNumber                             `xml:"SeqNum,attr"`
	OtherIssuerCountryText           *fincen.RestrictString2                      `xml:"OtherIssuerCountryText,omitempty" json:",omitempty"`
	OtherPartyIdentificationTypeText *fincen.RestrictString50                     `xml:"OtherPartyIdentificationTypeText,omitempty" json:",omitempty"`
	PartyIdentificationNumberText    *fincen.RestrictString25                     `xml:"PartyIdentificationNumberText,omitempty" json:",omitempty"`
	PartyIdentificationTypeCode      *ValidateActivityPartyIdentificationCodeType `xml:"PartyIdentificationTypeCode,omitempty" json:",omitempty"`
}

func (r PartyIdentificationType) fieldInclusion(typeCode string) error {
	if r.OtherIssuerCountryText == nil && typeCode == "15" {
		return fincen.NewErrFieldRequired("OtherIssuerCountryText")
	}

	if r.OtherPartyIdentificationTypeText == nil && typeCode == "15" {
		return fincen.NewErrFieldRequired("OtherPartyIdentificationTypeText")
	}

	if r.PartyIdentificationNumberText == nil && fincen.CheckInvolved(typeCode, "35", "15", "56") {
		return fincen.NewErrFieldRequired("PartyIdentificationNumberText")
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

	if typeCode == "35" {
		if r.PartyIdentificationTypeCode != nil {
			if !fincen.CheckInvolved(string(*r.PartyIdentificationTypeCode), "4", "28") {
				return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
			}
		}
	}

	if typeCode == "15" {
		if r.PartyIdentificationTypeCode != nil {
			if !fincen.CheckInvolved(string(*r.PartyIdentificationTypeCode), "2", "1", "9", "6", "999") {
				return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
			}
		}
	}

	if typeCode == "57" {
		if r.PartyIdentificationTypeCode != nil {
			if !fincen.CheckInvolved(string(*r.PartyIdentificationTypeCode), "1", "9", "31") {
				return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
			}
		}
	}

	if typeCode == "56" {
		if r.PartyIdentificationTypeCode != nil {
			if !fincen.CheckInvolved(string(*r.PartyIdentificationTypeCode), "2", "9") {
				return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
			}
		}
	}

	return fincen.Validate(&r, args...)
}

type PartyNameType struct {
	XMLName                     xml.Name                  `xml:"PartyName"`
	SeqNum                      fincen.SeqNumber          `xml:"SeqNum,attr"`
	PartyNameTypeCode           ValidatePartyNameCodeType `xml:"PartyNameTypeCode"`
	RawEntityIndividualLastName *fincen.RestrictString150 `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName      *fincen.RestrictString35  `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName     *fincen.RestrictString35  `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualNameSuffixText *fincen.RestrictString35  `xml:"RawIndividualNameSuffixText,omitempty" json:",omitempty"`
	RawIndividualTitleText      *fincen.RestrictString20  `xml:"RawIndividualTitleText,omitempty" json:",omitempty"`
	RawPartyFullName            *fincen.RestrictString150 `xml:"RawPartyFullName,omitempty" json:",omitempty"`
}

func (r PartyNameType) fieldInclusion(typeCode string, individualIndicator bool) error {

	if r.RawEntityIndividualLastName != nil && !fincen.CheckInvolved(typeCode, "15", "57") {
		return fincen.NewErrFiledOmitted("RawEntityIndividualLastName")
	}

	if r.RawIndividualFirstName != nil && !((typeCode == "15" && individualIndicator) || typeCode == "57") {
		return fincen.NewErrFiledOmitted("RawIndividualFirstName")
	}

	if r.RawIndividualMiddleName != nil && !((typeCode == "15" && individualIndicator) || typeCode == "57") {
		return fincen.NewErrFiledOmitted("RawIndividualMiddleName")
	}

	if r.RawIndividualNameSuffixText != nil && !(typeCode == "15" && individualIndicator) {
		return fincen.NewErrFiledOmitted("RawIndividualNameSuffixText")
	}

	if r.RawIndividualTitleText != nil && !(typeCode == "15" && individualIndicator) {
		return fincen.NewErrFiledOmitted("RawIndividualTitleText")
	}

	if r.RawPartyFullName == nil && fincen.CheckInvolved(typeCode, "35", "57", "56") {
		return fincen.NewErrFieldRequired("RawPartyFullName")
	}

	return nil
}

func (r PartyNameType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	var individualIndicator bool
	if len(args) > 1 {
		individualIndicator = true
	}
	if err := r.fieldInclusion(typeCode, individualIndicator); err != nil {
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
	if r.PhoneNumberExtensionText != nil && typeCode != "57" {
		return fincen.NewErrFiledOmitted("PhoneNumberExtensionText")
	}

	if r.PhoneNumberText == nil && fincen.CheckInvolved(typeCode, "35", "57") {
		return fincen.NewErrFieldRequired("PhoneNumberText")
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

type AccountPartyNameType struct {
	XMLName                     xml.Name                  `xml:"PartyName"`
	SeqNum                      fincen.SeqNumber          `xml:"SeqNum,attr"`
	PartyNameTypeCode           ValidatePartyNameCodeType `xml:"PartyNameTypeCode"`
	RawEntityIndividualLastName *fincen.RestrictString150 `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName      *fincen.RestrictString35  `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName     *fincen.RestrictString35  `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualNameSuffixText *fincen.RestrictString35  `xml:"RawIndividualNameSuffixText,omitempty" json:",omitempty"`
	RawIndividualTitleText      *fincen.RestrictString20  `xml:"RawIndividualTitleText,omitempty" json:",omitempty"`
	RawPartyFullName            *fincen.RestrictString150 `xml:"RawPartyFullName,omitempty" json:",omitempty"`
}

func (r AccountPartyNameType) fieldInclusion(typeCode string) error {

	if r.RawEntityIndividualLastName != nil && !fincen.CheckInvolved(typeCode, "42", "43") {
		return fincen.NewErrFiledOmitted("RawEntityIndividualLastName")
	}

	if r.RawIndividualFirstName != nil && !fincen.CheckInvolved(typeCode, "42", "43") {
		return fincen.NewErrFiledOmitted("RawIndividualFirstName")
	}

	if r.RawIndividualMiddleName != nil && !fincen.CheckInvolved(typeCode, "42", "43") {
		return fincen.NewErrFiledOmitted("RawIndividualMiddleName")
	}

	if r.RawIndividualNameSuffixText != nil && !fincen.CheckInvolved(typeCode, "42", "43") {
		return fincen.NewErrFiledOmitted("RawIndividualNameSuffixText")
	}

	if r.RawIndividualTitleText != nil && typeCode != "43" {
		return fincen.NewErrFiledOmitted("RawIndividualTitleText")
	}

	if r.RawPartyFullName != nil && !fincen.CheckInvolved(typeCode, "41", "44") {
		return fincen.NewErrFiledOmitted("RawPartyFullName")
	}

	return nil
}

func (r AccountPartyNameType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	if err := r.fieldInclusion(typeCode); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type AccountAddressType struct {
	XMLName               xml.Name                 `xml:"Address"`
	SeqNum                fincen.SeqNumber         `xml:"SeqNum,attr"`
	RawCityText           fincen.RestrictString50  `xml:"RawCityText"`
	RawCountryCodeText    fincen.RestrictString2   `xml:"RawCountryCodeText"`
	RawStateCodeText      fincen.RestrictString3   `xml:"RawStateCodeText"`
	RawStreetAddress1Text fincen.RestrictString100 `xml:"RawStreetAddress1Text"`
	RawZIPCode            fincen.RestrictString9   `xml:"RawZIPCode"`
}

func (r AccountAddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AccountPartyIdentificationType struct {
	XMLName                       xml.Name                                    `xml:"PartyIdentification"`
	SeqNum                        fincen.SeqNumber                            `xml:"SeqNum,attr"`
	PartyIdentificationNumberText *fincen.RestrictString25                    `xml:"PartyIdentificationNumberText,omitempty" json:",omitempty"`
	PartyIdentificationTypeCode   *ValidateAccountPartyIdentificationCodeType `xml:"PartyIdentificationTypeCode,omitempty" json:",omitempty"`
}

func (r AccountPartyIdentificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
