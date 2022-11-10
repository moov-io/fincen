// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package financial_accounts

import (
	"reflect"

	"github.com/moov-io/fincen"
)

// May be one of 15, 35, 37, 56, 57
type ValidateActivityPartyCodeType string

func (r ValidateActivityPartyCodeType) Validate() error {
	for _, vv := range []string{
		"15", "35", "37", "56", "57",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ActivityPartyCode")
}

// May be one of 41, 42, 43, 44
type ValidateAccountPartyCodeType string

func (r ValidateAccountPartyCodeType) Validate() error {
	for _, vv := range []string{
		"41", "42", "43", "44",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("AccountPartyCode")
}

// May be one of 1, 2, 4, 6, 9, 28, 31, 999
type ValidateActivityPartyIdentificationCodeType string

func (r ValidateActivityPartyIdentificationCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "4", "6", "9", "28", "31", "999",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ActivityPartyIdentificationCode")
}

// May be one of -2, 1, 2, 9
type ValidateAccountPartyIdentificationCodeType string

func (r ValidateAccountPartyIdentificationCodeType) Validate() error {
	for _, vv := range []string{
		"-2", "1", "2", "9",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("AccountPartyIdentificationCode")
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
	return fincen.NewErrValueInvalid("PartyNameCode")
}

// May be one of Y, N,
type ValidateIndicatorYNType string

func (r ValidateIndicatorYNType) Validate() error {
	for _, vv := range []string{
		"Y", "N",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("IndicatorYNType")
}
