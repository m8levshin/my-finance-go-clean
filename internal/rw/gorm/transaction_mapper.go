package gorm

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/transaction_group"
)

func mapTransactionToModel(entity *transaction) *domainasset.Transaction {
	return &domainasset.Transaction{
		Id:                 domain.Id(entity.Id),
		CreatedAt:          entity.CreatedAt,
		AssetId:            domain.Id(entity.AssetId),
		Volume:             entity.Volume,
		TransactionGroupId: domain.Id(entity.TransactionGroupId),
	}
}

func mapTransactionsToModels(entities []*transaction) []*domainasset.Transaction {
	result := make([]*domainasset.Transaction, 0, len(entities))
	for _, entity := range entities {
		result = append(result, mapTransactionToModel(entity))
	}
	return result
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

func mapTransactionsToEntities(domains []*domainasset.Transaction) []*transaction {
	result := make([]*transaction, 0, len(domains))
	for _, domainItem := range domains {
		result = append(result, mapTransactionToEntity(domainItem))
	}
	return result
}

func mapTransactionGroupToEntity(trxGroup *transaction_group.TransactionGroup) *transactionGroup {
	return &transactionGroup{
		Base: Base{
			Id: uuid.UUID(trxGroup.Id),
		},
		ParentId: (*uuid.UUID)(trxGroup.ParentId),
		UserId:   uuid.UUID(trxGroup.UserId),
		Name:     trxGroup.Name,
	}
}

func mapTransactionGroupsToEntities(trxGroups []*transaction_group.TransactionGroup) []*transactionGroup {
	result := make([]*transactionGroup, 0, len(trxGroups))
	for _, trxGroup := range trxGroups {
		result = append(result, mapTransactionGroupToEntity(trxGroup))
	}
	return result
}

func mapTransactionGroupToDomain(trxGroup *transactionGroup) *transaction_group.TransactionGroup {
	return &transaction_group.TransactionGroup{
		Id:       domain.Id(trxGroup.Id),
		ParentId: (*domain.Id)(trxGroup.ParentId),
		UserId:   domain.Id(trxGroup.UserId),
		Name:     trxGroup.Name,
	}
}

func mapTransactionGroupsToDomains(trxGroups []*transactionGroup) []*transaction_group.TransactionGroup {
	result := make([]*transaction_group.TransactionGroup, 0, len(trxGroups))
	for _, trxGroup := range trxGroups {
		result = append(result, mapTransactionGroupToDomain(trxGroup))
	}
	return result
}
