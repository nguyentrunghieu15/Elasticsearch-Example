package main

import (
	"github.com/gin-gonic/gin"
	elasticservice "github.com/nguyentrunghieu15/Elasticsearch-Example.git/pkg/elasticservice"
)

func InitRoutes(r *gin.Engine, e *elasticservice.ElasticService) {
	InitElasticRouter(r, e)
}
