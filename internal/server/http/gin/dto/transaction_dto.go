package dto

import (
	"github.com/google/uuid"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"time"
)

type TransactionDto struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Volume    float64   `json:"volume"`
}

type AddNewTransactionRequest struct {
	Volume float64 `json:"volume" binding:"required"`
}

func MapTransactionDomainToDto(domain *domainasset.Transaction) *TransactionDto {
	return &TransactionDto{
		Id:        uuid.UUID(domain.Id),
		CreatedAt: domain.CreatedAt,
		Volume:    domain.Volume,
	}
}
