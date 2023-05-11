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

	transactionGroups, err := rH.ucHandler.GetTransactionGroupsByUser(userUUID)
	if err != nil {
		c.Error(err)
		return
	}

	transactionGroupDtos := make([]dto.TransactionGroupDto, 0, len(transactionGroups))
	for _, transactionGroup := range transactionGroups {
		transactionGroupDtos = append(transactionGroupDtos, *dto.MapTransactionGroupDomainToDto(transactionGroup))
	}
	c.JSON(http.StatusOK, transactionGroupDtos)
}

func (rH *RouterHandler) createGroup(c *gin.Context) {
	userUUIDParam := c.Param("uuid")
	userUUID, err := uuid.Parse(userUUIDParam)
	if err != nil {
		c.Error(err)
		return
	}

	createRequest := dto.CreateTransactionGroupRequest{}
	if err := c.BindJSON(&createRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newTransactionGroup, err := rH.ucHandler.CreateNewTransactionGroup(userUUID, createRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, dto.MapTransactionGroupDomainToDto(newTransactionGroup))
}
