package service

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
)

type AssetDomainService interface {
	CreateAsset(opts ...func(u *model.Asset) error) (*model.Asset, error)
}

type assetDomainService struct {
	logger domain.Logger
}

func NewAssetService(l *domain.Logger) AssetDomainService {
	return &assetDomainService{
		logger: *l,
	}
}

func (s *assetDomainService) CreateAsset(opts ...func(u *model.Asset) error) (*model.Asset, error) {
	newAsset := model.Asset{
		Id:      domain.NewID(),
		Balance: 0.0,
	}
	for _, f := range opts {
		err := f(&newAsset)
		if err != nil {
			return nil, err
		}
	}
	err := model.ValidateAssetForCreateAndUpdate(&newAsset)
	if err != nil {
		return nil, err
	}
	return &newAsset, nil
}
