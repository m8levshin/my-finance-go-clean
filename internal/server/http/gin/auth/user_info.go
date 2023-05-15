package auth

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
	"github.com/patrickmn/go-cache"
	"time"
)

type UserInfo struct {
	Id uuid.UUID
}

type UserAuthInfoRW interface {
	GetUserAuthInfoByEmail(email string) (*UserInfo, error)
}

type inMemoryCachedUserAuthInfoRW struct {
	cache  *cache.Cache
	userRW rw.UserRW
}

func NewInMemoryCachedUserAuthInfoRW(userRW rw.UserRW) UserAuthInfoRW {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &inMemoryCachedUserAuthInfoRW{userRW: userRW, cache: c}
}

func (u *inMemoryCachedUserAuthInfoRW) GetUserAuthInfoByEmail(email string) (*UserInfo, error) {
	var userInfo *UserInfo
	cacheValue, exist := u.cache.Get(email)
	if exist {
		userInfo = cacheValue.(*UserInfo)
		return userInfo, nil
	}

	user, err := u.userRW.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	userInfo = &UserInfo{uuid.UUID(user.Id)}
	u.cache.Set(email, userInfo, 0)
	return userInfo, err
}
