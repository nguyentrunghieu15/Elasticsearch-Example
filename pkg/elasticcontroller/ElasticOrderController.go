package elasticcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nguyentrunghieu15/Elasticsearch-Example.git/pkg/elasticservice"
)

type ElasticOrderController struct {
	OrderService *elasticservice.ElasticService
}

func (v *ElasticOrderController) GetAll(ctx *gin.Context) {
	res, err := v.OrderService.OrderIndex.FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err})
	}
	defer res.Body.Close()
	r := map[string]interface{}{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatalf("Error parsing the response body: %s", err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": r})
}

func (v *ElasticOrderController) SearchIds(ctx *gin.Context) {
	if ids, ok := ctx.GetQuery("ids"); ok {
		res, err := v.OrderService.OrderIndex.SearchIds(strings.Split(ids, ","))
		if err != nil {
			ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		}
		defer res.Body.Close()
		r := map[string]interface{}{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, gin.H{"data": r})
	} else {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": fmt.Errorf("Ids is required")})
	}
}

func (v *ElasticOrderController) StatisticOrderAmount(ctx *gin.Context) {
	res, err := v.OrderService.OrderIndex.StatsOrder()
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err})
	}
	defer res.Body.Close()
	r := map[string]interface{}{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": r})
}
