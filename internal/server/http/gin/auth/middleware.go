package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/mlevshin/my-finance-go-clean/config"
	"net/http"
	"strings"
)

type AdditionalValidation func(context.Context, jwt.Token) jwt.ValidationError

func (f *factory) GetMiddleware(additionalValidations ...AdditionalValidation) func(c *gin.Context) {

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

		validateOptions := generateValidationOptions(keySet, f.config, additionalValidations)
		jwtToken, err := jwt.ParseString(authHeaderParts[1], validateOptions...)

		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}
		f.resolveUserAndFillContextByUserInfo(c, &jwtToken)
		c.Next()
	}
}

func generateValidationOptions(keySet jwk.Set, c config.AuthConfig, additionalValidations []AdditionalValidation) []jwt.ParseOption {
	validateOptions := make([]jwt.ParseOption, 0)
	validateOptions = append(validateOptions, jwt.WithKeySet(keySet))
	validateOptions = append(validateOptions, jwt.WithVerify(true))
	validateOptions = append(validateOptions, jwt.WithAudience(c.Audience))
	for _, validation := range additionalValidations {
		validateOptions = append(validateOptions, jwt.WithValidator(jwt.ValidatorFunc(validation)))
	}
	return validateOptions
}

func (f *factory) resolveUserAndFillContextByUserInfo(c *gin.Context, jwtToken *jwt.Token) bool {
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

func WithUserGroupValidation(group string) AdditionalValidation {
	return func(ctx context.Context, t jwt.Token) jwt.ValidationError {
		groupsValue, exist := t.Get("groups")
		if !exist {
			return jwt.NewValidationError(fmt.Errorf("token is not valid"))

		}

		ok := containsGroup(groupsValue, group)
		if !ok {
			return jwt.NewValidationError(fmt.Errorf("token is not valid"))
		}

		return nil
	}
}

func containsGroup(slice any, checkingRole string) bool {
	groups := slice.([]interface{})
	for _, v := range groups {
		if str, ok := v.(string); ok {
			if ok && str == checkingRole {
				return true
			}
		}
	}
	return false
}
