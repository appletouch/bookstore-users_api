package users

// INTERACTS WITH THE PERSISTANCY LAYER

import (
	"github.com/appletouch/bookstore-users_api/datasources/mysql/users_db"
	"github.com/appletouch/bookstore-users_api/utils/dates"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"strings"
)

const (
	updateQuery = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	insertQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	selectquery = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

//var (
//	userDB = make(map[int64]*User)
//)

// Get a user based on the user id
func (user *User) Get() *errors.RestErr {
	statement, err := users_db.Client.Prepare(selectquery)
	if err != nil {
		return errors.New(500, err.Error())
	}
	defer statement.Close()

	// if you use the Query function you get back *rows.
	// you then need to close the connection. Somthing you don't need to do withQueryRow
	//EXAMPLE
	// results, _ statement.Query(user.Id)
	// defer results.Close()
	// You then need to interact with the results via de scan function.

	result := statement.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		//check if the error is a 404 not found
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.New(404)
		}
		return errors.New(500, err.Error())
	}
	return nil
}

// save the user to the persistency layer...
func (user *User) Save() *errors.RestErr {

	//set the current time to the user.
	user.DateCreated = dates.GetDateString()

	statement, err := users_db.Client.Prepare(insertQuery)
	if err != nil {
		return errors.New(500, err.Error())
	}
	//after use you need to close the connection. (defer is stacked )
	defer statement.Close()

	//execute prepated statement
	insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		//check if the error is related to a double registration and give a specific message to user.
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.New(400, "Email has already been registered.")
		}

		return errors.New(500, err.Error())
	}
	//ALTERNATIVE DIRECT ON DB WITHOUT PREPARE STATEMENT
	//insertResult2, err := users_db.Client.Exec(insertQuery,user.FirstName, user.LastName, user.Email, user.DateCreated )

	//Get last userid if succesful
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.New(400, err.Error())
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {

	statement, err := users_db.Client.Prepare(updateQuery)
	if err != nil {
		return errors.New(500)
	}
	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.New(500)
	}
	return nil

}
