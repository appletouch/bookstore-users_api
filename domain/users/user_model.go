package users

//MODEL CONTAIN THE DOMAIN DEFINITION AND THE METHODS

import (
	"github.com/appletouch/bookstore-users_api/utils/emails"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"net/http"
	"strings"
)

// user model and json mappings
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"on:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//adding the validate method to the user struct
func (user *User) Validate() *errors.RestErr {
	emailToLower := strings.TrimSpace(strings.ToLower(user.Email))
	if err := emails.ValidateFormat(emailToLower); err != nil {
		return errors.New(http.StatusBadRequest, "Invalid email-address")
	}
	return nil
}
