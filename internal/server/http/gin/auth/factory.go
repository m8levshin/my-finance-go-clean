package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
)

type OAuth2MiddlewareFactory interface {
	GetMiddleware() func(c *gin.Context)
}

type factory struct {
	cache      jwk.Cache
	userInfoRw userInfoRW
}

const GinContextUserInfoKey = "user_info"
const certsUrl = `https://auth.mlevsh.in/realms/project-base/protocol/openid-connect/certs`

func CreateOAuth2ResourceServerMiddlewareFactory(rw rw.UserRW) OAuth2MiddlewareFactory {
	return &factory{
		cache: *initJWKSCache(), userInfoRw: newUserInfoRW(rw),
	}
}
