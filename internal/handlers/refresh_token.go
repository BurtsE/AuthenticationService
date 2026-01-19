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

	accessToken, refreshToken, err := h.service.RefreshToken(c.Request.Context(), refreshToken, getClientFingerprint(c))
	if err != nil {
		h.log.WithError(err).WithField("token", refreshToken).Error("Error refreshing token")
		createErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
