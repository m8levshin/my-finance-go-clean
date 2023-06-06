package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
)

func (k *handler) GetTransactionsByAssetId(assetId uuid.UUID) ([]*model.Transaction, error) {
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

func (k *handler) AddNewTransaction(assetUUID uuid.UUID,
	req *dto.AddNewTransactionRequest, userUUID uuid.UUID, isAdmin bool) (*model.Transaction, error) {

	asset, err := k.assetRw.FindById(domain.Id(assetUUID))
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, errors.New("asset is not found")
	}

	if !isAdmin && uuid.UUID(asset.UserId) != userUUID {
		return nil, errors.New("no permission for adding of a new transaction")
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

	assetState, err := k.assetStateService.CreateNewAssetState(asset, transactions)
	if err != nil {
		return nil, err
	}

	newTransaction, err := assetState.AddTransaction(req.Volume, transactionGroup)
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
