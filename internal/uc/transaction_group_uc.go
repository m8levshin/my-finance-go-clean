package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
)

func (k *keeper) GetTransactionGroupsByUser(userId uuid.UUID) ([]*domainasset.TransactionGroup, error) {
	user, err := k.userRw.FindById(domain.Id(userId))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user is not found")
	}

	groups, err := k.transactionGroupRw.GetTransactionGroupsByUserId(user.Id)
	return groups, err
}
