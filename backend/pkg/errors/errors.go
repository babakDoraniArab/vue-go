package errors

import (
	"errors"

	"github.com/go-playground/validator"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFromErros(err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, err := range validationErrors {
			errorsList[err.Field()] = err.Tag()
		}
	}
}
