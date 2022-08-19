// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package suspicious_activity

import (
	"reflect"

	"github.com/moov-io/fincen"
)

// May be one of 1, 2, 3, 4, 5
type ValidateActivityNarrativeSequenceNumber int

func (r ValidateActivityNarrativeSequenceNumber) Validate() error {
	for _, vv := range []string{
		"1", "2", "3", "4", "5",
	} {
		if reflect.DeepEqual(string(r), vv) {
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
	return fincen.NewErrValueInvalid("ValidateActivityPartyCodeType")
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
	return fincen.NewErrValueInvalid("ValidatePartyIdentificationCodeType")
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
	return fincen.NewErrValueInvalid("ValidatePartyNameCodeType")
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
	return fincen.NewErrValueInvalid("ValidateFederalRegulatorCodeType")
}
