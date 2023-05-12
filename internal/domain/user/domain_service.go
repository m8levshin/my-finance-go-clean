package user

import "github.com/mlevshin/my-finance-go-clean/internal/domain"

type UserDomainService interface {
	CreateUser(opts ...func(u *User) error) (*User, error)
	UpdateUser(initial *User, opts ...*func(u *User) error) (*User, error)
}

type userDomainService struct {
	logger domain.Logger
}

func CreateUserService(l *domain.Logger) UserDomainService {
	return &userDomainService{
		logger: *l,
	}
}

func (s *userDomainService) CreateUser(opts ...func(u *User) error) (*User, error) {

	newUser := User{
		Id: domain.NewID(),
	}
	for _, f := range opts {
		err := f(&newUser)
		if err != nil {
			return nil, err
		}
	}
	err := validateForCreateAndUpdate(&newUser)
	if err != nil {
		return nil, err

	}

	return &newUser, nil
}

func (s *userDomainService) UpdateUser(initial *User, opts ...*func(u *User) error) (*User, error) {
	for _, function := range opts {
		f := *function
		err := f(initial)
		if err != nil {
			return nil, err
		}
	}
	if err := validateForCreateAndUpdate(initial); err != nil {
		return nil, err
	}
	return initial, nil
}
