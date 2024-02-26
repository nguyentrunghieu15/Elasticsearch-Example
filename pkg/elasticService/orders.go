package elasticservice

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const INDEX_NAME = "orders"

type OrderIndex *elasticsearch.TypedClient

func (es OrderIndex) findAll() {
	res, err := es.Search().Index(INDEX_NAME).Request(&search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}).Do(context.Background())
	return res.Hits, err
}
