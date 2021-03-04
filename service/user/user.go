package user

import (
	"fmt"
	"reblog-server/domain/model"
	"reblog-server/store"
	"reblog-server/utils"
	"reblog-server/utils/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
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

type Token struct {
}

func (c *userService) VerifyToken(bearerToken string) bool {
	// bearerToken := strings.Split(authorizationHeader, " ")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(config.App.Auth.JWTSecret), nil
	})

	if err != nil {
		utils.Error("Error: %s", err)
	}

	i, err := strconv.ParseInt(claims["exp"].(string), 10, 64)
	if err != nil {
		utils.Error("failed to parse unix string to int64. err: %s", err)
		return false
	}

	tm := time.Unix(i, 0)
	return time.Now().Before(tm)
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
	atClaims["exp"] = strconv.FormatInt(time.Now().Add(30*time.Second).Unix(), 10)
	utils.Info("Exp: %s", atClaims["exp"])
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
