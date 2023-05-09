package asset

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"time"
)

type Transaction struct {
	Id        domain.Id
	CreatedAt time.Time
	Volume    float64
}
