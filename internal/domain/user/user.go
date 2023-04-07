package domainuser

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

type User struct {
	Id           domain.Id
	Name         string
	Email        string
	PasswordHash []byte
}

type UserUpdatableProperty uint8

var (
	validator = userValidator{}
)

const (
	Name UserUpdatableProperty = iota
	Email
	Password
)

func CreateUser(opts ...*func(u *User) error) (*User, error) {
	newUser := User{
		Id: domain.Id(uuid.New()),
	}

	for _, f := range opts {
		err := (*f)(&newUser)
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
