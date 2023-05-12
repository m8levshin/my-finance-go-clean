package transaction_group

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

type TransactionGroupDomainService interface {
	CreateNewTransactionGroup(user *user.User, parentTG *TransactionGroup, name string) (*TransactionGroup, error)
}

type transactionGroupDomainService struct {
	logger domain.Logger
}

func NewTransactionGroupService(l *domain.Logger) TransactionGroupDomainService {
	return &transactionGroupDomainService{logger: *l}
}

func (*transactionGroupDomainService) CreateNewTransactionGroup(user *user.User,
	parentTG *TransactionGroup, name string) (*TransactionGroup, error) {

	var parentTGId *domain.Id
	if parentTG != nil {
		parentTGId = &parentTG.Id
	}

	return &TransactionGroup{
		Id:       domain.NewID(),
		ParentId: parentTGId,
		UserId:   user.Id,
		Name:     name,
	}, nil
}
