package configs

import (
	"github.com/Valgard/godotenv"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Elastic ElasticConf
}

func Init() *Config {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}
	var conf Config
	err := envconfig.Init(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
