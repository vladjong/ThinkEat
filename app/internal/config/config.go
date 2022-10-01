package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	Listen        struct {
		Type       string `env:"LISTEN_TYPE" env-default:"port"`
		BindIP     string `env:"IS_DEV" env-default:"0.0.0.0"`
		Port       string `env:"PORT" env-default:"1234"`
		SocketFile string `env:"SOCKET_FILE" env-default:"app.sock"`
	}
	AppConfig struct {
		LogLevel  string `env:"LOG_LEVEL" env-default:"trace"`
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-default:"admin"`
			Password string `env:"ADMIN_PWD" env-default:"admin"`
		}
	}
	MongoDB struct {
		Host       string `env:"HOST" env-default:"localhost"`
		Port       string `env:"PORT" env-default:"27017"`
		Database   string `env:"DATABASE" env-default:"user-service"`
		Auth_db    string `env:"AUTH_DB" env-default:"admin"`
		Username   string `env:"USERNAME" env-default:""`
		Password   string `env:"AUTH_DB" env-default:""`
		Collection string `env:"AUTH_DB" env-default:"users"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("Read application configuration")
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Think Eat - Monolith system"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
