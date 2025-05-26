package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggingMiddleware(log Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		userID, exists := c.Get("userID")
		if !exists {
			userID = "anonymous"
		}

		latency := time.Since(start)
		if len(c.Errors) > 0 {
			log.WithFields(map[string]any{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       c.Request.URL.Path,
				"ip":         c.ClientIP(),
				"request_id": uuid.New().String(),
				"user_id":    userID,
				"lantency":   latency,
				"errors":     c.Errors[0].Err.Error(),
			}).Error("Request failed")
		} else {
			log.WithFields(map[string]any{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       c.Request.URL.Path,
				"ip":         c.ClientIP(),
				"request_id": uuid.New().String(),
				"user_id":    userID,
				"lantency":   latency,
				"errors":     "none",
			}).Info("Request completed")
		}
	}
}
