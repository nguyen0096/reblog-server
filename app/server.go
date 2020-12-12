package app

import (
	"log"
	"net"
	"net/http"
	"time"

	api "reblog-server/apiv2"
	"reblog-server/config"
	"reblog-server/domain/interactor"

	"github.com/gorilla/mux"
)

// ServerIface ...
type ServerIface interface {
	Start()
}

// server ...
type server struct {
	Router     *mux.Router
	Config     *config.Config
	Interactor interactor.Interactor
}

// NewServer initializes instances of dependencies
func NewServer() ServerIface {
	conf := config.NewConfig()
	router := mux.NewRouter()
	iter := interactor.NewInteractor()

	return &server{
		Config:     conf,
		Router:     router,
		Interactor: iter,
	}
}

// Start ...
func (s *server) Start() {
	log.Println("Starting server...")

	api.Init(s.Router)

	srv := &http.Server{
		Handler:      AddContext(s.Router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

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
