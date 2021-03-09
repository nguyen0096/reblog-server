package main

import (
	"encoding/json"
	"fmt"
	"reblog-server/utils"
	"reblog-server/utils/config"
	"reflect"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID        null.Int    `json:"-"`
	UserID    null.String `json:"user_id"`
	Password  null.String `json:"first_name"`
	LastName  null.String `json:"last_name"`
	Email     null.String `json:"email"`
	CreatedAt null.Time   `json:"created_at"`
	UpdatedAt null.Time   `json:"updated_at"`
}

func TestMain(m *testing.M) {
	var err error

	config.InitConfig()

	// pgConn, err := database.NewPostgresSqlxConn()
	// if err != nil {
	// 	utils.Error("failed to start db connection. err: %s", err)
	// }

	userJSON := "{\"user_id\": \"nguyen\", \"password\": \"123\"}"

	user := &User{}
	json.Unmarshal([]byte(userJSON), user)

	// Make sure password and userid are available

	// Hash password
	var hashedPw []byte
	if hashedPw, err = bcrypt.GenerateFromPassword([]byte(user.Password.String), config.App.Controller.HashCost); err != nil {
		utils.Error("failed to hash password. err: %s", err)
	}

	user.Password.String = string(hashedPw)

	query := GenerateQuery("user", user)
	fmt.Println(query)

	// queryString := `INSERT INTO rb_core.user (id, username, password) VALUES ($1, $2, $3);`
	// pgConn.Exec(queryString)
}

func GenerateQuery(tableName string, obj interface{}) string {
	// var query string
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {

		fmt.Printf("Field: %v - Value: %v - Tag: %s\n", t.Field(i).Name, v.Field(i), t.Field(i).Tag.Get("json"))
	}

	fmt.Printf("Type: %v\n", t)
	castedTime, ok := v.FieldByName("UpdatedAt").Interface().(time.Time)
	if ok {
		fmt.Printf("Time: %s\n", castedTime)
	}
	fmt.Printf("Value: %v\n", v)

	var nguyen string
	fmt.Printf("%v\n", utils.IsZeroValue(nguyen))
	return ""
}
