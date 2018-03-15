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
	"github.com/go-redis/redis"
	"strconv"
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
	rsclient := c.MustGet("RSClient").(*redis.Client)

	var review *models.Review
	if err := c.BindJSON(&review); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	getErr := rsclient.Get(strconv.Itoa(review.Id)).Err()
	if getErr == nil {
		// Key is already existed
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	parsedText, err := mlclient.ParseKNP(context.Background(), review)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	parsedTextJson, err := json.Marshal(parsedText)
	if err != nil {
		panic(err)
	}

	setErr := rsclient.Set(strconv.Itoa(review.Id), parsedTextJson, 0).Err()
	if setErr != nil {
		panic(setErr)
	}

	c.JSON(http.StatusOK, parsedText)
}
