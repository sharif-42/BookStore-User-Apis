package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sharif-42/BookStore-User-Apis/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()

	logger.Info("Application is about to start.......")

	router.Run(":8080")
}
