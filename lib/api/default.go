package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Default(
	addr string,
) (API, error) {
	logger := zerolog.New(os.Stdout)

	api := NewAPI(addr, logger)
	api.AddRouter(func(router gin.IRouter) {

		RegisterRoutes(router)
	})

	return api, nil
}
