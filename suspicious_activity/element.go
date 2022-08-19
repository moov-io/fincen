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
	AccountNumberText       *fincen.RestrictString40    `xml:"AccountNumberText,omitempty" json:",omitempty"`
	PartyAccountAssociation PartyAccountAssociationType `xml:"PartyAccountAssociation"`
	SeqNum                  int64                       `xml:"SeqNum,attr"`
}

func (r Account) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AccountType struct {
	AccountNumberText *fincen.RestrictString40 `xml:"AccountNumberText,omitempty" json:",omitempty"`
	SeqNum            int64                    `xml:"SeqNum,attr"`
}

func (r AccountType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Activity struct {
	EFilingPriorDocumentNumber    int64                              `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                fincen.DateYYYYMMDDType            `xml:"FilingDateText"`
	FilingInstitutionNotetoFinCEN *fincen.RestrictString50           `xml:"FilingInstitutionNotetoFinCEN,omitempty" json:",omitempty"`
	ActivityAssociation           ActivityAssociationType            `xml:"ActivityAssociation"`
	ActivitySupportDocument       *ActivitySupportDocumentType       `xml:"ActivitySupportDocument,omitempty" json:",omitempty"`
	Party                         []string                           `xml:"Party"`
	SuspiciousActivity            string                             `xml:"SuspiciousActivity"`
	ActivityIPAddress             []ActivityIPAddressType            `xml:"ActivityIPAddress,omitempty" json:",omitempty"`
	CyberEventIndicators          []CyberEventIndicatorsType         `xml:"CyberEventIndicators,omitempty" json:",omitempty"`
	Assets                        []AssetsTableType                  `xml:"Assets,omitempty" json:",omitempty"`
	AssetsAttribute               []AssetsAttributeType              `xml:"AssetsAttribute,omitempty" json:",omitempty"`
	ActivityNarrativeInformation  []ActivityNarrativeInformationType `xml:"ActivityNarrativeInformation"`
	SeqNum                        int64                              `xml:"SeqNum,attr"`
}

func (r Activity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	ContinuingActivityReportIndicator  *fincen.ValidateIndicatorType `xml:"ContinuingActivityReportIndicator,omitempty" json:",omitempty"`
	CorrectsAmendsPriorReportIndicator *fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator,omitempty" json:",omitempty"`
	InitialReportIndicator             *fincen.ValidateIndicatorType `xml:"InitialReportIndicator,omitempty" json:",omitempty"`
	JointReportIndicator               *fincen.ValidateIndicatorType `xml:"JointReportIndicator,omitempty" json:",omitempty"`
	SeqNum                             int64                         `xml:"SeqNum,attr"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityIPAddressType struct {
	ActivityIPAddressDateText      *fincen.DateYYYYMMDDOrBlankType     `xml:"ActivityIPAddressDateText,omitempty" json:",omitempty"`
	ActivityIPAddressTimeStampText *fincen.ValidateTimeDataOrBlankType `xml:"ActivityIPAddressTimeStampText,omitempty" json:",omitempty"`
	IPAddressText                  fincen.RestrictString39             `xml:"IPAddressText"`
	SeqNum                         int64                               `xml:"SeqNum,attr"`
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
	EFilingPriorDocumentNumber    int64                    `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                fincen.DateYYYYMMDDType  `xml:"FilingDateText"`
	FilingInstitutionNotetoFinCEN *fincen.RestrictString50 `xml:"FilingInstitutionNotetoFinCEN,omitempty" json:",omitempty"`
	SeqNum                        int64                    `xml:"SeqNum,attr"`
}

func (r ActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AddressType struct {
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
	SeqNum                        int64                         `xml:"SeqNum,attr"`
}

func (r AddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AssociationParty struct {
	XMLName                                           xml.Name                           `xml:"Party"`
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType      `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    *fincen.ValidateIndicatorType      `xml:"AdmissionConfessionNoIndicator,omitempty" json:",omitempty"`
	AdmissionConfessionYesIndicator                   *fincen.ValidateIndicatorType      `xml:"AdmissionConfessionYesIndicator,omitempty" json:",omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator *fincen.ValidateIndicatorType      `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty" json:",omitempty"`
	BirthDateUnknownIndicator                         *fincen.ValidateIndicatorType      `xml:"BirthDateUnknownIndicator,omitempty" json:",omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          *fincen.ValidateIndicatorType      `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty" json:",omitempty"`
	ContactDateText                                   *fincen.DateYYYYMMDDOrBlankType    `xml:"ContactDateText,omitempty" json:",omitempty"`
	FemaleGenderIndicator                             *fincen.ValidateIndicatorType      `xml:"FemaleGenderIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                           *fincen.DateYYYYMMDDOrBlankTypeDOB `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	LossToFinancialAmountText                         *fincen.RestrictString15           `xml:"LossToFinancialAmountText,omitempty" json:",omitempty"`
	MaleGenderIndicator                               *fincen.ValidateIndicatorType      `xml:"MaleGenderIndicator,omitempty" json:",omitempty"`
	NoBranchActivityInvolvedIndicator                 *fincen.ValidateIndicatorType      `xml:"NoBranchActivityInvolvedIndicator,omitempty" json:",omitempty"`
	NoKnownAccountInvolvedIndicator                   *fincen.ValidateIndicatorType      `xml:"NoKnownAccountInvolvedIndicator,omitempty" json:",omitempty"`
	NonUSFinancialInstitutionIndicator                *fincen.ValidateIndicatorType      `xml:"NonUSFinancialInstitutionIndicator,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator                *fincen.ValidateIndicatorType      `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PayeeReceiverIndicator                            *fincen.ValidateIndicatorType      `xml:"PayeeReceiverIndicator,omitempty" json:",omitempty"`
	PayLocationIndicator                              *fincen.ValidateIndicatorType      `xml:"PayLocationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode                          *ValidateFederalRegulatorCodeType  `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	PurchaserSenderIndicator                          *fincen.ValidateIndicatorType      `xml:"PurchaserSenderIndicator,omitempty" json:",omitempty"`
	SellingLocationIndicator                          *fincen.ValidateIndicatorType      `xml:"SellingLocationIndicator,omitempty" json:",omitempty"`
	SellingPayingLocationIndicator                    *fincen.ValidateIndicatorType      `xml:"SellingPayingLocationIndicator,omitempty" json:",omitempty"`
	UnknownGenderIndicator                            *fincen.ValidateIndicatorType      `xml:"UnknownGenderIndicator,omitempty" json:",omitempty"`
	Address                                           AddressType                        `xml:"Address"`
	PartyIdentification                               *PartyIdentificationType           `xml:"PartyIdentification,omitempty" json:",omitempty"`
	SeqNum                                            int64                              `xml:"SeqNum,attr"`
}

func (r AssociationParty) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AccountAssociationParty struct {
	XMLName                                           xml.Name                           `xml:"Party"`
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType      `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    *fincen.ValidateIndicatorType      `xml:"AdmissionConfessionNoIndicator,omitempty" json:",omitempty"`
	AdmissionConfessionYesIndicator                   *fincen.ValidateIndicatorType      `xml:"AdmissionConfessionYesIndicator,omitempty" json:",omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator *fincen.ValidateIndicatorType      `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty" json:",omitempty"`
	BirthDateUnknownIndicator                         *fincen.ValidateIndicatorType      `xml:"BirthDateUnknownIndicator,omitempty" json:",omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          *fincen.ValidateIndicatorType      `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty" json:",omitempty"`
	ContactDateText                                   *fincen.DateYYYYMMDDOrBlankType    `xml:"ContactDateText,omitempty" json:",omitempty"`
	FemaleGenderIndicator                             *fincen.ValidateIndicatorType      `xml:"FemaleGenderIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                           *fincen.DateYYYYMMDDOrBlankTypeDOB `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	LossToFinancialAmountText                         *fincen.RestrictString15           `xml:"LossToFinancialAmountText,omitempty" json:",omitempty"`
	MaleGenderIndicator                               *fincen.ValidateIndicatorType      `xml:"MaleGenderIndicator,omitempty" json:",omitempty"`
	NoBranchActivityInvolvedIndicator                 *fincen.ValidateIndicatorType      `xml:"NoBranchActivityInvolvedIndicator,omitempty" json:",omitempty"`
	NoKnownAccountInvolvedIndicator                   *fincen.ValidateIndicatorType      `xml:"NoKnownAccountInvolvedIndicator,omitempty" json:",omitempty"`
	NonUSFinancialInstitutionIndicator                *fincen.ValidateIndicatorType      `xml:"NonUSFinancialInstitutionIndicator,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator                *fincen.ValidateIndicatorType      `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PayeeReceiverIndicator                            *fincen.ValidateIndicatorType      `xml:"PayeeReceiverIndicator,omitempty" json:",omitempty"`
	PayLocationIndicator                              *fincen.ValidateIndicatorType      `xml:"PayLocationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode                          *ValidateFederalRegulatorCodeType  `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	PurchaserSenderIndicator                          *fincen.ValidateIndicatorType      `xml:"PurchaserSenderIndicator,omitempty" json:",omitempty"`
	SellingLocationIndicator                          *fincen.ValidateIndicatorType      `xml:"SellingLocationIndicator,omitempty" json:",omitempty"`
	SellingPayingLocationIndicator                    *fincen.ValidateIndicatorType      `xml:"SellingPayingLocationIndicator,omitempty" json:",omitempty"`
	UnknownGenderIndicator                            *fincen.ValidateIndicatorType      `xml:"UnknownGenderIndicator,omitempty" json:",omitempty"`
	PartyIdentification                               *PartyIdentificationType           `xml:"PartyIdentification,omitempty" json:",omitempty"`
	Account                                           []Account                          `xml:"Account,omitempty" json:",omitempty"`
	SeqNum                                            int64                              `xml:"SeqNum,attr"`
}

func (r AccountAssociationParty) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AssetsAttributeType struct {
	AssetAttributeDescriptionText *fincen.RestrictString50                    `xml:"AssetAttributeDescriptionText,omitempty" json:",omitempty"`
	AssetAttributeTypeID          fincen.ValidateAssetAttributeTypeIDTypeCode `xml:"AssetAttributeTypeID"`
	SeqNum                        int64                                       `xml:"SeqNum,attr"`
}

func (r AssetsAttributeType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AssetsTableType struct {
	AssetSubtypeID        fincen.ValidateAssetSubtypeIDTypeCode `xml:"AssetSubtypeID"`
	AssetTypeID           fincen.ValidateAssetTypeIDTypeCode    `xml:"AssetTypeID"`
	OtherAssetSubtypeText *fincen.RestrictString50              `xml:"OtherAssetSubtypeText,omitempty" json:",omitempty"`
	SeqNum                int64                                 `xml:"SeqNum,attr"`
}

func (r AssetsTableType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CyberEventIndicatorsType struct {
	CyberEventDateText           *fincen.DateYYYYMMDDOrBlankType             `xml:"CyberEventDateText,omitempty" json:",omitempty"`
	CyberEventIndicatorsTypeCode fincen.ValidateCyberEventIndicatorsTypeCode `xml:"CyberEventIndicatorsTypeCode"`
	CyberEventTimeStampText      *fincen.ValidateTimeDataOrBlankType         `xml:"CyberEventTimeStampText,omitempty" json:",omitempty"`
	CyberEventTypeOtherText      *fincen.RestrictString50                    `xml:"CyberEventTypeOtherText,omitempty" json:",omitempty"`
	EventValueText               *fincen.RestrictString4000                  `xml:"EventValueText"`
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
	OrganizationSubtypeID        *fincen.ValidateOrganizationSubtypeCodeSarType `xml:"OrganizationSubtypeID,omitempty" json:",omitempty"`
	OrganizationTypeID           fincen.ValidateOrganizationCodeType            `xml:"OrganizationTypeID"`
	OtherOrganizationSubTypeText *fincen.RestrictString50                       `xml:"OtherOrganizationSubTypeText,omitempty" json:",omitempty"`
	OtherOrganizationTypeText    *fincen.RestrictString50                       `xml:"OtherOrganizationTypeText,omitempty" json:",omitempty"`
	SeqNum                       int64                                          `xml:"SeqNum,attr"`
}

func (r OrganizationClassificationTypeSubtypeType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Party struct {
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType               `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    *fincen.ValidateIndicatorType               `xml:"AdmissionConfessionNoIndicator,omitempty" json:",omitempty"`
	AdmissionConfessionYesIndicator                   *fincen.ValidateIndicatorType               `xml:"AdmissionConfessionYesIndicator,omitempty" json:",omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator *fincen.ValidateIndicatorType               `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty" json:",omitempty"`
	BirthDateUnknownIndicator                         *fincen.ValidateIndicatorType               `xml:"BirthDateUnknownIndicator,omitempty" json:",omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          *fincen.ValidateIndicatorType               `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty" json:",omitempty"`
	ContactDateText                                   *fincen.DateYYYYMMDDOrBlankType             `xml:"ContactDateText,omitempty" json:",omitempty"`
	FemaleGenderIndicator                             *fincen.ValidateIndicatorType               `xml:"FemaleGenderIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                           *fincen.DateYYYYMMDDOrBlankTypeDOB          `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	LossToFinancialAmountText                         *fincen.RestrictString15                    `xml:"LossToFinancialAmountText,omitempty" json:",omitempty"`
	MaleGenderIndicator                               *fincen.ValidateIndicatorType               `xml:"MaleGenderIndicator,omitempty" json:",omitempty"`
	NoBranchActivityInvolvedIndicator                 *fincen.ValidateIndicatorType               `xml:"NoBranchActivityInvolvedIndicator,omitempty" json:",omitempty"`
	NoKnownAccountInvolvedIndicator                   *fincen.ValidateIndicatorType               `xml:"NoKnownAccountInvolvedIndicator,omitempty" json:",omitempty"`
	NonUSFinancialInstitutionIndicator                *fincen.ValidateIndicatorType               `xml:"NonUSFinancialInstitutionIndicator,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator                *fincen.ValidateIndicatorType               `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PayeeReceiverIndicator                            *fincen.ValidateIndicatorType               `xml:"PayeeReceiverIndicator,omitempty" json:",omitempty"`
	PayLocationIndicator                              *fincen.ValidateIndicatorType               `xml:"PayLocationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode                          *ValidateFederalRegulatorCodeType           `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	PurchaserSenderIndicator                          *fincen.ValidateIndicatorType               `xml:"PurchaserSenderIndicator,omitempty" json:",omitempty"`
	SellingLocationIndicator                          *fincen.ValidateIndicatorType               `xml:"SellingLocationIndicator,omitempty" json:",omitempty"`
	SellingPayingLocationIndicator                    *fincen.ValidateIndicatorType               `xml:"SellingPayingLocationIndicator,omitempty" json:",omitempty"`
	UnknownGenderIndicator                            *fincen.ValidateIndicatorType               `xml:"UnknownGenderIndicator,omitempty" json:",omitempty"`
	PartyName                                         []PartyNameType                             `xml:"PartyName,omitempty" json:",omitempty"`
	Address                                           []AddressType                               `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber                                       []PhoneNumberType                           `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification                               []PartyIdentificationType                   `xml:"PartyIdentification,omitempty" json:",omitempty"`
	OrganizationClassificationTypeSubtype             []OrganizationClassificationTypeSubtypeType `xml:"OrganizationClassificationTypeSubtype,omitempty" json:",omitempty"`
	PartyOccupationBusiness                           PartyOccupationBusinessType                 `xml:"PartyOccupationBusiness,omitempty" json:",omitempty"`
	ElectronicAddress                                 []ElectronicAddressType                     `xml:"ElectronicAddress,omitempty" json:",omitempty"`
	PartyAssociation                                  []string                                    `xml:"PartyAssociation,omitempty" json:",omitempty"`
	PartyAccountAssociation                           string                                      `xml:"PartyAccountAssociation,omitempty" json:",omitempty"`
	SeqNum                                            int64                                       `xml:"SeqNum,attr"`
}

func (r Party) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAccountAssociation struct {
	AccountClosedIndicator          *fincen.ValidateIndicatorType                  `xml:"AccountClosedIndicator,omitempty" json:",omitempty"`
	PartyAccountAssociationTypeCode fincen.ValidatePartyAccountAssociationCodeType `xml:"PartyAccountAssociationTypeCode"`
	Party                           []AccountAssociationParty                      `xml:"Party"`
	SeqNum                          int64                                          `xml:"SeqNum,attr"`
}

func (r PartyAccountAssociation) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAccountAssociationType struct {
	AccountClosedIndicator          *fincen.ValidateIndicatorType                  `xml:"AccountClosedIndicator,omitempty" json:",omitempty"`
	PartyAccountAssociationTypeCode fincen.ValidatePartyAccountAssociationCodeType `xml:"PartyAccountAssociationTypeCode"`
	SeqNum                          int64                                          `xml:"SeqNum,attr"`
}

func (r PartyAccountAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAssociation struct {
	AccountantIndicator                            *fincen.ValidateIndicatorType   `xml:"AccountantIndicator,omitempty" json:",omitempty"`
	ActionTakenDateText                            *fincen.DateYYYYMMDDOrBlankType `xml:"ActionTakenDateText,omitempty" json:",omitempty"`
	AgentIndicator                                 *fincen.ValidateIndicatorType   `xml:"AgentIndicator,omitempty" json:",omitempty"`
	AppraiserIndicator                             *fincen.ValidateIndicatorType   `xml:"AppraiserIndicator,omitempty" json:",omitempty"`
	AttorneyIndicator                              *fincen.ValidateIndicatorType   `xml:"AttorneyIndicator,omitempty" json:",omitempty"`
	BorrowerIndicator                              *fincen.ValidateIndicatorType   `xml:"BorrowerIndicator,omitempty" json:",omitempty"`
	CustomerIndicator                              *fincen.ValidateIndicatorType   `xml:"CustomerIndicator,omitempty" json:",omitempty"`
	DirectorIndicator                              *fincen.ValidateIndicatorType   `xml:"DirectorIndicator,omitempty" json:",omitempty"`
	EmployeeIndicator                              *fincen.ValidateIndicatorType   `xml:"EmployeeIndicator,omitempty" json:",omitempty"`
	NoRelationshipToInstitutionIndicator           *fincen.ValidateIndicatorType   `xml:"NoRelationshipToInstitutionIndicator,omitempty" json:",omitempty"`
	OfficerIndicator                               *fincen.ValidateIndicatorType   `xml:"OfficerIndicator,omitempty" json:",omitempty"`
	OtherPartyAssociationTypeText                  *fincen.RestrictString50        `xml:"OtherPartyAssociationTypeText,omitempty" json:",omitempty"`
	OtherRelationshipIndicator                     *fincen.ValidateIndicatorType   `xml:"OtherRelationshipIndicator,omitempty" json:",omitempty"`
	OwnerShareholderIndicator                      *fincen.ValidateIndicatorType   `xml:"OwnerShareholderIndicator,omitempty" json:",omitempty"`
	RelationshipContinuesIndicator                 *fincen.ValidateIndicatorType   `xml:"RelationshipContinuesIndicator,omitempty" json:",omitempty"`
	ResignedIndicator                              *fincen.ValidateIndicatorType   `xml:"ResignedIndicator,omitempty" json:",omitempty"`
	SubjectRelationshipFinancialInstitutionTINText *fincen.RestrictString25        `xml:"SubjectRelationshipFinancialInstitutionTINText,omitempty" json:",omitempty"`
	SuspendedBarredIndicator                       *fincen.ValidateIndicatorType   `xml:"SuspendedBarredIndicator,omitempty" json:",omitempty"`
	TerminatedIndicator                            *fincen.ValidateIndicatorType   `xml:"TerminatedIndicator,omitempty" json:",omitempty"`
	Party                                          []AssociationParty              `xml:"Party,omitempty" json:",omitempty"`
	SeqNum                                         int64                           `xml:"SeqNum,attr"`
}

func (r PartyAssociation) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyAssociationType struct {
	AccountantIndicator                            *fincen.ValidateIndicatorType   `xml:"AccountantIndicator,omitempty" json:",omitempty"`
	ActionTakenDateText                            *fincen.DateYYYYMMDDOrBlankType `xml:"ActionTakenDateText,omitempty" json:",omitempty"`
	AgentIndicator                                 *fincen.ValidateIndicatorType   `xml:"AgentIndicator,omitempty" json:",omitempty"`
	AppraiserIndicator                             *fincen.ValidateIndicatorType   `xml:"AppraiserIndicator,omitempty" json:",omitempty"`
	AttorneyIndicator                              *fincen.ValidateIndicatorType   `xml:"AttorneyIndicator,omitempty" json:",omitempty"`
	BorrowerIndicator                              *fincen.ValidateIndicatorType   `xml:"BorrowerIndicator,omitempty" json:",omitempty"`
	CustomerIndicator                              *fincen.ValidateIndicatorType   `xml:"CustomerIndicator,omitempty" json:",omitempty"`
	DirectorIndicator                              *fincen.ValidateIndicatorType   `xml:"DirectorIndicator,omitempty" json:",omitempty"`
	EmployeeIndicator                              *fincen.ValidateIndicatorType   `xml:"EmployeeIndicator,omitempty" json:",omitempty"`
	NoRelationshipToInstitutionIndicator           *fincen.ValidateIndicatorType   `xml:"NoRelationshipToInstitutionIndicator,omitempty" json:",omitempty"`
	OfficerIndicator                               *fincen.ValidateIndicatorType   `xml:"OfficerIndicator,omitempty" json:",omitempty"`
	OtherPartyAssociationTypeText                  *fincen.RestrictString50        `xml:"OtherPartyAssociationTypeText,omitempty" json:",omitempty"`
	OtherRelationshipIndicator                     *fincen.ValidateIndicatorType   `xml:"OtherRelationshipIndicator,omitempty" json:",omitempty"`
	OwnerShareholderIndicator                      *fincen.ValidateIndicatorType   `xml:"OwnerShareholderIndicator,omitempty" json:",omitempty"`
	RelationshipContinuesIndicator                 *fincen.ValidateIndicatorType   `xml:"RelationshipContinuesIndicator,omitempty" json:",omitempty"`
	ResignedIndicator                              *fincen.ValidateIndicatorType   `xml:"ResignedIndicator,omitempty" json:",omitempty"`
	SubjectRelationshipFinancialInstitutionTINText *fincen.RestrictString25        `xml:"SubjectRelationshipFinancialInstitutionTINText,omitempty" json:",omitempty"`
	SuspendedBarredIndicator                       *fincen.ValidateIndicatorType   `xml:"SuspendedBarredIndicator,omitempty" json:",omitempty"`
	TerminatedIndicator                            *fincen.ValidateIndicatorType   `xml:"TerminatedIndicator,omitempty" json:",omitempty"`
	SeqNum                                         int64                           `xml:"SeqNum,attr"`
}

func (r PartyAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyIdentificationType struct {
	IdentificationPresentUnknownIndicator *fincen.ValidateIndicatorType        `xml:"IdentificationPresentUnknownIndicator,omitempty" json:",omitempty"`
	OtherIssuerCountryText                *fincen.RestrictString2              `xml:"OtherIssuerCountryText,omitempty" json:",omitempty"`
	OtherIssuerStateText                  *fincen.RestrictString3              `xml:"OtherIssuerStateText,omitempty" json:",omitempty"`
	OtherPartyIdentificationTypeText      *fincen.RestrictString50             `xml:"OtherPartyIdentificationTypeText,omitempty" json:",omitempty"`
	PartyIdentificationNumberText         *fincen.RestrictString25             `xml:"PartyIdentificationNumberText,omitempty" json:",omitempty"`
	PartyIdentificationTypeCode           *ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode,omitempty" json:",omitempty"`
	TINUnknownIndicator                   *fincen.ValidateIndicatorType        `xml:"TINUnknownIndicator,omitempty" json:",omitempty"`
	SeqNum                                int64                                `xml:"SeqNum,attr"`
}

func (r PartyIdentificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyNameType struct {
	EntityLastNameUnknownIndicator *fincen.ValidateIndicatorType `xml:"EntityLastNameUnknownIndicator,omitempty" json:",omitempty"`
	FirstNameUnknownIndicator      *fincen.ValidateIndicatorType `xml:"FirstNameUnknownIndicator,omitempty" json:",omitempty"`
	PartyNameTypeCode              *ValidatePartyNameCodeType    `xml:"PartyNameTypeCode"`
	RawEntityIndividualLastName    *fincen.RestrictString150     `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName         *fincen.RestrictString35      `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName        *fincen.RestrictString35      `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualNameSuffixText    *fincen.RestrictString35      `xml:"RawIndividualNameSuffixText,omitempty" json:",omitempty"`
	RawPartyFullName               *fincen.RestrictString150     `xml:"RawPartyFullName,omitempty" json:",omitempty"`
	SeqNum                         int64                         `xml:"SeqNum,attr"`
}

func (r PartyNameType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyOccupationBusinessType struct {
	NAICSCode              *fincen.RestrictString6  `xml:"NAICSCode,omitempty" json:",omitempty"`
	OccupationBusinessText *fincen.RestrictString50 `xml:"OccupationBusinessText,omitempty" json:",omitempty"`
	SeqNum                 int64                    `xml:"SeqNum,attr"`
}

func (r PartyOccupationBusinessType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyType struct {
	ActivityPartyTypeCode                             ValidateActivityPartyCodeType      `xml:"ActivityPartyTypeCode"`
	AdmissionConfessionNoIndicator                    *fincen.ValidateIndicatorType      `xml:"AdmissionConfessionNoIndicator,omitempty" json:",omitempty"`
	AdmissionConfessionYesIndicator                   *fincen.ValidateIndicatorType      `xml:"AdmissionConfessionYesIndicator,omitempty" json:",omitempty"`
	AllCriticalSubjectInformationUnavailableIndicator *fincen.ValidateIndicatorType      `xml:"AllCriticalSubjectInformationUnavailableIndicator,omitempty" json:",omitempty"`
	BirthDateUnknownIndicator                         *fincen.ValidateIndicatorType      `xml:"BirthDateUnknownIndicator,omitempty" json:",omitempty"`
	BothPurchaserSenderPayeeReceiveIndicator          *fincen.ValidateIndicatorType      `xml:"BothPurchaserSenderPayeeReceiveIndicator,omitempty" json:",omitempty"`
	ContactDateText                                   *fincen.DateYYYYMMDDOrBlankType    `xml:"ContactDateText,omitempty" json:",omitempty"`
	FemaleGenderIndicator                             *fincen.ValidateIndicatorType      `xml:"FemaleGenderIndicator,omitempty" json:",omitempty"`
	IndividualBirthDateText                           *fincen.DateYYYYMMDDOrBlankTypeDOB `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	LossToFinancialAmountText                         *fincen.RestrictString15           `xml:"LossToFinancialAmountText,omitempty" json:",omitempty"`
	MaleGenderIndicator                               *fincen.ValidateIndicatorType      `xml:"MaleGenderIndicator,omitempty" json:",omitempty"`
	NoBranchActivityInvolvedIndicator                 *fincen.ValidateIndicatorType      `xml:"NoBranchActivityInvolvedIndicator,omitempty" json:",omitempty"`
	NoKnownAccountInvolvedIndicator                   *fincen.ValidateIndicatorType      `xml:"NoKnownAccountInvolvedIndicator,omitempty" json:",omitempty"`
	NonUSFinancialInstitutionIndicator                *fincen.ValidateIndicatorType      `xml:"NonUSFinancialInstitutionIndicator,omitempty" json:",omitempty"`
	PartyAsEntityOrganizationIndicator                *fincen.ValidateIndicatorType      `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PayeeReceiverIndicator                            *fincen.ValidateIndicatorType      `xml:"PayeeReceiverIndicator,omitempty" json:",omitempty"`
	PayLocationIndicator                              *fincen.ValidateIndicatorType      `xml:"PayLocationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode                          *ValidateFederalRegulatorCodeType  `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	PurchaserSenderIndicator                          *fincen.ValidateIndicatorType      `xml:"PurchaserSenderIndicator,omitempty" json:",omitempty"`
	SellingLocationIndicator                          *fincen.ValidateIndicatorType      `xml:"SellingLocationIndicator,omitempty" json:",omitempty"`
	SellingPayingLocationIndicator                    *fincen.ValidateIndicatorType      `xml:"SellingPayingLocationIndicator,omitempty" json:",omitempty"`
	UnknownGenderIndicator                            *fincen.ValidateIndicatorType      `xml:"UnknownGenderIndicator,omitempty" json:",omitempty"`
	SeqNum                                            int64                              `xml:"SeqNum,attr"`
}

func (r PartyType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PhoneNumberType struct {
	PhoneNumberExtensionText *fincen.RestrictString6             `xml:"PhoneNumberExtensionText,omitempty" json:",omitempty"`
	PhoneNumberText          *fincen.RestrictString16            `xml:"PhoneNumberText,omitempty" json:",omitempty"`
	PhoneNumberTypeCode      *fincen.ValidatePhoneNumberCodeType `xml:"PhoneNumberTypeCode,omitempty" json:",omitempty"`
	SeqNum                   int64                               `xml:"SeqNum,attr"`
}

func (r PhoneNumberType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type SuspiciousActivity struct {
	AmountUnknownIndicator             *fincen.ValidateIndicatorType          `xml:"AmountUnknownIndicator,omitempty" json:",omitempty"`
	CumulativeTotalViolationAmountText *fincen.RestrictString15               `xml:"CumulativeTotalViolationAmountText,omitempty" json:",omitempty"`
	NoAmountInvolvedIndicator          *fincen.ValidateIndicatorType          `xml:"NoAmountInvolvedIndicator,omitempty" json:",omitempty"`
	SuspiciousActivityFromDateText     *fincen.DateYYYYMMDDType               `xml:"SuspiciousActivityFromDateText"`
	SuspiciousActivityToDateText       *fincen.DateYYYYMMDDOrBlankType        `xml:"SuspiciousActivityToDateText,omitempty" json:",omitempty"`
	TotalSuspiciousAmountText          *fincen.RestrictString15               `xml:"TotalSuspiciousAmountText,omitempty" json:",omitempty"`
	SuspiciousActivityClassification   []SuspiciousActivityClassificationType `xml:"SuspiciousActivityClassification"`
	SeqNum                             int64                                  `xml:"SeqNum,attr"`
}

func (r SuspiciousActivity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type SuspiciousActivityClassificationType struct {
	OtherSuspiciousActivityTypeText *fincen.RestrictString50                   `xml:"OtherSuspiciousActivityTypeText,omitempty" json:",omitempty"`
	SuspiciousActivitySubtypeID     fincen.ValidateSuspiciousActivitySubtypeID `xml:"SuspiciousActivitySubtypeID"`
	SuspiciousActivityTypeID        fincen.ValidateSuspiciousActivityTypeID    `xml:"SuspiciousActivityTypeID"`
	SeqNum                          int64                                      `xml:"SeqNum,attr"`
}

func (r SuspiciousActivityClassificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type SuspiciousActivityType struct {
	AmountUnknownIndicator             *fincen.ValidateIndicatorType   `xml:"AmountUnknownIndicator,omitempty" json:",omitempty"`
	CumulativeTotalViolationAmountText *fincen.RestrictString15        `xml:"CumulativeTotalViolationAmountText,omitempty" json:",omitempty"`
	NoAmountInvolvedIndicator          *fincen.ValidateIndicatorType   `xml:"NoAmountInvolvedIndicator,omitempty" json:",omitempty"`
	SuspiciousActivityFromDateText     *fincen.DateYYYYMMDDType        `xml:"SuspiciousActivityFromDateText"`
	SuspiciousActivityToDateText       *fincen.DateYYYYMMDDOrBlankType `xml:"SuspiciousActivityToDateText,omitempty" json:",omitempty"`
	TotalSuspiciousAmountText          *fincen.RestrictString15        `xml:"TotalSuspiciousAmountText,omitempty" json:",omitempty"`
	SeqNum                             int64                           `xml:"SeqNum,attr"`
}

func (r SuspiciousActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
