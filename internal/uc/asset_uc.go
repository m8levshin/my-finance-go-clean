package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/utils"
	"time"
)

func (k *handler) GetAssetsByUserId(userUUID uuid.UUID) ([]*model.Asset, error) {
	assets, err := k.assetRw.FindByUserId(domain.Id(userUUID))
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (k *handler) CreateNewAsset(userUUID uuid.UUID, newAssetFields map[domain.UpdatableProperty]any) (*model.Asset, error) {

	userId := domain.Id(userUUID)
	user, err := k.userRw.FindById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user is not found")
	}

	name := (newAssetFields[model.AssetNameField]).(*string)
	currency := (newAssetFields[model.AssetCurrencyField]).(*string)
	limit := (newAssetFields[model.AssetLimitField]).(*float64)

	assetTypeName := (newAssetFields[model.AssetTypeField]).(*string)
	assetType := utils.ResolveAssetTypeByName(*assetTypeName)
	if assetType == nil {
		return nil, errors.New("can't recognize asset type")
	}

	newAsset, err := k.assetService.CreateAsset(
		model.SetAssetName(*name),
		model.SetAssetUserId(userId),
		model.SetAssetCurrency(model.Currency(*currency)),
		model.SetAssetType(*assetType),
		model.SetAssetLimit(*limit),
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

func (k *handler) GetAssetById(assetId uuid.UUID) (*model.Asset, error) {
	asset, err := k.assetRw.FindById(domain.Id(assetId))
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (k *handler) GetBalanceStateHistory(userId, assetId uuid.UUID, from time.Time,
	to time.Time, tz *time.Location, isAdmin bool) ([]*model.BalanceState, error) {

	asset, err := k.assetRw.FindById(domain.Id(assetId))
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, errors.New("asset is not found")
	}

	if !isAdmin && uuid.UUID(asset.UserId) != userId {
		return nil, errors.New("no permission")
	}

	transactions, err := k.transactionRw.GetTransactionsByAsset(asset.Id)
	if err != nil {
		return nil, err
	}

	assetState, err := k.assetStateService.CreateNewAssetState(asset, transactions)
	if err != nil {
		return nil, err
	}

	return assetState.CalculateBalanceInRange(from, to, tz)
}
