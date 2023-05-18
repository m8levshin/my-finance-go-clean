package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
)

type TransactionRW interface {
	GetTransactionsByAsset(assetId domain.Id) ([]*domainasset.Transaction, error)
	AddTransaction(assetId domain.Id, transaction domainasset.Transaction) error
}
