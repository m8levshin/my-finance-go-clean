package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Validator interface {
	validateForCreateAndUpdate(u *User) error
	checkRawPassword(pwd string) error
}

type userValidator struct{}

func (v *userValidator) validateForCreateAndUpdate(u *User) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.Name, validation.Required, validation.Length(3, 0)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.PasswordHash, validation.Required),
	)
}

func (v *userValidator) checkRawPassword(pwd string) error {
	return validation.Validate(pwd, validation.Length(10, 100))
}
