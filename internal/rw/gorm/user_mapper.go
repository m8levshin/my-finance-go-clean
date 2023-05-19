package gorm

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

func mapUserToDomain(e *user) *domainuser.User {
	return &domainuser.User{
		Id:           domain.Id(e.Id),
		Name:         e.Name,
		Email:        e.Email,
		PasswordHash: e.PasswordHash,
	}
}

func mapUserToEntity(u *domainuser.User) *user {
	userUUID := uuid.UUID(u.Id)
	return &user{
		Base:         Base{Id: userUUID},
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}
