package dto

import (
	"github.com/google/uuid"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"time"
)

type TransactionDto struct {
	Id                 uuid.UUID `json:"id"`
	CreatedAt          time.Time `json:"createdAt"`
	Volume             float64   `json:"volume"`
	TransactionGroupId uuid.UUID `json:"transactionGroupId"`
}

type AddNewTransactionRequest struct {
	Volume             float64   `json:"volume" binding:"required"`
	TransactionGroupId uuid.UUID `json:"transactionGroupId" binding:"required"`
}

func MapTransactionDomainToDto(domainEntity *domainasset.Transaction) *TransactionDto {
	return &TransactionDto{
		Id:                 uuid.UUID(domainEntity.Id),
		CreatedAt:          domainEntity.CreatedAt,
		Volume:             domainEntity.Volume,
		TransactionGroupId: uuid.UUID(domainEntity.TransactionGroup.Id),
	}
}
