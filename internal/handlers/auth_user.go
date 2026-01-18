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

	accessToken, refreshToken, err := h.service.AuthorizeUser(c.Request.Context(), request)
	if err != nil {
		h.log.WithError(err).WithField("user", request).Error("error authorizing user")
		createErrorResponse(c, err)
		return
	}

	c.SetCookie(refreshTokenCookieName, refreshToken, cookieMaxAge, "/", "", false, true)

	c.JSON(http.StatusOK, dto.TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
