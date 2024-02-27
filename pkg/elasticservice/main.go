package elasticservice

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticService struct {
	OrderIndex *OrderIndex
}

func Init(elasticClient *elasticsearch.Client) *ElasticService {
	return &ElasticService{
		OrderIndex: InitOrderIndex(elasticClient),
	}
}
