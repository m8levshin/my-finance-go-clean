package service

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
)

type AssetStateDomainService interface {
	CreateNewAssetState(asset *model.Asset, transactions []*model.Transaction) (*model.AssetState, error)
}

type assetStateService struct {
	logger domain.Logger
}

func NewAssetStateService(l *domain.Logger) AssetStateDomainService {
	return &assetStateService{
		logger: *l,
	}
}

func (s *assetStateService) CreateNewAssetState(asset *model.Asset, transactions []*model.Transaction) (*model.AssetState, error) {
	assetState := model.AssetState{
		Asset:        asset,
		Transactions: transactions,
	}
	err := assetState.CheckTransaction()

	if err != nil {
		return nil, err
	}
	return &assetState, nil
}
