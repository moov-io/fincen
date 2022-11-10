// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// FinCEN Suspicious Activity Report (FinCEN SAR) (FinCEN Report 111)

package suspicious_activity

import (
	"encoding/xml"
	"strconv"

	"github.com/moov-io/fincen"
)

func NewActivity() *ActivityType {
	return &ActivityType{}
}

type ActivityType struct {
	XMLName                       xml.Name                            `xml:"Activity"`
	SeqNum                        fincen.SeqNumber                    `xml:"SeqNum,attr"`
	EFilingPriorDocumentNumber    *EFilingPriorDocumentNumberType     `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                fincen.DateYYYYMMDDType             `xml:"FilingDateText"`
	FilingInstitutionNotetoFinCEN *fincen.RestrictString50            `xml:"FilingInstitutionNotetoFinCEN,omitempty" json:",omitempty"`
	ActivityAssociation           *ActivityAssociationType            `xml:"ActivityAssociation"`
	ActivitySupportDocument       *ActivitySupportDocumentType        `xml:"ActivitySupportDocument,omitempty" json:",omitempty"`
	Party                         []*PartyType                        `xml:"Party"`
	SuspiciousActivity            *SuspiciousActivityType             `xml:"SuspiciousActivity"`
	ActivityIPAddress             []*ActivityIPAddressType            `xml:"ActivityIPAddress,omitempty" json:",omitempty"`
	CyberEventIndicators          []*CyberEventIndicatorsType         `xml:"CyberEventIndicators,omitempty" json:",omitempty"`
	Assets                        []*AssetsTableType                  `xml:"Assets,omitempty" json:",omitempty"`
	AssetsAttribute               []*AssetsAttributeType              `xml:"AssetsAttribute,omitempty" json:",omitempty"`
	ActivityNarrativeInformation  []*ActivityNarrativeInformationType `xml:"ActivityNarrativeInformation"`
}

func (r ActivityType) FormTypeCode() string {
	return fincen.Report111
}

func (r ActivityType) TotalAmount() float64 {
	// The sum of all <TotalSuspiciousAmountText> element amounts
	// recorded in the batch

	var amount float64

	if r.SuspiciousActivity.TotalSuspiciousAmountText != nil {
		valueStr := string(*r.SuspiciousActivity.TotalSuspiciousAmountText)
		if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
			amount += value
		}
	}

	return amount
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

	if r.ActivityAssociation == nil {
		return fincen.NewErrFieldRequired("ActivityAssociation")
	}

	if r.SuspiciousActivity == nil {
		return fincen.NewErrFieldRequired("SuspiciousActivity")
	}

	if r.EFilingPriorDocumentNumber == nil {
		if r.ActivityAssociation.CorrectsAmendsPriorReportIndicator != nil || r.ActivityAssociation.ContinuingActivityReportIndicator != nil {
			// The element is not recorded and one or both of the following elements contains a “Y” value:
			//  <CorrectsAmendsPriorReportIndicator>
			//  <ContinuingActivityReportIndicator>
			return fincen.NewErrFieldRequired("EFilingPriorDocumentNumber")
		}
	}
	if r.EFilingPriorDocumentNumber != nil {
		if r.ActivityAssociation.CorrectsAmendsPriorReportIndicator == nil && r.ActivityAssociation.ContinuingActivityReportIndicator == nil {
			// The element is recorded (with a valid value) and none of the of the following elements
			// contains a “Y” value:
			//  <CorrectsAmendsPriorReportIndicator>
			//  <ContinuingActivityReportIndicator>
			return fincen.NewErrValueInvalid("EFilingPriorDocumentNumber")
		}
	}

	if len(r.Party) < 6 || len(r.Party) > 1203 {
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
	if _, ok := existed["30"]; !ok {
		return fincen.NewErrFieldRequired("Party(type 30)")
	}
	if _, ok := existed["8"]; !ok {
		return fincen.NewErrFieldRequired("Party(type 8)")
	}
	if _, ok := existed["34"]; !ok {
		return fincen.NewErrFieldRequired("Party(type 34)")
	}
	if _, ok := existed["33"]; !ok {
		return fincen.NewErrFieldRequired("Party(type 33)")
	}

	if cnt, ok := existed["34"]; ok && cnt > 99 {
		return fincen.NewErrMinMaxRange("Party(type 34)")
	}
	if cnt, ok := existed["33"]; ok && cnt > 999 {
		return fincen.NewErrMinMaxRange("Party(type 33)")
	}

	return fincen.Validate(&r)
}

type ActivityAssociationType struct {
	XMLName                            xml.Name                          `xml:"ActivityAssociation"`
	SeqNum                             fincen.SeqNumber                  `xml:"SeqNum,attr"`
	ContinuingActivityReportIndicator  *fincen.ValidateIndicatorNullType `xml:"ContinuingActivityReportIndicator,omitempty" json:",omitempty"`
	CorrectsAmendsPriorReportIndicator *fincen.ValidateIndicatorNullType `xml:"CorrectsAmendsPriorReportIndicator,omitempty" json:",omitempty"`
	InitialReportIndicator             *fincen.ValidateIndicatorNullType `xml:"InitialReportIndicator,omitempty" json:",omitempty"`
	JointReportIndicator               *fincen.ValidateIndicatorNullType `xml:"JointReportIndicator,omitempty" json:",omitempty"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	if r.InitialReportIndicator == nil && r.CorrectsAmendsPriorReportIndicator == nil && r.ContinuingActivityReportIndicator == nil {
		return fincen.NewErrValueInvalid("ActivityAssociation")
	}
	if r.InitialReportIndicator != nil && (r.CorrectsAmendsPriorReportIndicator != nil || r.ContinuingActivityReportIndicator != nil) {
		return fincen.NewErrValueInvalid("ActivityAssociation")
	}

	return fincen.Validate(&r)
}

