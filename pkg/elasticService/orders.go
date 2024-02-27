package elasticservice

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

const INDEX_NAME = "orders"

type OrderIndex struct {
	es *elasticsearch.Client
}

func InitOrderIndex(es *elasticsearch.Client) *OrderIndex {
	return &OrderIndex{es: es}
}

func (idx *OrderIndex) FindAll() (*esapi.Response, error) {
	query := `{"query":{"match_all":{}}}`
	res, err := idx.es.Search(idx.es.Search.WithIndex(INDEX_NAME), idx.es.Search.WithBody(strings.NewReader(query)))
	return res, err
}

func (idx *OrderIndex) SearchIds(ids []string) (*esapi.Response, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"ids": map[string]interface{}{
				"values": ids,
			},
		},
	}
	jsonQuery, _ := json.Marshal(query)
	res, err := idx.es.Search(idx.es.Search.WithIndex(INDEX_NAME), idx.es.Search.WithBody(bytes.NewReader(jsonQuery)))
	return res, err
}

func (idx *OrderIndex) StatsOrder() (*esapi.Response, error) {
	query := `{"size": 0,"aggs": {"stats_order": {"stats": {"field": "total_amount"}}}}`
	res, err := idx.es.Search(idx.es.Search.WithIndex(INDEX_NAME), idx.es.Search.WithBody(strings.NewReader(query)))
	return res, err
}
