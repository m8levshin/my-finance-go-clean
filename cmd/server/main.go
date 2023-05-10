package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/datasource/memory_rw"
	"github.com/mlevshin/my-finance-go-clean/internal/logger"
	server "github.com/mlevshin/my-finance-go-clean/internal/server/http/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
)

func main() {
	engine := gin.Default()
	userRw := memory_rw.NewMemoryUserRW()
	assetRw := memory_rw.NewMemoryAssetRW(&userRw)
	transactionRw := memory_rw.NewMemoryTransactionRW()

	server.
		NewRouter(
			uc.HandlerBuilder{UserRw: userRw, AssetRw: assetRw, TransactionRw: transactionRw}.Build(),
			logger.NewLogger("debug", "text"),
		).
		SetRoutes(engine).
		Run()
}
