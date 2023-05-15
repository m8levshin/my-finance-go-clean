package auth

import "github.com/gin-gonic/gin"

func GetUserInfoFromGinContext(c *gin.Context) *UserInfo {
	value, _ := c.Get(ginContextUserInfoKey)
	userInfo := value.(*UserInfo)
	return userInfo
}
