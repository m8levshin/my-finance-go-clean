package memory_rw

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

func memoryAssetToDomain(asset *memoryAsset) *domainasset.Asset {
	return &domainasset.Asset{
		Id:   domain.Id(asset.Id),
		Type: domainasset.Type(asset.Type),
		Name: asset.Name,
		Owner: &domainuser.User{
			Id: domain.Id(asset.Id),
		},
		Currency:     domainasset.Currency(asset.Currency),
		Balance:      asset.Balance,
		Limit:        asset.Limit,
		Transactions: nil,
	}
}

func domainAssetToMemory(asset *domainasset.Asset) *memoryAsset {
	return &memoryAsset{
		Id:       uuid.UUID(asset.Id),
		Type:     int(asset.Type),
		Name:     asset.Name,
		Owner:    uuid.UUID(asset.Owner.Id),
		Currency: string(asset.Currency),
		Balance:  asset.Balance,
		Limit:    asset.Limit,
	}

}
