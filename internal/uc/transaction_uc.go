package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
)

func (k *handler) GetTransactionsByAssetId(assetId uuid.UUID) ([]*asset.Transaction, error) {
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

func (k *handler) AddNewTransaction(assetUUID uuid.UUID, req *dto.AddNewTransactionRequest) (*asset.Transaction, error) {
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
	if transactionGroup == nil {
		return nil, errors.New("transaction group is not found")
	}

	newTransaction, err := k.assetService.AddTransaction(asset, transactions, req.Volume, transactionGroup)
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
