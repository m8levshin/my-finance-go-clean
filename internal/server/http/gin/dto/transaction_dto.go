package dto

import (
	"github.com/google/uuid"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"time"
)

type TransactionDto struct {
	Id        uuid.UUID
	CreatedAt time.Time
	Volume    float64
}

func MapTransactionDomainToDto(domain *domainasset.Transaction) *TransactionDto {
	return &TransactionDto{
		Id:        uuid.UUID(domain.Id),
		CreatedAt: domain.CreatedAt,
		Volume:    domain.Volume,
	}
}
