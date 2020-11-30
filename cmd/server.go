package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"reblog-server/api"
	"reblog-server/config"
)

// Server ...
type Server struct {
	Router *mux.Router

	Config *config.Config

	Server *http.Server
}

func main() {
	s := NewServer()

	s.Start()
}

// NewServer ...
func NewServer() *Server {
	rootRouter := mux.NewRouter()

	appConfig := config.NewConfig()

	return &Server{
		Router: rootRouter,
		Config: appConfig,
	}
}

// Start ...
func (s *Server) Start() {
	log.Println("Starting server...")

	api.Init(s.Router)

	s.Server = &http.Server{
		Handler: s.Router,
		// Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// log.Fatal(s.Server.ListenAndServe())

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen on port 8080. error: %v", err)
	}

	log.Printf("Server is listening on %v \n", listener.Addr().String())

	err = s.Server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve http server. error: %v", err)
	}
}
