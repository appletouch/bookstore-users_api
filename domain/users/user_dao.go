package users

// INTERACTS WITH THE PERSISTANCY LAYER

import (
	"github.com/appletouch/bookstore-users_api/datasources/mysql/users_db"
	"github.com/appletouch/bookstore-users_api/logger"
	"github.com/appletouch/bookstore-users_api/utils/errors"
	"strings"
)

const (
	deleteQuery  = "DELETE FROM users WHERE id=?"
	updateQuery  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	insertQuery  = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	selectQuery  = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	findByStatus = "SELECT first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	findByEmail  = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE email=?;"
)

//var (
//	userDB = make(map[int64]*User)
//)

//GET a user based on the user id
func (user *User) Get() *errors.RestErr {
	statement, err := users_db.Client.Prepare(selectQuery)
	if err != nil {
		logger.Error("error while trying to prepare get user statement", err)
		return errors.New(500)
	}
	defer statement.Close()

	// if you use the Query function you get back *rows.
	// you then need to close the connection. Somthing you don't need to do withQueryRow
	//EXAMPLE
	// results, _ statement.Query(user.Id)
	// defer results.Close()
	// You then need to interact with the results via de scan function.

	result := statement.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		//check if the error is a 404 not found
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.New(404)
		}
		logger.Error("error while scanning rows in get result", err)
		return errors.New(500)
	}
	return nil
}

//CREATE save the user to the persistency layer...
func (user *User) Save() *errors.RestErr {

	statement, err := users_db.Client.Prepare(insertQuery)
	if err != nil {
		logger.Error("error while trying to prepare get user statement", err)
		return errors.New(500)
	}
	//after use you need to close the connection. (defer is stacked )
	defer statement.Close()

	//execute prepared statement
	insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		//check if the error is related to a double registration and give a specific message to user.
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.New(400, "Email has already been registered.")
		}
		logger.Error("error while trying to insert user", err)
		return errors.New(500)
	}
	//ALTERNATIVE DIRECT ON DB WITHOUT PREPARE STATEMENT
	//insertResult2, err := users_db.Client.Exec(insertQuery,user.FirstName, user.LastName, user.Email, user.DateCreated )

	//Get last userid if succesful
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error while trying to get last inserted userid", err)
		return errors.New(400)
	}
	user.Id = userId

	return nil
}

//UPDATE
func (user *User) Update() *errors.RestErr {

	statement, err := users_db.Client.Prepare(updateQuery)
	if err != nil {
		logger.Error("error while trying to prepare update user statement", err)
		return errors.New(500)
	}
	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error while trying to execute update statement", err)
		return errors.New(500)
	}
	return nil

}

//DELETE
func (user *User) Delete() *errors.RestErr {
	statement, err := users_db.Client.Prepare(deleteQuery)
	if err != nil {
		logger.Error("error while trying to prepare delete user statement", err)
		return errors.New(500)
	}
	defer statement.Close()

	_, err = statement.Exec(user.Id)
	if err != nil {
		logger.Error("error while trying to execute delete user statement", err)
		return errors.New(500)
	}
	return nil

}

//SEARCH
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {

	//prepare a sql statement and open the connection.. don't forget to close it!!
	statement, err := users_db.Client.Prepare(findByStatus)
	if err != nil {
		logger.Error("error while trying to prepare find user statement", err)
		return nil, errors.New(500)
	}
	// only defere if you have a valid result.
	defer statement.Close()

	//Execute statement with te value from the user.
	foundRows, err := statement.Query(status)
	if err != nil {
		return nil, errors.New(404, "No users where found with this status")
	}

	// only defere if you have a valid result.
	defer foundRows.Close()

	//process the foundrows
	var foundUsers []User
	for foundRows.Next() {
		var user User
		err := foundRows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			return nil, errors.New(500)
		}
		foundUsers = append(foundUsers, user)
	}

	if len(foundUsers) == 0 {
		return nil, errors.New(404, " no users found with this status.")
	}
	return foundUsers, nil

}

//GET a user based by email and Password
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	statement, err := users_db.Client.Prepare(findByEmail)
	if err != nil {
		logger.Error("error while trying to prepare get user by email and password ", err)
		return errors.New(500)
	}
	defer statement.Close()

	result := statement.QueryRow(user.Email)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
		//check if the error is a 404 not found
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.New(404)
		}
		logger.Error("error while scanning rows in get result", err)
		return errors.New(500)
	}

	return nil
}
