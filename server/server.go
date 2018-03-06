package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func Run(elasticClient elastic.Client) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("Elasticsearch", elasticClient)
		c.Next()
	})
	setRoutes(r)
	r.Run(":8080")
}
