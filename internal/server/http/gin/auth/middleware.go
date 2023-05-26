package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/mlevshin/my-finance-go-clean/config"
	"net/http"
	"strings"
)

type AdditionalValidation func(context.Context, jwt.Token) jwt.ValidationError

var (
	tokenValidationError = jwt.NewValidationError(fmt.Errorf("token is not valid"))
)

func (f *factory) GetMiddleware(additionalValidations ...AdditionalValidation) func(c *gin.Context) {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(strings.TrimSpace(authHeader)) == 0 {
			c.AbortWithError(http.StatusUnauthorized, tokenValidationError)
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 && authHeaderParts[0] != "Bearer" {
			c.AbortWithError(http.StatusBadRequest, tokenValidationError)
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
		ok := f.resolveUserAndFillContextByUserInfo(c, &jwtToken)
		if !ok {
			return
		}
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
		c.AbortWithError(http.StatusForbidden, tokenValidationError)
		return false
	}

	email := emailClaim.(string)
	user, err := f.userAuthInfoService.GetUserAuthInfoByEmail(email)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.Join(err, tokenValidationError))
		return false
	}

	if user == nil {
		nameClaim, exist := (*jwtToken).Get("name")
		if !exist {
			c.AbortWithError(http.StatusBadRequest, tokenValidationError)
			return false
		}
		name := nameClaim.(string)
		newUser, err := f.userAuthInfoService.CreateNewUser(email, name)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return false
		}
		user = newUser
	}
	c.Set(ginContextUserInfoKey, user)
	return true
}

func WithUserGroupValidation(group string) AdditionalValidation {
	return func(ctx context.Context, t jwt.Token) jwt.ValidationError {
		groupsValue, exist := t.Get("groups")
		if !exist {
			return tokenValidationError

		}

		ok := containsGroup(groupsValue, group)
		if !ok {
			return tokenValidationError
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
