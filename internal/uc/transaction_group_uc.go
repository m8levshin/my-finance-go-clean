package uc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/transaction_group"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
)

func (k *handler) GetTransactionGroupsByUser(userId uuid.UUID) ([]*transaction_group.TransactionGroup, error) {
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

func (k *handler) CreateNewTransactionGroup(
	userId uuid.UUID,
	req dto.CreateTransactionGroupRequest,
) (*transaction_group.TransactionGroup, error) {

	user, err := k.userRw.FindById(domain.Id(userId))
	if err != nil {
		return nil, err
	}

	var parentTransactionGroup *transaction_group.TransactionGroup
	if req.ParentId != nil {
		parentTransactionGroup, err = k.transactionGroupRw.GetTransactionGroupById(domain.Id(*req.ParentId))
		if err != nil {
			return nil, err
		}
	}

	newTransactionGroup, err := k.transactionGroupService.CreateNewTransactionGroup(user,
		parentTransactionGroup, req.Name)
	if err != nil {
		return nil, err
	}

	newTransactionGroup, err = k.transactionGroupRw.Save(newTransactionGroup)
	if err != nil {
		return nil, err
	}

	return newTransactionGroup, nil
}
