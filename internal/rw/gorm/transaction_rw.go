package gorm

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/transaction_group"
	"gorm.io/gorm"
	"time"
)

type transaction struct {
	Base
	CreatedAt          time.Time `gorm:"not null;"`
	AssetId            uuid.UUID `gorm:"not null;"`
	Asset              asset
	Volume             float64   `gorm:"not null;"`
	TransactionGroupId uuid.UUID `gorm:"not null;"`
	TransactionGroup   *transactionGroup
}

type transactionGroup struct {
	Base
	ParentId *uuid.UUID
	Parent   *transactionGroup `gorm:"foreignKey:ParentId"`
	UserId   uuid.UUID
	User     user
	Name     string
}

type transactionRW struct {
	db *gorm.DB
}

func NewTransactionRW(db *gorm.DB) (*transactionRW, error) {
	err := db.AutoMigrate(&transaction{}, &transactionGroup{})
	if err != nil {
		return nil, err
	}
	return &transactionRW{
		db: db,
	}, nil
}

func (t *transactionRW) GetTransactionGroupsByIds(groupIds []domain.Id) ([]*transaction_group.TransactionGroup, error) {
	var trxGroups []*transactionGroup

	groupUUIDs := make([]*uuid.UUID, 0, len(groupIds))
	for _, groupId := range groupIds {
		groupUUID := uuid.UUID(groupId)
		groupUUIDs = append(groupUUIDs, &groupUUID)
	}

	err := t.db.Where("user_id IN ?", groupUUIDs).Find(&trxGroups).Error
	if err != nil {
		return nil, err
	}
	return mapList(trxGroups, mapTransactionGroupToDomain), nil
}

func (t *transactionRW) GetTransactionGroupById(groupId domain.Id) (*transaction_group.TransactionGroup, error) {
	var trxGroup *transactionGroup
	err := t.db.Where("id = ?", uuid.UUID(groupId)).First(&trxGroup).Error
	if err != nil {
		return nil, err
	}
	return mapTransactionGroupToDomain(trxGroup), nil
}

func (t *transactionRW) Save(trxGroup *transaction_group.TransactionGroup) (*transaction_group.TransactionGroup, error) {
	entity := mapTransactionGroupToEntity(trxGroup)

	err := t.db.Save(entity).Error
	if err != nil {
		return nil, err
	}
	return mapTransactionGroupToDomain(entity), nil
}

func (t *transactionRW) GetTransactionGroupsByUserId(userId domain.Id) ([]*transaction_group.TransactionGroup, error) {
	var trxGroups []*transactionGroup
	err := t.db.Where("user_id = ?", uuid.UUID(userId)).Find(&trxGroups).Error
	if err != nil {
		return nil, err
	}
	return mapList(trxGroups, mapTransactionGroupToDomain), nil
}

func (t *transactionRW) GetTransactionsByAsset(assetId domain.Id) ([]*domainasset.Transaction, error) {
	a := asset{}
	if tx := t.db.Preload("Transactions").Where("id = ?", uuid.UUID(assetId)).First(&a); tx.Error != nil {
		return nil, tx.Error
	}
	return mapList(a.Transactions, mapTransactionToDomain), nil
}

func (t *transactionRW) AddTransaction(assetId domain.Id, trx domainasset.Transaction) error {
	entity := mapTransactionToEntity(&trx)
	entity.AssetId = uuid.UUID(assetId)
	err := t.db.Save(entity).Error
	if err != nil {
		return err
	}
	return nil
}
