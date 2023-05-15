package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/auth"
)

func getUserInfoFromContext(c *gin.Context) *auth.UserInfo {
	value, _ := c.Get(auth.GinContextUserInfoKey)
	userInfo := value.(*auth.UserInfo)
	return userInfo
}
