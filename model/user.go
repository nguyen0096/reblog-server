package model

import (
	"encoding/json"
	"io"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password,omitempty" form:"password" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
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
