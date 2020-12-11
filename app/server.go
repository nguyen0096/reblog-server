package app

import (
	"log"
	"net"
	"net/http"
	"time"

	api "reblog-server/apiv2"
	"reblog-server/config"
	"reblog-server/domain/interactor"
)

// Server ...
type Server struct {
	Router     *http.ServeMux
	Config     *config.Config
	Interactor interactor.Interactor
}

// NewServer ...
func NewServer() *Server {
	appConfig := config.NewConfig()
	rootRouter := http.NewServeMux()
	iter := interactor.NewInteractor()

	return &Server{
		Config:     appConfig,
		Router:     rootRouter,
		Interactor: iter,
	}
}

// Start ...
func (s *Server) Start() {
	log.Println("Starting server...")

	api.Init(s.Router)

	// s.Router.Run()
	srv := &http.Server{
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

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve http server. error: %v", err)
	}
}
