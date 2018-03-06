package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func Run(r *gin.Engine) {
	r.Run(":8080")
}

func SetupServer(ESClient *elastic.Client) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("ESClient", ESClient)
		c.Next()
	})

	setRoutes(r)

	return r
}