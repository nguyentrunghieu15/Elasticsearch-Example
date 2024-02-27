package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/nguyentrunghieu15/Elasticsearch-Example.git/configs"
	elasticservice "github.com/nguyentrunghieu15/Elasticsearch-Example.git/pkg/elasticservice"
)

func main() {
	conf := configs.Init()
	routes := gin.Default()
	v, err := os.ReadFile(conf.Elastic.CertCAPath) //read the content of file
	if err != nil {
		fmt.Println(err)
		return
	}
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{conf.Elastic.Addresses},
		Username:  conf.Elastic.Username,
		Password:  conf.Elastic.Password,
		CACert:    v,
	})

	if err != nil {
		log.Panic("Error to create client Elastic")
	} else {
		log.Println("Created client Elastic")
	}

	esService := elasticservice.Init(esClient)

	InitRoutes(routes, esService)
	routes.Run()
}
