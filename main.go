package main

import (
	"log"
	"os"
	"os/signal"
	"reblog-server/api"
	"reblog-server/app/controller"
	"reblog-server/config"
	"reblog-server/store/sqlstore"
	"syscall"
)

// func main() {
// srv := app.NewServer()

// db := infra.NewPostgresConnection(srv.Config())
// srv.SetDatabaseConnection(db)

// store := sqlstore.New(db)
// srv.SetStore(store)

// res, err := store.User().GetUserById(`1`)
// if err != nil {
// 	log.Fatalf("Failed to get user by id")
// }

// iter := interactor.New(store)
// srv.SetInteractor(iter)

// log.Printf("Result: %v", res)

// // result, err := srv.Database().Exec("INSERT INTO rb_core.user (username, first_name, last_name, address) VALUES ($1, $2, $3, $4)", "jmoiron", "Jason", "Moiron", "jmoiron@jmoiron.net")
// // log.Printf("error: %v", err)
// // log.Printf("result: %v", result)
// srv.Start()
// }

func main() {
	config.InitConfig()

	db := config.NewPostgresSQLConnection()

	store := sqlstore.New(db)

	u, err := store.User().Get(`1`)
	if err != nil {
		log.Printf("failed to get user. error: %v", err)
	}

	log.Printf("Result: %v", u)

	ctrl := controller.New()

	router := api.Init(ctrl)
	router.Run()
	defer router.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
