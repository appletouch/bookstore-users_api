package users

//CONTOLLERS ARE THE INTERACTION POINT WITH THE REQUESTER
//CONTROLLERS PROCESS THE REQUEST AND PRODUCE THE RESPONSE

import (
	"fmt"
	"github.com/appletouch/bookstore-users_api/domain/users"
	"github.com/appletouch/bookstore-users_api/services"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {

	//!!!!!!This whole block of code can be replaced with the ShouldBindJSON function!!!!
	//bytes, err := ioutil.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	handle the errors read errors
	//}
	//err= json.Unmarshal(bytes,&user)
	//if err != nil {
	//	handle json errors
	//	return
	//}

	var user users.User
	//using ginEngine context to parse and bind json request to user struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errors.New(http.StatusBadRequest))
		return
	}
	// call the servce to create the user
	result, saveError := services.UserService.CreateUser(user)
	if saveError != nil {
		fmt.Println(saveError)
		ctx.JSON(saveError.Status, saveError)
	}

	// if user was succesfully created pass the user in json to as response.
	ctx.JSON(http.StatusCreated, result)
}

func Get(ctx *gin.Context) {

	//Get the user from te request and parse to string
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 10)

	//if user id is invalid (not a int) in the request return a error
	if userErr != nil {
		err := errors.New(http.StatusBadRequest, "UserID is not valid")
		ctx.JSON(http.StatusBadRequest, err)
	}

	// call the servce to GET the user
	user, getError := services.UserService.GetUser(userId)
	if getError != nil {
		fmt.Println(getError)
		ctx.JSON(getError.Status, getError)
		return
	}
	ctx.JSON(http.StatusOK, user.Marshal(ctx.GetHeader("X-Public") == "true"))
}

func Update(ctx *gin.Context) {

	//Get the user from te request and parse to string
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 10)

	//if user id is invalid (not a int) in the request return a error
	if userErr != nil {
		err := errors.New(http.StatusBadRequest, "UserID is not valid")
		ctx.JSON(http.StatusBadRequest, err)
	}

	var user users.User
	//using ginEngine context to parse and bind json request to user struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errors.New(http.StatusBadRequest))
		return
	}
	user.Id = userId

	isPartial := ctx.Request.Method == http.MethodPatch

	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, result)

}

func Delete(ctx *gin.Context) {
	//Get the user from te request and parse to string
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 10)
	if userErr != nil {
		ctx.JSON(http.StatusBadRequest, userErr)
	}
	if errDelete := services.UserService.DeleteUser(userId); errDelete != nil {
		ctx.JSON(http.StatusBadRequest, errDelete)
	}
	ctx.JSON(http.StatusOK, map[string]string{"title": "Record deleted", "status": "200", "detail": fmt.Sprintf("User %d has been deleted", userId)})

}

func SearchUsers(ctx *gin.Context) {
	fmt.Println("STARTNG SEARCH")

}

func GetActiveUsers(ctx *gin.Context) {
	fmt.Println("Getting active users")

}
