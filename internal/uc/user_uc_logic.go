package uc

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

func (k *keeper) GetAllUsers() (users []*domainuser.User, err error) {
	return (k.userRw).FindAll()
}

func (k *keeper) GetUserById(uuid uuid.UUID) (user *domainuser.User, err error) {
	return (k.userRw).FindById(domain.Id(uuid))
}

func (k *keeper) CreateNewUser(
	newUserFields map[domainuser.UserUpdatableProperty]*string,
) (user *domainuser.User, err error) {
	createdUser, err := domainuser.CreateUser(
		domainuser.SetName(newUserFields[domainuser.Name]),
		domainuser.SetEmail(newUserFields[domainuser.Email]),
		domainuser.SetPassword(newUserFields[domainuser.Password]),
	)
	if err != nil {
		return nil, err
	}
	err = k.userRw.Save(*createdUser)
	if err != nil {
		return nil, err
	}
	return createdUser, nil

}
