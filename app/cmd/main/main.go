package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/ThinkEat/config"
	"github.com/vladjong/ThinkEat/internal/app"
)

func main() {
	log.Print("Config initializing")
	cfg := config.GetConfig()
	logrus.Info("env variables initializing")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	logrus.Info("running service")
	service, err := app.NewService(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := service.Run(); err != nil {
		logrus.Fatal(err)
	}
}
