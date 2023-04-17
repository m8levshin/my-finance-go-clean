package memory_rw

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
)

type assetRW struct {
	store         *sync.Map
	userRw        *rw.UserRW
	transactionRw *rw.TransactionRW
}

type memoryAsset struct {
	Id       uuid.UUID
	Type     int
	Name     string
	Owner    uuid.UUID
	Currency string
	Balance  float64
	Limit    float64
}

func NewMemoryAssetRW(userRW *rw.UserRW) rw.AssetRW {
	return assetRW{
		store:  &sync.Map{},
		userRw: userRW,
	}
}

func (a assetRW) FindByOwnerId(ownerId domain.Id) ([]*domainasset.Asset, error) {
	ownerUUID := uuid.UUID(ownerId)
	user, err := (*a.userRw).FindById(ownerId)
	if err != nil {
		return nil, err
	}

	foundAsset := make([]*memoryAsset, 0)
	a.store.Range(func(key, value any) bool {
		memoryAsset := value.(memoryAsset)
		if memoryAsset.Owner == ownerUUID {
			foundAsset = append(foundAsset, &memoryAsset)
		}
		return true
	})

	mappedDomainAssets := make([]*domainasset.Asset, 0)
	for _, asset := range foundAsset {
		domainAsset := memoryAssetToDomain(asset)
		err := a.fillTransactions(domainAsset)
		if err != nil {
			return nil, err
		}
		domainAsset.Owner = user
		mappedDomainAssets = append(mappedDomainAssets, domainAsset)
	}
	return mappedDomainAssets, nil
}

func (a assetRW) Save(asset domainasset.Asset) error {
	memoryAsset := domainAssetToMemory(&asset)
	a.store.Swap(asset.Id, memoryAsset)
	err := (*a.transactionRw).SaveTransactionsByAsset(asset.Id, asset.Transactions)
	if err != nil {
		return err
	}
	return nil
}

func (a assetRW) fillTransactions(asset *domainasset.Asset) error {
	value, err := (*a.transactionRw).GetTransactionsByAsset(asset.Id)
	if err == nil {
		return err
	}
	var resultTransactions []domainasset.Transaction
	for _, transaction := range value {
		resultTransactions = append(resultTransactions, *transaction)
	}
	asset.Transactions = resultTransactions
	return nil
}
