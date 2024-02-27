package elasticservice

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

const INDEX_NAME = "orders"

type OrderIndex struct {
	es *elasticsearch.Client
}

type OrderItemModel struct {
	ProductId int     `json:"product_id"`
	Amount    float64 `json:"amount"`
	Quantity  int     `json:"quantity"`
}

type OrderModel struct {
	PurchaseAt  time.Time        `json:"purchased_at"`
	Lines       []OrderItemModel `json:"lines"`
	TotalAmount float64          `json:"total_amount"`
	SalesMan    struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"salesman"`
	SalesChanel string `json:"sales_chanel"`
	Status      string `json:"status"`
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

func (idx *OrderIndex) UpdateOrderById(orderId string, data OrderModel) (*esapi.Response, error) {
	query, _ := json.Marshal(data)
	res, err := idx.es.Update(INDEX_NAME, orderId, bytes.NewReader(query))
	return res, err
}

func (idx *OrderIndex) CreateOrder(orderId string, data OrderModel) (*esapi.Response, error) {
	query, _ := json.Marshal(data)
	res, err := idx.es.Create(INDEX_NAME, orderId, bytes.NewReader(query))
	return res, err
}

func (idx *OrderIndex) DeleteOrder(orderId string) (*esapi.Response, error) {
	res, err := idx.es.Delete(INDEX_NAME, orderId)
	return res, err
}
