package services

//SERVICES CONTAIN THE BUSINESS LOGIC

import (
	"github.com/appletouch/bookstore-users_api/domain/users"
	"github.com/appletouch/bookstore-users_api/utils/errors"
)

// The creat user function contains the business logic that does the validation and saves the user.
func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	//Validat the user or return a error.
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// save the user or return a error
	if err := user.Save(); err != nil {
		return nil, err
	}
	//If validat was succesful and user is stored then return the user
	return &user, nil
}

// The Get user function retrieves a user based on the user_id(int) from the persistancy layer.
func GetUser(userid int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userid}

	// call the method get on the user struct
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}