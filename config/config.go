package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Port     int `env:"PORT" envDefault:"8080"`
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DbConnectionString  string `env:"DB_CONNECTION_STRING"`
	DbConnectionTimeout int    `env:"DB_CONNECTION_TIMEOUT"`
}

var Cfg Config

func Get() Config {

	if (Cfg == Config{}) {
		if err := env.Parse(&Cfg); err != nil {
			panic(err)
		}
	}

	return Cfg
}
