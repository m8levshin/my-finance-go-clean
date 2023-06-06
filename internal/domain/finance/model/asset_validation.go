package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

var (
	NotAllowedAssetTypeError = domain.NewError("not allowed type of the asset")
	DebitNotAllowedError     = domain.NewError("debit is not allowed")
	ReachedLimitsError       = domain.NewError("you've reached limits")
)

const ()

func ValidateBalanceAndLimitForTransaction(a *Asset, trx *Transaction) error {
	resultBalance := a.Balance + trx.Volume

	if resultBalance < 0 && !AllowDebit[a.Type] {
		return DebitNotAllowedError
	} else if AllowDebit[a.Type] {
		if resultBalance+a.Limit < 0 {
			return ReachedLimitsError
		}
	}
	return nil
}

func ValidateAssetForCreateAndUpdate(a *Asset) error {

	err := validation.ValidateStruct(
		a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Currency, validation.Required),
		validation.Field(&a.UserId, validation.Required),
		validation.Field(&a.Type, validation.NotNil),
	)

	if err != nil {
		return domain.ConvertValidationError(err)
	}

	if !AllowedTypes[a.Type] {
		return NotAllowedAssetTypeError
	}

	return nil
}
