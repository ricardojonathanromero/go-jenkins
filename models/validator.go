package models

import "gopkg.in/go-playground/validator.v9"

// Validator struct
type Validator struct {
	Validator *validator.Validate
}

// Validate func is used for validate structs
func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
