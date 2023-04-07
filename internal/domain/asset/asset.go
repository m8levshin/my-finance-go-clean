package domainasset

import (
	"errors"
	. "github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"time"
)

type Currency string

type AssetType uint8

var (
	validator = assetValidator{}
)

const (
	Account AssetType = iota
	Credit
)

var (
	allowedTypes = map[AssetType]bool{
		Account: true,
		Credit:  true,
	}

	allowDebit = map[AssetType]bool{
		Account: false,
		Credit:  true,
	}
)

type Asset struct {
	Id           Id
	Type         AssetType
	Name         string
	Owner        *domainuser.User
	Currency     Currency
	Balance      float64
	Limit        float64
	Transactions []Transaction
}

type Transaction struct {
	Id        Id
	CreatedAt time.Time
	Volume    float64
}

func CreateNewAsset(opts ...*func(a *Asset) error) (*Asset, error) {
	newAsset := Asset{}
	newAsset.Id = NewID()
	for _, opt := range opts {
		err := (*opt)(&newAsset)
		if err != nil {
			return nil, err
		}
	}

	if err := validator.validateAssetForCreateAndUpdate(&newAsset); err != nil {
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

func (a *Asset) CheckTransactions() error {
	var finalBalance float64
	for _, transaction := range a.Transactions {
		finalBalance = finalBalance + transaction.Volume
	}

	if a.Balance != finalBalance {
		return errors.New("incorrect balance")
	}

	return nil
}

func (a *Asset) AddTransaction(trx Transaction) error {

	err := validator.validateBalanceAndLimitForTransaction(a, &trx)
	if err != nil {
		return err
	}

	trxId := NewID()
	trx.Id = trxId

	a.Transactions = append(a.Transactions, trx)
	a.Balance = a.Balance - trx.Volume

	return a.CheckTransactions()
}
