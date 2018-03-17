package route

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"RestaurantSearcherAPI/ml"
	"RestaurantSearcherAPI/config"
	"github.com/go-redis/redis"
)


func setRoutes(r *gin.Engine) {
	r.GET("/ping", Ping)
	r.GET("/restaurant/:id", GetRestaurantById)
	r.GET("/review/:id", GetReviewById)
	r.POST("/review/parse", ParseReview)
	//r.GET("/review/:size", route.GetSizeOfReviews)
	//r.POST("/restaurants/:text", route.ParseTextsByKNP)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	setRoutes(r)

	return r
}

func SetupContext(
	r *gin.Engine,
	conf *config.Config,
	ESClient *elastic.Client,
	MLClient *ml.Client,
	RSClient *redis.Client,
	) *gin.Engine {
	r.Use(func(c *gin.Context) {
		c.Set("Config", conf)
		c.Set("ESClient", ESClient)
		c.Set("MLClient", MLClient)
		c.Set("RSClient", RSClient)
		c.Next()
	})
	return r
}
