package uc

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
)

func (k *keeper) GetAssetsByUserId(userUUID uuid.UUID) ([]*domainasset.Asset, error) {
	assets, err := k.assetRw.FindByOwnerId(domain.Id(userUUID))
	if err != nil {
		return nil, err
	}
	return assets, nil
}
