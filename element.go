package fincen

// Element defines interface of complex xml elements
type Element interface {
	Validate(args ...string) error
}
