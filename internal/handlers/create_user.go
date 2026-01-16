package handlers

import (
	"AuthenticationService/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserHandler) CreateUser(c *gin.Context) {
	var request dto.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		createErrorResponse(c, err)
		return
	}

	err := u.service.CreateUser(c.Request.Context(), request)
	if err != nil {
		u.log.WithError(err).WithField("user", request).Error("Error creating user")
		createErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
