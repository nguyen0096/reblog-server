package api

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Handler struct {
	ctx     *Context
	Handler func(ctx *Context) error
}

func (c *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ctx.Request = r
	c.ctx.Writer = w

	err := c.Handler(c.ctx)
	if err != nil {
		log.Printf("Failed to create new user")
	}
}
