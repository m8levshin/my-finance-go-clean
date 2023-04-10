package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RouterHandler) getAllUsers(c *gin.Context) {
	users, err := rH.ucHandler.GetAllUsers()
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, users)
}
