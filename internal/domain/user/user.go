package user

import (
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
)

type User struct {
	Id           domain.Id
	Name         string
	Email        string
	PasswordHash []byte
}
