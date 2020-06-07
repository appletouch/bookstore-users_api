package app

import (
	"github.com/appletouch/bookstore-users_api/logger"
	"github.com/gin-gonic/gin"
)

// the router is only available in the app package
var (
	// by using the default option you are creating an Engine instance with the Logger and Recovery middleware already attached
	ginEngine = gin.Default()
)

//start application with urls mapped
func StartApplication() {
	mapUrls()

	logger.Info("About to start application...")
	ginEngine.Run(":3000")

}
