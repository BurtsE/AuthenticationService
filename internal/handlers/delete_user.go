package handlers

import (
	"AuthenticationService/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (u *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		u.log.WithError(err).WithField("id", idStr).Error("Invalid user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrParsingRequest.Error()})
		return
	}

	err = u.service.DeleteUser(c.Request.Context(), model.UserID(id))
	if err != nil {
		u.log.WithError(err).WithField("id", id).Error("Error deleting user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
}
