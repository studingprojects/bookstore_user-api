package app

import (
	"github.com/gin-gonic/gin"
	"github.com/studingprojects/bookstore_utils-go/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	initRoutes()
	logger.Info("about to start application...")
	router.Run(":8088")
}
