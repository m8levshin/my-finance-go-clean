package uc

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/config"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/rw"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/service"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
)

type Handler interface {
	UserLogic
	AssetLogic
	TransactionLogic
	TransactionGroupLogic
}

type handler struct {
	config                  config.Configuration
	userRw                  domainuser.UserRW
	assetRw                 rw.AssetRW
	transactionRw           rw.TransactionRW
	transactionGroupRw      rw.TransactionGroupRW
	userService             domainuser.UserDomainService
	assetService            service.AssetDomainService
	assetStateService       service.AssetStateDomainService
	transactionGroupService service.TransactionGroupDomainService
}

func NewHandler(
	config config.Configuration,
	userRw domainuser.UserRW, assetRw rw.AssetRW,
	transactionRw rw.TransactionRW,
	transactionGroupRw rw.TransactionGroupRW,
	userService domainuser.UserDomainService,
	assetService service.AssetDomainService,
	assetStateService service.AssetStateDomainService,
	transactionGroupService service.TransactionGroupDomainService,
) *handler {
	return &handler{config: config, userRw: userRw, assetRw: assetRw, transactionRw: transactionRw,
		transactionGroupRw: transactionGroupRw, userService: userService, assetService: assetService,
		transactionGroupService: transactionGroupService, assetStateService: assetStateService}
}

type UserLogic interface {
	GetAllUsers() (users []*domainuser.User, err error)
	GetUserById(uuid uuid.UUID) (user *domainuser.User, err error)

	GetUserByEmail(email string) (user *domainuser.User, err error)
	CreateNewUser(newUserFields map[domain.UpdatableProperty]any) (user *domainuser.User, err error)
}

type AssetLogic interface {
	GetAssetsByUserId(userUUID uuid.UUID) ([]*model.Asset, error)
	GetTransactionsByAssetId(assetId uuid.UUID) ([]*model.Transaction, error)
	CreateNewAsset(userId uuid.UUID, newAssetFields map[domain.UpdatableProperty]any) (*model.Asset, error)
	GetAssetById(assetId uuid.UUID) (*model.Asset, error)
}

type TransactionLogic interface {
	AddNewTransaction(assetUUID uuid.UUID, d *dto.AddNewTransactionRequest, userUUID uuid.UUID, isAdmin bool) (*model.Transaction, error)
}

type TransactionGroupLogic interface {
	GetTransactionGroupsByUser(userId uuid.UUID) ([]*model.TransactionGroup, error)
	CreateNewTransactionGroup(userId uuid.UUID, req dto.CreateTransactionGroupRequest) (*model.TransactionGroup, error)
}
