// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// FinCEN Designation of Exempt Person (FinCEN DOEP | Form 110)

package exempt_designation

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

const (
	PartyTransmitter          = "35"
	PartyTransmitterContact   = "37"
	PartyAuthorizedOfficial   = "3"
	PartyExemptParty          = "11"
	PartyExemptFilerBank      = "45"
	PartyExemptAffiliatedBank = "12"
)

func NewActivity() *ActivityType {
	return &ActivityType{}
}

type ActivityType struct {
	XMLName                            xml.Name                          `xml:"Activity"`
	SeqNum                             fincen.SeqNumber                  `xml:"SeqNum,attr"`
	DesignatedMoreThanOneBankIndicator *fincen.ValidateIndicatorNullType `xml:"DesignatedMoreThanOneBankIndicator,omitempty" json:",omitempty"`
	EFilingPriorDocumentNumber         *EFilingPriorDocumentNumberType   `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                     fincen.DateYYYYMMDDType           `xml:"FilingDateText"`
	ActivityAssociation                *ActivityAssociationType          `xml:"ActivityAssociation"`
	Party                              []*PartyType                      `xml:"Party"`
	DesignationExemptActivity          *DesignationExemptActivityType    `xml:"DesignationExemptActivity"`
}

func (r ActivityType) FormTypeCode() string {
	return fincen.Report110
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
	if len(r.Party) < 4 || len(r.Party) > 104 {
		return fincen.NewErrMinMaxRange("Party")
	}

	if r.ActivityAssociation == nil {
		return fincen.NewErrFieldRequired("ActivityAssociation")
	}

	if r.DesignationExemptActivity == nil {
		return fincen.NewErrFieldRequired("DesignationExemptActivity")
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

	if _, ok := existed[PartyTransmitter]; !ok {
		return fincen.NewErrFieldRequired("Party type(transmitter)")
	}
	if _, ok := existed[PartyTransmitterContact]; !ok {
		return fincen.NewErrFieldRequired("Party type(transmitter contact)")
	}
	if _, ok := existed[PartyExemptParty]; !ok {
		return fincen.NewErrFieldRequired("Party type(exempt party)")
	}
	if _, ok := existed[PartyExemptFilerBank]; !ok {
		return fincen.NewErrFieldRequired("Party type(exempt filer bank)")
	}
	if cnt, ok := existed[PartyExemptAffiliatedBank]; ok && cnt > 99 {
		return fincen.NewErrMinMaxRange("Party type(exempt affiliated bank)")
	}

	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	XMLName                     xml.Name                          `xml:"ActivityAssociation"`
	SeqNum                      fincen.SeqNumber                  `xml:"SeqNum,attr"`
	ExemptionAmendedIndicator   *fincen.ValidateIndicatorNullType `xml:"ExemptionAmendedIndicator,omitempty" json:",omitempty"`
	ExemptionRevokedIndicator   *fincen.ValidateIndicatorNullType `xml:"ExemptionRevokedIndicator,omitempty" json:",omitempty"`
	InitialDesignationIndicator *fincen.ValidateIndicatorNullType `xml:"InitialDesignationIndicator,omitempty" json:",omitempty"`
}

func (r ActivityAssociationType) Validate(args ...string) error {

	indicatorCnt := 0

	if r.ExemptionAmendedIndicator != nil {
		indicatorCnt++
	}

	if r.ExemptionRevokedIndicator != nil {
		indicatorCnt++
	}

	if r.InitialDesignationIndicator != nil {
		indicatorCnt++
	}

	if indicatorCnt != 1 {
		return fincen.NewErrValueInvalid("ActivityAssociation")
	}

	return fincen.Validate(&r, args...)
}

type PartyType struct {
	XMLName                            xml.Name                          `xml:"Party"`
	SeqNum                             fincen.SeqNumber                  `xml:"SeqNum,attr"`
	ActivityPartyTypeCode              ValidateActivityPartyCodeType     `xml:"ActivityPartyTypeCode"`
	PartyAsEntityOrganizationIndicator *fincen.ValidateIndicatorNullType `xml:"PartyAsEntityOrganizationIndicator,omitempty" json:",omitempty"`
	PrimaryRegulatorTypeCode           *ValidateFederalRegulatorCodeType `xml:"PrimaryRegulatorTypeCode,omitempty" json:",omitempty"`
	PartyName                          []*PartyNameType                  `xml:"PartyName,omitempty" json:",omitempty"`
	Address                            *AddressType                      `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber                        *PhoneNumberType                  `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification                []*PartyIdentificationType        `xml:"PartyIdentification,omitempty" json:",omitempty"`
	PartyOccupationBusiness            *PartyOccupationBusinessType      `xml:"PartyOccupationBusiness,omitempty" json:",omitempty"`
	ElectronicAddress                  *ElectronicAddressType            `xml:"ElectronicAddress,omitempty" json:",omitempty"`
}

func (r PartyType) fieldInclusion() error {
	typeCode := string(r.ActivityPartyTypeCode)

	if typeCode == PartyTransmitter {
		if len(r.PartyName) != 1 {
			return fincen.NewErrFieldRequired("PartyName")
		}
		if r.Address == nil {
			return fincen.NewErrFieldRequired("Address")
		}
		if r.PhoneNumber == nil {
			return fincen.NewErrFieldRequired("PhoneNumber")
		}
		if len(r.PartyIdentification) != 1 {
			return fincen.NewErrFieldRequired("PartyIdentification")
		}
	}

	if typeCode == PartyTransmitterContact {
		if len(r.PartyName) != 1 {
			return fincen.NewErrFieldRequired("PartyName")
		}
	}

	if typeCode == PartyExemptParty {
		if len(r.PartyName) < 1 || len(r.PartyName) > 2 {
			return fincen.NewErrMinMaxRange("PartyName")
		}
		if r.Address == nil {
			return fincen.NewErrFieldRequired("Address")
		}
		if len(r.PartyIdentification) != 1 {
			return fincen.NewErrFieldRequired("PartyIdentification")
		}
	}

	if typeCode == PartyExemptFilerBank || typeCode == PartyExemptAffiliatedBank {
		if len(r.PartyName) != 1 {
			return fincen.NewErrFieldRequired("PartyName")
		}
		if r.PrimaryRegulatorTypeCode == nil {
			return fincen.NewErrFieldRequired("PrimaryRegulatorTypeCode")
		}
		if r.Address == nil {
			return fincen.NewErrFieldRequired("Address")
		}
		if len(r.PartyIdentification) < 1 || len(r.PartyIdentification) > 2 {
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
	RawIndividualNameSuffixText *fincen.RestrictString35   `xml:"RawIndividualNameSuffixText,omitempty" json:",omitempty"`
	RawIndividualTitleText      *fincen.RestrictString35   `xml:"RawIndividualTitleText,omitempty" json:",omitempty"`
	RawPartyFullName            *fincen.RestrictString150  `xml:"RawPartyFullName,omitempty" json:",omitempty"`
}

func (r PartyNameType) fieldInclusion(typeCode string) error {

	if fincen.CheckInvolved(typeCode, PartyTransmitter, PartyTransmitterContact) {
		if r.PartyNameTypeCode == nil {
			return fincen.NewErrFieldRequired("PartyNameTypeCode")
		}

		if *r.PartyNameTypeCode == fincen.IndicateLegalName && r.RawPartyFullName == nil {
			return fincen.NewErrFieldRequired("RawPartyFullName")
		}
	}

	if typeCode == PartyExemptParty {
		if r.PartyNameTypeCode == nil {
			return fincen.NewErrFieldRequired("PartyNameTypeCode")
		}

		if *r.PartyNameTypeCode == fincen.IndicateLegalName && r.RawEntityIndividualLastName == nil {
			return fincen.NewErrFieldRequired("RawEntityIndividualLastName")
		}

		if *r.PartyNameTypeCode == fincen.IndicateDoingBusiness && r.RawPartyFullName == nil {
			return fincen.NewErrFieldRequired("RawPartyFullName")
		}
	}

	if fincen.CheckInvolved(typeCode, PartyExemptFilerBank, PartyExemptAffiliatedBank) {
		if r.RawPartyFullName == nil {
			return fincen.NewErrFieldRequired("RawPartyFullName")
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
	XMLName               xml.Name                 `xml:"Address"`
	SeqNum                fincen.SeqNumber         `xml:"SeqNum,attr"`
	RawCityText           fincen.RestrictString50  `xml:"RawCityText"`
	RawCountryCodeText    *fincen.RestrictString2  `xml:"RawCountryCodeText,omitempty" json:",omitempty"`
	RawStateCodeText      fincen.RestrictString3   `xml:"RawStateCodeText"`
	RawStreetAddress1Text fincen.RestrictString100 `xml:"RawStreetAddress1Text"`
	RawZIPCode            fincen.RestrictString9   `xml:"RawZIPCode"`
}

func (r AddressType) fieldInclusion(typeCode string) error {

	if typeCode == PartyTransmitter {
		if r.RawCountryCodeText == nil {
			return fincen.NewErrFieldRequired("RawCountryCodeText")
		}
	}

	if !fincen.CheckInvolved(typeCode, PartyTransmitter, PartyExemptParty, PartyExemptFilerBank, PartyExemptAffiliatedBank) {
		return fincen.NewErrFiledOmitted("Address")
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

type DesignationExemptActivityType struct {
	XMLName                      xml.Name                           `xml:"DesignationExemptActivity"`
	SeqNum                       fincen.SeqNumber                   `xml:"SeqNum,attr"`
	ExemptBasisTypeCode          fincen.ValidateExemptBasisTypeCode `xml:"ExemptBasisTypeCode"`
	ExemptEffectiveBeginDateText fincen.DateYYYYMMDDType            `xml:"ExemptEffectiveBeginDateText"`
}

func (r DesignationExemptActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ElectronicAddressType struct {
	XMLName               xml.Name                 `xml:"ElectronicAddress"`
	SeqNum                fincen.SeqNumber         `xml:"SeqNum,attr"`
	ElectronicAddressText fincen.RestrictString100 `xml:"ElectronicAddressText"`
}

func (r ElectronicAddressType) fieldInclusion(typeCode string) error {
	if !fincen.CheckInvolved(typeCode, PartyExemptParty) {
		return fincen.NewErrFiledOmitted("ElectronicAddress")
	}

	return nil
}

func (r ElectronicAddressType) Validate(args ...string) error {
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
	PartyIdentificationNumberText fincen.RestrictString25             `xml:"PartyIdentificationNumberText"`
	PartyIdentificationTypeCode   ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode"`
}

