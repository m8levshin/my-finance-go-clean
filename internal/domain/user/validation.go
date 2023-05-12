package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func validateForCreateAndUpdate(u *User) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.Name, validation.Required, validation.Length(3, 0)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.PasswordHash, validation.Required),
	)
}

func checkRawPassword(pwd string) error {
	return validation.Validate(pwd, validation.Length(6, 100))
}
