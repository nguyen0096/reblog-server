package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type API interface {
	GetHandler() gin.HandlerFunc
}

func New() API {
	return &api{}
}

type api struct {
}

func (a *api) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	}
}
