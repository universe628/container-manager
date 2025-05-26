package router

import (
	"container-manager/internal/handler"
	"container-manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRootRouter(logger middleware.Logger, authHandler *handler.AuthHandler, testHandler *handler.TestHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.LoggingMiddleware(logger), middleware.ErrorHandlingMiddleware())

	createAuthRouter(router, authHandler)
	createTestRouter(router, testHandler)

	// api := router.Group("/api")
	// {
	// 	api.GET("/registe", func(c *gin.Context) {

	// 		c.JSON(200, gin.H{
	// 			"message": "Register endpoint",
	// 		})
	// 	})
	// }

	return router
}
