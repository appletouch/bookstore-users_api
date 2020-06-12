package services

//SERVICES CONTAIN THE BUSINESS LOGIC
//THEY INTERACT WITH DATAPROVIDERS AND EXTERNAL APIS

import (
	"github.com/appletouch/bookstore-users_api/domain/users"
	"github.com/appletouch/bookstore-users_api/utils/cryptos"
	"github.com/appletouch/bookstore-users_api/utils/dates"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type userService struct {
}

//define a var that is a interfaces and contains a instance of userServices struct.
var (
	UserService userServiceInterface = &userService{}
)

// interactie is always on the interface
// ALWAYS START WITH DEFINING THE INTERFACE
type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	search(string) ([]users.User, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

// The creat user function contains the business logic that does the validation and saves the user.
func (us *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {

	//set the current time to the user.
	user.DateCreated = dates.GetNowDBDate()
	user.Status = "active" //default status of new users

	//Validate the user or return a error.
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.ValidatePassword(); err != nil {
		return nil, err
	}

	encryptedPWD := cryptos.HashAndSaltPassword(user.Password)
	user.Password = encryptedPWD

	// save the user or return a error
	if err := user.Save(); err != nil {
		return nil, err
	}
	//If validat was succesful and user is stored then return the user
	return &user, nil
}

// The Get user function retrieves a user based on the user_id(int) from the persistancy layer.
func (us *userService) GetUser(userid int64) (*users.User, *errors.RestErr) {

	//set the userid of the user to retieve
	result := &users.User{Id: userid}

	// call the method get on the user struct
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//The update users can be called via a put( replaces all user info) or a patch (replaces parts of user info).
func (us *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

	//get the user from the database
	current, err := UserService.GetUser(user.Id)
	if err != nil {
		return nil, err // if not found return a err not found
	}

	// if it is a patch and not a put then only a part is being updated
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	//Validate the the user to be updated before actual update
	if err := current.Validate(); err != nil {
		return nil, err
	}

	//call the datalayer to update
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil

}

// calls the delete method on the user object
func (us *userService) DeleteUser(userId int64) *errors.RestErr {

	//create a user and always interact with user object.
	var userToDelete = &users.User{Id: userId}
	userToDelete.Delete()
	return nil
}

func (us *userService) search(status string) ([]users.User, *errors.RestErr) {

	var user = &users.User{}
	return user.FindByStatus(status)
}

func (us *userService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {

	dao := &users.User{
		Email:    request.Email,
		Password: request.Password,
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	//check if account has been blocked.
	if dao.Status == "blocked" {
		return nil, errors.New(http.StatusUnauthorized, "Account has been blocked")
	}
	//check password
	err := bcrypt.CompareHashAndPassword([]byte(dao.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New(http.StatusNotFound, "Invalid credentials")
	}

	dao.Password = "OK"

	return dao, nil
}
