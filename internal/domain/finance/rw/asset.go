package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
)

type AssetRW interface {
	FindByUserId(userId domain.Id) ([]*model.Asset, error)
	FindById(assetId domain.Id) (*model.Asset, error)
	Save(asset model.Asset) error
}
