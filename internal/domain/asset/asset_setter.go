package asset

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

const (
	NameField domain.UpdatableProperty = iota
	TypeField
	OwnerField
	LimitField
	CurrencyField
)

func SetName(name string) func(a *Asset) error {
	return func(a *Asset) error {
		a.Name = name
		return nil
	}
}

func SetType(t Type) func(a *Asset) error {
	return func(a *Asset) error {
		a.Type = t
		return nil
	}
}

func SetOwner(u domain.Id) func(a *Asset) error {
	return func(a *Asset) error {
		a.OwnerId = u
		return nil
	}
}

func SetLimit(limit float64) func(a *Asset) error {
	return func(a *Asset) error {
		a.Limit = limit
		return nil
	}
}

func SetCurrency(c Currency) func(a *Asset) error {
	return func(a *Asset) error {
		a.Currency = c
		return nil
	}
}
