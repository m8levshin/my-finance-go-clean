package memory

import (
	"errors"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
	"sync"
)

type transactionGroupRw struct {
	//key: domain.Id; value: domainasset.TransactionGroup
	store *sync.Map
}

func NewMemoryTransactionGroupRW() rw.TransactionGroupRW {
	return &transactionGroupRw{
		store: &sync.Map{},
	}
}

func (t *transactionGroupRw) GetTransactionGroupsByIds(groupIds []domain.Id) ([]*domainasset.TransactionGroup, error) {
	transactionGroups := make([]*domainasset.TransactionGroup, 0)
	for _, groupId := range groupIds {
		group, _ := t.GetTransactionGroupById(groupId)
		if group != nil {
			transactionGroups = append(transactionGroups, group)
		} else {
			return nil, errors.New("can't find transaction group with id" + groupId.String())
		}
	}
	return transactionGroups, nil
}

func (t *transactionGroupRw) GetTransactionGroupById(groupId domain.Id) (*domainasset.TransactionGroup, error) {
	value, found := t.store.Load(groupId)
	if found {
		trxGroup := value.(domainasset.TransactionGroup)
		return &trxGroup, nil
	} else {
		return nil, nil
	}
}

func (t *transactionGroupRw) Save(
	transactionGroup *domainasset.TransactionGroup,
) (*domainasset.TransactionGroup, error) {
	t.store.Swap(transactionGroup.Id, *transactionGroup)
	return transactionGroup, nil
}

func (t *transactionGroupRw) GetTransactionGroupsByUserId(userId domain.Id) ([]*domainasset.TransactionGroup, error) {
	userTransactionGroups := make([]*domainasset.TransactionGroup, 0)
	t.store.Range(func(key, value any) bool {
		trxGroup := value.(domainasset.TransactionGroup)
		if trxGroup.UserId == userId {
			userTransactionGroups = append(userTransactionGroups, &trxGroup)
		}
		return true
	})
	return userTransactionGroups, nil
}
