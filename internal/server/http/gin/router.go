package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/config"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/auth"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
)

type RouterHandler struct {
	config            config.Configuration
	ucHandler         uc.Handler
	mutualMiddlewares []gin.HandlerFunc
}

func NewRouterHandler(ucHandler uc.Handler, config config.Configuration) *RouterHandler {
	return &RouterHandler{
		config:    config,
		ucHandler: ucHandler,
		mutualMiddlewares: []gin.HandlerFunc{
			createErrorHandlerMiddleware(),
		},
	}
}

func (rH *RouterHandler) SetRoutes(r *gin.Engine, authMiddlewareFactory auth.OAuth2MiddlewareFactory) *gin.Engine {
	api := r.Group("/api")
	api.Use(rH.mutualMiddlewares...)
	rH.usersRoutes(api)
	rH.assetsRoutes(api)
	return r
}

func (rH *RouterHandler) usersRoutes(api *gin.RouterGroup) {

	usersApi := api.Group("/users")
	//usersApi.GET("/me", rH.getCurrentUser)            //доступно для любого пользователя
	usersApi.GET("", rH.getAllUsers)                  //доступно для админа
	usersApi.GET("/:uuid", rH.getUserById)            //доступно для админа
	usersApi.POST("", rH.createUser)                  //доступно для админа
	usersApi.GET("/:uuid/assets", rH.getAssetsByUser) //доступно для админа, либо для текущего пользователя
	usersApi.GET("/:uuid/groups", rH.getGroupsByUser) //доступно для админа, либо для текущего пользователя
	usersApi.POST("/:uuid/groups", rH.createGroup)    //доступно для админа, либо для текущего пользователя

}

func (rH *RouterHandler) assetsRoutes(api *gin.RouterGroup) {
	assetsApi := api.Group("/assets")
	assetsApi.GET("/:uuid", rH.getAssetById) //доступно для админа, либо для владельца
	//assetsApi.GET("/:uuid/balance_tracking", rH.getBalanceTracking)   //доступно для админа, либо для владельца
	assetsApi.GET("/:uuid/transactions", rH.getTransactionsByAssetId) //доступно для админа, либо для владельца
	assetsApi.POST("/:uuid/transactions", rH.addNewTransaction)       //доступно для админа, либо для владельца
	assetsApi.POST("", rH.postAsset)                                  //доступно для админа, либо для текущего пользователя
}

func createErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(500, gin.H{"error": c.Errors.Last().Error()})
		}

	}
}
