package uc

import (
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
)

type keeper struct {
	userRw rw.UserRW
}

type HandlerBuilder struct {
	UserRw rw.UserRW
}

func (b HandlerBuilder) Build() Handler {
	return &keeper{userRw: b.UserRw}
}
