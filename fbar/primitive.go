// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fbar

import (
	"reflect"

	"github.com/moov-io/fincen"
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
