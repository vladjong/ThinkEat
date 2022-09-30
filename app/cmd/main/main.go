package main

import (
	"log"

	"github.com/vladjong/ThinkEat/internal/app"
	"github.com/vladjong/ThinkEat/internal/config"
	"github.com/vladjong/ThinkEat/pkg/logging"
)

func main() {
	log.Print("Config initializing")
	cfg := config.GetConfig()

	log.Print("Logger initializing")
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	app, err := app.NewApp(cfg, &logger)

	if err != nil {
		logger.Fatal(err)
	}

	log.Print("Running application")
	app.Run()
}
