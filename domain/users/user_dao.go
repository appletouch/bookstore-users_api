package users

// INTERACTS WITH THE PERSISTANCY LAYER

import (
	"fmt"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"net/http"
)

var (
	userDB = make(map[int64]*User)
)

// Get a user based on the user id
func (user *User) Get() *errors.RestErr {
	if result := userDB[user.Id]; result != nil {

		// as the user is a pointer we can modify the user
		user.Id = result.Id
		user.FirstName = result.FirstName
		user.LastName = result.LastName
		user.Email = result.Email
		user.DateCreated = result.DateCreated
		return nil
	} else {
		return errors.New(http.StatusBadRequest, fmt.Sprintf("User %d not found", user.Id))
	}
}

// save the user to the persistency layer...
func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.New(http.StatusConflict, fmt.Sprintf("%s has already been registered and is not available", user.Email))
		}
		// in this case we already have this userid
		return errors.New(http.StatusConflict, fmt.Sprintf("The user with id %d already exsists", user.Id))
	}
	userDB[user.Id] = user

	return nil
}
