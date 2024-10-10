package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mariobgr/pack-shipment-exercise/internal/application/config"
	"github.com/mariobgr/pack-shipment-exercise/internal/application/service"
	handler "github.com/mariobgr/pack-shipment-exercise/internal/infra/http"
	"github.com/mariobgr/pack-shipment-exercise/internal/infra/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// init loading config
	doneCh := make(chan bool)
	config.LoadConfigContinuously(doneCh)

	// init the custom logger
	appLogger := logger.NewLogger()

	// init the calculator service
	calculatorService := service.NewCalculatorService(config.NewPacksGetter())

	shipmentHandler := handler.NewPacksShipmentHandler(calculatorService, appLogger)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: shipmentHandler.Routes(),
	}

	go func() {
		appLogger.Info("server started", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			appLogger.Info("server stopped", err)
		}
	}()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-sigCh
		cancel()
	}()

	<-ctx.Done()
	doneCh <- true

}
