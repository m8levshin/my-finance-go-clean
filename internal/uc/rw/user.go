package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

type UserRW interface {
	FindAll() ([]*domainuser.User, error)
	FindById(id domain.Id) (*domainuser.User, error)
	Save(user domainuser.User) error
}