package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" env-default:"bis"`
	IsDebug     bool   `env:"IS_DEBUG" env-default:"true"`
	Listen      struct {
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"8081"`
	}
	Redis struct {
		Address  string `env:"REDIS_ADDRESS" env-default:"redis:6379"`
		Password string `env:"REDIS_PASSWORD" env-default:"my-password"`
	}

	SomeClient struct {
		Host string `env:"SOME_CLIENT_HOST" env-default:"https://someclient.com"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			log.Fatal(err)
		}

	})
	return instance
}
