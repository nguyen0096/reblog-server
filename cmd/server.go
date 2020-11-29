package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"

	"reblog-server/config"
)

type Server struct {
	Router *mux.Router

	Config *config.Config

	Server *http.Server
}

func main() {
	s := NewServer()

	s.Start()
}

func NewServer() *Server {
	rootRouter := mux.NewRouter()

	appConfig := config.NewConfig()

	return &Server{
		Router: rootRouter,
		Config: appConfig,
	}
}

func (s *Server) Start() {
	log.Println("Starting server...")
	var handler http.Handler = s.Router

	s.Server = &http.Server{
		Handler: handler,
	}

	listener, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatalf("failed to listen on port 6969. error: %v", err)
	}

	log.Printf("Server is listening on %v \n", listener.Addr().String())

	err = s.Server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve http server. error: %v", err)
	}
}
