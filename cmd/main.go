package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/AxulReich/kitchen/cmd/app"
	"github.com/AxulReich/kitchen/internal/config"
	"github.com/AxulReich/kitchen/internal/pkg/logger"
)

func main() {
	ctx := context.Background()
	logger.Error(ctx, "Starting the service...")

	cfg, err := config.FromEnv()
	if err != nil {
		logger.Fatal(ctx, err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	shutdown := make(chan struct{}, 1)

	application, err := app.NewApplication(ctx, cfg, shutdown)
	if err != nil {
		logger.Fatal(ctx, err)
	}

	if err = application.Run(ctx); err != nil {
		logger.Fatal(ctx, err)
	}

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			logger.Errorf(ctx, "Got SIGINT...")
		case syscall.SIGTERM:
			logger.Errorf(ctx, "Got SIGTERM...")
		}
	case <-shutdown:
		logger.Errorf(ctx, "Got an error...")
	}

	if err = application.Close(); err != nil {
		logger.Errorf(ctx, "app closed with err: %s", err.Error())
	}
}
