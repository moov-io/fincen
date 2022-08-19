// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package suspicious_activity

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

type EFilingBatchXML struct {
	XMLName                 xml.Name `xml:"EFilingBatchXML"`
	FormTypeCode            string   `xml:"FormTypeCode"`
	Activity                []string `xml:"Activity"`
	TotalAmount             float64  `xml:"TotalAmount,attr"`
	PartyCount              int64    `xml:"PartyCount,attr"`
	ActivityCount           int64    `xml:"ActivityCount,attr"`
	ActivityAttachmentCount int64    `xml:"ActivityAttachmentCount,attr"`
	AttachmentCount         int64    `xml:"AttachmentCount,attr"`
}

func (r EFilingBatchXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Account struct {
	AccountNumberText       fincen.RestrictString40     `xml:"AccountNumberText,omitempty"`
	PartyAccountAssociation PartyAccountAssociationType `xml:"PartyAccountAssociation"`
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
	EFilingPriorDocumentNumber    int64                              `xml:"EFilingPriorDocumentNumber,omitempty"`
	FilingDateText                fincen.DateYYYYMMDDType            `xml:"FilingDateText"`
	FilingInstitutionNotetoFinCEN fincen.RestrictString50            `xml:"FilingInstitutionNotetoFinCEN,omitempty"`
	ActivityAssociation           ActivityAssociationType            `xml:"ActivityAssociation"`
	ActivitySupportDocument       ActivitySupportDocumentType        `xml:"ActivitySupportDocument,omitempty"`
	Party                         []string                           `xml:"Party"`
	SuspiciousActivity            string                             `xml:"SuspiciousActivity"`
	ActivityIPAddress             []ActivityIPAddressType            `xml:"ActivityIPAddress,omitempty"`
	CyberEventIndicators          []CyberEventIndicatorsType         `xml:"CyberEventIndicators,omitempty"`
	Assets                        []AssetsTableType                  `xml:"Assets,omitempty"`
	AssetsAttribute               []AssetsAttributeType              `xml:"AssetsAttribute,omitempty"`
	ActivityNarrativeInformation  []ActivityNarrativeInformationType `xml:"ActivityNarrativeInformation"`
	SeqNum                        int64                              `xml:"SeqNum,attr"`
}

func (r Activity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	ContinuingActivityReportIndicator  fincen.ValidateIndicatorType `xml:"ContinuingActivityReportIndicator,omitempty"`
	CorrectsAmendsPriorReportIndicator fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator,omitempty"`
	InitialReportIndicator             fincen.ValidateIndicatorType `xml:"InitialReportIndicator,omitempty"`
	JointReportIndicator               fincen.ValidateIndicatorType `xml:"JointReportIndicator,omitempty"`
	SeqNum                             int64                        `xml:"SeqNum,attr"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityIPAddressType struct {
	ActivityIPAddressDateText      fincen.DateYYYYMMDDOrBlankType     `xml:"ActivityIPAddressDateText,omitempty"`
	ActivityIPAddressTimeStampText fincen.ValidateTimeDataOrBlankType `xml:"ActivityIPAddressTimeStampText,omitempty"`
	IPAddressText                  fincen.RestrictString39            `xml:"IPAddressText"`
	SeqNum                         int64                              `xml:"SeqNum,attr"`
}

func (r ActivityIPAddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityNarrativeInformationType struct {
	ActivityNarrativeSequenceNumber ValidateActivityNarrativeSequenceNumber `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString4000               `xml:"ActivityNarrativeText"`
	SeqNum                          int64                                   `xml:"SeqNum,attr"`
}

func (r ActivityNarrativeInformationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivitySupportDocumentType struct {
	OriginalAttachmentFileName fincen.RestrictString150 `xml:"OriginalAttachmentFileName"`
	SeqNum                     int64                    `xml:"SeqNum,attr"`
}

func (r ActivitySupportDocumentType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityType struct {
	EFilingPriorDocumentNumber    int64                   `xml:"EFilingPriorDocumentNumber,omitempty"`
	FilingDateText                fincen.DateYYYYMMDDType `xml:"FilingDateText"`
	FilingInstitutionNotetoFinCEN fincen.RestrictString50 `xml:"FilingInstitutionNotetoFinCEN,omitempty"`
	SeqNum                        int64                   `xml:"SeqNum,attr"`
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

type Anon1 struct {
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    fincen.ValidateIndicatorType      `xml:"AdmissionConfessionNoIndicator,omitempty"`
	AdmissionConfessionYesIndicator                   fincen.ValidateIndicatorType      `xml:"AdmissionConfessionYesIndicator,omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator fincen.ValidateIndicatorType      `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty"`
	BirthDateUnknownIndicator                         fincen.ValidateIndicatorType      `xml:"BirthDateUnknownIndicator,omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          fincen.ValidateIndicatorType      `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty"`
	ContactDateText                                   fincen.DateYYYYMMDDOrBlankType    `xml:"ContactDateText,omitempty"`
	FemaleGenderIndicator                             fincen.ValidateIndicatorType      `xml:"FemaleGenderIndicator,omitempty"`
	IndividualBirthDateText                           fincen.DateYYYYMMDDOrBlankTypeDOB `xml:"IndividualBirthDateText,omitempty"`
	LossToFinancialAmountText                         fincen.RestrictString15           `xml:"LossToFinancialAmountText,omitempty"`
	MaleGenderIndicator                               fincen.ValidateIndicatorType      `xml:"MaleGenderIndicator,omitempty"`
	NoBranchActivityInvolvedIndicator                 fincen.ValidateIndicatorType      `xml:"NoBranchActivityInvolvedIndicator,omitempty"`
	NoKnownAccountInvolvedIndicator                   fincen.ValidateIndicatorType      `xml:"NoKnownAccountInvolvedIndicator,omitempty"`
	NonUSFinancialInstitutionIndicator                fincen.ValidateIndicatorType      `xml:"NonUSFinancialInstitutionIndicator,omitempty"`
	PartyAsEntityOrganizationIndicator                fincen.ValidateIndicatorType      `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	PayeeReceiverIndicator                            fincen.ValidateIndicatorType      `xml:"PayeeReceiverIndicator,omitempty"`
	PayLocationIndicator                              fincen.ValidateIndicatorType      `xml:"PayLocationIndicator,omitempty"`
	PrimaryRegulatorTypeCode                          ValidateFederalRegulatorCodeType  `xml:"PrimaryRegulatorTypeCode,omitempty"`
	PurchaserSenderIndicator                          fincen.ValidateIndicatorType      `xml:"PurchaserSenderIndicator,omitempty"`
	SellingLocationIndicator                          fincen.ValidateIndicatorType      `xml:"SellingLocationIndicator,omitempty"`
	SellingPayingLocationIndicator                    fincen.ValidateIndicatorType      `xml:"SellingPayingLocationIndicator,omitempty"`
	UnknownGenderIndicator                            fincen.ValidateIndicatorType      `xml:"UnknownGenderIndicator,omitempty"`
	Address                                           AddressType                       `xml:"Address"`
	PartyIdentification                               PartyIdentificationType           `xml:"PartyIdentification,omitempty"`
	SeqNum                                            int64                             `xml:"SeqNum,attr"`
}

func (r Anon1) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Anon2 struct {
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    fincen.ValidateIndicatorType      `xml:"AdmissionConfessionNoIndicator,omitempty"`
	AdmissionConfessionYesIndicator                   fincen.ValidateIndicatorType      `xml:"AdmissionConfessionYesIndicator,omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator fincen.ValidateIndicatorType      `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty"`
	BirthDateUnknownIndicator                         fincen.ValidateIndicatorType      `xml:"BirthDateUnknownIndicator,omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          fincen.ValidateIndicatorType      `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty"`
	ContactDateText                                   fincen.DateYYYYMMDDOrBlankType    `xml:"ContactDateText,omitempty"`
	FemaleGenderIndicator                             fincen.ValidateIndicatorType      `xml:"FemaleGenderIndicator,omitempty"`
	IndividualBirthDateText                           fincen.DateYYYYMMDDOrBlankTypeDOB `xml:"IndividualBirthDateText,omitempty"`
	LossToFinancialAmountText                         fincen.RestrictString15           `xml:"LossToFinancialAmountText,omitempty"`
	MaleGenderIndicator                               fincen.ValidateIndicatorType      `xml:"MaleGenderIndicator,omitempty"`
	NoBranchActivityInvolvedIndicator                 fincen.ValidateIndicatorType      `xml:"NoBranchActivityInvolvedIndicator,omitempty"`
	NoKnownAccountInvolvedIndicator                   fincen.ValidateIndicatorType      `xml:"NoKnownAccountInvolvedIndicator,omitempty"`
	NonUSFinancialInstitutionIndicator                fincen.ValidateIndicatorType      `xml:"NonUSFinancialInstitutionIndicator,omitempty"`
	PartyAsEntityOrganizationIndicator                fincen.ValidateIndicatorType      `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	PayeeReceiverIndicator                            fincen.ValidateIndicatorType      `xml:"PayeeReceiverIndicator,omitempty"`
	PayLocationIndicator                              fincen.ValidateIndicatorType      `xml:"PayLocationIndicator,omitempty"`
	PrimaryRegulatorTypeCode                          ValidateFederalRegulatorCodeType  `xml:"PrimaryRegulatorTypeCode,omitempty"`
	PurchaserSenderIndicator                          fincen.ValidateIndicatorType      `xml:"PurchaserSenderIndicator,omitempty"`
	SellingLocationIndicator                          fincen.ValidateIndicatorType      `xml:"SellingLocationIndicator,omitempty"`
	SellingPayingLocationIndicator                    fincen.ValidateIndicatorType      `xml:"SellingPayingLocationIndicator,omitempty"`
	UnknownGenderIndicator                            fincen.ValidateIndicatorType      `xml:"UnknownGenderIndicator,omitempty"`
	PartyIdentification                               PartyIdentificationType           `xml:"PartyIdentification,omitempty"`
	Account                                           []Account                         `xml:"Account,omitempty"`
	SeqNum                                            int64                             `xml:"SeqNum,attr"`
}

func (r Anon2) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AssetsAttributeType struct {
	AssetAttributeDescriptionText fincen.RestrictString50                     `xml:"AssetAttributeDescriptionText,omitempty"`
	AssetAttributeTypeID          fincen.ValidateAssetAttributeTypeIDTypeCode `xml:"AssetAttributeTypeID"`
	SeqNum                        int64                                       `xml:"SeqNum,attr"`
}

func (r AssetsAttributeType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AssetsTableType struct {
	AssetSubtypeID        fincen.ValidateAssetSubtypeIDTypeCode `xml:"AssetSubtypeID"`
	AssetTypeID           fincen.ValidateAssetTypeIDTypeCode    `xml:"AssetTypeID"`
	OtherAssetSubtypeText fincen.RestrictString50               `xml:"OtherAssetSubtypeText,omitempty"`
	SeqNum                int64                                 `xml:"SeqNum,attr"`
}

func (r AssetsTableType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CyberEventIndicatorsType struct {
	CyberEventDateText           fincen.DateYYYYMMDDOrBlankType              `xml:"CyberEventDateText,omitempty"`
	CyberEventIndicatorsTypeCode fincen.ValidateCyberEventIndicatorsTypeCode `xml:"CyberEventIndicatorsTypeCode"`
	CyberEventTimeStampText      fincen.ValidateTimeDataOrBlankType          `xml:"CyberEventTimeStampText,omitempty"`
	CyberEventTypeOtherText      fincen.RestrictString50                     `xml:"CyberEventTypeOtherText,omitempty"`
	EventValueText               fincen.RestrictString4000                   `xml:"EventValueText"`
	SeqNum                       int64                                       `xml:"SeqNum,attr"`
}

func (r CyberEventIndicatorsType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ElectronicAddressType struct {
	ElectronicAddressText     fincen.RestrictString517                 `xml:"ElectronicAddressText"`
	ElectronicAddressTypeCode fincen.ValidateElectronicAddressTypeCode `xml:"ElectronicAddressTypeCode"`
	SeqNum                    int64                                    `xml:"SeqNum,attr"`
}

func (r ElectronicAddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type OrganizationClassificationTypeSubtypeType struct {
	OrganizationSubtypeID        fincen.ValidateOrganizationSubtypeCodeSarType `xml:"OrganizationSubtypeID,omitempty"`
	OrganizationTypeID           fincen.ValidateOrganizationCodeType           `xml:"OrganizationTypeID"`
	OtherOrganizationSubTypeText fincen.RestrictString50                       `xml:"OtherOrganizationSubTypeText,omitempty"`
	OtherOrganizationTypeText    fincen.RestrictString50                       `xml:"OtherOrganizationTypeText,omitempty"`
	SeqNum                       int64                                         `xml:"SeqNum,attr"`
}

func (r OrganizationClassificationTypeSubtypeType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Party struct {
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType               `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    fincen.ValidateIndicatorType                `xml:"AdmissionConfessionNoIndicator,omitempty"`
	AdmissionConfessionYesIndicator                   fincen.ValidateIndicatorType                `xml:"AdmissionConfessionYesIndicator,omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator fincen.ValidateIndicatorType                `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty"`
	BirthDateUnknownIndicator                         fincen.ValidateIndicatorType                `xml:"BirthDateUnknownIndicator,omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          fincen.ValidateIndicatorType                `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty"`
	ContactDateText                                   fincen.DateYYYYMMDDOrBlankType              `xml:"ContactDateText,omitempty"`
	FemaleGenderIndicator                             fincen.ValidateIndicatorType                `xml:"FemaleGenderIndicator,omitempty"`
	IndividualBirthDateText                           fincen.DateYYYYMMDDOrBlankTypeDOB           `xml:"IndividualBirthDateText,omitempty"`
	LossToFinancialAmountText                         fincen.RestrictString15                     `xml:"LossToFinancialAmountText,omitempty"`
	MaleGenderIndicator                               fincen.ValidateIndicatorType                `xml:"MaleGenderIndicator,omitempty"`
	NoBranchActivityInvolvedIndicator                 fincen.ValidateIndicatorType                `xml:"NoBranchActivityInvolvedIndicator,omitempty"`
	NoKnownAccountInvolvedIndicator                   fincen.ValidateIndicatorType                `xml:"NoKnownAccountInvolvedIndicator,omitempty"`
	NonUSFinancialInstitutionIndicator                fincen.ValidateIndicatorType                `xml:"NonUSFinancialInstitutionIndicator,omitempty"`
	PartyAsEntityOrganizationIndicator                fincen.ValidateIndicatorType                `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	PayeeReceiverIndicator                            fincen.ValidateIndicatorType                `xml:"PayeeReceiverIndicator,omitempty"`
	PayLocationIndicator                              fincen.ValidateIndicatorType                `xml:"PayLocationIndicator,omitempty"`
	PrimaryRegulatorTypeCode                          ValidateFederalRegulatorCodeType            `xml:"PrimaryRegulatorTypeCode,omitempty"`
	PurchaserSenderIndicator                          fincen.ValidateIndicatorType                `xml:"PurchaserSenderIndicator,omitempty"`
	SellingLocationIndicator                          fincen.ValidateIndicatorType                `xml:"SellingLocationIndicator,omitempty"`
	SellingPayingLocationIndicator                    fincen.ValidateIndicatorType                `xml:"SellingPayingLocationIndicator,omitempty"`
	UnknownGenderIndicator                            fincen.ValidateIndicatorType                `xml:"UnknownGenderIndicator,omitempty"`
	PartyName                                         []PartyNameType                             `xml:"PartyName,omitempty"`
	Address                                           []AddressType                               `xml:"Address,omitempty"`
	PhoneNumber                                       []PhoneNumberType                           `xml:"PhoneNumber,omitempty"`
	PartyIdentification                               []PartyIdentificationType                   `xml:"PartyIdentification,omitempty"`
	OrganizationClassificationTypeSubtype             []OrganizationClassificationTypeSubtypeType `xml:"OrganizationClassificationTypeSubtype,omitempty"`
	PartyOccupationBusiness                           PartyOccupationBusinessType                 `xml:"PartyOccupationBusiness,omitempty"`
	ElectronicAddress                                 []ElectronicAddressType                     `xml:"ElectronicAddress,omitempty"`
	PartyAssociation                                  []string                                    `xml:"PartyAssociation,omitempty"`
	PartyAccountAssociation                           string                                      `xml:"PartyAccountAssociation,omitempty"`
	SeqNum                                            int64                                       `xml:"SeqNum,attr"`
}

func (r Party) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAccountAssociation struct {
	AccountClosedIndicator          fincen.ValidateIndicatorType                   `xml:"AccountClosedIndicator,omitempty"`
	PartyAccountAssociationTypeCode fincen.ValidatePartyAccountAssociationCodeType `xml:"PartyAccountAssociationTypeCode"`
	Party                           []Anon2                                        `xml:"Party"`
	SeqNum                          int64                                          `xml:"SeqNum,attr"`
}

func (r PartyAccountAssociation) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAccountAssociationType struct {
	AccountClosedIndicator          fincen.ValidateIndicatorType                   `xml:"AccountClosedIndicator,omitempty"`
	PartyAccountAssociationTypeCode fincen.ValidatePartyAccountAssociationCodeType `xml:"PartyAccountAssociationTypeCode"`
	SeqNum                          int64                                          `xml:"SeqNum,attr"`
}

func (r PartyAccountAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAssociation struct {
	AccountantIndicator                            fincen.ValidateIndicatorType   `xml:"AccountantIndicator,omitempty"`
	ActionTakenDateText                            fincen.DateYYYYMMDDOrBlankType `xml:"ActionTakenDateText,omitempty"`
	AgentIndicator                                 fincen.ValidateIndicatorType   `xml:"AgentIndicator,omitempty"`
	AppraiserIndicator                             fincen.ValidateIndicatorType   `xml:"AppraiserIndicator,omitempty"`
	AttorneyIndicator                              fincen.ValidateIndicatorType   `xml:"AttorneyIndicator,omitempty"`
	BorrowerIndicator                              fincen.ValidateIndicatorType   `xml:"BorrowerIndicator,omitempty"`
	CustomerIndicator                              fincen.ValidateIndicatorType   `xml:"CustomerIndicator,omitempty"`
	DirectorIndicator                              fincen.ValidateIndicatorType   `xml:"DirectorIndicator,omitempty"`
	EmployeeIndicator                              fincen.ValidateIndicatorType   `xml:"EmployeeIndicator,omitempty"`
	NoRelationshipToInstitutionIndicator           fincen.ValidateIndicatorType   `xml:"NoRelationshipToInstitutionIndicator,omitempty"`
	OfficerIndicator                               fincen.ValidateIndicatorType   `xml:"OfficerIndicator,omitempty"`
	OtherPartyAssociationTypeText                  fincen.RestrictString50        `xml:"OtherPartyAssociationTypeText,omitempty"`
	OtherRelationshipIndicator                     fincen.ValidateIndicatorType   `xml:"OtherRelationshipIndicator,omitempty"`
	OwnerShareholderIndicator                      fincen.ValidateIndicatorType   `xml:"OwnerShareholderIndicator,omitempty"`
	RelationshipContinuesIndicator                 fincen.ValidateIndicatorType   `xml:"RelationshipContinuesIndicator,omitempty"`
	ResignedIndicator                              fincen.ValidateIndicatorType   `xml:"ResignedIndicator,omitempty"`
	SubjectRelationshipFinancialInstitutionTINText fincen.RestrictString25        `xml:"SubjectRelationshipFinancialInstitutionTINText,omitempty"`
	SuspendedBarredIndicator                       fincen.ValidateIndicatorType   `xml:"SuspendedBarredIndicator,omitempty"`
	TerminatedIndicator                            fincen.ValidateIndicatorType   `xml:"TerminatedIndicator,omitempty"`
	Party                                          []Anon1                        `xml:"Party,omitempty"`
	SeqNum                                         int64                          `xml:"SeqNum,attr"`
}

func (r PartyAssociation) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAssociationType struct {
	AccountantIndicator                            fincen.ValidateIndicatorType   `xml:"AccountantIndicator,omitempty"`
	ActionTakenDateText                            fincen.DateYYYYMMDDOrBlankType `xml:"ActionTakenDateText,omitempty"`
	AgentIndicator                                 fincen.ValidateIndicatorType   `xml:"AgentIndicator,omitempty"`
	AppraiserIndicator                             fincen.ValidateIndicatorType   `xml:"AppraiserIndicator,omitempty"`
	AttorneyIndicator                              fincen.ValidateIndicatorType   `xml:"AttorneyIndicator,omitempty"`
	BorrowerIndicator                              fincen.ValidateIndicatorType   `xml:"BorrowerIndicator,omitempty"`
	CustomerIndicator                              fincen.ValidateIndicatorType   `xml:"CustomerIndicator,omitempty"`
	DirectorIndicator                              fincen.ValidateIndicatorType   `xml:"DirectorIndicator,omitempty"`
	EmployeeIndicator                              fincen.ValidateIndicatorType   `xml:"EmployeeIndicator,omitempty"`
	NoRelationshipToInstitutionIndicator           fincen.ValidateIndicatorType   `xml:"NoRelationshipToInstitutionIndicator,omitempty"`
	OfficerIndicator                               fincen.ValidateIndicatorType   `xml:"OfficerIndicator,omitempty"`
	OtherPartyAssociationTypeText                  fincen.RestrictString50        `xml:"OtherPartyAssociationTypeText,omitempty"`
	OtherRelationshipIndicator                     fincen.ValidateIndicatorType   `xml:"OtherRelationshipIndicator,omitempty"`
	OwnerShareholderIndicator                      fincen.ValidateIndicatorType   `xml:"OwnerShareholderIndicator,omitempty"`
	RelationshipContinuesIndicator                 fincen.ValidateIndicatorType   `xml:"RelationshipContinuesIndicator,omitempty"`
	ResignedIndicator                              fincen.ValidateIndicatorType   `xml:"ResignedIndicator,omitempty"`
	SubjectRelationshipFinancialInstitutionTINText fincen.RestrictString25        `xml:"SubjectRelationshipFinancialInstitutionTINText,omitempty"`
	SuspendedBarredIndicator                       fincen.ValidateIndicatorType   `xml:"SuspendedBarredIndicator,omitempty"`
	TerminatedIndicator                            fincen.ValidateIndicatorType   `xml:"TerminatedIndicator,omitempty"`
	SeqNum                                         int64                          `xml:"SeqNum,attr"`
}

func (r PartyAssociationType) Validate(args ...string) error {
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
	PartyNameTypeCode              ValidatePartyNameCodeType    `xml:"PartyNameTypeCode"`
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
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    fincen.ValidateIndicatorType      `xml:"AdmissionConfessionNoIndicator,omitempty"`
	AdmissionConfessionYesIndicator                   fincen.ValidateIndicatorType      `xml:"AdmissionConfessionYesIndicator,omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator fincen.ValidateIndicatorType      `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty"`
	BirthDateUnknownIndicator                         fincen.ValidateIndicatorType      `xml:"BirthDateUnknownIndicator,omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          fincen.ValidateIndicatorType      `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty"`
	ContactDateText                                   fincen.DateYYYYMMDDOrBlankType    `xml:"ContactDateText,omitempty"`
	FemaleGenderIndicator                             fincen.ValidateIndicatorType      `xml:"FemaleGenderIndicator,omitempty"`
	IndividualBirthDateText                           fincen.DateYYYYMMDDOrBlankTypeDOB `xml:"IndividualBirthDateText,omitempty"`
	LossToFinancialAmountText                         fincen.RestrictString15           `xml:"LossToFinancialAmountText,omitempty"`
	MaleGenderIndicator                               fincen.ValidateIndicatorType      `xml:"MaleGenderIndicator,omitempty"`
	NoBranchActivityInvolvedIndicator                 fincen.ValidateIndicatorType      `xml:"NoBranchActivityInvolvedIndicator,omitempty"`
	NoKnownAccountInvolvedIndicator                   fincen.ValidateIndicatorType      `xml:"NoKnownAccountInvolvedIndicator,omitempty"`
	NonUSFinancialInstitutionIndicator                fincen.ValidateIndicatorType      `xml:"NonUSFinancialInstitutionIndicator,omitempty"`
	PartyAsEntityOrganizationIndicator                fincen.ValidateIndicatorType      `xml:"PartyAsEntityOrganizationIndicator,omitempty"`
	PayeeReceiverIndicator                            fincen.ValidateIndicatorType      `xml:"PayeeReceiverIndicator,omitempty"`
	PayLocationIndicator                              fincen.ValidateIndicatorType      `xml:"PayLocationIndicator,omitempty"`
	PrimaryRegulatorTypeCode                          ValidateFederalRegulatorCodeType  `xml:"PrimaryRegulatorTypeCode,omitempty"`
	PurchaserSenderIndicator                          fincen.ValidateIndicatorType      `xml:"PurchaserSenderIndicator,omitempty"`
	SellingLocationIndicator                          fincen.ValidateIndicatorType      `xml:"SellingLocationIndicator,omitempty"`
	SellingPayingLocationIndicator                    fincen.ValidateIndicatorType      `xml:"SellingPayingLocationIndicator,omitempty"`
	UnknownGenderIndicator                            fincen.ValidateIndicatorType      `xml:"UnknownGenderIndicator,omitempty"`
	SeqNum                                            int64                             `xml:"SeqNum,attr"`
}

func (r PartyType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PhoneNumberType struct {
	PhoneNumberExtensionText fincen.RestrictString6             `xml:"PhoneNumberExtensionText,omitempty"`
	PhoneNumberText          fincen.RestrictString16            `xml:"PhoneNumberText,omitempty"`
	PhoneNumberTypeCode      fincen.ValidatePhoneNumberCodeType `xml:"PhoneNumberTypeCode,omitempty"`
	SeqNum                   int64                              `xml:"SeqNum,attr"`
}

func (r PhoneNumberType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type SuspiciousActivity struct {
	AmountUnknownIndicator             fincen.ValidateIndicatorType           `xml:"AmountUnknownIndicator,omitempty"`
	CumulativeTotalViolationAmountText fincen.RestrictString15                `xml:"CumulativeTotalViolationAmountText,omitempty"`
	NoAmountInvolvedIndicator          fincen.ValidateIndicatorType           `xml:"NoAmountInvolvedIndicator,omitempty"`
	SuspiciousActivityFromDateText     fincen.DateYYYYMMDDType                `xml:"SuspiciousActivityFromDateText"`
	SuspiciousActivityToDateText       fincen.DateYYYYMMDDOrBlankType         `xml:"SuspiciousActivityToDateText,omitempty"`
	TotalSuspiciousAmountText          fincen.RestrictString15                `xml:"TotalSuspiciousAmountText,omitempty"`
	SuspiciousActivityClassification   []SuspiciousActivityClassificationType `xml:"SuspiciousActivityClassification"`
	SeqNum                             int64                                  `xml:"SeqNum,attr"`
}

func (r SuspiciousActivity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type SuspiciousActivityClassificationType struct {
	OtherSuspiciousActivityTypeText fincen.RestrictString50                    `xml:"OtherSuspiciousActivityTypeText,omitempty"`
	SuspiciousActivitySubtypeID     fincen.ValidateSuspiciousActivitySubtypeID `xml:"SuspiciousActivitySubtypeID"`
	SuspiciousActivityTypeID        fincen.ValidateSuspiciousActivityTypeID    `xml:"SuspiciousActivityTypeID"`
	SeqNum                          int64                                      `xml:"SeqNum,attr"`
}

func (r SuspiciousActivityClassificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type SuspiciousActivityType struct {
	AmountUnknownIndicator             fincen.ValidateIndicatorType   `xml:"AmountUnknownIndicator,omitempty"`
	CumulativeTotalViolationAmountText fincen.RestrictString15        `xml:"CumulativeTotalViolationAmountText,omitempty"`
	NoAmountInvolvedIndicator          fincen.ValidateIndicatorType   `xml:"NoAmountInvolvedIndicator,omitempty"`
	SuspiciousActivityFromDateText     fincen.DateYYYYMMDDType        `xml:"SuspiciousActivityFromDateText"`
	SuspiciousActivityToDateText       fincen.DateYYYYMMDDOrBlankType `xml:"SuspiciousActivityToDateText,omitempty"`
	TotalSuspiciousAmountText          fincen.RestrictString15        `xml:"TotalSuspiciousAmountText,omitempty"`
	SeqNum                             int64                          `xml:"SeqNum,attr"`
}

func (r SuspiciousActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
