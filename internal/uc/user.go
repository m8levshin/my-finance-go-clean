package uc

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

func (k *keeper) GetAllUsers() (users []*domainuser.User, err error) {
	return (k.userRw).FindAll()
}

func (k *keeper) GetUserById(id *domain.Id) (user *domainuser.User, err error) {
	return (k.userRw).FindById(*id)
}
