package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
	"net/http"
)

func (rH *RouterHandler) getGroupsByUser(c *gin.Context) {
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
