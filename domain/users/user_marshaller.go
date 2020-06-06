package users

//Missing all personal information
type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// is only missing password
type InternalUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"on:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (User *User) Marshal(isPubli bool) {

}
