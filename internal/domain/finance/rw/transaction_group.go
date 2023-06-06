package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
)

type TransactionGroupRW interface {
	GetTransactionGroupsByIds(groupIds []domain.Id) ([]*model.TransactionGroup, error)
	GetTransactionGroupById(groupId domain.Id) (*model.TransactionGroup, error)
	Save(transactionGroup *model.TransactionGroup) (*model.TransactionGroup, error)
	GetTransactionGroupsByUserId(userId domain.Id) ([]*model.TransactionGroup, error)
}
