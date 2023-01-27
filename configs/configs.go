package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug bool `envconfig:"debug" default:"false"`
	Port  int  `envconfig:"PORT" default:"8080"`
}

func GetDefault() *Config {
	var c Config
	err := envconfig.Process("fiberg", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &c
}
