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

type userInfoRW interface {
	GetUserByEmail(email string) (*UserInfo, error)
}

type inMemoryCachedUserInfoRW struct {
	cache  *cache.Cache
	userRW rw.UserRW
}

func newUserInfoRW(userRW rw.UserRW) userInfoRW {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &inMemoryCachedUserInfoRW{userRW: userRW, cache: c}
}

func (u *inMemoryCachedUserInfoRW) GetUserByEmail(email string) (*UserInfo, error) {
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
