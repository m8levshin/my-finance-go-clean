package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/config"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/auth"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
	"golang.org/x/exp/slices"
	"net/http"
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

	adminMiddleware := authMiddlewareFactory.GetMiddleware(auth.WithUserGroupValidation("admin"))
	userMiddleware := authMiddlewareFactory.GetMiddleware()

	api := r.Group("/api")
	api.Use(rH.mutualMiddlewares...)

	usersApi := api.Group("/users")
	usersApi.GET("/me", userMiddleware, rH.getCurrentUser)
	usersApi.GET("", adminMiddleware, rH.getAllUsers)
	usersApi.GET("/:uuid", adminMiddleware, rH.getUserById)

	usersApi.GET("/:uuid/assets", userMiddleware, rH.getAssetsByUser)
	usersApi.GET("/:uuid/groups", userMiddleware, rH.getGroupsByUser)
	usersApi.POST("/:uuid/groups", userMiddleware, rH.createGroup)

	assetsApi := api.Group("/assets")
	assetsApi.GET("/:uuid", userMiddleware, rH.getAssetById)
	//assetsApi.GET("/:uuid/balance_tracking", rH.getBalanceTracking)
	assetsApi.GET("/:uuid/transactions", userMiddleware, rH.getTransactionsByAssetId)
	assetsApi.POST("/:uuid/transactions", userMiddleware, rH.addNewTransaction)
	assetsApi.POST("", userMiddleware, rH.postAsset)

	return r
}

func createErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(500, gin.H{"error": c.Errors.Last().Error()})
		}
	}
}

func isAdmin(info *auth.UserInfo) bool {
	return slices.Contains(info.Roles, auth.AdminGroup)
}

func verifyUserOwnershipOrAdminAccess(c *gin.Context, userUUID uuid.UUID) bool {
	authInfo := auth.GetUserInfoFromGinContext(c)
	if userUUID != authInfo.Id && isAdmin(authInfo) {
		c.AbortWithStatus(http.StatusForbidden)
		return false
	}
	return true
}
