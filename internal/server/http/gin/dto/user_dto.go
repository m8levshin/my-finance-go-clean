package dto

import (
	"github.com/google/uuid"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (r *CreateUserRequest) MapToUpdatableFields() *map[domainuser.UserUpdatableProperty]*string {
	return &map[domainuser.UserUpdatableProperty]*string{
		domainuser.Name:     &r.Name,
		domainuser.Email:    &r.Email,
		domainuser.Password: &r.Password,
	}
}

type UserDto struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func MapUserDomainToDto(r *domainuser.User) *UserDto {
	return &UserDto{
		Id:    uuid.UUID(r.Id),
		Name:  r.Name,
		Email: r.Email,
	}
}
