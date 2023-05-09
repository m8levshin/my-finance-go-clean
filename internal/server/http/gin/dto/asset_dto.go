package dto

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
)

type CreateAssetRequest struct {
	Name     string    `json:"name" binding:"required"`
	Type     string    `json:"type" binding:"required"`
	Limit    string    `json:"limit" binding:"required"`
	OwnerId  uuid.UUID `json:"ownerId"`
	Currency string    `json:"currency" binding:"required"`
}

func (r *CreateAssetRequest) MapToUpdatableFields() map[domain.UpdatableProperty]any {
	createUserFields := map[domain.UpdatableProperty]any{}
	createUserFields[domainasset.NameField] = &(r.Name)
	createUserFields[domainasset.TypeField] = &(r.Type)
	createUserFields[domainasset.LimitField] = &(r.Limit)
	createUserFields[domainasset.OwnerField] = &(r.OwnerId)
	createUserFields[domainasset.CurrencyField] = &(r.Currency)

	return createUserFields
}

type AssetDto struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Limit    float64   `json:"limit"`
	OwnerId  uuid.UUID `json:"ownerId"`
	Currency string    `json:"currency"`
}

func MapAssetDomainToDto(r *domainasset.Asset) *AssetDto {
	return &AssetDto{
		Id:       uuid.UUID(r.Id),
		Name:     r.Name,
		Type:     domainasset.TypeNames[r.Type],
		Limit:    r.Limit,
		OwnerId:  uuid.UUID(r.OwnerId),
		Currency: string(r.Currency),
	}
}
