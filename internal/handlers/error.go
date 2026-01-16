package handlers

import (
	"AuthenticationService/internal/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createErrorResponse(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidEmail):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrWeakPassword):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrUserAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	return
}
