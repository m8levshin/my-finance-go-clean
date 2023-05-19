package gorm

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	domaincurrency "github.com/mlevshin/my-finance-go-clean/internal/domain/currency"
)

func mapAssetToDomain(e *asset) *domainasset.Asset {
	return &domainasset.Asset{
		Id:       domain.Id(e.Id),
		Type:     domainasset.Type(e.Type),
		Name:     e.Name,
		UserId:   domain.Id(e.UserId),
		Currency: domaincurrency.Currency(e.CurrencyName),
		Balance:  e.Balance,
		Limit:    e.Limit,
	}
}

func mapAssetToEntity(d *domainasset.Asset) *asset {
	return &asset{
		Base: Base{
			Id: uuid.UUID(d.Id),
		},
		Type: uint8(d.Type),
		Name: d.Name,
		Currency: currency{
			Name: string(d.Currency),
		},
		Balance: d.Balance,
		Limit:   d.Limit,
		UserId:  uuid.UUID(d.UserId),
	}
}
