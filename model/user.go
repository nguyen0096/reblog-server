package model

import (
	"encoding/json"
	"io"
)

type User struct {
	ID        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username" form:"username" binding:"required"`
	Password  string `json:"password,omitempty" db:"password" form:"password" binding:"required"`
	Email     string `json:"email" form:"email"`
	FirstName string `json:"first_name" form:"firstname"`
	LastName  string `json:"last_name" form:"lastname"`
}

func UserFromJson(data io.Reader) (*User, error) {
	var user *User
	err := json.NewDecoder(data).Decode(&user)
	return user, err
}

func UserFromFormData(data io.Reader) (*User, error) {
	return nil, nil
}
