package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
)

type AssetRW interface {
	FindByUserId(userId domain.Id) ([]*domainasset.Asset, error)
	FindById(assetId domain.Id) (*domainasset.Asset, error)
	Save(asset domainasset.Asset) error
}
