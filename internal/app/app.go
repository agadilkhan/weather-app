// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/evrone/go-clean-template/pkg/mongodb"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/evrone/go-clean-template/config"
	v1 "github.com/evrone/go-clean-template/internal/controller/http/v1"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/internal/usecase/repo"
	"github.com/evrone/go-clean-template/internal/usecase/webapi"
	"github.com/evrone/go-clean-template/pkg/httpserver"
	"github.com/evrone/go-clean-template/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	mongo, err := mongodb.New(cfg.MongoDB.URI)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mongodb.New: %v", err))
	}

	weatherUseCase := usecase.New1(
		repo.New(mongo.Client.Database("service")),
		webapi.New(cfg.WebAPI.APIKey),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, weatherUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
