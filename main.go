package main

import (
	"os"
	"os/signal"
	"reblog-server/api/http"
	"reblog-server/service"
	"reblog-server/store/sqlstore"
	"reblog-server/utils/config"
	"reblog-server/utils/database"
	"syscall"
)

func main() {
	config.InitConfig()

	db := database.InitPostgres()

	store := sqlstore.New(db)

	ctrl := service.New(store)

	router := http.Init(ctrl)
	router.Run()
	defer router.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
