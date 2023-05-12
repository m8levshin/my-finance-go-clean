package asset

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/transaction_group"
	"time"
)

type AssetDomainService interface {
	CreateAsset(opts ...func(u *Asset) error) (*Asset, error)
	UpdateAsset(initial *Asset, opts ...*func(u *Asset) error) (*Asset, error)
	AddTransaction(a *Asset, assetTransactions []*Transaction, volume float64,
		transactionGroup *transaction_group.TransactionGroup) (*Transaction, error)
}

type assetDomainService struct {
	logger domain.Logger
}

func NewAssetService(l *domain.Logger) AssetDomainService {
	return &assetDomainService{
		logger: *l,
	}
}

func (s *assetDomainService) CreateAsset(opts ...func(u *Asset) error) (*Asset, error) {
	newAsset := Asset{
		Id:      domain.NewID(),
		Balance: 0.0,
	}
	for _, f := range opts {
		err := f(&newAsset)
		if err != nil {
			return nil, err
		}
	}
	err := validateAssetForCreateAndUpdate(&newAsset)
	if err != nil {
		return nil, err
	}
	return &newAsset, nil
}

func (s *assetDomainService) UpdateAsset(initial *Asset, opts ...*func(u *Asset) error) (*Asset, error) {
	for _, function := range opts {
		f := *function
		err := f(initial)
		if err != nil {
			return nil, err
		}
	}

	if err := validateAssetForCreateAndUpdate(initial); err != nil {
		return nil, err
	}
	return initial, nil
}

func (s *assetDomainService) AddTransaction(
	a *Asset,
	assetTransactions []*Transaction,
	volume float64,
	transactionGroup *transaction_group.TransactionGroup,
) (*Transaction, error) {

	newTransaction := Transaction{
		Id:                 domain.NewID(),
		CreatedAt:          time.Now(),
		AssetId:            a.Id,
		Volume:             volume,
		TransactionGroupId: transactionGroup.Id,
	}

	err := validateBalanceAndLimitForTransaction(a, &newTransaction)
	if err != nil {
		return nil, err
	}

	assetTransactions = append(assetTransactions, &newTransaction)
	a.Balance = a.Balance + newTransaction.Volume

	err = a.CheckTransaction(assetTransactions)
	if err != nil {
		return nil, err
	}

	return &newTransaction, nil
}
