package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
)

func (k *keeper) GetTransactionsByAssetId(assetId uuid.UUID) ([]*domainasset.Transaction, error) {
	asset, err := k.assetRw.FindById(domain.Id(assetId))
	if err != nil {
		return nil, err
	}
	transactions, err := k.transactionRw.GetTransactionsByAsset(asset.Id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (k *keeper) AddNewTransaction(assetUUID uuid.UUID, req *dto.AddNewTransactionRequest) (*domainasset.Transaction, error) {
	asset, err := k.assetRw.FindById(domain.Id(assetUUID))
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, errors.New("asset is not found")
	}

	transactions, err := k.transactionRw.GetTransactionsByAsset(asset.Id)
	if err != nil {
		return nil, err
	}

	transactionGroup, err := k.transactionGroupRw.GetTransactionGroupById(domain.Id(req.TransactionGroupId))
	if err != nil {
		return nil, err
	}
	if transactionGroup != nil {
		return nil, errors.New("transaction group is not found")
	}

	newTransaction, err := domainasset.AddTransaction(asset, transactions, req.Volume, transactionGroup)
	if err != nil {
		return nil, err
	}

	err = k.assetRw.Save(*asset)
	if err != nil {
		return nil, err
	}

	err = k.transactionRw.AddTransaction(asset.Id, *newTransaction)
	if err != nil {
		return nil, err
	}

	return newTransaction, err
}
