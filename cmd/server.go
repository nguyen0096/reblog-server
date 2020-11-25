package main

import (
	"fmt"

	"reblog-auth/config"
	"reblog-auth/dependency"
	"reblog-auth/infrastructure/persistence/database"
)

func main() {
	config.InitConfig()

	pgConn, _ := dependency.NewPostgreSQLConnection()

	userRepo := database.NewUserRepository(pgConn)

	user, err := userRepo.GetUserByID("1")
	if err != nil {
		panic(err)
	}

	fmt.Printf("App started! %v", user)
}
