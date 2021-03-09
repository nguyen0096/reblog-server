package user

import (
	"fmt"
	"reblog-server/domain/model"
	"reblog-server/store"
	"reblog-server/utils"
	"reblog-server/utils/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUserFromSignUp(user *model.User) error
	CreateToken(user *model.User) (string, error)
	VerifyToken(token string) bool
}

type userService struct {
	store store.UserStore
}

func NewUserService(store store.UserStore) UserService {
	return &userService{
		store: store,
	}
}

func (c *userService) VerifyToken(bearerToken string) bool {
	// bearerToken := strings.Split(authorizationHeader, " ")
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		utils.Error("Error: %s", err)
	}

	var exp time.Duration
	mapstructure.Decode(token.Claims, &exp)

	utils.Info("Exp: %v", exp)
	return false
}

func (c *userService) CreateToken(user *model.User) (string, error) {
	var err error

	queryUser, err := c.store.GetByUsername(user.Username)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(queryUser.Password), []byte(user.Password)); err != nil {
		utils.Error("Wrong password", err)
		return "", err
	}

	atClaims := jwt.MapClaims{}
	// atClaims["authorized"] = true
	atClaims["user_id"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.App.Auth.JWTSecret))
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func (c *userService) CreateUserFromSignUp(user *model.User) error {
	var hashedPw []byte
	var id uuid.UUID
	var err error

	if hashedPw, err = bcrypt.GenerateFromPassword([]byte(user.Password), config.App.Controller.HashCost); err != nil {
		return err
	}

	if id, err = uuid.NewUUID(); err != nil {
		return err
	}

	user.Password = string(hashedPw)
	user.ID = id.String()

	err = c.store.Create(user)
	return err
}
