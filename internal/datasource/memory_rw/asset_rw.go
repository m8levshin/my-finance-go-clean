package memory_rw

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
)

type assetRW struct {
	//key: domain.Id, value: memoryAsset
	store  *sync.Map
	userRw *rw.UserRW
}

type memoryAsset struct {
	Id       uuid.UUID
	Type     int
	Name     string
	OwnerId  uuid.UUID
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

func (a assetRW) FindById(assetId domain.Id) (*domainasset.Asset, error) {
	value, _ := a.store.Load(assetId)
	memoryAsset := value.(*memoryAsset)
	domainAsset := memoryAssetToDomain(memoryAsset)
	return domainAsset, nil
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
		if memoryAsset.OwnerId == ownerUUID {
			foundAsset = append(foundAsset, &memoryAsset)
		}
		return true
	})

	mappedDomainAssets := make([]*domainasset.Asset, 0)
	for _, asset := range foundAsset {
		domainAsset := memoryAssetToDomain(asset)
		domainAsset.OwnerId = user.Id
		mappedDomainAssets = append(mappedDomainAssets, domainAsset)
	}
	return mappedDomainAssets, nil
}

func (a assetRW) Save(asset domainasset.Asset) error {
	memoryAsset := domainAssetToMemory(&asset)
	a.store.Swap(asset.Id, memoryAsset)
	return nil
}
