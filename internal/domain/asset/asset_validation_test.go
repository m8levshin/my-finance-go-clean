package domainasset

import (
	"errors"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var (
	TestAssetName          = "TestName"
	TestCurrency  Currency = "USD"
	TestOwner              = domainuser.User{Id: domain.NewID()}
)

func TestCreateNewAsset(t *testing.T) {

	t.Run("ok behavior", func(t *testing.T) {
		mockAssetValidator := NewMockAssetValidator(t)
		validator = mockAssetValidator
		mockAssetValidator.
			On("validateAssetForCreateAndUpdate", mock.Anything).
			Return(nil)

		asset, err := CreateNewAsset(
			SetName(TestAssetName),
			SetCurrency(TestCurrency),
			SetOwner(&TestOwner),
		)
		assert.Nil(t, err)
		assert.Equal(t, TestAssetName, asset.Name)
		assert.Equal(t, TestCurrency, asset.Currency)
		assert.Equal(t, TestOwner.Id, asset.Owner.Id)
	})

	t.Run("validation error", func(t *testing.T) {
		mockAssetValidator := NewMockAssetValidator(t)
		validator = mockAssetValidator

		mockAssetValidator.
			On("validateAssetForCreateAndUpdate", mock.Anything).
			Return(errors.New("validation error"))

		asset, err := CreateNewAsset(
			SetName(TestAssetName),
			SetCurrency(TestCurrency),
			SetOwner(&TestOwner),
		)
		assert.Error(t, err, "validation error")
		assert.Nil(t, asset)
	})
}

func TestCheckBalanceAndLimitForTransaction(t *testing.T) {
	t.Run("debt is not allowed", func(t *testing.T) {
		asset := Asset{Type: Account, Balance: 1000}
		trx := Transaction{Volume: 1200, CreatedAt: time.Now()}

		err := validator.validateBalanceAndLimitForTransaction(&asset, &trx)

		assert.Error(t, err, DebitNotAllowed)
	})

	t.Run("limit is reached", func(t *testing.T) {
		asset := Asset{Type: Credit, Balance: 0, Limit: 1000}
		trx := Transaction{Volume: 1200, CreatedAt: time.Now()}

		err := validator.validateBalanceAndLimitForTransaction(&asset, &trx)

		assert.Error(t, err, ReachedLimits)
	})
}
