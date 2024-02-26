package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/nguyentrunghieu15/Elasticsearch-Example.git/configs"
)

func main() {
	conf := configs.Init()

	elasticClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		APIKey: conf.Elastic.API_KEY,
	})
}
