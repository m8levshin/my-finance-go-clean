package asset

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

func AddTransaction(a *Asset, assetTransactions []*Transaction, trx Transaction) (*Asset, []*Transaction, error) {

	err := validator.validateBalanceAndLimitForTransaction(a, &trx)
	if err != nil {
		return nil, nil, err
	}

	trx.Id = domain.NewID()
	assetTransactions = append(assetTransactions, &trx)
	a.Balance = a.Balance - trx.Volume

	err = a.checkTransaction(assetTransactions)
	if err != nil {
		return nil, nil, err
	}

	return a, assetTransactions, nil
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
