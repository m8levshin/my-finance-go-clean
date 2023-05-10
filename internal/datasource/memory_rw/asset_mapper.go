package memory_rw

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
)

func memoryAssetToDomain(asset *memoryAsset) *domainasset.Asset {
	return &domainasset.Asset{
		Id:       domain.Id(asset.Id),
		Type:     domainasset.Type(asset.Type),
		Name:     asset.Name,
		UserId:   domain.Id(asset.UserId),
		Currency: domainasset.Currency(asset.Currency),
		Balance:  asset.Balance,
		Limit:    asset.Limit,
	}
}

func domainAssetToMemory(asset *domainasset.Asset) memoryAsset {
	return memoryAsset{
		Id:       uuid.UUID(asset.Id),
		Type:     int(asset.Type),
		Name:     asset.Name,
		UserId:   uuid.UUID(asset.UserId),
		Currency: string(asset.Currency),
		Balance:  asset.Balance,
		Limit:    asset.Limit,
	}

}
