package model

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

type Type uint8

const (
	Account Type = iota
	Credit
)

var (
	TypeNames = map[Type]string{
		Account: "Account",
		Credit:  "Credit",
	}

	AllowedTypes = map[Type]bool{
		Account: true,
		Credit:  true,
	}

	AllowDebit = map[Type]bool{
		Account: false,
		Credit:  true,
	}
)

type Asset struct {
	Id       domain.Id
	Type     Type
	Name     string
	UserId   domain.Id
	Currency Currency
	Balance  float64
	Limit    float64
}

func (s *Asset) UpdateAsset(initial *Asset, opts ...*func(u *Asset) error) (*Asset, error) {
	for _, function := range opts {
		f := *function
		err := f(initial)
		if err != nil {
			return nil, err
		}
	}

	if err := ValidateAssetForCreateAndUpdate(initial); err != nil {
		return nil, err
	}
	return initial, nil
}
