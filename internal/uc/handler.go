package uc

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

type Handler interface {
	UserLogic
}

type UserLogic interface {
	GetAllUsers() (users []*domainuser.User, err error)
	GetUserById(id *domain.Id) (user *domainuser.User, err error)
}
