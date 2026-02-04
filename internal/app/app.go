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
	db2 "github.com/paincake00/inventory-management-service/internal/infrastructure/db"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	config Config
	logger *zap.SugaredLogger
	router *gin.Engine
	db     *gorm.DB
}

func NewApp(config Config, logger *zap.SugaredLogger) *App {
	app := &App{}

	app.config = config
	app.logger = logger

	app.router = app.InitRouter()

	db, err := db2.ConnectDB(config.db.address, config.db.maxOpenConn, config.db.maxIdleConn, config.db.maxConnLifetime)
	if err != nil {
		app.logger.Fatal(err)
	}
	app.db = db

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

	// closing db connection
	sqlDB, err := app.db.DB()
	if err != nil {
		app.logger.Errorf("Database connection error: %v", err)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		app.logger.Errorf("Database close error: %v", err)
		return err
	}

	return nil
}
