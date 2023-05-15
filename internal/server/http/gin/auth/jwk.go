package auth

import (
	"context"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/mlevshin/my-finance-go-clean/config"
	"github.com/mlevshin/my-finance-go-clean/internal/log"
	"time"
)

func initJWKSCache(config config.AuthConfig) *jwk.Cache {
	ctx, _ := context.WithCancel(context.Background())

	c := jwk.NewCache(ctx)
	err := c.Register(config.JwksUrl, jwk.WithMinRefreshInterval(15*time.Minute))
	if err != nil {
		return nil
	}

	_, err = c.Refresh(ctx, config.JwksUrl)
	if err != nil {
		log.Fatal("Failed to refresh JWKS: %s\n", err)
		return nil
	}

	return c
}
