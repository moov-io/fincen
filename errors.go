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
