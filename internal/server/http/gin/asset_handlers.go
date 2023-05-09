package gin

import (
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
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

func (rH *RouterHandler) getTransactionsByAssetId(c *gin.Context) {
	assetUUIDParam := c.Param("uuid")
	assetUUID, err := uuid.Parse(assetUUIDParam)
	if err != nil {
		c.Error(err)
		return
	}
	transactions, err := rH.ucHandler.GetTransactionsByAssetId(assetUUID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, transactions)

}

func (rH *RouterHandler) postAssetForUser(c *gin.Context) {
	userUUIDParam := c.Param("uuid")
	userUUID, err := uuid.Parse(userUUIDParam)

	if err != nil {
		c.Error(err)
		return
	}

	body := dto.CreateAssetRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	body.OwnerId = userUUID

	asset, err := rH.ucHandler.CreateNewAsset(body.MapToUpdatableFields())
	if asset != nil && err == nil {
		c.JSON(http.StatusCreated, dto.MapAssetDomainToDto(asset))
		return
	}
	if err != nil {
		c.Status(500)
		c.Error(err)
	}
}
