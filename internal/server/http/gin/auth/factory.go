package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/mlevshin/my-finance-go-clean/config"
)

const ginContextUserInfoKey = "user_info"

type OAuth2MiddlewareFactory interface {
	GetMiddleware(additionalValidations ...AdditionalValidation) func(c *gin.Context)
}

type factory struct {
	config              config.AuthConfig
	cache               jwk.Cache
	userAuthInfoService UserAuthInfoService
}

func CreateOAuth2ResourceServerMiddlewareFactory(
	c config.Configuration,
	userAuthInfoService UserAuthInfoService,
) (OAuth2MiddlewareFactory, error) {

	cache, err := initJWKSCache(c.Auth)
	if err != nil {
		return nil, err
	}

	return &factory{
		config:              c.Auth,
		cache:               *cache,
		userAuthInfoService: userAuthInfoService,
	}, nil
}
