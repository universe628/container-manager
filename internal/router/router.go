package router

import (
	"container-manager/internal/handler"
	"container-manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRootRouter(logger middleware.Logger, authHandler *handler.AuthHandler, testHandler *handler.TestHandler, fileHandler *handler.FileHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.RequestIdMiddleware(), middleware.LoggingMiddleware(logger), middleware.ErrorHandlingMiddleware(), middleware.ContextMiddleware())

	createAuthRouter(router, authHandler)
	createTestRouter(router, testHandler)
	createFileRouter(router, fileHandler)

	return router
}
