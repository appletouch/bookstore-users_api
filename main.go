package main

import "github.com/appletouch/bookstore-users_api/app"

func main() {
	//We don't start the application here, but we call a function to start the application.
	//This is done so the StartApplication can also be called by the testframework and we don't call the main function.
	app.StartApplication()
}
