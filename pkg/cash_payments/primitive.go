// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package cash_payments

import (
	"reflect"

	"github.com/moov-io/fincen"
)

// May be one of 1
type ValidateActivityNarrativeSequenceNumber int

func (r ValidateActivityNarrativeSequenceNumber) Validate() error {
	for _, vv := range []int{
		1,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ActivityNarrativeSequenceNumber")
}

// May be one of 35, 37, 16, 23, 4, 8, 3
type ValidateActivityPartyCodeType string

func (r ValidateActivityPartyCodeType) Validate() error {
	for _, vv := range []string{
		"35", "37", "16", "23", "4", "8", "3",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ActivityPartyCode")
}

// May be one of 1, 2, 5, 6, 7, 28, 999
type ValidatePartyIdentificationCodeType string

func (r ValidatePartyIdentificationCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "5", "6", "7", "28", "999",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("PartyIdentificationCode")
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
	return fincen.NewErrValueInvalid("PartyNameCode")
}

// May be one of 1, 2, 3, 4, 5, 6, 7, 8, 9, 999
type ValidateCurrencyTransactionActivityDetailCodeType string

func (r ValidateCurrencyTransactionActivityDetailCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "999",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("CurrencyTransactionActivityDetailCodeType")
}
