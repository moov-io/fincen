// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fincen

import (
	"fmt"
)

// NewErrTextLength returns an error that the length of value is invalid
func NewErrValueInvalid(typeStr string) error {
	return fmt.Errorf("The %s has invalid value", typeStr)
}

// NewErrFieldRequired returns an error when a field is required
func NewErrFieldRequired(typeStr string) error {
	return fmt.Errorf("The %s is a required field", typeStr)
}

// NewErrFiledOmitted returns an error that the field should be omitted
func NewErrFiledOmitted(typeStr string) error {
	return fmt.Errorf("The %s should be omitted", typeStr)
}

// NewErrMinMaxRange returns an error that the field has min/max element range
func NewErrMinMaxRange(typeStr string) error {
	return fmt.Errorf("The %s has invalid min & max range", typeStr)
}
