package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`
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
