package handlers

import (
	"AuthenticationService/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) AuthorizeUser(c *gin.Context) {
	var request dto.AuthorizeUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.WithError(err).Error("invalid request body")
		createErrorResponse(c, ErrInvalidRequestBody)
		return
	}

	token, err := h.service.AuthorizeUser(c.Request.Context(), request)
	if err != nil {
		h.log.WithError(err).Error("invalid request body")
		createErrorResponse(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
