package models

// define the model for users (subject to change)
type User struct {
	ID       int    `json:id`
	Username string `json:username`
	Password string `json:password`
	Email    string `json:email`
}
