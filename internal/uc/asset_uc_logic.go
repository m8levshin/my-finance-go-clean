package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/utils"
)

func (k *keeper) GetAssetsByUserId(userUUID uuid.UUID) ([]*domainasset.Asset, error) {
	assets, err := k.assetRw.FindByOwnerId(domain.Id(userUUID))
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (k *keeper) GetTransactionsByAssetId(assetId uuid.UUID) ([]*domainasset.Transaction, error) {
	asset, err := k.assetRw.FindById(domain.Id(assetId))
	if err != nil {
		return nil, err
	}
	transactions, err := k.transactionRw.GetTransactionsByAsset(asset.Id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (k *keeper) CreateNewAsset(
	ownerId *uuid.UUID,
	newAssetFields map[domain.UpdatableProperty]any,
) (*domainasset.Asset, error) {

	ownerDomainId := domain.Id(*ownerId)
	owner, err := k.userRw.FindById(ownerDomainId)
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, errors.New("owner is not found")
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
		domainasset.SetOwner(ownerDomainId),
		domainasset.SetCurrency(domainasset.Currency(*currency)),
		domainasset.SetType(*assetType),
		domainasset.SetLimit(*limit),
	)
	err = k.assetRw.Save(*newAsset)
	if err != nil {
		return nil, err
	}
	return newAsset, err
}
