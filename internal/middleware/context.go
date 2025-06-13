package middleware

import (
	errs "container-manager/internal/errors"
	"container-manager/internal/infra/config"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeout := config.GetConfig().TimeoutSecond
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(timeout)*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		switch {
		case ctx.Err() == context.DeadlineExceeded:
			c.Error(errs.ErrDeadlineExceeded)
			c.Abort()
		case ctx.Err() == context.Canceled:
			c.Error(errs.ErrCancelled)
			c.Abort()
		default:
		}
	}
}
