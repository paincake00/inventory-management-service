package main

import (
	"github.com/joho/godotenv"
	"github.com/paincake00/inventory-management-service/internal/app"
	"github.com/paincake00/inventory-management-service/internal/utils/logs"
)

func main() {
	logger := logs.NewLogger()
	defer func() {
		_ = logger.Sync()
	}()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file", err)
	}

	config := app.InitConfig()
	application := app.NewApp(config, logger)

	if err = application.Run(); err != nil {
		logger.Fatal(err)
	}
}
