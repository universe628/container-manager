package router

import (
	"container-manager/internal/handler"
	"container-manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func createFileRouter(router *gin.Engine, fileHandler *handler.FileHandler) {
	file := router.Group("/file")
	file.Use(middleware.JWTAuthMiddleware())
	{
		file.POST(("/upload"), fileHandler.UploadFile)
	}

}
