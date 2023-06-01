package uc

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/config"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/transaction_group"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
)

type Handler interface {
	UserLogic
	AssetLogic
	TransactionLogic
	TransactionGroupLogic
}

type handler struct {
	config                  config.Configuration
	userRw                  rw.UserRW
	assetRw                 rw.AssetRW
	transactionRw           rw.TransactionRW
	transactionGroupRw      rw.TransactionGroupRW
	userService             domainuser.UserDomainService
	assetService            domainasset.AssetDomainService
	transactionGroupService transaction_group.TransactionGroupDomainService
}

func NewHandler(
	config config.Configuration,
	userRw rw.UserRW, assetRw rw.AssetRW,
	transactionRw rw.TransactionRW,
	transactionGroupRw rw.TransactionGroupRW,
	userService domainuser.UserDomainService,
	assetService domainasset.AssetDomainService,
	transactionGroupService transaction_group.TransactionGroupDomainService,
) *handler {
	return &handler{config: config, userRw: userRw, assetRw: assetRw, transactionRw: transactionRw,
		transactionGroupRw: transactionGroupRw, userService: userService, assetService: assetService,
		transactionGroupService: transactionGroupService}
}

type UserLogic interface {
	GetAllUsers() (users []*domainuser.User, err error)
	GetUserById(uuid uuid.UUID) (user *domainuser.User, err error)

	GetUserByEmail(email string) (user *domainuser.User, err error)
	CreateNewUser(newUserFields map[domain.UpdatableProperty]any) (user *domainuser.User, err error)
}

type AssetLogic interface {
	GetAssetsByUserId(userUUID uuid.UUID) ([]*domainasset.Asset, error)
	GetTransactionsByAssetId(assetId uuid.UUID) ([]*domainasset.Transaction, error)
	CreateNewAsset(userId uuid.UUID, newAssetFields map[domain.UpdatableProperty]any) (*domainasset.Asset, error)
	GetAssetById(assetId uuid.UUID) (*domainasset.Asset, error)
}

type TransactionLogic interface {
	AddNewTransaction(assetUUID uuid.UUID, d *dto.AddNewTransactionRequest, userUUID uuid.UUID, isAdmin bool) (*domainasset.Transaction, error)
}

type TransactionGroupLogic interface {
	GetTransactionGroupsByUser(userId uuid.UUID) ([]*transaction_group.TransactionGroup, error)
	CreateNewTransactionGroup(userId uuid.UUID, req dto.CreateTransactionGroupRequest) (*transaction_group.TransactionGroup, error)
}
