package uc

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
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
	CreateNewUser(newUserFields map[domain.UpdatableProperty]any) (user *domainuser.User, err error)
}

type AssetLogic interface {
	GetAssetsByUserId(userUUID uuid.UUID) ([]*domainasset.Asset, error)
	GetTransactionsByAssetId(assetId uuid.UUID) ([]*domainasset.Transaction, error)
	CreateNewAsset(ownerId *uuid.UUID, newAssetFields map[domain.UpdatableProperty]any) (*domainasset.Asset, error)
}
