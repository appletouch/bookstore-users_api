package users

//CONTOLLERS ARE THE INTERACTION POINT WITH THE REQUESTER

import (
	"fmt"
	"github.com/appletouch/bookstore-users_api/domain/users"
	"github.com/appletouch/bookstore-users_api/services"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {

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
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		fmt.Println(saveError)
		ctx.JSON(saveError.Status, saveError)
	}

	// if user was succesfully created pass the user in json to as response.
	ctx.JSON(http.StatusCreated, result)
}

func GetUser(ctx *gin.Context) {

	//Get the user from te request and parse to string
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 10)

	//if user id is invalid (not a int) in the request return a error
	if userErr != nil {
		err := errors.New(http.StatusBadRequest, "UserID is not valid")
		ctx.JSON(http.StatusBadRequest, err)
	}

	// call the servce to GET the user
	user, getError := services.GetUser(userId)
	if getError != nil {
		fmt.Println(getError)
		ctx.JSON(getError.Status, getError)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {

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

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, result)

}
