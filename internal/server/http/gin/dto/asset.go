package dto

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
)

type CreateAssetRequest struct {
	Name     string    `json:"name" binding:"required"`
	Type     string    `json:"type" binding:"required"`
	Limit    float64   `json:"limit"`
	UserId   uuid.UUID `json:"userId" binding:"required"`
	Currency string    `json:"currency" binding:"required"`
}

func (r *CreateAssetRequest) MapToUpdatableFields() map[domain.UpdatableProperty]any {
	createUserFields := map[domain.UpdatableProperty]any{}
	createUserFields[model.AssetNameField] = &(r.Name)
	createUserFields[model.AssetTypeField] = &(r.Type)
	createUserFields[model.AssetLimitField] = &(r.Limit)
	createUserFields[model.AssetUserIdField] = &(r.UserId)
	createUserFields[model.AssetCurrencyField] = &(r.Currency)

	return createUserFields
}

type AssetDto struct {
	Id       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Limit    float64    `json:"limit"`
	UserId   *uuid.UUID `json:"userId"`
	Currency string     `json:"currency"`
	Balance  float64    `json:"balance"`
}

func MapAssetDomainToDto(r *model.Asset) *AssetDto {
	userId := uuid.UUID(r.UserId)
	return &AssetDto{
		Id:       uuid.UUID(r.Id),
		Name:     r.Name,
		Type:     model.TypeNames[r.Type],
		Limit:    r.Limit,
		UserId:   &userId,
		Currency: string(r.Currency),
		Balance:  r.Balance,
	}
}
