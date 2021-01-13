package api

import (
	"log"
	"net/http"

	"reblog-server/model"

	"github.com/gin-gonic/gin"
)

func (c *APIServer) initAuthAPI() {
	c.Auth = c.Root.Group("/auth")

	c.Auth.POST("/login", c.handleLogin)
}

func (api *APIServer) handleLogin(c *gin.Context) {
	var u model.User
	var err error

	if err = c.ShouldBindJSON(&u); err != nil {
		log.Printf("failed to parse json. err: %v", err)
		api.error(c.Writer, http.StatusBadRequest, err)
		return
	}

	token, err := api.Controller.User().CreateToken(&u)
	if err != nil {
		api.error(c.Writer, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, token)
	return
}
