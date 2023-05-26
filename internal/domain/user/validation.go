package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

func validateForCreateAndUpdate(u *User) error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.Name, validation.Required, validation.Length(3, 0)),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
	return domain.ConvertValidationError(err)
}

func checkRawPassword(pwd string) error {
	err := validation.Validate(pwd, validation.Length(6, 100))
	return domain.ConvertValidationError(err)
}
