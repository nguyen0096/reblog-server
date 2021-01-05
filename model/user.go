package model

type User struct {
	ID        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}
