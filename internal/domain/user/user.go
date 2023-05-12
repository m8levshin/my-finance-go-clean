package user

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

type User struct {
	Id           domain.Id
	Name         string
	Email        string
	PasswordHash []byte
}

var (
	validator = userValidator{}
)

func CreateUser(opts ...func(u *User) error) (*User, error) {

	newUser := User{
		Id: domain.NewID(),
	}
	for _, f := range opts {
		err := f(&newUser)
		if err != nil {
			return nil, err
		}
	}
	err := validator.validateForCreateAndUpdate(&newUser)
	if err != nil {
		return nil, err

	}

	return &newUser, nil
}

func UpdateUser(initial *User, opts ...*func(u *User) error) (*User, error) {
	for _, function := range opts {
		f := *function
		err := f(initial)
		if err != nil {
			return nil, err
		}
	}
	if err := validator.validateForCreateAndUpdate(initial); err != nil {
		return nil, err
	}
	return initial, nil
}
