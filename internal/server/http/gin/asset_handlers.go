package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (rH *RouterHandler) getAssetsByUser(c *gin.Context) {
	userUUIDParam := c.Param("uuid")
	userUUID, err := uuid.Parse(userUUIDParam)

	if err != nil {
		c.Error(err)
		return
	}

	assets, err := rH.ucHandler.GetAssetsByUserId(userUUID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, assets)

}
