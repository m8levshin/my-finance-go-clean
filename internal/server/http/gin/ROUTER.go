package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/uc"
)

type RouterHandler struct {
	ucHandler uc.Handler
}

func NewRouter(ucHandler uc.Handler) RouterHandler {
	return RouterHandler{
		ucHandler: ucHandler,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	rH.usersRoutes(api)
}

func (rH RouterHandler) usersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	users.GET("", rH.getAllUsers)
}
