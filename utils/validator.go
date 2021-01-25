package utils

import (
	"github.com/go-playground/validator/v10"
)

type Validate struct {
	Validator *validator.Validate
}

func (cv *Validate) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func ValidateRequest(i interface{}) error {
	v := &Validate{Validator: validator.New()}
	err := v.Validate(i)
	if err != nil {
		return err
	}
	return nil
}


