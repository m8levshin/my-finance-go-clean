package user

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

const (
	NameField domain.UpdatableProperty = iota
	EmailField
	PasswordField
)

func SetName(name *string) func(u *User) error {
	return func(u *User) error {
		u.Name = *name
		return nil
	}
}

func SetEmail(email *string) func(u *User) error {
	return func(u *User) error {
		u.Email = *email
		return nil
	}
}

func SetPassword(password *string) func(u *User) error {
	return func(u *User) error {
		if err := validator.checkRawPassword(*password); err != nil {
			return err
		}
		newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.PasswordHash = newPasswordHash
		return nil
	}
}
