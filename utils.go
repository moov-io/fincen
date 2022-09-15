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

	SequenceFieldName = "SeqNum"

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
					if len(errStr) > 0 && string(errStr[len(errStr)-1:]) == ")" {
						errStr = errStr[:len(errStr)-1] + ", " + typeName + ")"
					} else {
						errStr = errStr + " (" + typeName + ")"
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

	seqMap := map[SeqNumber]bool{}
	checkSeqNumber := func(data reflect.Value) error {
		num := getSeqNum(data)
		if num < 0 {
			return nil
		}

		_, ok := seqMap[num]
		if ok {
			return fmt.Errorf("%s has dupulicated sequence number", getTypeName(data.String()))
		}

		seqMap[num] = true
		return nil
	}

	var err error
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldData := fields.Field(i)
		kind := fieldData.Kind()
		if kind == reflect.Slice {
			for i := 0; i < fieldData.Len(); i++ {
				elm := fieldData.Index(i)
				err = validateCallbackByValue(elm, args...)
				if err != nil {
					return err
				}
				if seqErr := checkSeqNumber(elm); seqErr != nil {
					return seqErr
				}
			}
		} else if kind == reflect.Map {
			for _, key := range fieldData.MapKeys() {
				elm := fieldData.MapIndex(key)
				err = validateCallbackByValue(fieldData.MapIndex(key), args...)
				if err != nil {
					return err
				}
				if seqErr := checkSeqNumber(elm); seqErr != nil {
					return seqErr
				}
			}
		} else if kind == reflect.Ptr {
			if fieldData.Pointer() != 0 {
				err = validateCallbackByValue(fieldData, args...)
				if err != nil {
					return err
				}
				if seqErr := checkSeqNumber(fieldData); seqErr != nil {
					return seqErr
				}
			}
		} else {
			err = validateCallbackByValue(fieldData, args...)
			if err != nil {
				return err
			}
			if seqErr := checkSeqNumber(fieldData); seqErr != nil {
				return seqErr
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

func getSeqNum(data reflect.Value) SeqNumber {

	ignoreValue := SeqNumber(-1)

	elm, ok := data.Interface().(Element)
	if !ok || elm == nil {
		return ignoreValue
	}

	kind := data.Kind()
	if kind == reflect.Ptr {
		data = reflect.Indirect(data)
	} else if kind == reflect.Interface {
		elm, ok = data.Interface().(ElementActivity)
		if !ok || elm == nil {
			return ignoreValue
		}
		data = reflect.Indirect(reflect.ValueOf(elm))
	}

	seqNum := data.FieldByName(SequenceFieldName)
	if !seqNum.IsValid() {
		return ignoreValue
	}

	return SeqNumber(seqNum.Int())

}

func setSeqNum(data reflect.Value, number SeqNumber) SeqNumber {

	elm, ok := data.Interface().(Element)
	if !ok || elm == nil {
		return number
	}

	kind := data.Kind()
	if kind == reflect.Interface {
		elm, ok = data.Interface().(ElementActivity)
		if !ok || elm == nil {
			return number
		}
		data = reflect.Indirect(reflect.ValueOf(elm))
	} else {
		data = reflect.Indirect(reflect.ValueOf(data))
	}

	seqNum := data.FieldByName(SequenceFieldName)
	if seqNum.Kind() != reflect.Int64 || !seqNum.IsValid() || !seqNum.CanSet() {
		return number
	}

	seqNum.SetInt(int64(number))
	return number + 1
}

// GenerateSeqNumbers generate unique sequence numbers
func GenerateSeqNumbers(r interface{}) error {

	seqNumber := SeqNumber(1)

	var fields reflect.Value
	if reflect.ValueOf(r).Kind() == reflect.Ptr {
		fields = reflect.ValueOf(r).Elem()
	} else {
		fields = reflect.ValueOf(r)
	}

	if fields.Kind() == reflect.Struct {
		fields = reflect.Indirect(reflect.ValueOf(r))
	} else {
		return nil
	}

	for i := 0; i < fields.NumField(); i++ {
		fieldData := fields.Field(i)
		kind := fieldData.Kind()
		if kind == reflect.Slice {
			for i := 0; i < fieldData.Len(); i++ {
				elm := fieldData.Index(i)
				seqNumber = setSeqNum(elm, seqNumber)
				if err := GenerateSeqNumbers(elm.Interface()); err != nil {
					return err
				}
			}
		} else if kind == reflect.Map {
			for _, key := range fieldData.MapKeys() {
				elm := fieldData.MapIndex(key)
				seqNumber = setSeqNum(elm, seqNumber)
				if err := GenerateSeqNumbers(elm.Interface()); err != nil {
					return err
				}
			}
		} else if kind == reflect.Ptr {
			if fieldData.Pointer() != 0 {
				seqNumber = setSeqNum(fieldData, seqNumber)
				if err := GenerateSeqNumbers(fieldData.Interface()); err != nil {
					return err
				}
			}
		} else if kind == reflect.Struct {
			seqNumber = setSeqNum(fieldData, seqNumber)
			if err := GenerateSeqNumbers(fieldData.Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}
