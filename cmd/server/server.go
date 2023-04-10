package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/datasource/memory/userRW"
	"github.com/mlevshin/my-finance-go-clean/internal/logger"
	server "github.com/mlevshin/my-finance-go-clean/internal/server/http/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
)

func main() {
	engine := gin.Default()
	userRw := userRW.NewMemoryUserRW()
	server.NewRouter(
		uc.HandlerBuilder{UserRw: userRw}.Build(),
		logger.NewLogger("debug", "text"),
	).SetRoutes(engine)

	engine.Run()
}
