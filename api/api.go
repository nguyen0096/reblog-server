package api

import (
	"fmt"
	"log"
	"net/http"
	"reblog-server/app"
	"reblog-server/app/middleware"
	"reblog-server/config"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Server http.Server

	Router *gin.Engine
	RouterGroups

	Controller app.Controller
}

type RouterGroups struct {
	Root *gin.RouterGroup

	User *gin.RouterGroup
}

func Init(ctrl app.Controller) *APIServer {
	api := &APIServer{
		Router:     gin.New(),
		Controller: ctrl,
	}
	api.RouterGroups.Root = api.Router.Group("/api")

	api.RouterGroups.Root.Use(middleware.Cors())

	api.InitUserAPI()

	return api
}

func (c *APIServer) Run() {
	log.Println("Starting API Server...")
	port := fmt.Sprintf(":%d", config.App.API.Port)

	srv := &http.Server{
		Addr:    port,
		Handler: c.Router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to listen on port %s. error: %s\n", port, err)
	}
}

func (c *APIServer) Close() {
	log.Println("Closing API Server...")
}
