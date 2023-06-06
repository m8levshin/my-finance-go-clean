package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
	"time"
)

type ExchangeRateRW interface {
	GetExchangeRate(
		base model.Currency,
		secondary model.Currency,
		from time.Time,
		to time.Time,
	) ([]*model.ExchangeRate, error)
	SaveExchangeRate(rate *model.ExchangeRate) error
	SaveExchangeRates(rates []*model.ExchangeRate) error
}
