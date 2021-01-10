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

	// Parse form data
	if err := c.ShouldBind(&form); err != nil {
		api.respond(c.Writer, http.StatusBadRequest, errorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid form data",
		})
	}

	err := api.Controller.User().CreateUserFromSignup(form)
	if err != nil {
		api.error(c.Writer, err)
		return
	}

	api.respond(c.Writer, http.StatusCreated, &response{
		StatusCode: http.StatusAccepted,
		Message:    "Created new user",
	})
}
