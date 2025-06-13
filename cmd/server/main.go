package main

import (
	"container-manager/internal/handler"
	"container-manager/internal/infra/auth"
	"container-manager/internal/infra/config"
	"container-manager/internal/infra/database"
	"container-manager/internal/infra/localstorage"
	"container-manager/internal/infra/logger"
	"container-manager/internal/repository/file-repo"
	postgresRepo "container-manager/internal/repository/postgres"
	"container-manager/internal/router"
	authservice "container-manager/internal/service/auth"
	fileservice "container-manager/internal/service/file"
	validationservice "container-manager/internal/service/validation"
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

	fileManager := localstorage.NewfileManager()
	fileRepo := filerepo.NewLocalStorageRepo(fileManager)
	fileService := fileservice.NewFileService(fileRepo)
	validationService := validationservice.NewValidationService()
	fileHandler := handler.NewFileHandler(fileService, validationService)

	router := router.CreateRootRouter(logger, authHandler, testHandler, fileHandler)

	router.Run(":8080")

}
