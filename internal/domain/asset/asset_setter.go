package asset

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
)

const (
	NameField domain.UpdatableProperty = iota
	TypeField
	UserIdField
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

func SetUserId(u domain.Id) func(a *Asset) error {
	return func(a *Asset) error {
		a.UserId = u
		return nil
	}
}

func SetLimit(limit float64) func(a *Asset) error {
	return func(a *Asset) error {
		a.Limit = limit
		return nil
	}
}

func SetCurrency(c currency.Currency) func(a *Asset) error {
	return func(a *Asset) error {
		a.Currency = c
		return nil
	}
}
