package auth

import (
	"container-manager/internal/infra/config"
	"container-manager/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct{}

func NewJwt() *Jwt {
	return &Jwt{}
}

func (j *Jwt) GenerateToken(userId int32) (string, error) {
	TokenExpireDuration := time.Hour * time.Duration((config.GetConfig().Jwt.ExpiresDurationHour))
	claims := jwt.RegisteredClaims{
		Subject:   utils.IntToString(int(userId)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		Issuer:    "sirius",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetConfig().Jwt.SecretKey))
}
