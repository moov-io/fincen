// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package currency_transaction

import (
	"reflect"

	"github.com/moov-io/fincen"
)

// May be one of 35, 37, 30, 34, 50, 17, 23, 58, 8
type ValidateActivityPartyCodeType string

func (r ValidateActivityPartyCodeType) Validate() error {
	for _, vv := range []string{
		"35", "37", "30", "34", "50", "17", "23", "58", "8",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidateActivityPartyCode")
}

// May be one of 1, 2, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 28, 999
type ValidatePartyIdentificationCodeType string

func (r ValidatePartyIdentificationCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "4", "5", "6", "7", "9", "10", "11", "12", "13", "14", "28", "999",
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

// May be one of 9, 1, 2, 7, 3, 4, 6, 14
type ValidateFederalRegulatorCodeType string

func (r ValidateFederalRegulatorCodeType) Validate() error {
	for _, vv := range []string{
		"9", "1", "2", "7", "3", "4", "6", "14",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidateFederalRegulatorCode")
}

// May be one of 8, 9
type ValidatePartyAccountAssociationCodeType string

func (r ValidatePartyAccountAssociationCodeType) Validate() error {
	for _, vv := range []string{
		"8", "9",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("ValidatePartyAccountAssociationCode")
}

// May be one of "55", "46", "23", "12", "14", "49", "18", "21", "25", "997", "53",
type ValidateCurrencyTransactionActivityDetailCodeType string

func (r ValidateCurrencyTransactionActivityDetailCodeType) Validate() error {
	for _, vv := range []string{
		"55", "46", "23", "12", "14", "49", "18", "21", "25", "997", "53",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return fincen.NewErrValueInvalid("CurrencyTransactionActivityDetailCodeType")
}
