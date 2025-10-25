package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs each request in the format: [METHOD] URL - STATUS_CODE
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()

		log.Printf("[%s] %s - %d", method, path, status)
	}
}

// SimpleAuthMiddleware checks for X-API-Key == "12345" and rejects otherwise.
func SimpleAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key != "12345" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
