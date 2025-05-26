package middleware

import (
	errs "container-manager/internal/errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var msg string
				switch x := r.(type) {
				case string:
					msg = x
				case error:
					msg = x.Error()
				default:
					msg = fmt.Sprintf("%v", x)
				}
				panic := fmt.Errorf("%w: %s", errs.ErrPanic, msg)

				c.Error(panic)
				c.Abort()
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			switch err {
			//4XX
			case errs.ErrUnknownToken:
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			case errs.ErrTokenInvalid, errs.ErrInvalidCredentials, errs.ErrUserNotFound:
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			case errs.ErrUserNameTaken:
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

			case errs.ErrPanic:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			c.Abort()
		}
	}
}
