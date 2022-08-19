// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package exempt_designation

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

type EFilingBatchXML struct {
	XMLName       xml.Name `xml:"EFilingBatchXML"`
	FormTypeCode  string   `xml:"FormTypeCode"`
	Activity      []string `xml:"Activity"`
	PartyCount    int64    `xml:"PartyCount,attr"`
	ActivityCount int64    `xml:"ActivityCount,attr"`
}

func (r EFilingBatchXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Activity struct {
	DesignatedMoreThanOneBankIndicator *fincen.ValidateIndicatorType `xml:"DesignatedMoreThanOneBankIndicator,omitempty" json:",omitempty"`
	EFilingPriorDocumentNumber         int64                         `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                     fincen.DateYYYYMMDDType       `xml:"FilingDateText"`
	ActivityAssociation                ActivityAssociationType       `xml:"ActivityAssociation"`
	Party                              []string                      `xml:"Party"`
	DesignationExemptActivity          DesignationExemptActivityType `xml:"DesignationExemptActivity"`
	SeqNum                             int64                         `xml:"SeqNum,attr"`
}

func (r Activity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	ExemptionAmendedIndicator   *fincen.ValidateIndicatorType `xml:"ExemptionAmendedIndicator,omitempty" json:",omitempty"`
	ExemptionRevokedIndicator   *fincen.ValidateIndicatorType `xml:"ExemptionRevokedIndicator,omitempty" json:",omitempty"`
	InitialDesignationIndicator *fincen.ValidateIndicatorType `xml:"InitialDesignationIndicator,omitempty" json:",omitempty"`
	SeqNum                      int64                         `xml:"SeqNum,attr"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityType struct {
	DesignatedMoreThanOneBankIndicator *fincen.ValidateIndicatorType `xml:"DesignatedMoreThanOneBankIndicator,omitempty" json:",omitempty"`
	EFilingPriorDocumentNumber         int64                         `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                     fincen.DateYYYYMMDDType       `xml:"FilingDateText"`
	SeqNum                             int64                         `xml:"SeqNum,attr"`
}

func (r ActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AddressType struct {
	RawCityText           fincen.RestrictString50  `xml:"RawCityText"`
	RawCountryCodeText    *fincen.RestrictString2  `xml:"RawCountryCodeText,omitempty" json:",omitempty"`
	RawStateCodeText      fincen.RestrictString3   `xml:"RawStateCodeText"`
	RawStreetAddress1Text fincen.RestrictString100 `xml:"RawStreetAddress1Text"`
	RawZIPCode            fincen.RestrictString9   `xml:"RawZIPCode"`
	SeqNum                int64                    `xml:"SeqNum,attr"`
}

func (r AddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type DesignationExemptActivityType struct {
	ExemptBasisTypeCode          fincen.ValidateExemptBasisTypeCode `xml:"ExemptBasisTypeCode"`
	ExemptEffectiveBeginDateText fincen.DateYYYYMMDDType            `xml:"ExemptEffectiveBeginDateText"`
	SeqNum                       int64                              `xml:"SeqNum,attr"`
}

func (r DesignationExemptActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ElectronicAddressType struct {
	ElectronicAddressText fincen.RestrictString100 `xml:"ElectronicAddressText"`
	SeqNum                int64                    `xml:"SeqNum,attr"`
}

func (r ElectronicAddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Party struct {
	ActivityPartyTypeCode              ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	PartyAsEntityOrganizationIndicator *fincen.ValidateIndicatorType     `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode           *ValidateFederalRegulatorCodeType `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	PartyName                          []PartyNameType                   `xml:"PartyName,omitempty" json:",omitempty"`
	Address                            *AddressType                      `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber                        *PhoneNumberType                  `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification                []PartyIdentificationType         `xml:"PartyIdentification,omitempty" json:",omitempty"`
	PartyOccupationBusiness            *PartyOccupationBusinessType      `xml:"PartyOccupationBusiness,omitempty" json:",omitempty"`
	ElectronicAddress                  *ElectronicAddressType            `xml:"ElectronicAddress,omitempty" json:",omitempty"`
	SeqNum                             int64                             `xml:"SeqNum,attr"`
}

func (r Party) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyIdentificationType struct {
	PartyIdentificationNumberText fincen.RestrictString25             `xml:"PartyIdentificationNumberText"`
	PartyIdentificationTypeCode   ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode"`
	SeqNum                        int64                               `xml:"SeqNum,attr"`
}

func (r PartyIdentificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyNameType struct {
	PartyNameTypeCode           *ValidatePartyNameCodeType `xml:"PartyNameTypeCode,omitempty" json:",omitempty"`
	RawEntityIndividualLastName *fincen.RestrictString150  `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName      *fincen.RestrictString35   `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName     *fincen.RestrictString35   `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualNameSuffixText *fincen.RestrictString35   `xml:"RawIndividualNameSuffixText,omitempty" json:",omitempty"`
	RawIndividualTitleText      *fincen.RestrictString35   `xml:"RawIndividualTitleText,omitempty" json:",omitempty"`
	RawPartyFullName            *fincen.RestrictString150  `xml:"RawPartyFullName,omitempty" json:",omitempty"`
	SeqNum                      int64                      `xml:"SeqNum,attr"`
}

func (r PartyNameType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyOccupationBusinessType struct {
	NAICSCode              *fincen.RestrictString6  `xml:"NAICSCode,omitempty" json:",omitempty"`
	OccupationBusinessText *fincen.RestrictString30 `xml:"OccupationBusinessText,omitempty" json:",omitempty"`
	SeqNum                 int64                    `xml:"SeqNum,attr"`
}

func (r PartyOccupationBusinessType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyType struct {
	ActivityPartyTypeCode              ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	PartyAsEntityOrganizationIndicator *fincen.ValidateIndicatorType     `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode           *ValidateFederalRegulatorCodeType `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	SeqNum                             int64                             `xml:"SeqNum,attr"`
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
