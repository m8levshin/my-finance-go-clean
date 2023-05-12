package gorm

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

func (e *userEntity) mapToDomain() *domainuser.User {
	return &domainuser.User{
		Id:           domain.Id(*(e.ID)),
		Name:         e.Name,
		Email:        e.Email,
		PasswordHash: e.PasswordHash,
	}
}

func mapToEntity(user *domainuser.User) *userEntity {
	userUUID := uuid.UUID(user.Id)
	return &userEntity{
		Base:         Base{ID: &userUUID},
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}
