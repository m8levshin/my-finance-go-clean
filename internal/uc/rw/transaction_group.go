package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/transaction_group"
)

type TransactionGroupRW interface {
	GetTransactionGroupsByIds(groupIds []domain.Id) ([]*transaction_group.TransactionGroup, error)
	GetTransactionGroupById(groupId domain.Id) (*transaction_group.TransactionGroup, error)
	Save(transactionGroup *transaction_group.TransactionGroup) (*transaction_group.TransactionGroup, error)
	GetTransactionGroupsByUserId(userId domain.Id) ([]*transaction_group.TransactionGroup, error)
}
