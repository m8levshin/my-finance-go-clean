package auth

import (
	"context"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/exp/slices"
)

func WithUserGroupValidation(group string) AdditionalValidation {
	return func(ctx context.Context, t jwt.Token) jwt.ValidationError {
		ok, err := containsGroup(&t, group)
		if err != nil || !ok {
			return tokenValidationError
		}
		return nil
	}
}

func getUserGroups(jwtToken *jwt.Token) ([]string, error) {
	groupsClaim, exist := (*jwtToken).Get("groups")
	if !exist {
		return nil, tokenValidationError
	}

	groups, ok := groupsClaim.([]interface{})
	if !ok {
		return nil, tokenValidationError
	}
	stringGroups := make([]string, 0)
	for _, v := range groups {
		if str, ok := v.(string); ok {
			if ok {
				stringGroups = append(stringGroups, str)
			}
		}
	}

	return stringGroups, nil
}

func containsGroup(jwtToken *jwt.Token, checkingRole string) (bool, error) {
	groups, err := getUserGroups(jwtToken)
	if err != nil {
		return false, err
	}
	return slices.Contains(groups, checkingRole), nil
}
