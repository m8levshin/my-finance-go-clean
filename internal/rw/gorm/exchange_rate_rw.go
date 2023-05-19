package gorm

import (
	domaincurrency "github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
	"gorm.io/gorm"
	"time"
)

type exchangeRate struct {
	BaseCurrencyName   string `gorm:"primaryKey;not null;"`
	BaseCurrency       currency
	TargetCurrencyName string `gorm:"primaryKey;not null;"`
	TargetCurrency     currency
	Date               time.Time `gorm:"primaryKey;not null;"`
	Value              float64   `gorm:"not null;"`
}

type exchangeRateRW struct {
	db *gorm.DB
}

func NewExchangeRateRW(db *gorm.DB) (rw.ExchangeRateRW, error) {
	err := db.AutoMigrate(&exchangeRate{})
	if err != nil {
		return nil, err
	}
	return &exchangeRateRW{
		db: db,
	}, nil
}

func (e *exchangeRateRW) GetExchangeRate(base domaincurrency.Currency, secondary domaincurrency.Currency,
	from time.Time, to time.Time) ([]*domaincurrency.ExchangeRate, error) {

	var rates []*exchangeRate
	err := e.db.Where(
		"base_currency_name = ? AND target_currency_name = ? AND "+
			"date <= ? AND date >= ?", string(base), string(secondary), to, from,
	).Find(&rates).Error
	if err != nil {
		return nil, err
	}
	return mapList(rates, mapExchangeRateToDomain), nil
}

func (e *exchangeRateRW) SaveExchangeRate(rate *domaincurrency.ExchangeRate) error {
	entity := mapExchangeRateToEntity(rate)
	err := e.db.Save(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *exchangeRateRW) SaveExchangeRates(rates []*domaincurrency.ExchangeRate) error {
	entities := mapList(rates, mapExchangeRateToEntity)
	err := e.db.Save(&entities).Error
	if err != nil {
		return err
	}
	return nil
}
