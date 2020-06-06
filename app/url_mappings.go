package app

import (
	"github.com/appletouch/bookstore-users_api/controllers/heartbeat"
	"github.com/appletouch/bookstore-users_api/controllers/users"
)

// define every route here en de function that has to be called when request on the path.
func mapUrls() {

	//heartbeat
	ginEngine.GET("/heartbeat", heartbeat.HeartBeat)

	//users
	ginEngine.POST("/users", users.Create)
	ginEngine.GET("/users/:user_id", users.Get)
	ginEngine.PUT("/users/:user_id", users.Update)
	ginEngine.PATCH("/users/:user_id", users.Update)
	ginEngine.DELETE("/users/:user_id", users.Delete)
	//ginEngine.GET("/users/search/status", users.SearchUsers)

}
