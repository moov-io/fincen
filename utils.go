// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fincen

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var (
	DefaultValidateFunction = "Validate"

	dateTextRegex = regexp.MustCompile(`^\d{4}(0?[1-9]|1[012])(0?[1-9]|[12][0-9]|3[01])$`)
)

// NumericStringField return number string with filling zero
func NumericStringField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[ln-max:]
	}
	s = strings.Repeat("0", int(max-ln)) + s
	return s
}

func getTypeName(value string) string {
	values := strings.Split(value, ".")
	if len(values) > 1 {
		values := strings.Split(values[1], " ")
		return values[0]
	} else {
		return values[0]
	}
}

func validateCallbackByValue(data reflect.Value, args ...string) error {
	method := data.MethodByName(DefaultValidateFunction)
	if method.IsValid() {

		var response []reflect.Value
		_, ok := method.Interface().(func(args ...string))
		if ok {
			var converted []reflect.Value
			for _, arg := range args {
				converted = append(converted, reflect.ValueOf(arg))
			}
			response = method.Call(converted)
		}

		response = method.Call(response)
		if len(response) > 0 {
			err := response[0]
			if !err.IsNil() {
				typeName := getTypeName(data.String())
				if len(typeName) > 0 {
					errStr := err.Interface().(error).Error()
					if !strings.Contains(errStr, ")") {
						errStr = errStr + " (" + typeName + ")"
					} else {
						errStr = errStr[:len(errStr)-1] + ", " + typeName + ")"
					}
					return errors.New(errStr)
				}
				return err.Interface().(error)
			}
		}
	}
	return nil
}

// Validate is trying to run Validate(...string) function of fields
func Validate(r interface{}, args ...string) error {
	var err error
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldData := fields.Field(i)
		kind := fieldData.Kind()
		if kind == reflect.Slice {
			for i := 0; i < fieldData.Len(); i++ {
				err = validateCallbackByValue(fieldData.Index(i), args...)
				if err != nil {
					return err
				}
			}
		} else if kind == reflect.Map {
			for _, key := range fieldData.MapKeys() {
				err = validateCallbackByValue(fieldData.MapIndex(key), args...)
				if err != nil {
					return err
				}
			}
		} else if kind == reflect.Ptr {
			if fieldData.Pointer() != 0 {
				err = validateCallbackByValue(fieldData, args...)
				if err != nil {
					return err
				}
			}
		} else {
			err = validateCallbackByValue(fieldData, args...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ValidateDateText check input string is date string (YYYYMMDD)
func ValidateDateText(str string) bool {
	return !dateTextRegex.MatchString(str)
}

// ValidateNumericCharacters check input string is numeric characters
func ValidateNumericCharacters(str string, min, max int) bool {
	fmtStr := fmt.Sprintf(`[0-9]{%d,%d}`, min, max)
	reg := regexp.MustCompile(fmtStr)
	return !reg.MatchString(str)
}

func CheckInvolved(element string, elements ...string) bool {
	for _, elm := range elements {
		if element == elm {
			return true
		}
	}
	return false
}
