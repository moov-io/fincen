// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package exempt_designation

import (
	"reflect"

	"github.com/moov-io/fincen"
)

// validating elements
var (
	_ fincen.Element = (*EFilingBatchXML)(nil)
	_ fincen.Element = (*Activity)(nil)
	_ fincen.Element = (*ActivityAssociationType)(nil)
	_ fincen.Element = (*ActivityType)(nil)
	_ fincen.Element = (*AddressType)(nil)
	_ fincen.Element = (*DesignationExemptActivityType)(nil)
	_ fincen.Element = (*ElectronicAddressType)(nil)
	_ fincen.Element = (*Party)(nil)
	_ fincen.Element = (*PartyIdentificationType)(nil)
	_ fincen.Element = (*PartyNameType)(nil)
	_ fincen.Element = (*PartyOccupationBusinessType)(nil)
	_ fincen.Element = (*PartyType)(nil)
	_ fincen.Element = (*PhoneNumberType)(nil)
)

// May be one of 35, 37, 11, 45, 12, 3
type ValidateActivityPartyCodeType string

func (r ValidateActivityPartyCodeType) Validate() error {
	for _, vv := range []string{
		"35", "37", "11", "45", "12", "3",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidateActivityPartyCodeType")
}

// May be one of 1, 2, 14, 28
type ValidatePartyIdentificationCodeType string

func (r ValidatePartyIdentificationCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "14", "28",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidatePartyIdentificationCodeType")
}

// May be one of L, DBA
type ValidatePartyNameCodeType string

func (r ValidatePartyNameCodeType) Validate() error {
	for _, vv := range []string{
		"L", "DBA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidatePartyNameCodeType")
}

// May be one of 1, 2, 7, 3, 4
type ValidateFederalRegulatorCodeType string

func (r ValidateFederalRegulatorCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "7", "3", "4",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidateFederalRegulatorCodeType")
}