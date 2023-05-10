package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
	"net/http"
)

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

	transactionDtos := make([]*dto.TransactionDto, 0, len(transactions))
	for _, transaction := range transactions {
		transactionDtos = append(transactionDtos, dto.MapTransactionDomainToDto(transaction))
	}
	c.JSON(http.StatusOK, transactions)
}
