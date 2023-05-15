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

	assetDtos := make([]dto.AssetDto, 0, len(assets))
	for _, asset := range assets {
		assetDtos = append(assetDtos, *dto.MapAssetDomainToDto(asset))
	}
	c.JSON(http.StatusOK, assetDtos)
}

func (rH *RouterHandler) postAsset(c *gin.Context) {
	body := dto.CreateAssetRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	asset, err := rH.ucHandler.CreateNewAsset(body.UserId, body.MapToUpdatableFields())
	if asset != nil && err == nil {
		c.JSON(http.StatusCreated, dto.MapAssetDomainToDto(asset))
		return
	}
	if err != nil {
		c.Status(500)
		c.Error(err)
	}
}

func (rH *RouterHandler) getAssetById(c *gin.Context) {
	assetIdParam := c.Param("uuid")
	assetId, err := uuid.Parse(assetIdParam)
	if err != nil {
		c.Error(err)
		return
	}

	asset, err := rH.ucHandler.GetAssetById(assetId)
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.MapAssetDomainToDto(asset))
}
