package handlers

import (
	"AuthenticationService/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) CreateUser(c *gin.Context) {
	var request dto.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.WithError(err).Error("invalid request body")
		createErrorResponse(c, ErrInvalidRequestBody)
		return
	}

	err := h.service.CreateUser(c.Request.Context(), request)
	if err != nil {
		h.log.WithError(err).WithField("user", request).Error("Error creating user")
		createErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
