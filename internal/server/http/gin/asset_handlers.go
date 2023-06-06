package gin

import (
	"errors"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/auth"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
	"net/http"
	"time"

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

	if !verifyUserOwnershipOrAdminAccess(c, userUUID) {
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

	if !verifyUserOwnershipOrAdminAccess(c, body.UserId) {
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

	if !verifyUserOwnershipOrAdminAccess(c, uuid.UUID(asset.UserId)) {
		return
	}

	c.JSON(http.StatusOK, dto.MapAssetDomainToDto(asset))
}

func (rH *RouterHandler) getBalanceTracking(c *gin.Context) {
	authInfo := auth.GetUserInfoFromGinContext(c)
	userId := authInfo.Id
	assetIdParam := c.Param("uuid")
	assetId, err := uuid.Parse(assetIdParam)
	if err != nil {
		c.Error(err)
		return
	}

	_, currencyMode := c.GetQuery("currency")
	fromParam, fromExist := c.GetQuery("from")
	toParam, toExist := c.GetQuery("to")
	tzParam, tzExist := c.GetQuery("tz")

	if !toExist || !fromExist || !tzExist {
		c.AbortWithError(http.StatusBadRequest, errors.New("missing input parameters"))
		return
	}

	from, err := time.Parse(time.DateOnly, fromParam)
	to, err := time.Parse(time.DateOnly, toParam)
	tz, err := time.LoadLocation(tzParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("incorrect date parameters"))
	}

	if currencyMode {
		c.AbortWithStatus(http.StatusNotImplemented)
	} else {
		balanceStateHistory, err := rH.ucHandler.GetBalanceStateHistory(userId, assetId, from, to, tz, isAdmin(authInfo))
		if err == nil {
			c.JSON(http.StatusOK, &balanceStateHistory)
		}
	}
}
