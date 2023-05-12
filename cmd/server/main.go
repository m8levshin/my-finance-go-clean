package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/rw/gorm"
	"github.com/mlevshin/my-finance-go-clean/internal/rw/memory"
	server "github.com/mlevshin/my-finance-go-clean/internal/server/http/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
	"log"
	_ "net/http/pprof"
)

func main() {

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

	engine := gin.Default()
	server.
		NewRouter(
			uc.HandlerBuilder{
				UserRw:             userRw,
				AssetRw:            assetRw,
				TransactionRw:      transactionRw,
				TransactionGroupRw: transactionGroupRw,
			}.Build(),
		).
		SetRoutes(engine).
		Run()
}
