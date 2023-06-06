package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
)

type TransactionRW interface {
	GetTransactionsByAsset(assetId domain.Id) ([]*model.Transaction, error)
	AddTransaction(assetId domain.Id, transaction model.Transaction) error
}
