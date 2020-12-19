package api

import (
	"log"
	"net/http"
	"reblog-server/dependency"
)

type Context struct {
	Server  dependency.IServer
	Request *http.Request
	Writer  http.ResponseWriter
}

type WrapHandler struct {
	ctx     *Context
	Handler func(ctx *Context) error
}

func (c *WrapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ctx.Request = r
	c.ctx.Writer = w

	err := c.Handler(c.ctx)
	if err != nil {
		log.Printf("Failed to create new user")
	}
}

func (a *API) NewWrapHandler(fn func(ctx *Context) error) *WrapHandler {

	ctx := &Context{
		Server: a.Server,
	}

	return &WrapHandler{
		ctx:     ctx,
		Handler: fn,
	}
}

func (c *API) initDummy() {
	log.Println("Init dummy api!")
	c.Routes.Dummy.Handle("/user", c.NewWrapHandler(createNewUser))
}

func createNewUser(ctx *Context) error {
	log.Println("createNewUser")
	return nil
}
