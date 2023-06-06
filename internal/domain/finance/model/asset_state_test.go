package model

import (
	"testing"
	"time"
)

func TestCalculateBalanceInRange_periodInPast(t *testing.T) {

	testAsset := &Asset{
		Balance:  1400,
		Currency: "USD",
	}
	location := time.UTC
	transactions := []*Transaction{
		{
			CreatedAt: time.Now().UTC(),
			Volume:    400,
		},
		{
			CreatedAt: time.Date(2022, time.January, 3, 0, 0, 0, 0, location),
			Volume:    400,
		},
		{
			CreatedAt: time.Date(2022, time.January, 2, 0, 0, 0, 0, location),
			Volume:    600,
		},
	}
	assetState := AssetState{
		Asset:        testAsset,
		Transactions: transactions,
	}

	from := time.Date(2022, time.January, 1, 0, 0, 0, 0, location)
	to := time.Date(2022, time.January, 4, 0, 0, 0, 0, location)

	result, err := assetState.CalculateBalanceInRange(from, to, *location)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedResult := map[int64]float64{
		time.Date(2022, time.January, 4, 0, 0, 0, 0, location).Unix(): 1000,
		time.Date(2022, time.January, 3, 0, 0, 0, 0, location).Unix(): 1000,
		time.Date(2022, time.January, 2, 0, 0, 0, 0, location).Unix(): 600,
		time.Date(2022, time.January, 1, 0, 0, 0, 0, location).Unix(): 0,
	}

	checkCalculateBalanceInRangeResult(expectedResult, result, t)
}

func TestCalculateBalanceInRangeInTargetCurrency_periodInPast(t *testing.T) {

	testAsset := &Asset{
		Balance:  1400,
		Currency: "USD",
	}

	location := time.UTC
	var targetCurrency Currency = "RUB"
	transactions := []*Transaction{
		{
			CreatedAt: time.Now().UTC(),
			Volume:    400,
		},
		{
			CreatedAt: time.Date(2022, time.January, 3, 0, 0, 0, 0, location),
			Volume:    400,
		},
		{
			CreatedAt: time.Date(2022, time.January, 2, 0, 0, 0, 0, location),
			Volume:    600,
		},
	}
	assetState := AssetState{
		Asset:        testAsset,
		Transactions: transactions,
	}

	exchangeRates := []*ExchangeRate{
		{
			Value:          2,
			BaseCurrency:   testAsset.Currency,
			TargetCurrency: targetCurrency,
			Date:           time.Date(2022, time.January, 3, 0, 0, 0, 0, location),
		},
		{
			Value:          2,
			BaseCurrency:   testAsset.Currency,
			TargetCurrency: targetCurrency,
			Date:           time.Date(2022, time.January, 2, 0, 0, 0, 0, location),
		},
		{
			Value:          2,
			BaseCurrency:   testAsset.Currency,
			TargetCurrency: targetCurrency,
			Date:           time.Date(2022, time.January, 1, 0, 0, 0, 0, location),
		},
	}

	from := time.Date(2022, time.January, 1, 0, 0, 0, 0, location)
	to := time.Date(2022, time.January, 4, 0, 0, 0, 0, location)

	result, err := assetState.CalculateBalanceInRangeInTargetCurrency(from, to, targetCurrency, exchangeRates, *location)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedResult := map[int64]float64{
		time.Date(2022, time.January, 4, 0, 0, 0, 0, location).Unix(): 2000,
		time.Date(2022, time.January, 3, 0, 0, 0, 0, location).Unix(): 2000,
		time.Date(2022, time.January, 2, 0, 0, 0, 0, location).Unix(): 1200,
		time.Date(2022, time.January, 1, 0, 0, 0, 0, location).Unix(): 0,
	}
	checkCalculateBalanceInRangeResult(expectedResult, result, t)
}

func TestCalculateBalanceInRangeInTargetCurrency_incorrectRates(t *testing.T) {

	testAsset := &Asset{
		Balance:  1400,
		Currency: "USD",
	}

	location := time.UTC
	var targetCurrency Currency = "RUB"
	transactions := []*Transaction{
		{
			CreatedAt: time.Date(2022, time.January, 2, 0, 0, 0, 0, location),
			Volume:    600,
		},
	}
	assetState := AssetState{
		Asset:        testAsset,
		Transactions: transactions,
	}

	exchangeRates := []*ExchangeRate{
		{
			Value:          2,
			BaseCurrency:   testAsset.Currency,
			TargetCurrency: "EUR",
			Date:           time.Date(2022, time.January, 2, 0, 0, 0, 0, location),
		},
	}

	from := time.Date(2022, time.January, 1, 0, 0, 0, 0, location)
	to := time.Date(2022, time.January, 4, 0, 0, 0, 0, location)

	_, err := assetState.CalculateBalanceInRangeInTargetCurrency(from, to, targetCurrency, exchangeRates, *location)

	if err == nil || err != ThereIsNotNecessaryExchangeRate {
		t.Errorf("Didn't meet expected error: %s", ThereIsNotNecessaryExchangeRate.Error())
	}
}

func checkCalculateBalanceInRangeResult(expected map[int64]float64, actual []*BalanceState, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("Unexpected result length: got %d, want %d", len(actual), len(expected))
	}
	for _, state := range actual {
		value, exist := expected[state.Date.Unix()]
		if !exist {
			t.Errorf("Can't find date %s in expected result", state.Date.Format("02-01-2006"))
		}
		if value != state.Value {
			t.Errorf("Incorrect value %f for date %s, expected value is %f",
				state.Value, state.Date.Format("02-01-2006"), value)
		}
	}
}
