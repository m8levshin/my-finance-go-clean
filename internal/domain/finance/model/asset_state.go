package model

import (
	"errors"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"time"
)

const (
	dayDuration = time.Hour * 24
)

var (
	ThereIsNotNecessaryExchangeRate = domain.NewError("there isn't exchange rate for calculating")
)

type AssetState struct {
	Asset        *Asset
	Transactions []*Transaction
}

type BalanceState struct {
	Date         time.Time
	Currency     Currency
	BaseCurrency Currency
	ExchangeRate float64
	Value        float64
}

func (s *AssetState) AddTransaction(volume float64, transactionGroup *TransactionGroup) (*Transaction, error) {

	newTransaction := Transaction{
		Id:                 domain.NewID(),
		CreatedAt:          time.Now(),
		AssetId:            s.Asset.Id,
		Volume:             volume,
		TransactionGroupId: transactionGroup.Id,
	}

	err := ValidateBalanceAndLimitForTransaction(s.Asset, &newTransaction)
	if err != nil {
		return nil, err
	}

	s.Transactions = append(s.Transactions, &newTransaction)
	s.Asset.Balance = s.Asset.Balance + newTransaction.Volume

	err = s.CheckTransaction()
	if err != nil {
		return nil, err
	}

	return &newTransaction, nil
}

func (s *AssetState) CheckTransaction() error {
	var finalBalance float64
	for _, t := range s.Transactions {
		finalBalance = finalBalance + t.Volume
	}

	if s.Asset.Balance != finalBalance {
		return errors.New("incorrect balance")
	}
	return nil
}

func (s *AssetState) CalculateBalanceInRange(from time.Time, to time.Time, tz *time.Location) ([]*BalanceState, error) {

	resultBalanceState := []*BalanceState{}
	transactionsByDays := groupTransactionsByDays(s.Transactions, tz)

	today := resetTime(time.Now(), tz)
	startDate := resetTime(from, tz)
	toDate := resetTime(to, tz)

	balance := s.Asset.Balance
	iterationDate := today
	for ; isDateAfterOrEqual(iterationDate, startDate); iterationDate = iterationDate.Add(-dayDuration) {

		if iterationDate != today {
			nextDay := iterationDate.Add(dayDuration)
			nextDayTransactions, exist := transactionsByDays[nextDay.Unix()]
			if exist {
				for _, t := range *nextDayTransactions {
					balance = balance - t.Volume
				}
			}
		}

		if isDateInRange(iterationDate, startDate, toDate) {
			resultBalanceState = append(resultBalanceState, &BalanceState{
				Date:     iterationDate,
				Currency: s.Asset.Currency,
				Value:    balance,
			})
		}
	}

	return resultBalanceState, nil
}

func (s *AssetState) CalculateBalanceInRangeInTargetCurrency(from time.Time, to time.Time, targetCurrency Currency,
	rates []*ExchangeRate, tz *time.Location) ([]*BalanceState, error) {

	balanceTrackingInAssetCurrency, err := s.CalculateBalanceInRange(from, to, tz)
	if err != nil {
		return nil, err
	}

	for _, balanceState := range balanceTrackingInAssetCurrency {
		exchangeRate := findNearestRate(balanceState.Date, rates, tz, s.Asset.Currency, targetCurrency)
		if exchangeRate == nil {
			return nil, ThereIsNotNecessaryExchangeRate
		}

		balanceState.ExchangeRate = exchangeRate.Value
		balanceState.Value = exchangeRate.Value * balanceState.Value
		balanceState.BaseCurrency = balanceState.Currency
		balanceState.Currency = targetCurrency
	}

	return balanceTrackingInAssetCurrency, nil
}

func isDateInRange(date, startDate, endDate time.Time) bool {
	return date.After(startDate) && date.Before(endDate) || date.Equal(startDate) || date.Equal(endDate)
}

func isDateAfterOrEqual(date, reference time.Time) bool {
	return date.After(reference) || date.Equal(reference)
}

func resetTime(timeValue time.Time, tz *time.Location) time.Time {
	t := timeValue.In(tz)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tz)
}

func groupTransactionsByDays(transactions []*Transaction, tz *time.Location) map[int64]*[]*Transaction {
	resultMap := map[int64]*[]*Transaction{}
	for _, transaction := range transactions {
		transactionDate := resetTime(transaction.CreatedAt, tz)
		transactionsForDay, exist := resultMap[transactionDate.Unix()]
		if !exist {
			resultMap[transactionDate.Unix()] = &[]*Transaction{transaction}
		} else {
			*transactionsForDay = append(*resultMap[transactionDate.Unix()], transaction)
		}
	}
	return resultMap
}

func findNearestRate(t time.Time, r []*ExchangeRate, tz *time.Location,
	baseCurrency Currency, targetCurrency Currency) *ExchangeRate {

	targetDate := resetTime(t, tz)
	// find nearest date
	var resultRate *ExchangeRate
	for _, rate := range r {

		rateDate := rate.Date
		resetRateDate := resetTime(rateDate, tz)

		if !resetRateDate.After(targetDate) &&
			rate.TargetCurrency == targetCurrency &&
			rate.BaseCurrency == baseCurrency {

			if resultRate == nil || resetRateDate.Before(resetRateDate) {
				resultRate = rate
			}

			if resetRateDate == targetDate {
				resultRate = rate
				break
			}
		}
	}

	return resultRate
}
