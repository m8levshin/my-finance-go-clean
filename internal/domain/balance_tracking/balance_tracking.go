package balance_tracking

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
	"time"
)

type BalanceState struct {
	Date     time.Time
	Currency currency.Currency
	Value    float64
}
