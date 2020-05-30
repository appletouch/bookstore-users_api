package app

import "github.com/gin-gonic/gin"

// the router is only available in the app package
var (
	// by using the default option you are creating an Engine instance with the Logger and Recovery middleware already attached
	ginEngine = gin.Default()
)

//start application with urls mapped
func StartApplication() {
	mapUrls()
	ginEngine.Run(":3000")

}
