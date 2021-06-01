package api

import (
	todoapi "reblog-server/lib/api/todo"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router gin.IRouter) {

	router.Group("/api/v1")

	v1 := router.Group("/api/v1")

	todoAPI := todoapi.New()
	todoRouter := v1.Group("/todo")
	todoRouter.GET(
		"",
		todoAPI.GetHandler(),
	)
}
