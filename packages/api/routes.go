package main

import (
	"github.com/gin-gonic/gin"
	storage "tribeverse.dev/editor/api/packages/api/storage"
)

func initRoutes(r *gin.Engine) {
	recordGroup := r.Group("/storage")
	{
		recordGroup.GET("/:key", storage.GenerateUploadURL)
		recordGroup.POST("/upload-token", storage.GenerateUploadURL)
	}
}
