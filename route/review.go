package route

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"RestaurantSearcherAPI/config"
	"context"
	"net/http"
	"encoding/json"
)

type Response struct {
	Id int `json:"id"`
	Body string `json:"body"`
}

func GetReviewById(c *gin.Context) {
	esclient := c.MustGet("ESClient").(*elastic.Client)
	conf := c.MustGet("Config").(*config.Config)

	id := c.Param("id")

	resp, err := esclient.Get().
		Index(conf.Elasticsearch.ReviewsIndexName).
		Type(conf.Elasticsearch.ReviewsTypeName).
		Id(id).
		Do(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	var r Response

	json.Unmarshal(*resp.Source, &r)

	c.JSON(http.StatusOK, r)
}