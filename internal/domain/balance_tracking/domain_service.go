package balance_tracking

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
	"time"
)

type BalanceTrackingDomainService interface {
	CalculateBalanceInRangeInTargetCurrency(
		asset *asset.Asset,
		transactions []*asset.Transaction,
		from time.Time,
		to time.Time,
		targetCurrency currency.Currency,
		rates []*currency.ExchangeRate,
	) ([]*BalanceState, error)

	CalculateBalanceInRange(
		asset *asset.Asset,
		transactions []*asset.Transaction,
		from time.Time,
		to time.Time,
	) ([]*BalanceState, error)
}

type balanceTrackingDomainService struct {
	logger domain.Logger
}

func NewCreateNewTransactionGroup(l *domain.Logger) BalanceTrackingDomainService {
	return &balanceTrackingDomainService{logger: *l}
}

func (b *balanceTrackingDomainService) CalculateBalanceInRangeInTargetCurrency(asset *asset.Asset, transactions []*asset.Transaction, from time.Time, to time.Time, targetCurrency currency.Currency, rates []*currency.ExchangeRate) ([]*BalanceState, error) {
	//TODO implement me
	panic("implement me")
}

func (b *balanceTrackingDomainService) CalculateBalanceInRange(asset *asset.Asset, transactions []*asset.Transaction, from time.Time, to time.Time) ([]*BalanceState, error) {
	//TODO implement me
	panic("implement me")
}
