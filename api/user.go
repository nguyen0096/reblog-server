package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (c *APIServer) InitUserAPI() {
	c.User = c.Root.Group("/user")

	c.User.GET("/", c.GetUser)
	c.User.POST("/", c.CreateUser)
}

func (api *APIServer) GetUser(c *gin.Context) {
	log.Println("GetUser hit!")
}

func (api *APIServer) CreateUser(c *gin.Context) {
	log.Println("CreateUser hit!")
}
