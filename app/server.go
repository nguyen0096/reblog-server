package app

import (
	"log"
	"net"
	"net/http"
	"time"

	"reblog-server/api"
	"reblog-server/config"
	"reblog-server/dependency"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// server ...
type server struct {
	Router *mux.Router
	config *config.Config
	DB     *sqlx.DB

	Store      dependency.IStore
	Interactor dependency.IInteractor
}

// NewServer initializes instances of dependencies
func NewServer() dependency.IServer {
	conf := config.NewConfig()
	router := mux.NewRouter()

	srv := &server{
		config: conf,
		Router: router,
	}

	api.Init(srv, router)

	return srv
}

// GETTERS

func (s *server) Config() dependency.IConfig {
	return s.config
}

func (s *server) Database() *sqlx.DB {
	return s.DB
}

// SETTERS
func (s *server) SetDatabaseConnection(db *sqlx.DB) {
	s.DB = db
}

func (s *server) SetStore(store dependency.IStore) {
	s.Store = store
}

func (s *server) SetInteractor(iter dependency.IInteractor) {
	s.Interactor = iter
}

// Start ...
func (s *server) Start() {
	log.Println("Starting server...")

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
