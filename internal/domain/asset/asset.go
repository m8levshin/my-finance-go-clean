package asset

import (
	"errors"
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

func (a *Asset) CheckTransaction(transactions []*Transaction) error {
	var finalBalance float64
	for _, t := range transactions {
		finalBalance = finalBalance + t.Volume
	}

	if a.Balance != finalBalance {
		return errors.New("incorrect balance")
	}
	return nil
}
