package configs

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Elastic ElasticConf
}

func Init() (*Config, error) {
	var conf Config
	err := envconfig.Init(&conf)
	if err != nil {
		fmt.Printf("err=%s\n", err)
	}
	return &conf, err
}
