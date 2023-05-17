package tracking

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
	"time"
)

type BalanceState struct {
	Date         time.Time
	Currency     currency.Currency
	BaseCurrency currency.Currency
	ExchangeRate float64
	Value        float64
}
