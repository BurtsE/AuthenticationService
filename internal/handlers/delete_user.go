package handlers

import (
	"AuthenticationService/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (u *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		u.log.WithError(err).WithField("id", idStr).Error("Invalid user id")
		createErrorResponse(c, err)
		return
	}

	err = u.service.DeleteUser(c.Request.Context(), domain.UserID(id))
	if err != nil {
		u.log.WithError(err).WithField("id", id).Error("Error deleting user")
		createErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
}
