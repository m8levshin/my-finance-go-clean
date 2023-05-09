package asset

import (
	"errors"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

type Type uint8

var (
	validator Validator = &assetValidator{}
)

const (
	Account Type = iota
	Credit
)

var (
	TypeNames = map[Type]string{
		Account: "Account",
		Credit:  "Credit",
	}

	allowedTypes = map[Type]bool{
		Account: true,
		Credit:  true,
	}

	allowDebit = map[Type]bool{
		Account: false,
		Credit:  true,
	}
)

type Asset struct {
	Id       domain.Id
	Type     Type
	Name     string
	OwnerId  domain.Id
	Currency Currency
	Balance  float64
	Limit    float64
}

func CreateAsset(opts ...func(u *Asset) error) (*Asset, error) {
	newAsset := Asset{
		Id:      domain.NewID(),
		Balance: 0.0,
	}
	for _, f := range opts {
		err := f(&newAsset)
		if err != nil {
			return nil, err
		}
	}
	err := validator.validateAssetForCreateAndUpdate(&newAsset)
	if err != nil {
		return nil, err
	}
	return &newAsset, nil
}

func UpdateAsset(initial *Asset, opts ...*func(u *Asset) error) (*Asset, error) {
	for _, function := range opts {
		f := *function
		err := f(initial)
		if err != nil {
			return nil, err
		}
	}

	if err := validator.validateAssetForCreateAndUpdate(initial); err != nil {
		return nil, err
	}
	return initial, nil
}

func (a *Asset) checkTransaction(transactions []*Transaction) error {
	var finalBalance float64
	for _, transaction := range transactions {
		finalBalance = finalBalance + transaction.Volume
	}

	if a.Balance != finalBalance {
		return errors.New("incorrect balance")
	}
	return nil
}
