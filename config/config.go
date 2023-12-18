package config

import (
	"log"

	"github.com/caarlos0/env/v10"
)

type HttpConfig struct {
	Host string `env:"HTTP_HOST" envDefault:"localhost"`
	Port int    `env:"HTTP_PORT" envDefault:"8090"`
}

type MongoDB struct {
	URI string `env:"MONGODB_URI" envDefault:"mongodb://root:root@localhost:27017/openchat?authSource=admin"`
}

type Config struct {
	HttpConfig HttpConfig
	MongoDB    MongoDB
}

func NewConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
