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

func NewRouter(ucHandler uc.Handler, logger domain.Logger) RouterHandler {
	return RouterHandler{
		ucHandler:         ucHandler,
		logger:            logger,
		mutualMiddlewares: []gin.HandlerFunc{},
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(rH.mutualMiddlewares...)
	rH.usersRoutes(api)
}

func (rH RouterHandler) usersRoutes(api *gin.RouterGroup) {
	usersApi := api.Group("/users")
	usersApi.GET("", rH.getAllUsers)
}
