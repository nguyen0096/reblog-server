package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"reblog-server/middleware"
	"reblog-server/service"
	"reblog-server/utils"
	"reblog-server/utils/config"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Server http.Server

	Mux *mux.Router
	RouterGroups
	Service service.App
}

type RouterGroups struct {
	Root *mux.Router

	Auth *mux.Router
	User *mux.Router
	Demo *mux.Router
	Todo *mux.Router
}

// Implement interface

func Init(sv service.App) *APIServer {
	api := &APIServer{
		Service: sv,
	}

	// Config for all routes
	api.Mux = mux.NewRouter()
	api.Mux.NotFoundHandler = api.notFoundHandler()

	api.Mux.Use(middleware.Logger)
	api.Mux.Use(middleware.Authenticator(sv.User()))

	// API
	api.Root = api.Mux
	api.initDemoAPI()
	api.initUserAPI()
	api.initAuthAPI()
	api.initTodoAPI()

	return api
}

func (c *APIServer) Run() {
	addr := fmt.Sprintf("%s:%d", "localhost", config.App.API.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: c.Mux,
	}

	utils.Info("Listening on http://%s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to listen on %s. error: %s\n", addr, err)
	}
}

func (c *APIServer) Close() {
	utils.Info("Closing API Server...")
}

func (c *APIServer) notFoundHandler() http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {
		utils.Error("Cannot find handler for %s", r.RequestURI)
	}

	return http.HandlerFunc(f)
}

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Error error `json:"error"`
}

// TODO: figure out best way to send error to client and log it down
func (c *APIServer) error(w http.ResponseWriter, code int, e error) {
	var body []byte
	var err error

	res := errorResponse{
		Error: e,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if body, err = json.Marshal(res); err != nil {
		errorBody := "{\"status\": 500, \"message\": \"Something happened wrong during generating response\"}"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorBody))
		return
	}

	w.WriteHeader(code)
	w.Write(body)
}

func (c *APIServer) respond(w http.ResponseWriter, code int, data interface{}) {
	var body []byte
	var err error

	res := response{
		Data: data,
	}

	if body, err = json.Marshal(res); err != nil {
		errorBody := "{\"status\": 500, \"message\": \"Something happened wrong during generating response\"}"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorBody))
		return
	}

	w.WriteHeader(code)
	w.Write(body)
}

func (c *APIServer) debug(format string, args ...interface{}) {
	// TODO implement debug handler for API
}

type APIError struct {
	Message string
}

func (c *APIError) Error() string {
	return c.Message
}
