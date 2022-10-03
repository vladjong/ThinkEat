package main

import (
	"context"
	"log"

	"github.com/vladjong/ThinkEat/internal/adapters/db/mongoDB"
	"github.com/vladjong/ThinkEat/internal/config"
	"github.com/vladjong/ThinkEat/internal/domain/entity"
	"github.com/vladjong/ThinkEat/pkg/logging"
	"github.com/vladjong/ThinkEat/pkg/mongodb"
)

func main() {
	log.Print("Config initializing")
	cfg := config.GetConfig()

	log.Print("Logger initializing")
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	mongoDBClient, err := mongodb.NewClient(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.Auth_db)
	if err != nil {
		panic(err)
	}
	storage := mongoDB.NewStorage(mongoDBClient, cfg.MongoDB.Collection, &logger)

	item1 := entity.Item{
		Name:     "vds Wok 12",
		Describe: "ddds",
		Category: []string{"lapsha"},
		Price:    123,
		Photo:    "dsdfsdf",
		Weight:   300,
	}
	id, err := storage.Create(context.Background(), &item1)
	if err != nil {
		panic(err)
	}
	logger.Info(id)
	logger.Info("Create succes")

	m, err := storage.GetName(context.Background(), "Wok")
	if err != nil {
		panic(err)
	}
	logger.Info(m[1].Name)
	logger.Info("Find succses")

	// m.Name = "Dima"
	// m, err := storage.FindID(context.Background(), "000000000000000000000000")
	// if err != nil {
	// 	panic(err)
	// }
	// logger.Info(m.Name)
	// logger.Info("Find succses")

	// m.Name = "Dima"

	// err = storage.Update(context.Background(), m)
	// if err != nil {
	// 	panic(err)
	// }
	// logger.Info("Update succses")

	// itemss, err := storage.FindAll(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// logger.Info("Find all succses")
	// fmt.Println(itemss)

	// err = storage.Delete(context.Background(), id)
	// if err != nil {
	// 	panic(err)
	// }
	// logger.Info("Delete succses")

	// app, err := app.NewApp(cfg, &logger)

	// if err != nil {
	// 	logger.Fatal(err)
	// }

	// log.Print("Running application")
	// app.Run()
}
