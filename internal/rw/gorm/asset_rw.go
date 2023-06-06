package gorm

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
	asset2 "github.com/mlevshin/my-finance-go-clean/internal/domain/finance/rw"
	"gorm.io/gorm"
)

type currency struct {
	Name                       string          `gorm:"column:name;primaryKey;"`
	Assets                     []*asset        `gorm:"foreignKey:CurrencyName;references:Name"`
	BaseCurrencyExchangeRate   []*exchangeRate `gorm:"foreignKey:BaseCurrencyName;references:Name"`
	TargetCurrencyExchangeRate []*exchangeRate `gorm:"foreignKey:TargetCurrencyName;references:Name"`
}

type asset struct {
	Base
	Type         uint8  `gorm:"column:type;not null;"`
	Name         string `gorm:"column:name;not null;"`
	Currency     currency
	CurrencyName string         `gorm:"not null;"`
	Balance      float64        `gorm:"column:balance;not null;"`
	Limit        float64        `gorm:"column:limit;not null;"`
	Transactions []*transaction `gorm:"foreignKey:AssetId;references:Id"`
	UserId       uuid.UUID      `gorm:"not null;"`
	User         user           `gorm:"foreignKey:UserId"`
}

func (*currency) TableName() string {
	return "currency"
}

func (*asset) TableName() string {
	return "asset"
}

type assetRw struct {
	db *gorm.DB
}

func NewAssetRw(db *gorm.DB) (asset2.AssetRW, error) {
	err := db.AutoMigrate(&currency{}, &asset{})
	if err != nil {
		return nil, err
	}
	return &assetRw{db: db}, nil
}

func (a *assetRw) FindByUserId(userId domain.Id) ([]*model.Asset, error) {
	var assets []*asset
	tx := a.db.Where("userId = ?", uuid.UUID(userId)).Find(&assets)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return mapList(assets, mapAssetToDomain), nil
}

func (a *assetRw) FindById(assetId domain.Id) (*model.Asset, error) {
	var foundAsset asset
	tx := a.db.Where("id = ?", uuid.UUID(assetId)).First(&foundAsset)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return mapAssetToDomain(&foundAsset), nil
}

func (a *assetRw) Save(asset model.Asset) error {
	entity := mapAssetToEntity(&asset)
	err := a.db.Preload("Currency").Save(entity).Error
	if err != nil {
		return err
	}
	return nil
}
