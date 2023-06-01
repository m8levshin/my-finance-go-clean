package auth

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
	"github.com/patrickmn/go-cache"
	"time"
)

type UserInfo struct {
	Id    uuid.UUID
	Roles []string
}

type UserAuthInfoService interface {
	GetUserAuthInfoByEmail(email string) (*UserInfo, error)
	CreateNewUser(email string, name string) (*UserInfo, error)
}

type inMemoryCachedUserAuthInfoRW struct {
	cache  *cache.Cache
	userUC uc.UserLogic
}

func (u *inMemoryCachedUserAuthInfoRW) CreateNewUser(email string, name string) (*UserInfo, error) {
	newUser, err := u.userUC.CreateNewUser(map[domain.UpdatableProperty]any{
		user.NameField:  &name,
		user.EmailField: &email,
	})
	if err != nil {
		return nil, err
	}

	return &UserInfo{Id: uuid.UUID(newUser.Id)}, nil
}

func NewInMemoryCachedUserAuthService(userLogic uc.UserLogic) UserAuthInfoService {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &inMemoryCachedUserAuthInfoRW{userUC: userLogic, cache: c}
}

func (u *inMemoryCachedUserAuthInfoRW) GetUserAuthInfoByEmail(email string) (*UserInfo, error) {
	var userInfo *UserInfo
	cacheValue, exist := u.cache.Get(email)
	if exist {
		userInfo = cacheValue.(*UserInfo)
		return userInfo, nil
	}

	authUser, err := u.userUC.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if authUser == nil {
		return nil, nil
	}

	userInfo = &UserInfo{Id: uuid.UUID(authUser.Id)}
	u.cache.Set(email, userInfo, 0)
	return userInfo, err
}
