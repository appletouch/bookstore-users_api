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
	ginEngine.POST("/users", users.CreateUser)
	ginEngine.GET("/users/:user_id", users.GetUser)
	//ginEngine.GET("/users/search", controllers.SearchUser)

}
