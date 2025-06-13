package middleware

import (
	"container-manager/internal/schema"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), schema.RequestIDKey, uuid.New().String()))

		c.Next()
	}
}
