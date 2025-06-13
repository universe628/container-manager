package middleware

import (
	"container-manager/internal/schema"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(log Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		userID, ok := c.Request.Context().Value(schema.UserIDKey).(string)
		if !ok {
			userID = "anonymous"
		}
		requestID, ok := c.Request.Context().Value(schema.RequestIDKey).(string)
		if !ok {
			requestID = "unknown"
		}

		latency := time.Since(start)
		if len(c.Errors) > 0 {
			log.WithFields(map[string]any{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       c.Request.URL.Path,
				"ip":         c.ClientIP(),
				"request_id": requestID,
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
				"request_id": requestID,
				"user_id":    userID,
				"lantency":   latency,
				"errors":     "none",
			}).Info("Request completed")
		}
	}
}
