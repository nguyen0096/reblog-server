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

	store := sqlstore.New()
	store.SetSqlxConn(database.NewPostgresSqlxConn())
	store.SetGormConn(database.NewPostgresGormConn())
	store.Migrate()

	ctrl := service.New(store)

	router := http.Init(ctrl)
	router.Run()
	defer router.Close()

	// grpcServer := grpc.NewGRPCServer(ctrl.Todo())
	// go grpcServer.Run()
	// defer grpcServer.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
