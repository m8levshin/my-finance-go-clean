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

func mapAssetsToDomains(e []*asset) []*domainasset.Asset {
	result := make([]*domainasset.Asset, 0, len(e))
	for _, entity := range e {
		result = append(result, mapAssetToDomain(entity))
	}
	return result
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

func mapAssetToEntities(d []*domainasset.Asset) []*asset {
	result := make([]*asset, 0, len(d))
	for _, domainItem := range d {
		result = append(result, mapAssetToEntity(domainItem))
	}
	return result
}
