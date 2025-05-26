package config

import (
	"container-manager/internal/utils"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	config     *Config
	onceConfig sync.Once
)

func getEnvConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtExpiresDurationHourStr := os.Getenv("JWT_EXPIRES_DURATION_HOUR")
	jwtExpiresDurationHour, err := utils.StringToInt(jwtExpiresDurationHourStr)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Pg: PostgresDataBase{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Name:     os.Getenv("POSTGRES_DB"),
		},
		Jwt: JwtConfig{
			SecretKey:           os.Getenv("JWT_SECRET_KEY"),
			ExpiresDurationHour: jwtExpiresDurationHour,
		},
		Database: os.Getenv("DATABASE"),
		Salt:     os.Getenv("SALT"),
	}
}

func GetConfig() *Config {
	onceConfig.Do(func() {
		config = getEnvConfig()
	})
	return config
}
