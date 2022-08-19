// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package financial_accounts

import (
	"reflect"

	"github.com/moov-io/fincen"
)

// validating elements
var (
	_ fincen.Element = (*EFilingBatchXML)(nil)
	_ fincen.Element = (*Account)(nil)
	_ fincen.Element = (*AccountType)(nil)
	_ fincen.Element = (*Activity)(nil)
	_ fincen.Element = (*ActivityAssociationType)(nil)
	_ fincen.Element = (*ActivityNarrativeInformation)(nil)
	_ fincen.Element = (*ActivityNarrativeInformationType)(nil)
	_ fincen.Element = (*ActivityType)(nil)
	_ fincen.Element = (*AddressType)(nil)
	_ fincen.Element = (*AccountParty)(nil)
	_ fincen.Element = (*ForeignAccountActivity)(nil)
	_ fincen.Element = (*ForeignAccountActivityType)(nil)
	_ fincen.Element = (*Party)(nil)
	_ fincen.Element = (*PartyIdentificationType)(nil)
	_ fincen.Element = (*PartyNameType)(nil)
	_ fincen.Element = (*PartyType)(nil)
	_ fincen.Element = (*PhoneNumberType)(nil)
)

// May be one of 15, 35, 37, 41, 42, 43, 44, 56, 57
type ValidateActivityPartyCodeType string

func (r ValidateActivityPartyCodeType) Validate() error {
	for _, vv := range []string{
		"15", "35", "37", "41", "42", "43", "44", "56", "57",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidateActivityPartyCodeType")
}

// May be one of -2, 1, 2, 4, 6, 9, 28, 31, 999
type ValidatePartyIdentificationCodeType string

func (r ValidatePartyIdentificationCodeType) Validate() error {
	for _, vv := range []string{
		"-2", "1", "2", "4", "6", "9", "28", "31", "999",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidatePartyIdentificationCodeType")
}

// May be one of L
type ValidatePartyNameCodeType string

func (r ValidatePartyNameCodeType) Validate() error {
	for _, vv := range []string{
		"L",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidatePartyNameCodeType")
}
