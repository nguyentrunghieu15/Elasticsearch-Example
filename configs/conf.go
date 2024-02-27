package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Elastic ElasticConf
}

func Init() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Default().Panic(err)
	}
	var conf Config
	err := envconfig.Init(&conf)
	if err != nil {
		log.Default().Panic(err)
	}
	return &conf
}
