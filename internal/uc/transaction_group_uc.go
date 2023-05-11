package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
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

func (k *keeper) CreateNewTransactionGroup(
	userId uuid.UUID,
	req dto.CreateTransactionGroupRequest,
) (*domainasset.TransactionGroup, error) {

	user, err := k.userRw.FindById(domain.Id(userId))
	if err != nil {
		return nil, err
	}

	var parentTransactionGroup *domainasset.TransactionGroup
	if req.ParentId != nil {
		parentTransactionGroup, err = k.transactionGroupRw.GetTransactionGroupById(domain.Id(*req.ParentId))
		if err != nil {
			return nil, err
		}
	}

	newTransactionGroup, err := domainasset.CreateNewTransactionGroup(user, parentTransactionGroup, req.Name)
	if err != nil {
		return nil, err
	}

	newTransactionGroup, err = k.transactionGroupRw.Save(newTransactionGroup)
	if err != nil {
		return nil, err
	}

	return newTransactionGroup, nil
}
