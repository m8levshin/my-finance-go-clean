package asset

import "github.com/mlevshin/my-finance-go-clean/internal/domain"

type TransactionGroup struct {
	Id       domain.Id
	ParentId domain.Id
	UserId   domain.Id
	Name     string
}
