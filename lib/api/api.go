package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Router func(gin.IRouter)

type Worker func(context.Context)

type API interface {
	Run(context.Context) error
	AddRouter(Router)
	AddWorker(string, Worker)
	AddNoRoute([]gin.HandlerFunc)
}

type api struct {
	addr            string
	logger          zerolog.Logger
	routers         []Router
	workers         []Worker
	noRouteHandlers []gin.HandlerFunc
	waitGroup       sync.WaitGroup
}

func NewAPI(addr string, logger zerolog.Logger) API {
	return &api{
		addr:   addr,
		logger: logger,
	}
}

func (a *api) Run(ctx context.Context) error {
	handler := a.registerRoutes()
	a.runWorkers(ctx)

	// Start HTTP server
	srv := http.Server{
		Addr:    a.addr,
		Handler: handler,
	}

	// Shutdown routine
	go func() {
		<-ctx.Done()
		shutdownCtx, shutdownCancel := context.WithCancel(context.Background())
		defer shutdownCancel()

		if err := srv.Shutdown(shutdownCtx); err != nil && err != http.ErrServerClosed {
			a.logger.Err(err).Msg("server forced to shutdown")
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.logger.Err(err).Msg("server failed to start")

		return err
	}

	a.waitGroup.Wait()

	return nil
}

func (a *api) AddRouter(router Router) {
	a.routers = append(a.routers, router)
}

func (a *api) AddWorker(key string, worker Worker) {

}

func (a *api) AddNoRoute(handlers []gin.HandlerFunc) {

}

func (a *api) registerRoutes() http.Handler {
	ginRouter := gin.New()
	ginRouter.Use(gin.Logger())

	for _, router := range a.routers {
		router(ginRouter)
	}

	if len(a.noRouteHandlers) > 0 {
		ginRouter.NoRoute(a.noRouteHandlers...)
	}

	return ginRouter
}

func (a *api) runWorkers(ctx context.Context) {
	for key := range a.workers {
		a.waitGroup.Add(1)
		f := a.workers[key]

		go StartGoroutine(
			fmt.Sprintf("worker.%v", key),
			func() {
				f(ctx)
			},
			&a.waitGroup,
			a.logger,
		)
	}
}

func StartGoroutine(
	key string,
	f func(),
	wg *sync.WaitGroup,
	logger zerolog.Logger,
) {

}
