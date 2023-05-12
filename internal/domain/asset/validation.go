package asset

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	NotAllowedAssetType = "not allowed type of the asset"
	DebitNotAllowed     = "debit is not allowed"
	ReachedLimits       = "you've reached limits"
)

const ()

func validateBalanceAndLimitForTransaction(a *Asset, trx *Transaction) error {
	resultBalance := a.Balance + trx.Volume

	if resultBalance < 0 && !AllowDebit[a.Type] {
		return errors.New(DebitNotAllowed)
	} else if AllowDebit[a.Type] {
		if resultBalance+a.Limit < 0 {
			return errors.New(ReachedLimits)
		}
	}
	return nil
}

func validateAssetForCreateAndUpdate(a *Asset) error {

	err := validation.ValidateStruct(
		a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Currency, validation.Required),
		validation.Field(&a.UserId, validation.Required),
		validation.Field(&a.Type, validation.NotNil),
	)
	if err != nil {
		return err
	}

	if !AllowedTypes[a.Type] {
		return errors.New(NotAllowedAssetType)
	}
	return nil
}
