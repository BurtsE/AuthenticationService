package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware returns a Gin middleware that logs HTTP requests using logrus.
func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		entry := logger.WithFields(logrus.Fields{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"status":      status,
			"duration_ms": duration.Seconds() * 1000,
			"client_ip":   c.ClientIP(),
			"user_agent":  c.Request.UserAgent(),
		})

		if requestId := c.GetString("request_id"); requestId != "" {
			entry = entry.WithField("request_id", requestId)
		}

		switch {
		case status >= 200 && status < 300:
			entry.Info("HTTP request")
		case status >= 300 && status < 400:
			entry.Warn("HTTP request")
		case status >= 400 && status < 500:
			entry.Error("HTTP request")
		default:
			entry.Info("HTTP request")
		}
	}
}
