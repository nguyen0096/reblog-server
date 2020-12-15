package main

import (
	"log"
	"reblog-server/app"
	"reblog-server/infra/database"
)

func main() {
	srv := app.NewServer()

	db := database.NewPostgresConnection(srv.Config())
	srv.SetDatabaseConnection(db)

	result, err := srv.Database().Exec("INSERT INTO rb_core.user (username, first_name, last_name, address) VALUES ($1, $2, $3, $4)", "jmoiron", "Jason", "Moiron", "jmoiron@jmoiron.net")

	log.Printf("error: %v", err)
	log.Printf("result: %v", result)
}
