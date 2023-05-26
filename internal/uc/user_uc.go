package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

func (k *handler) GetAllUsers() (users []*domainuser.User, err error) {
	return (k.userRw).FindAll()
}

func (k *handler) GetUserById(uuid uuid.UUID) (user *domainuser.User, err error) {
	return (k.userRw).FindById(domain.Id(uuid))
}

func (k *handler) GetUserByEmail(email string) (user *domainuser.User, err error) {
	return (k.userRw).FindByEmail(email)
}

func (k *handler) CreateNewUser(
	newUserFields map[domain.UpdatableProperty]any,
) (user *domainuser.User, err error) {

	var name = (newUserFields[domainuser.NameField]).(*string)
	var email = (newUserFields[domainuser.EmailField]).(*string)

	user, err = k.userRw.FindByEmail(*email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("user with that email is already exist")
	}

	createdUser, err := k.userService.CreateUser(
		domainuser.SetName(name),
		domainuser.SetEmail(email),
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
