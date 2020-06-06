package users

//MODEL CONTAIN THE DOMAIN DEFINITION AND THE METHODS

import (
	"github.com/appletouch/bookstore-users_api/utils/emails"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"net/http"
	"strings"
)

// user model and json mappings (Password is never shown in json response)
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"on:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//adding the validate method to the user struct
func (user *User) Validate() *errors.RestErr {
	//clean up username
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	// check email
	emailToLower := strings.TrimSpace(strings.ToLower(user.Email))
	if err := emails.ValidateFormat(emailToLower); err != nil {
		return errors.New(http.StatusBadRequest, "Invalid email-address")
	}
	//check password
	user.Password = strings.TrimSpace(user.Password)
	if len(user.Password) < 5 {
		return errors.New(http.StatusBadRequest, "Password needs to be 6 character long. ")
	}

	return nil
}
