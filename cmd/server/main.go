package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/config"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/rw"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/finance/service"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"github.com/mlevshin/my-finance-go-clean/internal/log"
	"github.com/mlevshin/my-finance-go-clean/internal/rw/gorm"
	server "github.com/mlevshin/my-finance-go-clean/internal/server/http/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/auth"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
	"go.uber.org/dig"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	log.InitLogger("info", "text")

	container := dig.New()

	container.Provide(config.InitAndReadConfig)
	container.Provide(log.NewDomainLogger)
	container.Provide(gorm.InitGorm)
	container.Provide(gorm.NewTransactionRW, dig.As(new(rw.TransactionRW), new(rw.TransactionGroupRW)))
	container.Provide(gorm.NewAssetRw)
	container.Provide(gorm.NewUserRw)
	container.Provide(gorm.NewExchangeRateRW)
	container.Provide(service.NewAssetService)
	container.Provide(service.NewAssetStateService)
	container.Provide(user.CreateUserService)
	container.Provide(service.NewTransactionGroupService)
	container.Provide(auth.NewInMemoryCachedUserAuthService)
	container.Provide(auth.CreateOAuth2ResourceServerMiddlewareFactory)
	container.Provide(uc.NewHandler, dig.As(new(uc.Handler), new(uc.UserLogic)))
	container.Provide(server.NewRouterHandler)

	err := container.Invoke(initAndRunServer)
	if err != nil {
		log.Fatal(err)
	}
}

func initAndRunServer(
	config config.Configuration,
	handler uc.Handler,
	authMiddlewareFactory auth.OAuth2MiddlewareFactory,
	e rw.ExchangeRateRW,
) {
	engine := gin.Default()
	routerHandler := server.NewRouterHandler(handler, config)
	routerHandler.SetRoutes(engine, authMiddlewareFactory)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Server.Port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}

	}()
	waitShutdownAndHandleIt(srv)
}

func waitShutdownAndHandleIt(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
