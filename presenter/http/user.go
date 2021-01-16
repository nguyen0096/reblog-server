package http

import (
	"net/http"

	"reblog-server/model"
	"reblog-server/utils"

	"github.com/gin-gonic/gin"
)

func (c *APIServer) initUserAPI() {
	c.User = c.Root.Group("/user")

	c.User.GET("", c.getUser)
	c.User.POST("", c.createUser)
}

func (api *APIServer) getUser(c *gin.Context) {
	utils.Info("GetUser hit!")
}

func (api *APIServer) createUser(c *gin.Context) {
	var user *model.User
	var err error

	if user, err = model.UserFromJson(c.Request.Body); err != nil {
		api.error(c.Writer, http.StatusBadRequest, err)
		return
	}

	if user == nil {
		api.error(c.Writer, http.StatusBadRequest, &APIError{Message: "Can't parse JSON data"})
		return
	}

	utils.Info("%v", user)

	err = api.Controller.User().CreateUserFromSignUp(user)
	if err != nil {
		api.error(c.Writer, http.StatusBadRequest, err)
		return
	}

	api.respond(c.Writer, http.StatusCreated, nil)
}
