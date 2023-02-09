package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	Listen        struct {
		Type       string `env:"LISTEN_TYPE" env-default:"false"`
		BindIP     string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port       string `env:"PORT" env-default:"10000"`
		SocketFile string `env:"SOCKET_FILE" env-default:"./app.sock"`
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-default:"admin"`
			Password string `env:"ADMIN_PWD" env-default:"admin"`
		}
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Info("gather config")
	})

	instance = &Config{}

	if err := cleanenv.ReadEnv(instance); err != nil {
		log.WithFields(log.Fields{
			"package":  "config",
			"function": "GetConfig",
		}).Fatal(err)
	}

	return instance
}
