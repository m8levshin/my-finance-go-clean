package gorm

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
)

func mapTransactionToDomain(entity *transaction) *domainasset.Transaction {
	return &domainasset.Transaction{
		Id:                 domain.Id(entity.Id),
		CreatedAt:          entity.CreatedAt,
		AssetId:            domain.Id(entity.AssetId),
		Volume:             entity.Volume,
		TransactionGroupId: domain.Id(entity.TransactionGroupId),
	}
}

func mapTransactionToEntity(domain *domainasset.Transaction) *transaction {
	return &transaction{
		Base: Base{
			Id: uuid.UUID(domain.Id),
		},
		CreatedAt:          domain.CreatedAt,
		AssetId:            uuid.UUID(domain.AssetId),
		Volume:             domain.Volume,
		TransactionGroupId: uuid.UUID(domain.TransactionGroupId),
	}
}

func mapTransactionGroupToEntity(trxGroup *domainasset.TransactionGroup) *transactionGroup {
	return &transactionGroup{
		Base: Base{
			Id: uuid.UUID(trxGroup.Id),
		},
		ParentId: (*uuid.UUID)(trxGroup.ParentId),
		UserId:   uuid.UUID(trxGroup.UserId),
		Name:     trxGroup.Name,
	}
}

func mapTransactionGroupToDomain(trxGroup *transactionGroup) *domainasset.TransactionGroup {
	return &domainasset.TransactionGroup{
		Id:       domain.Id(trxGroup.Id),
		ParentId: (*domain.Id)(trxGroup.ParentId),
		UserId:   domain.Id(trxGroup.UserId),
		Name:     trxGroup.Name,
	}
}
