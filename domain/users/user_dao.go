package users

// INTERACTS WITH THE PERSISTANCY LAYER

import (
	"fmt"
	"github.com/appletouch/bookstore-users_api/datasources/mysql/users_db"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"net/http"
)

const (
	insertQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	userDB = make(map[int64]*User)
)

// Get a user based on the user id
func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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

	//set the current time to the user.
	//user.DateCreated = dates.GetDateString()

	statement, err := users_db.Client.Prepare(insertQuery)
	if err != nil {
		return errors.New(500, err.Error())
	}
	//after use you need to close the connection. (defer is stacked )
	defer statement.Close()

	//execute prepated statement
	insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.New(500, "Error while trying to execute statement")
	}
	//ALTERNATIVE DIRECT ON DB WITHOUT PREPARE STATEMENT
	//insertResult2, err := users_db.Client.Exec(insertQuery,user.FirstName, user.LastName, user.Email, user.DateCreated )

	//Get last userid if succesful
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.New(500, err.Error())
	}
	user.Id = userId

	return nil
}
