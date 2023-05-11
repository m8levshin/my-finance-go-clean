package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/utils"
)

func (k *keeper) GetAssetsByUserId(userUUID uuid.UUID) ([]*domainasset.Asset, error) {
	assets, err := k.assetRw.FindByUserId(domain.Id(userUUID))
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (k *keeper) CreateNewAsset(userId uuid.UUID, newAssetFields map[domain.UpdatableProperty]any) (*domainasset.Asset, error) {

	userDomainId := domain.Id(userId)
	user, err := k.userRw.FindById(userDomainId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user is not found")
	}

	name := (newAssetFields[domainasset.NameField]).(*string)
	currency := (newAssetFields[domainasset.CurrencyField]).(*string)
	limit := (newAssetFields[domainasset.LimitField]).(*float64)

	assetTypeName := (newAssetFields[domainasset.TypeField]).(*string)
	assetType := utils.ResolveAssetTypeByName(*assetTypeName)
	if assetType == nil {
		return nil, errors.New("can't recognize asset type")
	}

	newAsset, err := domainasset.CreateAsset(
		domainasset.SetName(*name),
		domainasset.SetUserId(userDomainId),
		domainasset.SetCurrency(domainasset.Currency(*currency)),
		domainasset.SetType(*assetType),
		domainasset.SetLimit(*limit),
	)
	if err != nil {
		return nil, err
	}

	err = k.assetRw.Save(*newAsset)
	if err != nil {
		return nil, err
	}

	return newAsset, err
}

func (k *keeper) GetAssetById(assetId uuid.UUID) (*domainasset.Asset, error) {
	asset, err := k.assetRw.FindById(domain.Id(assetId))
	if err != nil {
		return nil, err
	}
	return asset, nil
}
