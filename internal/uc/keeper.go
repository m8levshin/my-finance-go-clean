package uc

import (
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
)

type keeper struct {
	userRw             rw.UserRW
	assetRw            rw.AssetRW
	transactionRw      rw.TransactionRW
	transactionGroupRw rw.TransactionGroupRW
}

type HandlerBuilder struct {
	UserRw             rw.UserRW
	AssetRw            rw.AssetRW
	TransactionRw      rw.TransactionRW
	TransactionGroupRw rw.TransactionGroupRW
}

func (b HandlerBuilder) Build() Handler {
	return &keeper{
		userRw:             b.UserRw,
		assetRw:            b.AssetRw,
		transactionRw:      b.TransactionRw,
		transactionGroupRw: b.TransactionGroupRw,
	}
}
