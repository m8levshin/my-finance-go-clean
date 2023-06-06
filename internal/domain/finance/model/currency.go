package model

import "time"

type Currency string

const (
	BaseServerCurrency Currency = "USD"
)

type ExchangeRate struct {
	BaseCurrency   Currency
	TargetCurrency Currency
	Date           time.Time
	Value          float64
}
