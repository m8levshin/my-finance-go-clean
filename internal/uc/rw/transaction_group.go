package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
)

type TransactionGroupRW interface {
	GetTransactionGroupsByIds(groupIds []*domain.Id) ([]*asset.TransactionGroup, error)
	GetTransactionGroupById(groupId *domain.Id) (*asset.TransactionGroup, error)
	Save(transactionGroup *asset.TransactionGroup) (*asset.TransactionGroup, error)
	GetTransactionGroupsByUserId(userId *domain.Id) ([]*asset.TransactionGroup, error)
}
