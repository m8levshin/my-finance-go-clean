package user

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

type UserRW interface {
	FindByEmail(email string) (*User, error)
	FindAll() ([]*User, error)
	FindById(id domain.Id) (*User, error)
	Save(user User) error
}
