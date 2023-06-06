package user

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

const (
	NameField domain.UpdatableProperty = iota
	EmailField
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
