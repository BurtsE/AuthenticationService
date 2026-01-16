package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PanicHandlerMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Warnf("Recovered from panic: %v", r)
				c.AbortWithStatus(500)
			}
		}()
	}
}
