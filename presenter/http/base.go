package http

import (
	"fmt"
	"log"
	"net/http"
	"reblog-server/config"
	"reblog-server/controller"
	"reblog-server/middleware"
	"reblog-server/utils"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Server http.Server

	Router *gin.Engine
	RouterGroups

	Controller controller.IController
}

type RouterGroups struct {
	Root *gin.RouterGroup

	Auth *gin.RouterGroup
	User *gin.RouterGroup
}

// Implement interface

func Init(ctrl controller.IController) *APIServer {
	api := &APIServer{
		Router:     gin.New(),
		Controller: ctrl,
	}
	api.RouterGroups.Root = api.Router.Group("/api")

	api.RouterGroups.Root.Use(middleware.Cors())

	api.initUserAPI()
	api.initAuthAPI()

	return api
}

func (c *APIServer) Run() {
	utils.Info("Starting API Server...")
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
	utils.Info("Closing API Server...")
}
