package uc

import (
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
)

type keeper struct {
	userRw  rw.UserRW
	assetRw rw.AssetRW
}

type HandlerBuilder struct {
	UserRw  rw.UserRW
	AssetRw rw.AssetRW
}

func (b HandlerBuilder) Build() Handler {
	return &keeper{
		userRw:  b.UserRw,
		assetRw: b.AssetRw,
	}
}
