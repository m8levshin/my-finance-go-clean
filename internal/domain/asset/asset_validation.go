package asset

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

//go:generate mockery --name AssetValidator
type AssetValidator interface {
	validateBalanceAndLimitForTransaction(a *Asset, trx *Transaction) error
	validateAssetForCreateAndUpdate(a *Asset) error
}

type assetValidator struct {
}

const (
	DebitNotAllowed     = "debit is not allowed"
	ReachedLimits       = "you've reached limits"
	NotAllowedAssetType = "not allowed type of the asset"
)

func (v *assetValidator) validateBalanceAndLimitForTransaction(a *Asset, trx *Transaction) error {
	resultBalance := a.Balance - trx.Volume
	if resultBalance < 0 && !allowDebit[a.Type] {

		return errors.New(DebitNotAllowed)
	} else if allowDebit[a.Type] {
		if resultBalance+a.Limit < 0 {
			return errors.New(ReachedLimits)
		}
	}
	return nil
}

func (v *assetValidator) validateAssetForCreateAndUpdate(a *Asset) error {

	err := validation.ValidateStruct(
		a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Currency, validation.Required),
		validation.Field(&a.Owner, validation.Required),
		validation.Field(&a.Owner.Id, validation.Required),
		validation.Field(&a.Type, validation.Required),
	)
	if err != nil {
		return err
	}

	if !allowedTypes[a.Type] {
		return errors.New(NotAllowedAssetType)
	}
	return nil
}
