package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyentrunghieu15/Elasticsearch-Example.git/pkg/elasticcontroller"
	elasticservice "github.com/nguyentrunghieu15/Elasticsearch-Example.git/pkg/elasticservice"
)

func InitElasticRouter(r *gin.Engine, e *elasticservice.ElasticService) {
	ctx := r.Group("/elastic")
	{
		InitOrderElasticRoute(ctx, e)
	}
}

func InitOrderElasticRoute(r *gin.RouterGroup, e *elasticservice.ElasticService) {
	con := &elasticcontroller.ElasticOrderController{OrderService: e}
	ctx := r.Group("/order")
	{
		ctx.GET("/search-all", con.GetAll)
		ctx.GET("/search-ids", con.SearchIds)
		ctx.GET("/stats-order", con.StatisticOrderAmount)
	}
}
