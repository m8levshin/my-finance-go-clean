package model

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

const (
	AssetNameField domain.UpdatableProperty = iota
	AssetTypeField
	AssetUserIdField
	AssetLimitField
	AssetCurrencyField
)

func SetAssetName(name string) func(a *Asset) error {
	return func(a *Asset) error {
		a.Name = name
		return nil
	}
}

func SetAssetType(t Type) func(a *Asset) error {
	return func(a *Asset) error {
		a.Type = t
		return nil
	}
}

func SetAssetUserId(u domain.Id) func(a *Asset) error {
	return func(a *Asset) error {
		a.UserId = u
		return nil
	}
}

func SetAssetLimit(limit float64) func(a *Asset) error {
	return func(a *Asset) error {
		a.Limit = limit
		return nil
	}
}

func SetAssetCurrency(c Currency) func(a *Asset) error {
	return func(a *Asset) error {
		a.Currency = c
		return nil
	}
}
