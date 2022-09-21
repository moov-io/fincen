// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package exempt_designation

import (
	"reflect"
	"regexp"

	"github.com/moov-io/fincen"
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
	return fincen.NewErrValueInvalid("ValidateActivityPartyCode")
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
	return fincen.NewErrValueInvalid("ValidatePartyIdentificationCode")
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
	return fincen.NewErrValueInvalid("ValidatePartyNameCode")
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
	return fincen.NewErrValueInvalid("ValidateFederalRegulatorCode")
}

// 14-digit numeric
type EFilingPriorDocumentNumberType string

func (r EFilingPriorDocumentNumberType) Validate() error {
	reg := regexp.MustCompile(`[0-9]{14}`)
	if !reg.MatchString(string(r)) {
		return fincen.NewErrValueInvalid("EFilingPriorDocumentNumber")
	}
	return nil
}

func (r EFilingPriorDocumentNumberType) String() string {
	return fincen.NumericStringField(string(r), 14)
}
