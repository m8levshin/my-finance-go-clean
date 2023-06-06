package dto

import (
	"github.com/google/uuid"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
)

type TransactionGroupDto struct {
	Id       uuid.UUID  `json:"id" binding:"required"`
	Name     string     `json:"name" binding:"required"`
	ParentId *uuid.UUID `json:"parentId"`
}

type CreateTransactionGroupRequest struct {
	Name     string     `json:"name" binding:"required"`
	ParentId *uuid.UUID `json:"parentId"`
}

func MapTransactionGroupDomainToDto(domainEntity *domainasset.TransactionGroup) *TransactionGroupDto {
	var parentId *uuid.UUID
	if domainEntity.ParentId != nil {
		value := uuid.UUID(*(domainEntity.ParentId))
		parentId = &value
	}

	return &TransactionGroupDto{
		Id:       uuid.UUID(domainEntity.Id),
		Name:     domainEntity.Name,
		ParentId: parentId,
	}
}
