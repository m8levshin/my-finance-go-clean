package auth

import (
	"context"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"time"
)

func initJWKSCache() *jwk.Cache {
	ctx, _ := context.WithCancel(context.Background())

	c := jwk.NewCache(ctx)
	err := c.Register(certsUrl, jwk.WithMinRefreshInterval(15*time.Minute))
	if err != nil {
		return nil
	}

	_, err = c.Refresh(ctx, certsUrl)
	if err != nil {
		fmt.Printf("Failed to refresh JWKS: %s\n", err)
		return nil
	}

	return c
}
