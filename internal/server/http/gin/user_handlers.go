package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/server/http/gin/dto"
	"net/http"
)

func (rH *RouterHandler) getAllUsers(c *gin.Context) {
	users, err := rH.ucHandler.GetAllUsers()
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}

	resultUserDtoItems := make([]*dto.UserDto, 0, 10)
	for _, user := range users {
		resultUserDtoItems = append(resultUserDtoItems, dto.MapUserDomainToDto(user))
	}

	c.JSON(http.StatusOK, resultUserDtoItems)
}

func (rH *RouterHandler) getUserById(c *gin.Context) {
	userUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}

	user, err := rH.ucHandler.GetUserById(userUUID)
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}

	if user != nil {
		c.JSON(http.StatusOK, dto.MapUserDomainToDto(user))
		return
	}
	c.Status(http.StatusNotFound)
}

func (rH *RouterHandler) createUser(c *gin.Context) {

	body := dto.CreateUserRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := rH.ucHandler.CreateNewUser(*body.MapToUpdatableFields())

	if user != nil && err == nil {
		c.JSON(http.StatusCreated, dto.MapUserDomainToDto(user))
		return
	}

	if err != nil {
		c.Status(500)
		c.Error(err)
	}
}
