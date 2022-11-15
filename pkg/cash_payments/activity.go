// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Report of Cash Payments Over $10,000 Received in a Trade or Business (FinCEN Form 8300)

package cash_payments

import (
	"encoding/xml"
	"strconv"

	"github.com/moov-io/fincen"
)

const (
	PartyTransmitter        = "35"
	PartyTransmitterContact = "37"
	PartyAuthorizedOfficial = "3"
	PartyContactAssistance  = "8"
	PartyIndividual         = "16"
	PartyPerson             = "23"
	PartyBusiness           = "4"
	PartyITCTCC             = "28"
	PartyITCEIN             = "2"
)

func NewActivity() *ActivityType {
	return &ActivityType{}
}

type ActivityType struct {
	XMLName                                     xml.Name                          `xml:"Activity"`
	SeqNum                                      fincen.SeqNumber                  `xml:"SeqNum,attr"`
	EFilingPriorDocumentNumber                  int64                             `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                              fincen.DateYYYYMMDDType           `xml:"FilingDateText"`
	MultipleSubjectsIndicator                   *fincen.ValidateIndicatorNullType `xml:"MultipleSubjectsIndicator,omitempty" json:",omitempty"`
	SuspiciousTransactionIndicator              *fincen.ValidateIndicatorNullType `xml:"SuspiciousTransactionIndicator,omitempty" json:",omitempty"`
	TransactionOnBehalfMultiplePersonsIndicator *fincen.ValidateIndicatorNullType `xml:"TransactionOnBehalfMultiplePersonsIndicator,omitempty" json:",omitempty"`
	ActivityAssociation                         *ActivityAssociationType          `xml:"ActivityAssociation"`
	Party                                       []*PartyType                      `xml:"Party"`
	CurrencyTransactionActivity                 *CurrencyTransactionActivityType  `xml:"CurrencyTransactionActivity"`
	ActivityNarrativeInformation                *ActivityNarrativeInformationType `xml:"ActivityNarrativeInformation,omitempty" json:",omitempty"`
}

func (r ActivityType) FormTypeCode() string {
	return fincen.Form8300
}

func (r ActivityType) TotalAmount() float64 {
	// The sum of all <DetailTransactionAmountText> element amounts

	var amount float64
	for _, currency := range r.CurrencyTransactionActivity.CurrencyTransactionActivityDetail {
		if currency.DetailTransactionAmountText == nil {
			continue
		}

		valueStr := string(*currency.DetailTransactionAmountText)
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
	if len(r.Party) < 4 || len(r.Party) > 203 {
		return fincen.NewErrMinMaxRange("Party")
	}

	if r.CurrencyTransactionActivity == nil {
		return fincen.NewErrFieldRequired("CurrencyTransactionActivity")
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
	if _, ok := existed[PartyAuthorizedOfficial]; !ok {
		return fincen.NewErrFieldRequired("Party type(authorized official)")
	}
	if _, ok := existed[PartyBusiness]; !ok {
		return fincen.NewErrFieldRequired("Party type(business that received cash)")
	}
	if cnt, ok := existed[PartyIndividual]; ok && cnt > 99 {
		return fincen.NewErrMinMaxRange("Party type(individual from whom the cash was received)")
	}
	if cnt, ok := existed[PartyPerson]; ok && cnt > 99 {
		return fincen.NewErrMinMaxRange("Party type(person on whose behalf transaction conducted)")
	}
	if cnt, ok := existed[PartyAuthorizedOfficial]; ok && cnt > 1 {
		return fincen.NewErrMinMaxRange("Party type(authorized official)")
	}
	if cnt, ok := existed[PartyBusiness]; ok && cnt > 1 {
		return fincen.NewErrMinMaxRange("Party type(business that received cash)")
	}

	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	XMLName                            xml.Name                          `xml:"ActivityAssociation"`
	SeqNum                             fincen.SeqNumber                  `xml:"SeqNum,attr"`
	CorrectsAmendsPriorReportIndicator *fincen.ValidateIndicatorNullType `xml:"CorrectsAmendsPriorReportIndicator,omitempty" json:",omitempty"`
	InitialReportIndicator             *fincen.ValidateIndicatorNullType `xml:"InitialReportIndicator,omitempty" json:",omitempty"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyType struct {
	XMLName                 xml.Name                        `xml:"Party"`
	SeqNum                  fincen.SeqNumber                `xml:"SeqNum,attr"`
	ActivityPartyTypeCode   ValidateActivityPartyCodeType   `xml:"ActivityPartyTypeCode"`
	IndividualBirthDateText *fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	PartyTypeCode           *fincen.ValidatePartyTypeCode   `xml:"PartyTypeCode,omitempty" json:",omitempty"`
	PartyName               []*PartyNameType                `xml:"PartyName,omitempty" json:",omitempty"`
	Address                 *AddressType                    `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber             *PhoneNumberType                `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification     []*PartyIdentificationType      `xml:"PartyIdentification,omitempty" json:",omitempty"`
	PartyOccupationBusiness *PartyOccupationBusinessType    `xml:"PartyOccupationBusiness,omitempty" json:",omitempty"`
}

func (r PartyType) fieldInclusion() error {
	typeCode := string(r.ActivityPartyTypeCode)

	if typeCode == PartyTransmitter {
		if r.PartyName == nil {
			return fincen.NewErrFieldRequired("PartyName")
		}
		if r.Address == nil {
			return fincen.NewErrFieldRequired("Address")
		}
		if r.PhoneNumber == nil {
			return fincen.NewErrFieldRequired("PhoneNumber")
		}
		if len(r.PartyIdentification) < 1 {
			return fincen.NewErrFieldRequired("PartyIdentification")
		}
		if len(r.PartyIdentification) > 1 {
			return fincen.NewErrMinMaxRange("PartyIdentification")
		}
	}

	if typeCode == PartyTransmitterContact || typeCode == PartyAuthorizedOfficial {
		if r.PartyName == nil {
			return fincen.NewErrFieldRequired("PartyName")
		}
	}

	if typeCode == PartyIndividual {
		if r.PartyName == nil {
			return fincen.NewErrFieldRequired("PartyName")
		}
		if r.Address == nil {
			return fincen.NewErrFieldRequired("Address")
		}
		if len(r.PartyIdentification) < 1 {
			return fincen.NewErrFieldRequired("PartyIdentification")
		}
		if len(r.PartyIdentification) > 2 {
			return fincen.NewErrMinMaxRange("PartyIdentification")
		}
	}

	if typeCode == PartyPerson {
		if r.PartyTypeCode == nil {
			return fincen.NewErrFieldRequired("PartyTypeCode")
		}
		if len(r.PartyName) > 2 {
			return fincen.NewErrMinMaxRange("PartyName")
		}
		if len(r.PartyIdentification) > 3 {
			return fincen.NewErrMinMaxRange("PartyIdentification")
		}
	}

	if typeCode == PartyBusiness {
		if r.PartyName == nil {
			return fincen.NewErrFieldRequired("PartyName")
		}
		if len(r.PartyIdentification) < 1 {
			return fincen.NewErrFieldRequired("PartyIdentification")
		}
		if len(r.PartyIdentification) > 2 {
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
	RawIndividualTitleText      *fincen.RestrictString35   `xml:"RawIndividualTitleText,omitempty" json:",omitempty"`
	RawPartyFullName            *fincen.RestrictString150  `xml:"RawPartyFullName,omitempty" json:",omitempty"`
}

func (r PartyNameType) fieldInclusion(typeCode string) error {
	if fincen.CheckInvolved(typeCode, PartyTransmitter, PartyTransmitterContact) {
		if r.PartyNameTypeCode == nil || *r.PartyNameTypeCode != fincen.IndicateLegalName {
			return fincen.NewErrValueInvalid("PartyNameTypeCode")
		}
		if r.RawPartyFullName == nil {
			return fincen.NewErrFieldRequired("RawPartyFullName")
		}
	}

	if fincen.CheckInvolved(typeCode, PartyPerson) {
		if r.PartyNameTypeCode != nil && *r.PartyNameTypeCode != fincen.IndicateDoingBusiness {
			if r.RawPartyFullName == nil {
				return fincen.NewErrFieldRequired("RawPartyFullName")
			}
		}
	}

	if fincen.CheckInvolved(typeCode, PartyBusiness) && r.RawPartyFullName == nil {
		return fincen.NewErrFieldRequired("RawPartyFullName")
	}

	if fincen.CheckInvolved(typeCode, PartyAuthorizedOfficial) && r.RawIndividualTitleText == nil {
		return fincen.NewErrFieldRequired("RawIndividualTitleText")
	}

	if fincen.CheckInvolved(typeCode, PartyContactAssistance) && r.RawPartyFullName == nil {
		return fincen.NewErrFieldRequired("RawPartyFullName")
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
	XMLName               xml.Name                  `xml:"Address"`
	SeqNum                fincen.SeqNumber          `xml:"SeqNum,attr"`
	RawCityText           *fincen.RestrictString50  `xml:"RawCityText,omitempty" json:",omitempty"`
	RawCountryCodeText    *fincen.RestrictString2   `xml:"RawCountryCodeText,omitempty" json:",omitempty"`
	RawStateCodeText      *fincen.RestrictString3   `xml:"RawStateCodeText,omitempty" json:",omitempty"`
	RawStreetAddress1Text *fincen.RestrictString100 `xml:"RawStreetAddress1Text,omitempty" json:",omitempty"`
	RawZIPCode            *fincen.RestrictString9   `xml:"RawZIPCode,omitempty" json:",omitempty"`
}

func (r AddressType) fieldInclusion(typeCode string) error {
	if fincen.CheckInvolved(typeCode, PartyTransmitter) {
		if r.RawCityText == nil {
			return fincen.NewErrFieldRequired("RawCityText")
		}
		if r.RawCountryCodeText == nil {
			return fincen.NewErrFieldRequired("RawCountryCodeText")
		}
		if r.RawStateCodeText == nil {
			return fincen.NewErrFieldRequired("RawStateCodeText")
		}
		if r.RawStreetAddress1Text == nil {
			return fincen.NewErrFieldRequired("RawStreetAddress1Text")
		}
		if r.RawZIPCode == nil {
			return fincen.NewErrFieldRequired("RawZIPCode")
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
	XMLName         xml.Name                 `xml:"PhoneNumber"`
	SeqNum          fincen.SeqNumber         `xml:"SeqNum,attr"`
	PhoneNumberText *fincen.RestrictString16 `xml:"PhoneNumberText,omitempty" json:",omitempty"`
}

func (r PhoneNumberType) fieldInclusion(typeCode string) error {
	if fincen.CheckInvolved(typeCode, PartyTransmitter, PartyContactAssistance) {
		if r.PhoneNumberText == nil {
			return fincen.NewErrFieldRequired("PhoneNumberText")
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
	XMLName                       xml.Name                            `xml:"PartyIdentification"`
	SeqNum                        fincen.SeqNumber                    `xml:"SeqNum,attr"`
	OtherIssuerStateText          *fincen.RestrictString3             `xml:"OtherIssuerStateText,omitempty" json:",omitempty"`
	PartyIdentificationNumberText *fincen.RestrictString25            `xml:"PartyIdentificationNumberText,omitempty" json:",omitempty"`
	PartyIdentificationTypeCode   ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode"`
}

func (r PartyIdentificationType) fieldInclusion(typeCode string) error {
	if fincen.CheckInvolved(typeCode, PartyTransmitter) {
		if r.PartyIdentificationNumberText == nil {
			return fincen.NewErrFieldRequired("PartyIdentificationNumberText")
		}
		if string(r.PartyIdentificationTypeCode) != PartyITCTCC && string(r.PartyIdentificationTypeCode) != PartyITCEIN {
			return fincen.NewErrValueInvalid("PartyIdentificationTypeCode")
		}
	}

	if fincen.CheckInvolved(typeCode, PartyBusiness, PartyPerson, PartyIndividual) {
		if r.PartyIdentificationNumberText == nil {
			return fincen.NewErrFieldRequired("PartyIdentificationNumberText")
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

type PartyOccupationBusinessType struct {
	XMLName                xml.Name                `xml:"PartyOccupationBusiness"`
	SeqNum                 fincen.SeqNumber        `xml:"SeqNum,attr"`
	OccupationBusinessText fincen.RestrictString30 `xml:"OccupationBusinessText"`
}

func (r PartyOccupationBusinessType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityType struct {
	XMLName                           xml.Name                                 `xml:"CurrencyTransactionActivity"`
	SeqNum                            fincen.SeqNumber                         `xml:"SeqNum,attr"`
	InstallmentPaymentOtherIndicator  *fincen.ValidateIndicatorNullType        `xml:"InstallmentPaymentOtherIndicator,omitempty" json:",omitempty"`
	Total100DollarBillInAmountText    string                                   `xml:"Total100DollarBillInAmountText,omitempty" json:",omitempty"`
	TotalCashInReceiveAmountText      string                                   `xml:"TotalCashInReceiveAmountText"`
	TotalPriceAmountText              string                                   `xml:"TotalPriceAmountText,omitempty" json:",omitempty"`
	TransactionDateText               fincen.DateYYYYMMDDType                  `xml:"TransactionDateText"`
	CurrencyTransactionActivityDetail []*CurrencyTransactionActivityDetailType `xml:"CurrencyTransactionActivityDetail"`
}

func (r CurrencyTransactionActivityType) fieldInclusion() error {
	if len(r.CurrencyTransactionActivityDetail) < 2 || len(r.CurrencyTransactionActivityDetail) > 9 {
		return fincen.NewErrMinMaxRange("CurrencyTransactionActivityDetail")
	}

	return nil
}

func (r CurrencyTransactionActivityType) Validate(args ...string) error {
	if err := r.fieldInclusion(); err != nil {
		return err
	}

	return fincen.Validate(&r, args...)
}

type ActivityNarrativeInformationType struct {
	XMLName                         xml.Name                                `xml:"ActivityNarrativeInformation"`
	SeqNum                          fincen.SeqNumber                        `xml:"SeqNum,attr"`
	ActivityNarrativeSequenceNumber ValidateActivityNarrativeSequenceNumber `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString750                `xml:"ActivityNarrativeText"`
}

func (r ActivityNarrativeInformationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityDetailType struct {
	XMLName                                   xml.Name                                           `xml:"CurrencyTransactionActivityDetail"`
	SeqNum                                    fincen.SeqNumber                                   `xml:"SeqNum,attr"`
	CurrencyTransactionActivityDetailTypeCode *ValidateCurrencyTransactionActivityDetailCodeType `xml:"CurrencyTransactionActivityDetailTypeCode,omitempty" json:",omitempty"`
	DetailTransactionAmountText               *fincen.RestrictString15                           `xml:"DetailTransactionAmountText,omitempty" json:",omitempty"`
	DetailTransactionDescription              string                                             `xml:"DetailTransactionDescription,omitempty" json:",omitempty"`
	InstrumentProductServiceTypeCode          *fincen.ValidateInstrumentProductServiceTypeCode   `xml:"InstrumentProductServiceTypeCode,omitempty" json:",omitempty"`
	IssuerNameText                            string                                             `xml:"IssuerNameText,omitempty" json:",omitempty"`
	OtherForeignCurrencyCountryText           *fincen.RestrictString2                            `xml:"OtherForeignCurrencyCountryText,omitempty" json:",omitempty"`
}

func (r CurrencyTransactionActivityDetailType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}
