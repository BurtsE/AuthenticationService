package handlers

import (
	"AuthenticationService/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserHandler) CreateUser(c *gin.Context) {
	var request dto.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrParsingRequest.Error(),
		})
		return
	}

	user := request.ToEntity()

	err := u.service.CreateUser(c.Request.Context(), &user)
	if err != nil {
		u.log.WithError(err).WithField("user", user).Error("Error creating user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": ErrInternalServer.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
