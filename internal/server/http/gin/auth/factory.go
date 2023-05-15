package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/mlevshin/my-finance-go-clean/config"
)

const ginContextUserInfoKey = "user_info"

type OAuth2MiddlewareFactory interface {
	GetMiddleware() func(c *gin.Context)
}

type factory struct {
	config         config.AuthConfig
	cache          jwk.Cache
	userAuthInfoRw UserAuthInfoRW
}

func CreateOAuth2ResourceServerMiddlewareFactory(
	c config.Configuration,
	rw UserAuthInfoRW,
) (OAuth2MiddlewareFactory, error) {

	cache, err := initJWKSCache(c.Auth)
	if err != nil {
		return nil, err
	}

	return &factory{
		config:         c.Auth,
		cache:          *cache,
		userAuthInfoRw: rw,
	}, nil
}