type ActivitySupportDocumentType struct {
	XMLName                    xml.Name                 `xml:"ActivitySupportDocument"`
	SeqNum                     fincen.SeqNumber         `xml:"SeqNum,attr"`
	OriginalAttachmentFileName fincen.RestrictString150 `xml:"OriginalAttachmentFileName"`
}

func (r ActivitySupportDocumentType) Validate(args ...string) error {
	return fincen.Validate(&r)
}

type PartyType struct {
	XMLName                                           xml.Name                                     `xml:"Party"`
	SeqNum                                            fincen.SeqNumber                             `xml:"SeqNum,attr"`
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType                `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    *fincen.ValidateIndicatorNullType            `xml:"AdmissionConfessionNoIndicator,omitempty" json:",omitempty"`
	AdmissionConfessionYesIndicator                   *fincen.ValidateIndicatorNullType            `xml:"AdmissionConfessionYesIndicator,omitempty" json:",omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator *fincen.ValidateIndicatorNullType            `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty" json:",omitempty"`
	BirthDateUnknownIndicator                         *fincen.ValidateIndicatorNullType            `xml:"BirthDateUnknownIndicator,omitempty" json:",omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          *fincen.ValidateIndicatorNullType            `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty" json:",omitempty"`
	ContactDateText                                   *fincen.DateYYYYMMDDOrBlankType              `xml:"ContactDateText,omitempty" json:",omitempty"`
	FemaleGenderIndicator                             *fincen.ValidateIndicatorNullType            `xml:"FemaleGenderIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                           *fincen.DateYYYYMMDDOrBlankTypeDOB           `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	LossToFinancialAmountText                         *fincen.RestrictString15                     `xml:"LossToFinancialAmountText,omitempty" json:",omitempty"`
	MaleGenderIndicator                               *fincen.ValidateIndicatorNullType            `xml:"MaleGenderIndicator,omitempty" json:",omitempty"`
	NoBranchActivityInvolvedIndicator                 *fincen.ValidateIndicatorNullType            `xml:"NoBranchActivityInvolvedIndicator,omitempty" json:",omitempty"`
	NoKnownAccountInvolvedIndicator                   *fincen.ValidateIndicatorNullType            `xml:"NoKnownAccountInvolvedIndicator,omitempty" json:",omitempty"`
	NonUSFinancialInstitutionIndicator                *fincen.ValidateIndicatorNullType            `xml:"NonUSFinancialInstitutionIndicator,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator                *fincen.ValidateIndicatorNullType            `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PayeeReceiverIndicator                            *fincen.ValidateIndicatorNullType            `xml:"PayeeReceiverIndicator,omitempty" json:",omitempty"`
	PayLocationIndicator                              *fincen.ValidateIndicatorNullType            `xml:"PayLocationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode                          *ValidateFederalRegulatorCodeType            `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	PurchaserSenderIndicator                          *fincen.ValidateIndicatorNullType            `xml:"PurchaserSenderIndicator,omitempty" json:",omitempty"`
	SellingLocationIndicator                          *fincen.ValidateIndicatorNullType            `xml:"SellingLocationIndicator,omitempty" json:",omitempty"`
	SellingPayingLocationIndicator                    *fincen.ValidateIndicatorNullType            `xml:"SellingPayingLocationIndicator,omitempty" json:",omitempty"`
	UnknownGenderIndicator                            *fincen.ValidateIndicatorNullType            `xml:"UnknownGenderIndicator,omitempty" json:",omitempty"`
	PartyName                                         []*PartyNameType                             `xml:"PartyName,omitempty" json:",omitempty"`
	Address                                           []*AddressType                               `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber                                       []*PhoneNumberType                           `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification                               []*PartyIdentificationType                   `xml:"PartyIdentification,omitempty" json:",omitempty"`
	OrganizationClassificationTypeSubtype             []*OrganizationClassificationTypeSubtypeType `xml:"OrganizationClassificationTypeSubtype,omitempty" json:",omitempty"`
	PartyOccupationBusiness                           *PartyOccupationBusinessType                 `xml:"PartyOccupationBusiness,omitempty" json:",omitempty"`
	ElectronicAddress                                 []*ElectronicAddressType                     `xml:"ElectronicAddress,omitempty" json:",omitempty"`
	PartyAssociation                                  []*PartyAssociationType                      `xml:"PartyAssociation,omitempty" json:",omitempty"`
	PartyAccountAssociation                           *PartyAccountAssociationType                 `xml:"PartyAccountAssociation,omitempty" json:",omitempty"`
}

