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

	for _, elem := range transactions {
		transactions = append(transactions, elem)
	}
	return transactions, nil
}

func (k *keeper) CreateNewAsset(newAssetFields map[domain.UpdatableProperty]any) (*domainasset.Asset, error) {

	name := (newAssetFields[domainasset.NameField]).(string)

	ownerId := (newAssetFields[domainasset.OwnerField]).(uuid.UUID)
	owner, err := k.userRw.FindById(domain.Id(ownerId))
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, errors.New("owner is not found")
	}

	currency := (newAssetFields[domainasset.CurrencyField]).(string)
	limit := (newAssetFields[domainasset.LimitField]).(float64)

	assetTypeName := (newAssetFields[domainasset.TypeField]).(*string)
	assetType := utils.ResolveAssetTypeByName(*assetTypeName)
	if assetType == nil {
		return nil, errors.New("can't recognize asset type")
	}

	newAsset, err := domainasset.CreateAsset(
		domainasset.SetName(name),
		domainasset.SetOwner(domain.Id(ownerId)),
		domainasset.SetCurrency(domainasset.Currency(currency)),
		domainasset.SetType(assetType),
		domainasset.SetLimit(limit),
	)
	return newAsset, err
}
