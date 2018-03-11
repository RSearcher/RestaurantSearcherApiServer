package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"RestaurantSearcherAPI/config"
	"RestaurantSearcherAPI/ml"
	"github.com/go-redis/redis"
)

func Run(r *gin.Engine) {
	r.Run(":8080")
}

func SetupServer(
	conf *config.Config,
	ESClient *elastic.Client,
	MLClient *ml.Client,
	RSClient *redis.Client,
	) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("Config", conf)
		c.Set("ESClient", ESClient)
		c.Set("MLClient", MLClient)
		c.Set("RSClient", RSClient)
		c.Next()
	})

	setRoutes(r)

	return r
}