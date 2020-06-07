package users

import "encoding/json"

//Missing all personal information
type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// is only missing password
type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"on:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (User *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          User.Id,
			DateCreated: User.DateCreated,
			Status:      User.Status,
		}
	}
	//private user has no password
	userJson, _ := json.Marshal(User)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
