package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"strings"
)

func (f *factory) GetMiddleware() func(c *gin.Context) {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(strings.TrimSpace(authHeader)) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 && authHeaderParts[0] != "Bearer" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		keySet, err := f.cache.Get(context.Background(), f.config.JwksUrl)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		jwtToken, err := jwt.ParseString(authHeaderParts[1], jwt.WithKeySet(keySet), jwt.WithVerify(true))
		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}
		f.resolveUserAndFillContestByUserInfo(c, &jwtToken)
		c.Next()
	}
}

func (f *factory) resolveUserAndFillContestByUserInfo(c *gin.Context, jwtToken *jwt.Token) bool {
	emailClaim, exist := (*jwtToken).Get("email")
	if !exist {
		c.AbortWithStatus(http.StatusForbidden)
		return false
	}

	email := emailClaim.(string)
	user, err := f.userAuthInfoRw.GetUserAuthInfoByEmail(email)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return false
	}

	if user == nil {
		//todo: create user, if user have been created outside the app
		c.AbortWithStatus(http.StatusNotImplemented)
		return false
	}
	c.Set(ginContextUserInfoKey, user)
	return true
}
