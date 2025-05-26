package main

import (
	"container-manager/internal/handler"
	"container-manager/internal/infra/auth"
	"container-manager/internal/infra/config"
	"container-manager/internal/infra/database"
	"container-manager/internal/infra/logger"
	postgresRepo "container-manager/internal/repository/postgres"
	"container-manager/internal/router"
	authservice "container-manager/internal/service/auth"
)

func main() {

	logger := logger.NewLogrusLogger()

	db := database.NewDatabase(config.GetConfig().Database)

	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Disconnect(dbConnection)

	repo := postgresRepo.NewPostgresUserRepository(dbConnection)
	authService := authservice.NewAuthService(repo, auth.NewJwt())
	authHandler := handler.NewAuthHandler(authService)

	testHandler := handler.NewTestHandler()

	router := router.CreateRootRouter(logger, authHandler, testHandler)

	router.Run(":8080")

}
