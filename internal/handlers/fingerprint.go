package handlers

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
)

func getClientFingerprint(c *gin.Context) string {
	ip := c.ClientIP()
	ua := c.GetHeader("User-agent")
	lang := c.GetHeader("Accept-Language")
	encoding := c.GetHeader("Accept-Encoding")

	raw := fmt.Sprintf("%s|%s|%s|%s", ip, ua, lang, encoding)

	hash := sha256.Sum256([]byte(raw))
	return fmt.Sprintf("%x", hash)
}
