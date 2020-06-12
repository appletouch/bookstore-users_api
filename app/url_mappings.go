package app

import (
	"fmt"
	"github.com/appletouch/bookstore-users_api/controllers/heartbeat"
	"github.com/appletouch/bookstore-users_api/controllers/users"
	"github.com/gin-gonic/gin"
	"strings"
)

// define every route here en de function that has to be called when request on the path.
func mapUrls() {

	//heartbeat
	ginEngine.GET("/heartbeat", heartbeat.HeartBeat)

	//users
	ginEngine.POST("/users", users.Create)
	ginEngine.GET("/users/user/:user_id", users.Get)
	ginEngine.PUT("/users/user/:user_id", users.Update)
	ginEngine.PATCH("/users/user/:user_id", users.Update)
	ginEngine.DELETE("/users/user/:user_id", users.Delete)

	//Solutions to prevent a Gin specic routing problem where wildcards conflict with other routes
	//default route
	ginEngine.GET("/users/search", users.SearchUsers)
	//wilde card route vs actions routes
	ginEngine.GET("/users/search/:user_id", func(ctx *gin.Context) {
		myUri := strings.HasPrefix(ctx.Request.RequestURI, "/users/search/active")
		fmt.Println(ctx.Request.RequestURI)
		if myUri {
			users.GetActiveUsers(ctx)
			return
		}
		users.SearchUsers(ctx)
	})

	//find by email and password
	ginEngine.POST("/users/user/login", users.Login)

}
