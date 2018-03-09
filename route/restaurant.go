package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/olivere/elastic"
	"context"
	"RestaurantSearcherAPI/config"
)

func GetRestaurantById(c *gin.Context) {
	client := c.MustGet("ESClient").(*elastic.Client)
	conf := c.MustGet("Config").(*config.Config)

	id := c.Param("id")

	resp, err := client.Get().
		Index(conf.Elasticsearch.RestaurantsIndexName).
		Type(conf.Elasticsearch.RestaurantsTypeName).
		Id(id).
		Do(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, resp)
}
