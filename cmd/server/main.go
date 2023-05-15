package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/config"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/transaction_group"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"github.com/mlevshin/my-finance-go-clean/internal/log"
	"github.com/mlevshin/my-finance-go-clean/internal/rw/gorm"
	"github.com/mlevshin/my-finance-go-clean/internal/rw/memory"
	server "github.com/mlevshin/my-finance-go-clean/internal/server/http/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/auth"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
)

func main() {
	appConfig := config.InitAndReadConfig()

	logger := log.InitLogger("info", "text")
	db, err := gorm.InitGorm()
	if err != nil {
		log.Fatal(err.Error())
	}

	userRw, err := gorm.NewUserRw(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	assetRw := memory.NewMemoryAssetRW(&userRw)
	transactionRw := memory.NewMemoryTransactionRW()
	transactionGroupRw := memory.NewMemoryTransactionGroupRW()

	assetService := asset.NewAssetService(&logger)
	userService := user.CreateUserService(&logger)
	transactionGroupService := transaction_group.NewTransactionGroupService(&logger)

	oAuth2MiddlewareFactory := auth.CreateOAuth2ResourceServerMiddlewareFactory(userRw)

	engine := gin.Default()
	server.
		NewRouter(
			uc.HandlerBuilder{
				Config:                  appConfig,
				UserRw:                  userRw,
				AssetRw:                 assetRw,
				TransactionRw:           transactionRw,
				TransactionGroupRw:      transactionGroupRw,
				AssetService:            assetService,
				UserService:             userService,
				TransactionGroupService: transactionGroupService,
			}.Build(),
			appConfig,
		).
		SetRoutes(engine, oAuth2MiddlewareFactory).
		Run()
}
