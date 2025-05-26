package router

import (
	"container-manager/internal/handler"

	"github.com/gin-gonic/gin"
)

func createAuthRouter(router *gin.Engine, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}
}
