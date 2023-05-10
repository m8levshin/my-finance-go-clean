package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
)

type RouterHandler struct {
	ucHandler         uc.Handler
	logger            domain.Logger
	mutualMiddlewares []gin.HandlerFunc
}

func NewRouter(ucHandler uc.Handler, logger domain.Logger) *RouterHandler {
	return &RouterHandler{
		ucHandler:         ucHandler,
		logger:            logger,
		mutualMiddlewares: []gin.HandlerFunc{createErrorHandlerMiddleware()},
	}
}

func (rH *RouterHandler) SetRoutes(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	api.Use(rH.mutualMiddlewares...)
	rH.usersRoutes(api)
	rH.assetsRoutes(api)
	return r
}

func (rH *RouterHandler) usersRoutes(api *gin.RouterGroup) {
	usersApi := api.Group("/users")
	usersApi.GET("", rH.getAllUsers)
	usersApi.GET("/:uuid", rH.getUserById)
	usersApi.POST("", rH.createUser)
	usersApi.GET("/:uuid/assets", rH.getAssetsByUser)
}

func (rH *RouterHandler) assetsRoutes(api *gin.RouterGroup) {
	assetsApi := api.Group("/assets")
	assetsApi.GET("/:uuid", rH.getAssetById)
	assetsApi.GET("/:uuid/transactions", rH.getTransactionsByAssetId)
	assetsApi.POST("/:uuid/transactions", rH.addNewTransaction)
	assetsApi.POST("", rH.postAsset)
}

func createErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(500, gin.H{"error": c.Errors.Last().Error()})
		}

	}
}
