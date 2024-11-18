package validator

import "github.com/go-playground/validator/v10"

// main

var Validate *validator.Validate

func Init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
