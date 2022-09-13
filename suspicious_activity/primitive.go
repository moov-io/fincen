// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package suspicious_activity

import (
	"reflect"
	"regexp"

	"github.com/moov-io/fincen"
)

// May be one of 1, 2, 3, 4, 5
type ValidateActivityNarrativeSequenceNumber int

func (r ValidateActivityNarrativeSequenceNumber) Validate() error {
	for _, vv := range []int{
		1, 2, 3, 4, 5,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidateActivityNarrativeSequenceNumber")
}

// May be one of 35, 37, 30, 33, 34, 8, 46, 18, 19, 41
type ValidateActivityPartyCodeType string

func (r ValidateActivityPartyCodeType) Validate() error {
	for _, vv := range []string{
		"35", "37", "30", "33", "34", "8", "46", "18", "19", "41",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidateActivityPartyCode")
}

// May be one of 1, 2, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 28, 32, 33, 29, 999
type ValidatePartyIdentificationCodeType string

func (r ValidatePartyIdentificationCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "4", "5", "6", "7", "9", "10", "11", "12", "13", "14", "28", "32", "33", "29", "999",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidatePartyIdentificationCode")
}

// May be one of L, AKA, DBA
type ValidatePartyNameCodeType string

func (r ValidatePartyNameCodeType) Validate() error {
	for _, vv := range []string{
		"L", "AKA", "DBA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidatePartyNameCode")
}

// May be one of 9, 1, 2, 7, 3, 4, 6, 13, 99
type ValidateFederalRegulatorCodeType string

func (r ValidateFederalRegulatorCodeType) Validate() error {
	for _, vv := range []string{
		"9", "1", "2", "7", "3", "4", "6", "13", "99",
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

// May be one of 5, 7
type ValidatePartyAccountAssociationCodeType string

func (r ValidatePartyAccountAssociationCodeType) Validate() error {
	for _, vv := range []string{
		"5", "7",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidatePartyAccountAssociationCode")
}
