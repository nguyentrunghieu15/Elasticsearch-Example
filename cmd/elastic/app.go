package main

import (
	"fmt"

	"github.com/nguyentrunghieu15/Elasticsearch-Example.git/configs"
)

func main() {
	conf, err := configs.Init()
	if err != nil {
		return
	}
	fmt.Println(conf.Elastic.API_KEY)
}
