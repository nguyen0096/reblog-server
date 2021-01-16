package main

import (
	"os"
	"os/signal"
	"reblog-server/config"
	"reblog-server/controller"
	"reblog-server/presenter/http"
	"reblog-server/store/sqlstore"
	"syscall"
)

func main() {
	config.InitConfig()

	db := config.NewPostgresSQLConnection()

	store := sqlstore.New(db)

	ctrl := controller.New(store)

	router := http.Init(ctrl)
	router.Run()
	defer router.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
