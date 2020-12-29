package api

import (
	"github.com/gin-gonic/gin"
)

func (c *APIServer) InitUserAPI() {
	c.User = c.Router.Group("/user")

	c.User.POST("/", c.CreateUser)
}

func (api *APIServer) CreateUser(c *gin.Context) {

}
