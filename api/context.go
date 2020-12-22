package api

import (
	"log"
	"net/http"
	"reblog-server/dependency"
)

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type Context struct {
	Server     dependency.IServer
	Request    *http.Request
	Writer     http.ResponseWriter
	TestString string
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

type MiddlwareFunc func(*WrapHandler) *WrapHandler

func firstCtxMiddleware(h *WrapHandler) *WrapHandler {
	return &WrapHandler{
		ctx: h.ctx,
		Handler: func(ctx *Context) error {

			log.Printf("Server: %v", ctx.Server)
			ctx.TestString = "nguyen1"
			log.Println("Middleware with context. First string set!")
			// log.Printf("Try getting database config: %v", ctx.Server.Config().GetDatabase())

			err := h.Handler(ctx)
			return err
		},
	}
}

func secondCtxMiddleware(h *WrapHandler) *WrapHandler {

	return &WrapHandler{
		ctx: h.ctx,
		Handler: func(ctx *Context) error {

			ctx.TestString = "nguyen2"
			log.Println("Middleware with context. Second string set!")

			err := h.Handler(ctx)
			return err
		},
	}
}

// func firstMiddleware(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("This is first middleware")
// 		handler.ServeHTTP(w, r) // call ServeHTTP on the original handler

// 	})
// }

// func secondMiddleware(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("This is second middleware")
// 		handler.ServeHTTP(w, r) // call ServeHTTP on the original handler
// 	})
// }

func (c *API) initDummy() {
	log.Println("Init dummy api!")
	c.Routes.Dummy.Handle("/user", firstCtxMiddleware(secondCtxMiddleware(c.NewWrapHandler(createNewUser))))
	// c.Routes.Dummy.Use(firstMiddleware)
	// c.Routes.Dummy.Use(secondMiddleware)
}

func createNewUser(ctx *Context) error {
	log.Println("createNewUser")
	log.Printf("Get TestString: %v", ctx.TestString)
	return nil
}
