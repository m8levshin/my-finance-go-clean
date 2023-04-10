package userRW

import (
	"errors"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
	"sync"
)

type userRW struct {
	store *sync.Map
}

func NewMemoryUserRW() rw.UserRW {
	return userRW{
		store: &sync.Map{},
	}
}

func (rw userRW) FindAll() ([]*domainuser.User, error) {
	toReturn := []*domainuser.User{}
	rw.store.Range(func(key, value any) bool {
		user, ok := value.(domainuser.User)
		if ok {
			toReturn = append(toReturn, &user)
		}
		return true
	})
	return toReturn, nil
}

func (rw userRW) FindById(id domain.Id) (*domainuser.User, error) {
	value, _ := rw.store.Load(id)
	if value == nil {
		return nil, errors.New("user not found")
	}
	user := value.(domainuser.User)
	return &user, nil
}

func (rw userRW) Save(user domainuser.User) error {
	value, _ := rw.store.Load(user.Id)
	if value == nil {
		return errors.New("user not found")
	}
	rw.store.Store(user.Id, user)
	return nil
}
