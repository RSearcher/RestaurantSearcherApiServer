package route

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"RestaurantSearcherAPI/config"
	"context"
	"net/http"
	"encoding/json"
	"RestaurantSearcherAPI/models"
	"RestaurantSearcherAPI/ml"
)

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

	var r models.Review

	json.Unmarshal(*resp.Source, &r)

	c.JSON(http.StatusOK, r)
}

func ParseReview(c *gin.Context) {
	mlclient := c.MustGet("MLClient").(*ml.Client)

	var review *models.Review
	if err := c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	parsedText, err := mlclient.ParseKNP(context.Background(), review)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
	}

	c.JSON(http.StatusOK, parsedText)
}
