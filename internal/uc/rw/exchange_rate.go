package rw

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
	"time"
)

type ExchangeRateRW interface {
	GetExchangeRate(
		base currency.Currency,
		secondary currency.Currency,
		from time.Time,
		to time.Time,
	) ([]*currency.ExchangeRate, error)
	SaveExchangeRate(rate *currency.ExchangeRate) error
	SaveExchangeRates(rates []*currency.ExchangeRate) error
}