func (r PartyType) fieldInclusion() error {

	typeCode := string(r.ActivityPartyTypeCode)

	existed := make(map[string]bool)
	for _, pi := range r.PartyIdentification {
		if pi.PartyIdentificationTypeCode == nil {
			continue
		}

		existed[string(*pi.PartyIdentificationTypeCode)] = true
	}

	switch typeCode {
	case "35":
		if len(r.PartyName) != 1 || len(r.Address) != 1 || len(r.PhoneNumber) != 1 {
			return fincen.NewErrValueInvalid("Party")
		}
		if len(r.PartyIdentification) < 1 || len(r.PartyIdentification) > 2 {
			return fincen.NewErrValueInvalid("PartyIdentification")
		}
		if !existed["28"] || !existed["4"] {
			return fincen.NewErrValueInvalid("PartyIdentification")
		}
	case "37":
		if len(r.PartyName) != 1 {
			return fincen.NewErrValueInvalid("Party")
		}
	case "30":
		if r.PrimaryRegulatorTypeCode == nil || (len(r.PartyName) < 1 || len(r.PartyName) > 2) || len(r.Address) != 1 ||
			(len(r.OrganizationClassificationTypeSubtype) < 1 || len(r.OrganizationClassificationTypeSubtype) > 15) {
			return fincen.NewErrValueInvalid("Party")
		}
		if len(r.PartyIdentification) < 1 || len(r.PartyIdentification) > 3 {
			return fincen.NewErrValueInvalid("PartyIdentification")
		}
		if !existed["4"] {
			return fincen.NewErrValueInvalid("PartyIdentification")
		}
	case "8":
		if len(r.PartyName) != 1 || len(r.PhoneNumber) != 1 {
			return fincen.NewErrValueInvalid("Party")
		}
	case "18":
		if len(r.PartyName) != 1 {
			return fincen.NewErrValueInvalid("Party")
		}
	case "34":
		if r.PrimaryRegulatorTypeCode == nil ||
			(len(r.PartyName) < 1 || len(r.PartyName) > 2) ||
			(len(r.OrganizationClassificationTypeSubtype) < 1 || len(r.OrganizationClassificationTypeSubtype) > 12) ||
			len(r.Address) != 1 {
			return fincen.NewErrValueInvalid("Party")
		}
		if len(r.PartyIdentification) < 1 || len(r.PartyIdentification) > 3 {
			return fincen.NewErrValueInvalid("PartyIdentification")
		}
		if !existed["4"] {
			return fincen.NewErrValueInvalid("PartyIdentification")
		}
	case "33":
		if (len(r.PartyName) < 1 || len(r.PartyName) > 100) ||
			(len(r.Address) < 1 || len(r.Address) > 99) {
			return fincen.NewErrValueInvalid("Party")
		}
		if len(r.PartyIdentification) > 100 {
			return fincen.NewErrValueInvalid("PartyIdentification")
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
	XMLName                        xml.Name                          `xml:"PartyName"`
	SeqNum                         fincen.SeqNumber                  `xml:"SeqNum,attr"`
	EntityLastNameUnknownIndicator *fincen.ValidateIndicatorNullType `xml:"EntityLastNameUnknownIndicator,omitempty" json:",omitempty"`
	FirstNameUnknownIndicator      *fincen.ValidateIndicatorNullType `xml:"FirstNameUnknownIndicator,omitempty" json:",omitempty"`
	PartyNameTypeCode              *ValidatePartyNameCodeType        `xml:"PartyNameTypeCode"`
	RawEntityIndividualLastName    *fincen.RestrictString150         `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName         *fincen.RestrictString35          `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName        *fincen.RestrictString35          `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualNameSuffixText    *fincen.RestrictString35          `xml:"RawIndividualNameSuffixText,omitempty" json:",omitempty"`
	RawPartyFullName               *fincen.RestrictString150         `xml:"RawPartyFullName,omitempty" json:",omitempty"`
}

func (r *PartyNameType) fieldInclusion(typeCode string) error {
	if typeCode == "35" || typeCode == "37" || typeCode == "8" || typeCode == "18" || typeCode == "19" {
		if r.PartyNameTypeCode == nil || *r.PartyNameTypeCode == "L" || r.RawPartyFullName == nil {
			return fincen.NewErrValueInvalid("PartyName")
		}
	}
	if typeCode == "30" {
		if r.PartyNameTypeCode == nil || r.RawPartyFullName == nil {
			return fincen.NewErrValueInvalid("PartyName")
		}
	}
	if typeCode == "34" || typeCode == "33" {
		if r.PartyNameTypeCode == nil {
			return fincen.NewErrValueInvalid("PartyName")
		}
		if *r.PartyNameTypeCode != "L" || r.RawPartyFullName == nil {
			return fincen.NewErrValueInvalid("PartyName")
		}
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
	XMLName                       xml.Name                          `xml:"Address"`
	SeqNum                        fincen.SeqNumber                  `xml:"SeqNum,attr"`
	CityUnknownIndicator          *fincen.ValidateIndicatorNullType `xml:"CityUnknownIndicator,omitempty" json:",omitempty"`
	CountryCodeUnknownIndicator   *fincen.ValidateIndicatorNullType `xml:"CountryCodeUnknownIndicator,omitempty" json:",omitempty"`
	RawCityText                   *fincen.RestrictString50          `xml:"RawCityText,omitempty" json:",omitempty"`
	RawCountryCodeText            *fincen.RestrictString2           `xml:"RawCountryCodeText,omitempty" json:",omitempty"`
	RawStateCodeText              *fincen.RestrictString3           `xml:"RawStateCodeText,omitempty" json:",omitempty"`
	RawStreetAddress1Text         *fincen.RestrictString100         `xml:"RawStreetAddress1Text,omitempty" json:",omitempty"`
	RawZIPCode                    *fincen.RestrictString9           `xml:"RawZIPCode,omitempty" json:",omitempty"`
	StateCodeUnknownIndicator     *fincen.ValidateIndicatorNullType `xml:"StateCodeUnknownIndicator,omitempty" json:",omitempty"`
	StreetAddressUnknownIndicator *fincen.ValidateIndicatorNullType `xml:"StreetAddressUnknownIndicator,omitempty" json:",omitempty"`
	ZIPCodeUnknownIndicator       *fincen.ValidateIndicatorNullType `xml:"ZIPCodeUnknownIndicator,omitempty" json:",omitempty"`
}

func (r AddressType) fieldInclusion(typeCode string) error {
	if typeCode == "35" || typeCode == "30" {
		if r.RawCityText == nil || r.RawCountryCodeText == nil || r.RawStateCodeText == nil ||
			r.RawStreetAddress1Text == nil || r.RawZIPCode == nil {
			return fincen.NewErrValueInvalid("Address")
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
	XMLName                  xml.Name                            `xml:"PhoneNumber"`
	SeqNum                   fincen.SeqNumber                    `xml:"SeqNum,attr"`
	PhoneNumberExtensionText *fincen.RestrictString6             `xml:"PhoneNumberExtensionText,omitempty" json:",omitempty"`
	PhoneNumberText          *fincen.RestrictString16            `xml:"PhoneNumberText,omitempty" json:",omitempty"`
	PhoneNumberTypeCode      *fincen.ValidatePhoneNumberCodeType `xml:"PhoneNumberTypeCode,omitempty" json:",omitempty"`
}

func (r PhoneNumberType) fieldInclusion(typeCode string) error {
	if typeCode == "35" || typeCode == "8" {
		if r.PhoneNumberText == nil {
			return fincen.NewErrValueInvalid("PhoneNumber")
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
	XMLName                               xml.Name                             `xml:"PartyIdentification"`
	SeqNum                                fincen.SeqNumber                     `xml:"SeqNum,attr"`
	IdentificationPresentUnknownIndicator *fincen.ValidateIndicatorNullType    `xml:"IdentificationPresentUnknownIndicator,omitempty" json:",omitempty"`
	OtherIssuerCountryText                *fincen.RestrictString2              `xml:"OtherIssuerCountryText,omitempty" json:",omitempty"`
	OtherIssuerStateText                  *fincen.RestrictString3              `xml:"OtherIssuerStateText,omitempty" json:",omitempty"`
	OtherPartyIdentificationTypeText      *fincen.RestrictString50             `xml:"OtherPartyIdentificationTypeText,omitempty" json:",omitempty"`
	PartyIdentificationNumberText         *fincen.RestrictString25             `xml:"PartyIdentificationNumberText,omitempty" json:",omitempty"`
	PartyIdentificationTypeCode           *ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode,omitempty" json:",omitempty"`
	TINUnknownIndicator                   *fincen.ValidateIndicatorNullType    `xml:"TINUnknownIndicator,omitempty" json:",omitempty"`
}

func (r PartyIdentificationType) fieldInclusion(typeCode string) error {
	if typeCode == "35" || typeCode == "30" || typeCode == "34" || typeCode == "41" {
		if r.PartyIdentificationNumberText == nil || r.PartyIdentificationTypeCode == nil {
			return fincen.NewErrValueInvalid("PartyIdentification")
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

type OrganizationClassificationTypeSubtypeType struct {
	XMLName                      xml.Name                                       `xml:"OrganizationClassificationTypeSubtype"`
	SeqNum                       fincen.SeqNumber                               `xml:"SeqNum,attr"`
	OrganizationSubtypeID        *fincen.ValidateOrganizationSubtypeCodeSarType `xml:"OrganizationSubtypeID,omitempty" json:",omitempty"`
	OrganizationTypeID           *fincen.ValidateOrganizationCodeType           `xml:"OrganizationTypeID"`
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
	XMLName                   xml.Name                                 `xml:"ElectronicAddress"`
	SeqNum                    fincen.SeqNumber                         `xml:"SeqNum,attr"`
	ElectronicAddressText     fincen.RestrictString517                 `xml:"ElectronicAddressText"`
	ElectronicAddressTypeCode fincen.ValidateElectronicAddressTypeCode `xml:"ElectronicAddressTypeCode"`
}

func (r ElectronicAddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAssociationType struct {
	XMLName                                        xml.Name                          `xml:"PartyAssociation"`
	SeqNum                                         fincen.SeqNumber                  `xml:"SeqNum,attr"`
	AccountantIndicator                            *fincen.ValidateIndicatorNullType `xml:"AccountantIndicator,omitempty" json:",omitempty"`
	ActionTakenDateText                            *fincen.DateYYYYMMDDOrBlankType   `xml:"ActionTakenDateText,omitempty" json:",omitempty"`
	AgentIndicator                                 *fincen.ValidateIndicatorNullType `xml:"AgentIndicator,omitempty" json:",omitempty"`
	AppraiserIndicator                             *fincen.ValidateIndicatorNullType `xml:"AppraiserIndicator,omitempty" json:",omitempty"`
	AttorneyIndicator                              *fincen.ValidateIndicatorNullType `xml:"AttorneyIndicator,omitempty" json:",omitempty"`
	BorrowerIndicator                              *fincen.ValidateIndicatorNullType `xml:"BorrowerIndicator,omitempty" json:",omitempty"`
	CustomerIndicator                              *fincen.ValidateIndicatorNullType `xml:"CustomerIndicator,omitempty" json:",omitempty"`
	DirectorIndicator                              *fincen.ValidateIndicatorNullType `xml:"DirectorIndicator,omitempty" json:",omitempty"`
	EmployeeIndicator                              *fincen.ValidateIndicatorNullType `xml:"EmployeeIndicator,omitempty" json:",omitempty"`
	NoRelationshipToInstitutionIndicator           *fincen.ValidateIndicatorNullType `xml:"NoRelationshipToInstitutionIndicator,omitempty" json:",omitempty"`
	OfficerIndicator                               *fincen.ValidateIndicatorNullType `xml:"OfficerIndicator,omitempty" json:",omitempty"`
	OtherPartyAssociationTypeText                  *fincen.RestrictString50          `xml:"OtherPartyAssociationTypeText,omitempty" json:",omitempty"`
	OtherRelationshipIndicator                     *fincen.ValidateIndicatorNullType `xml:"OtherRelationshipIndicator,omitempty" json:",omitempty"`
	OwnerShareholderIndicator                      *fincen.ValidateIndicatorNullType `xml:"OwnerShareholderIndicator,omitempty" json:",omitempty"`
	RelationshipContinuesIndicator                 *fincen.ValidateIndicatorNullType `xml:"RelationshipContinuesIndicator,omitempty" json:",omitempty"`
	ResignedIndicator                              *fincen.ValidateIndicatorNullType `xml:"ResignedIndicator,omitempty" json:",omitempty"`
	SubjectRelationshipFinancialInstitutionTINText *fincen.RestrictString25          `xml:"SubjectRelationshipFinancialInstitutionTINText,omitempty" json:",omitempty"`
	SuspendedBarredIndicator                       *fincen.ValidateIndicatorNullType `xml:"SuspendedBarredIndicator,omitempty" json:",omitempty"`
	TerminatedIndicator                            *fincen.ValidateIndicatorNullType `xml:"TerminatedIndicator,omitempty" json:",omitempty"`
	Party                                          []*AssociationParty               `xml:"Party,omitempty" json:",omitempty"`
}

func (r PartyAssociationType) fieldInclusion(typeCode string) error {
	if typeCode == "34" {
		if len(r.Party) > 99 {
			return fincen.NewErrMinMaxRange("Party")
		}
	}

	return nil
}

func (r PartyAssociationType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	if err := r.fieldInclusion(typeCode); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type AssociationParty struct {
	XMLName                        xml.Name                          `xml:"Party"`
	SeqNum                         fincen.SeqNumber                  `xml:"SeqNum,attr"`
	ActivityPartyTypeCode          ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	PayLocationIndicator           *fincen.ValidateIndicatorNullType `xml:"PayLocationIndicator,omitempty" json:",omitempty"`
	SellingLocationIndicator       *fincen.ValidateIndicatorNullType `xml:"SellingLocationIndicator,omitempty" json:",omitempty"`
	SellingPayingLocationIndicator *fincen.ValidateIndicatorNullType `xml:"SellingPayingLocationIndicator,omitempty" json:",omitempty"`
	Address                        *AddressType                      `xml:"Address"`
	PartyIdentification            *PartyIdentificationType          `xml:"PartyIdentification,omitempty" json:",omitempty"`
}

func (r AssociationParty) Validate(args ...string) error {

	if r.Address == nil {
		return fincen.NewErrFieldRequired("AddressType")
	}

	return fincen.Validate(&r, string(r.ActivityPartyTypeCode))
}

type PartyAccountAssociationType struct {
	XMLName                         xml.Name                                `xml:"PartyAccountAssociation"`
	SeqNum                          fincen.SeqNumber                        `xml:"SeqNum,attr"`
	Party                           []*AccountAssociationParty              `xml:"Party"`
	AccountClosedIndicator          *fincen.ValidateIndicatorNullType       `xml:"AccountClosedIndicator,omitempty" json:",omitempty"`
	PartyAccountAssociationTypeCode ValidatePartyAccountAssociationCodeType `xml:"PartyAccountAssociationTypeCode"`
}

func (r PartyAccountAssociationType) fieldInclusion(typeCode string) error {
	if typeCode == "33" {
		if len(r.Party) < 1 || len(r.Party) > 99 {
			return fincen.NewErrMinMaxRange("Party")
		}
	}

	return nil
}

func (r PartyAccountAssociationType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	if err := r.fieldInclusion(typeCode); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type AccountAssociationParty struct {
	XMLName                            xml.Name                          `xml:"Party"`
	SeqNum                             fincen.SeqNumber                  `xml:"SeqNum,attr"`
	ActivityPartyTypeCode              ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	NonUSFinancialInstitutionIndicator *fincen.ValidateIndicatorNullType `xml:"NonUSFinancialInstitutionIndicator,omitempty" json:",omitempty"`
	PartyIdentification                *PartyIdentificationType          `xml:"PartyIdentification,omitempty" json:",omitempty"`
	Account                            []*AccountType                    `xml:"Account,omitempty" json:",omitempty"`
}

func (r AccountAssociationParty) Validate(args ...string) error {
	return fincen.Validate(&r, string(r.ActivityPartyTypeCode))
}

type AccountType struct {
	XMLName                 xml.Name                    `xml:"Account"`
	SeqNum                  fincen.SeqNumber            `xml:"SeqNum,attr"`
	AccountNumberText       *fincen.RestrictString40    `xml:"AccountNumberText,omitempty" json:",omitempty"`
	PartyAccountAssociation PartyAccountAssociationType `xml:"PartyAccountAssociation"`
}

func (r AccountType) fieldInclusion(typeCode string) error {

	if typeCode == "41" {
		if r.AccountNumberText == nil {
			return fincen.NewErrValueInvalid("Account")
		}
	}

	return nil
}

func (r AccountType) Validate(args ...string) error {
	if len(args) == 0 {
		return fincen.Validate(&r, args...)
	}

	typeCode := args[0]
	if err := r.fieldInclusion(typeCode); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type ActivityIPAddressType struct {
	XMLName                        xml.Name                            `xml:"ActivityIPAddress"`
	SeqNum                         fincen.SeqNumber                    `xml:"SeqNum,attr"`
	ActivityIPAddressDateText      *fincen.DateYYYYMMDDOrBlankType     `xml:"ActivityIPAddressDateText,omitempty" json:",omitempty"`
	ActivityIPAddressTimeStampText *fincen.ValidateTimeDataOrBlankType `xml:"ActivityIPAddressTimeStampText,omitempty" json:",omitempty"`
	IPAddressText                  fincen.RestrictString39             `xml:"IPAddressText"`
}

func (r ActivityIPAddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityNarrativeInformationType struct {
	XMLName                         xml.Name                                `xml:"ActivityNarrativeInformation"`
	SeqNum                          fincen.SeqNumber                        `xml:"SeqNum,attr"`
	ActivityNarrativeSequenceNumber ValidateActivityNarrativeSequenceNumber `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString4000               `xml:"ActivityNarrativeText"`
}

func (r ActivityNarrativeInformationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AssetsAttributeType struct {
	XMLName                       xml.Name                                    `xml:"AssetsAttribute"`
	SeqNum                        fincen.SeqNumber                            `xml:"SeqNum,attr"`
	AssetAttributeDescriptionText *fincen.RestrictString50                    `xml:"AssetAttributeDescriptionText,omitempty" json:",omitempty"`
	AssetAttributeTypeID          fincen.ValidateAssetAttributeTypeIDTypeCode `xml:"AssetAttributeTypeID"`
}

func (r AssetsAttributeType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AssetsTableType struct {
	XMLName               xml.Name                              `xml:"AssetsTable"`
	SeqNum                fincen.SeqNumber                      `xml:"SeqNum,attr"`
	AssetSubtypeID        fincen.ValidateAssetSubtypeIDTypeCode `xml:"AssetSubtypeID"`
	AssetTypeID           fincen.ValidateAssetTypeIDTypeCode    `xml:"AssetTypeID"`
	OtherAssetSubtypeText *fincen.RestrictString50              `xml:"OtherAssetSubtypeText,omitempty" json:",omitempty"`
}

func (r AssetsTableType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CyberEventIndicatorsType struct {
	XMLName                      xml.Name                                    `xml:"CyberEventIndicators"`
	SeqNum                       fincen.SeqNumber                            `xml:"SeqNum,attr"`
	CyberEventDateText           *fincen.DateYYYYMMDDOrBlankType             `xml:"CyberEventDateText,omitempty" json:",omitempty"`
	CyberEventIndicatorsTypeCode fincen.ValidateCyberEventIndicatorsTypeCode `xml:"CyberEventIndicatorsTypeCode"`
	CyberEventTimeStampText      *fincen.ValidateTimeDataOrBlankType         `xml:"CyberEventTimeStampText,omitempty" json:",omitempty"`
	CyberEventTypeOtherText      *fincen.RestrictString50                    `xml:"CyberEventTypeOtherText,omitempty" json:",omitempty"`
	EventValueText               *fincen.RestrictString4000                  `xml:"EventValueText"`
}

func (r CyberEventIndicatorsType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type SuspiciousActivityType struct {
	XMLName                            xml.Name                                `xml:"SuspiciousActivity"`
	SeqNum                             fincen.SeqNumber                        `xml:"SeqNum,attr"`
	AmountUnknownIndicator             *fincen.ValidateIndicatorNullType       `xml:"AmountUnknownIndicator,omitempty" json:",omitempty"`
	CumulativeTotalViolationAmountText *fincen.RestrictString15                `xml:"CumulativeTotalViolationAmountText,omitempty" json:",omitempty"`
	NoAmountInvolvedIndicator          *fincen.ValidateIndicatorNullType       `xml:"NoAmountInvolvedIndicator,omitempty" json:",omitempty"`
	SuspiciousActivityFromDateText     fincen.DateYYYYMMDDType                 `xml:"SuspiciousActivityFromDateText"`
	SuspiciousActivityToDateText       *fincen.DateYYYYMMDDOrBlankType         `xml:"SuspiciousActivityToDateText,omitempty" json:",omitempty"`
	TotalSuspiciousAmountText          *fincen.RestrictString15                `xml:"TotalSuspiciousAmountText,omitempty" json:",omitempty"`
	SuspiciousActivityClassification   []*SuspiciousActivityClassificationType `xml:"SuspiciousActivityClassification"`
}

func (r SuspiciousActivityType) fieldInclusion() error {
	if len(r.SuspiciousActivityClassification) < 1 || len(r.SuspiciousActivityClassification) > 99 {
		return fincen.NewErrMinMaxRange("SuspiciousActivity")
	}

	return nil
}

func (r SuspiciousActivityType) Validate(args ...string) error {
	if err := r.fieldInclusion(); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type SuspiciousActivityClassificationType struct {
	XMLName                         xml.Name                                   `xml:"SuspiciousActivityClassification"`
	SeqNum                          fincen.SeqNumber                           `xml:"SeqNum,attr"`
	OtherSuspiciousActivityTypeText *fincen.RestrictString50                   `xml:"OtherSuspiciousActivityTypeText,omitempty" json:",omitempty"`
	SuspiciousActivitySubtypeID     fincen.ValidateSuspiciousActivitySubtypeID `xml:"SuspiciousActivitySubtypeID"`
	SuspiciousActivityTypeID        fincen.ValidateSuspiciousActivityTypeID    `xml:"SuspiciousActivityTypeID"`
}

func (r SuspiciousActivityClassificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
