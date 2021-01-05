package api

import (
	"log"
	"net/http"

	"reblog-server/dto"

	"github.com/gin-gonic/gin"
)

func (c *APIServer) initUserAPI() {
	c.User = c.Root.Group("/user")

	c.User.GET("", c.getUser)
	c.User.POST("", c.createUser)
}

func (api *APIServer) getUser(c *gin.Context) {
	log.Println("GetUser hit!")
}

func (api *APIServer) createUser(c *gin.Context) {
	var form dto.LoginForm
	if err := c.ShouldBind(&form); err != nil {
		api.respond(c.Writer, http.StatusBadRequest, "Invalid form data")
	}

	// if err := api.Controller.User().CreateUser(form); err != nil {
	// 	c.JSON(http.StatusConflict, )
	// }
}
