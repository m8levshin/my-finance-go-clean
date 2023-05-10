package asset

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"time"
)

func AddTransaction(a *Asset, assetTransactions []*Transaction, volume float64) (*Transaction, error) {

	newTransaction := Transaction{
		Id:        domain.NewID(),
		CreatedAt: time.Now(),
		AssetId:   a.Id,
		Volume:    volume,
	}

	err := validator.validateBalanceAndLimitForTransaction(a, &newTransaction)
	if err != nil {
		return nil, err
	}

	assetTransactions = append(assetTransactions, &newTransaction)
	a.Balance = a.Balance + newTransaction.Volume

	err = a.checkTransaction(assetTransactions)
	if err != nil {
		return nil, err
	}

	return &newTransaction, nil
}

func CreateNewAsset(opts ...func(a *Asset) error) (*Asset, error) {
	newAsset := Asset{}
	newAsset.Id = domain.NewID()
	for _, opt := range opts {
		err := opt(&newAsset)
		if err != nil {
			return nil, err
		}
	}

	if err := validator.validateAssetForCreateAndUpdate(&newAsset); err != nil {
		return nil, err
	}

	return &newAsset, nil
}
