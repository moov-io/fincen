// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fincen

import (
	"reflect"
	"regexp"
)

// Must match the pattern \S+( +\S+)*|
type RestrictString100 string

func (r RestrictString100) Validate() error {
	if len(string(r)) > 100 {
		return NewErrValueInvalid("RestrictString100")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString15 string

func (r RestrictString15) Validate() error {
	if len(string(r)) > 15 {
		return NewErrValueInvalid("RestrictString15")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString150 string

func (r RestrictString150) Validate() error {
	if len(string(r)) > 150 {
		return NewErrValueInvalid("RestrictString150")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString16 string

func (r RestrictString16) Validate() error {
	if len(string(r)) > 16 {
		return NewErrValueInvalid("RestrictString16")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString2 string

func (r RestrictString2) Validate() error {
	if len(string(r)) > 2 {
		return NewErrValueInvalid("RestrictString2")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString25 string

func (r RestrictString25) Validate() error {
	if len(string(r)) > 25 {
		return NewErrValueInvalid("RestrictString25")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString3 string

func (r RestrictString3) Validate() error {
	if len(string(r)) > 3 {
		return NewErrValueInvalid("RestrictString3")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString30 string

func (r RestrictString30) Validate() error {
	if len(string(r)) > 30 {
		return NewErrValueInvalid("RestrictString30")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString35 string

func (r RestrictString35) Validate() error {
	if len(string(r)) > 35 {
		return NewErrValueInvalid("RestrictString35")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString39 string

func (r RestrictString39) Validate() error {
	if len(string(r)) > 39 {
		return NewErrValueInvalid("RestrictString39")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString40 string

func (r RestrictString40) Validate() error {
	if len(string(r)) > 40 {
		return NewErrValueInvalid("RestrictString40")
	}
	return nil
}

// May be no more than 4000 items long
type RestrictString4000 string

func (r RestrictString4000) Validate() error {
	if len(string(r)) > 4000 {
		return NewErrValueInvalid("RestrictString4000")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString50 string

func (r RestrictString50) Validate() error {
	if len(string(r)) > 50 {
		return NewErrValueInvalid("RestrictString50")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString517 string

func (r RestrictString517) Validate() error {
	if len(string(r)) > 517 {
		return NewErrValueInvalid("RestrictString517")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString6 string

func (r RestrictString6) Validate() error {
	if len(string(r)) > 6 {
		return NewErrValueInvalid("RestrictString6")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString9 string

func (r RestrictString9) Validate() error {
	if len(string(r)) > 9 {
		return NewErrValueInvalid("RestrictString9")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString750 string

func (r RestrictString750) Validate() error {
	if len(string(r)) > 750 {
		return NewErrValueInvalid("RestrictString750")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString4 string

func (r RestrictString4) Validate() error {
	if len(string(r)) > 4 {
		return NewErrValueInvalid("RestrictString4")
	}
	return nil
}

// Must match the pattern \S+( +\S+)*|
type RestrictString20 string

func (r RestrictString20) Validate() error {
	if len(string(r)) > 20 {
		return NewErrValueInvalid("RestrictString20")
	}
	return nil
}

// Must match the pattern ([0-1][0-9]|(2[0-3])):[0-5][0-9]:[0-5][0-9]|
type ValidateTimeDataOrBlankType string

func (r ValidateTimeDataOrBlankType) Validate() error {
	reg := regexp.MustCompile(`([0-1][0-9]|(2[0-3])):[0-5][0-9]:[0-5][0-9]|`)
	if !reg.MatchString(string(r)) {
		return NewErrValueInvalid("ValidateTimeDataOrBlankType")
	}
	return nil
}

// Must match the pattern (19|20)[0-9][0-9](0[1-9]|1[0-2])(0[1-9]|1[0-9]|2[0-9]|3[01])|
type DateYYYYMMDDOrBlankType string

func (r DateYYYYMMDDOrBlankType) Validate() error {
	reg := regexp.MustCompile(`(19|20)[0-9][0-9](0[1-9]|1[0-2])(0[1-9]|1[0-9]|2[0-9]|3[01])|`)
	if !reg.MatchString(string(r)) {
		return NewErrValueInvalid("DateYYYYMMDDOrBlankType")
	}
	return nil
}

// Must match the pattern (19|20)[0-9][0-9](0[0-9]|1[0-2])(0[0-9]|1[0-9]|2[0-9]|3[01])|
type DateYYYYMMDDOrBlankTypeDOB string

func (r DateYYYYMMDDOrBlankTypeDOB) Validate() error {
	reg := regexp.MustCompile(`(19|20)[0-9][0-9](0[0-9]|1[0-2])(0[0-9]|1[0-9]|2[0-9]|3[01])|`)
	if !reg.MatchString(string(r)) {
		return NewErrValueInvalid("DateYYYYMMDDOrBlankTypeDOB")
	}
	return nil
}

// Must match the pattern (19|20)[0-9][0-9](0[1-9]|1[0-2])(0[1-9]|1[0-9]|2[0-9]|3[01])
type DateYYYYMMDDType string

func (r DateYYYYMMDDType) Validate() error {
	reg := regexp.MustCompile(`(19|20)[0-9][0-9](0[1-9]|1[0-2])(0[1-9]|1[0-9]|2[0-9]|3[01])`)
	if !reg.MatchString(string(r)) {
		return NewErrValueInvalid("DateYYYYMMDDType")
	}
	return nil
}

// Must match the pattern (19|20)[0-9][0-9]
type DateYYYYType string

func (r DateYYYYType) Validate() error {
	reg := regexp.MustCompile(`(19|20)[0-9][0-9]`)
	if !reg.MatchString(string(r)) {
		return NewErrValueInvalid("DateYYYYType")
	}
	return nil
}

// May be one of 1, 2, 3, 4
type ValidateAssetAttributeTypeIDTypeCode int64

func (r ValidateAssetAttributeTypeIDTypeCode) Validate() error {
	for _, vv := range []int{
		1, 2, 3, 4,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateAssetAttributeTypeIDTypeCode")
}

// May be one of 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16, 17, 18, 19, 20, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 41, 46, 47
type ValidateAssetSubtypeIDTypeCode int64

func (r ValidateAssetSubtypeIDTypeCode) Validate() error {
	for _, vv := range []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16, 17, 18, 19, 20, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 41, 46, 47,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateAssetSubtypeIDTypeCode")
}

// May be one of 5, 6
type ValidateAssetTypeIDTypeCode int64

func (r ValidateAssetTypeIDTypeCode) Validate() error {
	for _, vv := range []int{
		5, 6,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateAssetTypeIDTypeCode")
}

// May be one of 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 999
type ValidateCyberEventIndicatorsTypeCode string

func (r ValidateCyberEventIndicatorsTypeCode) Validate() error {
	for _, vv := range []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "999",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateCyberEventIndicatorsTypeCode")
}

// May be one of E, U
type ValidateElectronicAddressTypeCode string

func (r ValidateElectronicAddressTypeCode) Validate() error {
	for _, vv := range []string{
		"E", "U",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateElectronicAddressTypeCode")
}

// May be one of Y,
type ValidateIndicatorType string

func (r ValidateIndicatorType) Validate() error {
	for _, vv := range []string{
		"Y",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateIndicatorType")
}

// May be one of 1, 2, 3, 4, 5, 11, 12, 999
type ValidateOrganizationCodeType int

func (r ValidateOrganizationCodeType) Validate() error {
	for _, vv := range []int{
		1, 2, 3, 4, 5, 11, 12, 999,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateOrganizationCodeType")
}

// May be one of 101, 102, 103, 503, 504, 508, 513, 514, 535, 528, 529, 533, 534, 539, 540, 541, 542, 1999, 5999
type ValidateOrganizationSubtypeCodeSarType int

func (r ValidateOrganizationSubtypeCodeSarType) Validate() error {
	for _, vv := range []int{
		101, 102, 103, 503, 504, 508, 513, 514, 535, 528, 529, 533, 534, 539, 540, 541, 542, 1999, 5999,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateOrganizationSubtypeCodeSarType")
}

// May be one of F, M, R, W
type ValidatePhoneNumberCodeType string

func (r ValidatePhoneNumberCodeType) Validate() error {
	for _, vv := range []string{
		"F", "M", "R", "W",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidatePhoneNumberCodeType")
}

// May be one of 106, 111, 112, 113, 114, 301, 304, 305, 308, 309, 310, 312, 320, 321, 322, 323, 324, 325, 401, 402, 403, 404, 405, 409, 501, 502, 504, 505, 506, 507, 601, 603, 604, 608, 609, 701, 801, 804, 805, 806, 807, 808, 809, 812, 820, 821, 822, 823, 824, 901, 903, 904, 905, 907, 908, 909, 910, 911, 913, 917, 920, 921, 922, 924, 925, 926, 927, 928, 1001, 1003, 1005, 1006, 1007, 1101, 1102, 1201, 1202, 1203, 1204, 1999, 3999, 4999, 5999, 6999, 7999, 8999, 9999, 10999, 11999, 12999
type ValidateSuspiciousActivitySubtypeID int

func (r ValidateSuspiciousActivitySubtypeID) Validate() error {
	for _, vv := range []int{
		106, 111, 112, 113, 114, 301, 304, 305, 308, 309, 310, 312, 320, 321, 322, 323, 324, 325, 401, 402, 403, 404, 405, 409, 501, 502, 504, 505, 506, 507, 601, 603, 604, 608, 609, 701, 801, 804, 805, 806, 807, 808, 809, 812, 820, 821, 822, 823, 824, 901, 903, 904, 905, 907, 908, 909, 910, 911, 913, 917, 920, 921, 922, 924, 925, 926, 927, 928, 1001, 1003, 1005, 1006, 1007, 1101, 1102, 1201, 1202, 1203, 1204, 1999, 3999, 4999, 5999, 6999, 7999, 8999, 9999, 10999, 11999, 12999,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateSuspiciousActivitySubtypeID")
}

// May be one of 1, 12, 3, 4, 5, 6, 7, 8, 9, 10, 11
type ValidateSuspiciousActivityTypeID int

func (r ValidateSuspiciousActivityTypeID) Validate() error {
	for _, vv := range []int{
		1, 12, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateSuspiciousActivityTypeID")
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
	return NewErrValueInvalid("ValidateCurrencyTransactionActivityDetailCodeType")
}

// May be one of 35, 16, 39, 26, 40, 34
type ValidateInstrumentProductServiceTypeCode int

func (r ValidateInstrumentProductServiceTypeCode) Validate() error {
	for _, vv := range []int{
		35, 16, 39, 26, 40, 34,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateInstrumentProductServiceTypeCode")
}

// May be one of I, O, U
type ValidatePartyTypeCode string

func (r ValidatePartyTypeCode) Validate() error {
	for _, vv := range []string{
		"I", "O", "U",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidatePartyTypeCode")
}

// May be one of 101, 102, 103, 1999
type ValidateOrganizationSubtypeCodeCtrType int

func (r ValidateOrganizationSubtypeCodeCtrType) Validate() error {
	for _, vv := range []int{
		101, 102, 103, 1999,
	} {
		if reflect.DeepEqual(int(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateOrganizationSubtypeCodeCtrType")
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
	return NewErrValueInvalid("ValidateIndicatorYNType")
}

// May be one of 1, 2, 3, 4, 5, 6, 7, 8, 9, 999
type ValidateLateFilingReasonCodeType string

func (r ValidateLateFilingReasonCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "999",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateLateFilingReasonCodeType")
}

// May be one of 141, 142, 143, 144
type ValidateEFilingAccountTypeCodeType string

func (r ValidateEFilingAccountTypeCodeType) Validate() error {
	for _, vv := range []string{
		"141", "142", "143", "144",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateEFilingAccountTypeCodeType")
}

// May be one of 1, 2, 999
type ValidateAccountTypeCodeType string

func (r ValidateAccountTypeCodeType) Validate() error {
	for _, vv := range []string{
		"1", "2", "999",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateAccountTypeCodeType")
}

// May be one of C, D, E, F
type ValidateExemptBasisTypeCode string

func (r ValidateExemptBasisTypeCode) Validate() error {
	for _, vv := range []string{
		"C", "D", "E", "F",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return NewErrValueInvalid("ValidateExemptBasisTypeCode")
}

// May be one of C, D, E, F
type SeqNumber int64

func (r SeqNumber) Validate() error {
	if int64(r) == 0 {
		return NewErrValueInvalid("SeqNumber")
	}
	return nil
}
