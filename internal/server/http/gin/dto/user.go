package dto

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) MapToUpdatableFields() *map[domain.UpdatableProperty]any {

	createUserFields := map[domain.UpdatableProperty]any{}
	createUserFields[domainuser.NameField] = &(r.Name)
	createUserFields[domainuser.EmailField] = &(r.Email)
	createUserFields[domainuser.PasswordField] = &(r.Password)

	return &createUserFields
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
