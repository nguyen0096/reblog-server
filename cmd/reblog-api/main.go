package main

import (
	"context"
	"os"
	"os/signal"
	libapi "reblog-server/lib/api"

	"github.com/rs/zerolog"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	logger := zerolog.New(os.Stdout)

	api, err := libapi.Default("127.0.0.1:14070")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create api")
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		cancel()
	}()

	err = api.Run(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start api")
	}
}
