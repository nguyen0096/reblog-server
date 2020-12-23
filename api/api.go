package api

import (
	"log"
	"reblog-server/dependency"

	"github.com/gorilla/mux"
)

type IApi interface {
	Init()
	GetRoutes() *Routes
	NewWrapHandler(fn func(ctx *Context) error) *Handler
}

// =================== Implementation ===================

type api struct {
	Server dependency.IServer
	Routes *Routes
}

func NewAPI(srv dependency.IServer, r *mux.Router) IApi {
	a := api{
		Server: srv,
		Routes: &Routes{
			Root: &Router{Mux: r},
		},
	}

	return a
}

func (a api) Init() {
	a.InitUser()
}

func (a api) NewWrapHandler(fn func(ctx *Context) error) *Handler {
	ctx := &Context{
		Server: a.Server,
	}
	return &Handler{
		ctx:     ctx,
		Handler: fn,
	}
}

func (a api) GetRoutes() *Routes {
	return a.Routes
}

func (a *api) InitUser() {
	a.Routes.User = &Router{Mux: a.Routes.Root.Mux.PathPrefix("/user").Subrouter()}

	a.Routes.User.Mux.Handle("", a.NewWrapHandler(createNewUser)).Methods("GET")
}

func createNewUser(ctx *Context) error {
	log.Println("createNewUser")
	log.Printf("Get TestString: %v", ctx.TestString)
	return nil
}
