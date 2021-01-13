package controller

import (
	"log"
	"reblog-server/config"
	"reblog-server/model"
	"reblog-server/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	base *baseController
}

func newUserController(base *baseController) *userController {
	return &userController{
		base: base,
	}
}

func (c *userController) CreateToken(user *model.User) (string, error) {
	var err error

	log.Println(user.Password)

	queryUser, err := c.base.store.User().GetByUsername(user.Username)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(queryUser.Password), []byte(user.Password)); err != nil {
		utils.Error("Wrong password", err)
		return "", err
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.App.Auth.JWTSecret))
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func (c *userController) CreateUserFromSignUp(user *model.User) error {
	var hashedPw []byte
	var id uuid.UUID
	var err error

	utils.Info("%v", user)

	if hashedPw, err = bcrypt.GenerateFromPassword([]byte(user.Password), config.App.Controller.HashCost); err != nil {
		return err
	}

	if id, err = uuid.NewUUID(); err != nil {
		return err
	}

	user.Password = string(hashedPw)
	user.ID = id.String()

	err = c.base.store.User().Create(user)
	return err
}
