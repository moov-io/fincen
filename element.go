// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fincen

// Element defines interface of complex xml elements
type Element interface {
	Validate(args ...string) error
}

// Element defines interface of complex xml elements
type ElementActivity interface {
	Validate(args ...string) error
	FormTypeCode() string
	TotalAmount() float64
	PartyCount(args ...string) int64
}
