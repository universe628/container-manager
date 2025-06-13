package middleware

import (
	errs "container-manager/internal/errors"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

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
			switch {
			//400 Bad Request
			case errors.Is(err, errs.ErrUnknownToken),
				errors.Is(err, http.ErrNotMultipart),
				errors.Is(err, http.ErrMissingBoundary),
				errors.Is(err, multipart.ErrMessageTooLarge),
				errors.Is(err, errs.ErrNoFileUploaded),
				errors.Is(err, errs.ErrFileTooLarge),
				errors.Is(err, errs.ErrInvalidFilePath),
				strings.Contains(err.Error(), "unexpected EOF"):

				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			//401 Unauthorized
			case errors.Is(err, errs.ErrTokenInvalid),
				errors.Is(err, errs.ErrInvalidCredentials),
				errors.Is(err, errs.ErrUserNotFound):

				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			//409 Conflict
			case errors.Is(err, errs.ErrUserNameTaken):

				c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

			//499 Client Closed Request
			case errors.Is(err, errs.ErrCancelled):

				c.JSON(http.StatusGatewayTimeout, gin.H{"error": err.Error()})

			//500 Internal Server Error
			case errors.Is(err, errs.ErrPanic):

				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

			//504 Gateway Timeout
			case errors.Is(err, errs.ErrDeadlineExceeded):

				c.JSON(http.StatusGatewayTimeout, gin.H{"error": err.Error()})

			default:

				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			c.Abort()
		}
	}
}
