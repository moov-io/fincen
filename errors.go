// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fincen

import "fmt"

// NewErrTextLength returns a error that the length of value is invalid
func NewErrValueInvalid(typeStr string) error {
	errStr := fmt.Sprintf("The %s is invalid", typeStr)
	return fmt.Errorf(errStr)
}

// NewErrFieldRequired returns a error when a field is required
func NewErrFieldRequired(typeStr string) error {
	errStr := fmt.Sprintf("The %s is a required field", typeStr)
	return fmt.Errorf(errStr)
}

// NewErrFiledNotAssociated returns a error that the field is associated
func NewErrFiledNotAssociated(typeStr string) error {
	errStr := fmt.Sprintf("The %s is not associated", typeStr)
	return fmt.Errorf(errStr)
}

// NewErrMinMaxRange returns a error that the field has min/max element range
func NewErrMinMaxRange(typeStr string) error {
	errStr := fmt.Sprintf("The %s has invalid min & max range", typeStr)
	return fmt.Errorf(errStr)
}
