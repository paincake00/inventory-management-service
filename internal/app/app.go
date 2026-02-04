package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	config Config
	logger *zap.SugaredLogger
	router *gin.Engine
}

func NewApp(config Config, logger *zap.SugaredLogger) *App {
	app := &App{}

	app.config = config
	app.logger = logger

	app.router = app.InitRouter()

	return app
}

func (app *App) Run() error {
	errCh := make(chan error, 10)

	serv := &http.Server{
		Addr:    app.config.addr,
		Handler: app.router,
	}

	shutdown := make(chan error, 1)
	// graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		app.logger.Infof("Got signal %s, exiting gracefully...", s)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdown <- serv.Shutdown(ctx)
	}()

	// starting server
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			errCh <- fmt.Errorf("http server listen err: %w", err)
		}
	}()

	app.logger.Infow("Server started", "addr", serv.Addr)

	select {
	case err := <-errCh:
		app.logger.Errorf("Server exited with error %v", err)
		return err
	case err := <-shutdown:
		if err != nil {
			app.logger.Errorf("Server shutdown with error: %v", err)
			return err
		}
	}

	app.logger.Infof("Service is shutting down...")

	// TODO closing gorm

	return nil
}
