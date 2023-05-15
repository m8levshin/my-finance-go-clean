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

type HandlerBuilder struct {
	Config                  config.Configuration
	UserRw                  rw.UserRW
	AssetRw                 rw.AssetRW
	TransactionRw           rw.TransactionRW
	TransactionGroupRw      rw.TransactionGroupRW
	UserService             domainuser.UserDomainService
	AssetService            domainasset.AssetDomainService
	TransactionGroupService transaction_group.TransactionGroupDomainService
}

func (b HandlerBuilder) Build() Handler {
	return &handler{
		config:                  b.Config,
		userRw:                  b.UserRw,
		assetRw:                 b.AssetRw,
		transactionRw:           b.TransactionRw,
		transactionGroupRw:      b.TransactionGroupRw,
		userService:             b.UserService,
		assetService:            b.AssetService,
		transactionGroupService: b.TransactionGroupService,
	}
}

type UserLogic interface {
	GetAllUsers() (users []*domainuser.User, err error)
	GetUserById(uuid uuid.UUID) (user *domainuser.User, err error)
	CreateNewUser(newUserFields map[domain.UpdatableProperty]any) (user *domainuser.User, err error)
}

type AssetLogic interface {
	GetAssetsByUserId(userUUID uuid.UUID) ([]*domainasset.Asset, error)
	GetTransactionsByAssetId(assetId uuid.UUID) ([]*domainasset.Transaction, error)
	CreateNewAsset(userId uuid.UUID, newAssetFields map[domain.UpdatableProperty]any) (*domainasset.Asset, error)
	GetAssetById(assetId uuid.UUID) (*domainasset.Asset, error)
}

type TransactionLogic interface {
	AddNewTransaction(assetUUID uuid.UUID, d *dto.AddNewTransactionRequest) (*domainasset.Transaction, error)
}

type TransactionGroupLogic interface {
	GetTransactionGroupsByUser(userId uuid.UUID) ([]*transaction_group.TransactionGroup, error)
	CreateNewTransactionGroup(userId uuid.UUID, req dto.CreateTransactionGroupRequest) (*transaction_group.TransactionGroup, error)
}
