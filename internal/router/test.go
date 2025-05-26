package router

import (
	"container-manager/internal/handler"
	"container-manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func createTestRouter(router *gin.Engine, testHandler *handler.TestHandler) {
	auth := router.Group("/test")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/ping", testHandler.Pong)
	}

}
