package main

import (
	"os"
	"os/signal"
	"reblog-server/api"
	"reblog-server/app/controller"
	"reblog-server/config"
	"reblog-server/store/sqlstore"
	"syscall"
)

func main() {
	config.InitConfig()

	db := config.NewPostgresSQLConnection()

	store := sqlstore.New(db)

	ctrl := controller.New(store)

	router := api.Init(ctrl)
	router.Run()
	defer router.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
