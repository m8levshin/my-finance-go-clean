package uc

import (
	"github.com/google/uuid"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

type Handler interface {
	UserLogic
	AssetLogic
}

type UserLogic interface {
	GetAllUsers() (users []*domainuser.User, err error)
	GetUserById(uuid uuid.UUID) (user *domainuser.User, err error)
	CreateNewUser(newUserFields map[domainuser.UserUpdatableProperty]*string) (user *domainuser.User, err error)
}

type AssetLogic interface {
	GetAssetsByUserId(userUUID uuid.UUID) ([]*domainasset.Asset, error)
}
