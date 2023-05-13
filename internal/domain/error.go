package domain

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Error error

type ValidationError struct {
	CauseError   *error
	FailedFields map[string]string
}

func (e ValidationError) Error() string {
	return (*e.CauseError).Error()
}

func NewError(msg string) Error {
	return Error(errors.New(msg))
}

func ConvertValidationError(e error) error {

	validationError, ok := e.(validation.Errors)

	if ok {
		errorFieldsErrors := map[string]error(validationError)
		errorFieldsDescription := map[string]string{}
		for fieldName, fieldError := range errorFieldsErrors {
			errorFieldsDescription[fieldName] = fieldError.Error()
		}
		return &ValidationError{
			CauseError:   &e,
			FailedFields: errorFieldsDescription,
		}

	}

	return e
}
