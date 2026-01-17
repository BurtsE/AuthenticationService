package handlers

import (
	"AuthenticationService/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.log.WithError(err).WithField("id", idStr).Error("Invalid user id")
		createErrorResponse(c, ErrInvalidRequestBody)
		return
	}

	err = h.service.DeleteUser(c.Request.Context(), domain.UserID(id))
	if err != nil {
		h.log.WithError(err).WithField("id", id).Error("Error deleting user")
		createErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
}
