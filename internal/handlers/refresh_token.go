package handlers

import (
	"AuthenticationService/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie(refreshTokenCookieName)
	if err != nil {
		h.log.WithError(err).Error("Refresh token cookie not found")
		createErrorResponse(c, ErrInvalidRequestBody)
		return
	}
	var request dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.WithError(err).Error("Invalid refresh token request")
		createErrorResponse(c, ErrInvalidRequestBody)
		return
	}

	accessToken, refreshToken, err := h.service.RefreshToken(c.Request.Context(), request)
	if err != nil {
		h.log.WithError(err).WithField("fingerprint", request).Error("Error refreshing token")
		createErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, dto.TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
