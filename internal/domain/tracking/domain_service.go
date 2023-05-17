package tracking

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
	"time"
)

const (
	dayDuration = time.Hour * 24
)

var (
	ThereIsNotNecessaryExchangeRate = domain.NewError("there isn't exchange rate for calculating")
)

type BalanceTrackingDomainService interface {
	CalculateBalanceInRangeInTargetCurrency(asset *asset.Asset,
		transactions []*asset.Transaction, from time.Time, to time.Time, targetCurrency currency.Currency,
		rates []*currency.ExchangeRate, tz time.Location) ([]*BalanceState, error)

	CalculateBalanceInRange(
		asset *asset.Asset,
		transactions []*asset.Transaction,
		from time.Time,
		to time.Time,
		tz time.Location,
	) ([]*BalanceState, error)
}

type balanceTrackingDomainService struct {
}

func NewBalanceTrackingDomainService() BalanceTrackingDomainService {
	return &balanceTrackingDomainService{}
}

func (b *balanceTrackingDomainService) CalculateBalanceInRangeInTargetCurrency(asset *asset.Asset,
	transactions []*asset.Transaction, from time.Time, to time.Time, targetCurrency currency.Currency,
	rates []*currency.ExchangeRate, tz time.Location) ([]*BalanceState, error) {

	balanceTrackingInAssetCurrency, err := b.CalculateBalanceInRange(asset, transactions, from, to, tz)
	if err != nil {
		return nil, err
	}

	for _, balanceState := range balanceTrackingInAssetCurrency {
		exchangeRate := findNearestRate(balanceState.Date, rates, tz, asset.Currency, targetCurrency)
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

// CalculateBalanceInRange calculates the balance state of an asset within the specified date range.
//
// This method takes an asset, a list of transactions associated with the asset, and the start and end dates of the range.
// It iterates over the dates in reverse order, starting from the current date and moving backward to the start date.
// For each date in the range, it adjusts the asset's balance considering the transactions that occurred on the next day.
// Then it creates a BalanceState object with the date, asset currency, and current balance value, and adds it to the result list.
//
// Parameters:
// - asset: The asset object for which to calculate the balance state.
// - transactions: The list of transactions associated with the asset.
// - from: The start date of the range.
// - to: The end date of the range.
//
// Returns:
// - []*BalanceState: The list of BalanceState objects representing the balance state for each date in the range.
// - error: An error if any issue occurs during the execution of the method.
func (b *balanceTrackingDomainService) CalculateBalanceInRange(asset *asset.Asset,
	transactions []*asset.Transaction, from time.Time, to time.Time, tz time.Location) ([]*BalanceState, error) {

	resultBalanceState := make([]*BalanceState, 0, 10)
	transactionsByDays := groupTransactionsByDays(transactions, tz)

	today := resetTime(time.Now(), tz)
	startDate := resetTime(from, tz)
	toDate := resetTime(to, tz)

	balance := asset.Balance
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
				Currency: asset.Currency,
				Value:    balance,
			})
		}
	}

	return resultBalanceState, nil
}

func isDateInRange(date, startDate, endDate time.Time) bool {
	return date.After(startDate) && date.Before(endDate) || date.Equal(startDate) || date.Equal(endDate)
}

func isDateAfterOrEqual(date, reference time.Time) bool {
	return date.After(reference) || date.Equal(reference)
}

func resetTime(timeValue time.Time, tz time.Location) time.Time {
	t := timeValue.In(&tz)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, &tz)
}

func groupTransactionsByDays(transactions []*asset.Transaction, tz time.Location) map[int64]*[]*asset.Transaction {
	resultMap := map[int64]*[]*asset.Transaction{}
	for _, transaction := range transactions {
		transactionDate := resetTime(transaction.CreatedAt, tz)
		transactionsForDay, exist := resultMap[transactionDate.Unix()]
		if !exist {
			resultMap[transactionDate.Unix()] = &[]*asset.Transaction{transaction}
		} else {
			*transactionsForDay = append(*resultMap[transactionDate.Unix()], transaction)
		}
	}
	return resultMap
}

func findNearestRate(t time.Time, r []*currency.ExchangeRate, tz time.Location,
	baseCurrency currency.Currency, targetCurrency currency.Currency) *currency.ExchangeRate {

	targetDate := resetTime(t, tz)
	// find nearest date
	var resultRate *currency.ExchangeRate
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
