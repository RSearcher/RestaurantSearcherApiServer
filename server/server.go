package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func Run(elasticClient elastic.Client) {
	r := setupRouter(elasticClient)
	r.Run(":8080")
}

func setupRouter(elasticClient elastic.Client) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("Elasticsearch", elasticClient)
		c.Next()
	})

	setRoutes(r)
	return r
}
