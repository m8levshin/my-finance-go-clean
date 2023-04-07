package domainasset

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
