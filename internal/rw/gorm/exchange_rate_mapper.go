package gorm

import (
	domaincurrency "github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
)

func mapExchangeRateToDomain(rate *exchangeRate) *domaincurrency.ExchangeRate {
	return &domaincurrency.ExchangeRate{
		BaseCurrency:   domaincurrency.Currency(rate.BaseCurrencyName),
		TargetCurrency: domaincurrency.Currency(rate.TargetCurrencyName),
		Date:           rate.Date,
		Value:          rate.Value,
	}
}

func mapExchangeRateToEntity(rate *domaincurrency.ExchangeRate) *exchangeRate {
	return &exchangeRate{
		BaseCurrencyName:   string(rate.BaseCurrency),
		TargetCurrencyName: string(rate.TargetCurrency),
		Date:               rate.Date,
		Value:              rate.Value,
	}
}