func (r PartyIdentificationType) fieldInclusion(typeCode string) error {
	if !fincen.CheckInvolved(typeCode, PartyTransmitter, PartyExemptFilerBank, PartyExemptParty, PartyExemptAffiliatedBank) {
		return fincen.NewErrFiledOmitted("PartyIdentification")
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
	XMLName                xml.Name                 `xml:"PartyOccupationBusiness"`
	SeqNum                 fincen.SeqNumber         `xml:"SeqNum,attr"`
	NAICSCode              *fincen.RestrictString6  `xml:"NAICSCode,omitempty" json:",omitempty"`
	OccupationBusinessText *fincen.RestrictString30 `xml:"OccupationBusinessText,omitempty" json:",omitempty"`
}

func (r PartyOccupationBusinessType) fieldInclusion(typeCode string) error {
	if !fincen.CheckInvolved(typeCode, PartyExemptParty) {
		return fincen.NewErrFiledOmitted("PartyOccupationBusiness")
	}

	return nil
}

func (r PartyOccupationBusinessType) Validate(args ...string) error {
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

	if typeCode == PartyTransmitter {
		if r.PhoneNumberText == nil {
			return fincen.NewErrFieldRequired("PhoneNumberText")
		}
	}

	if typeCode == PartyExemptParty {
		if r.PhoneNumberText == nil {
			return fincen.NewErrFieldRequired("PhoneNumberText")
		}
	}

	if typeCode == PartyAuthorizedOfficial {
		if r.PhoneNumberText == nil {
			return fincen.NewErrFieldRequired("PhoneNumberText")
		}
	}

	if !fincen.CheckInvolved(typeCode, PartyTransmitter, PartyAuthorizedOfficial, PartyExemptParty) {
		return fincen.NewErrFiledOmitted("PhoneNumber")
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
