package middleware

import (
	errs "container-manager/internal/errors"
	"container-manager/internal/infra/config"
	"container-manager/internal/schema"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	secretKey := config.GetConfig().Jwt.SecretKey
	keyBytes := []byte(secretKey)

	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return keyBytes, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return &jwt.RegisteredClaims{
			Issuer:    claims.Issuer,
			Subject:   claims.Subject,
			ExpiresAt: claims.ExpiresAt,
		}, nil
	}
	return nil, jwt.ErrTokenSignatureInvalid
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)
		if authHeader == "" || len(parts) != 2 || parts[0] != "Bearer" {
			c.Error(errs.ErrUnknownToken)
			c.Abort()
			return
		}

		claims, err := ParseToken(parts[1])
		if err != nil {
			c.Error(errs.ErrTokenInvalid)
			c.Abort()
			return
		}

		uId := claims.Subject

		ctx := context.WithValue(c.Request.Context(), schema.UserIDKey, uId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
