package gorm

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
	currency2 "github.com/mlevshin/my-finance-go-clean/internal/domain/finance/rw"
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

func NewExchangeRateRW(db *gorm.DB) (currency2.ExchangeRateRW, error) {
	err := db.AutoMigrate(&exchangeRate{})
	if err != nil {
		return nil, err
	}
	return &exchangeRateRW{
		db: db,
	}, nil
}

func (e *exchangeRateRW) GetExchangeRate(base model.Currency, secondary model.Currency,
	from time.Time, to time.Time) ([]*model.ExchangeRate, error) {

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

func (e *exchangeRateRW) SaveExchangeRate(rate *model.ExchangeRate) error {
	entity := mapExchangeRateToEntity(rate)
	err := e.db.Save(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *exchangeRateRW) SaveExchangeRates(rates []*model.ExchangeRate) error {
	entities := mapList(rates, mapExchangeRateToEntity)
	err := e.db.Save(&entities).Error
	if err != nil {
		return err
	}
	return nil
}
