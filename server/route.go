package server

import (
	"RestaurantSearcherAPI/route"
	"github.com/gin-gonic/gin"
)

func setRoutes(r *gin.Engine) {
	r.GET("/ping", route.Ping)
	r.GET("/restaurant/:id", route.GetRestaurantById)
	r.GET("/review/:id", route.GetReviewById)
	r.POST("/review/parse", route.ParseReview)
	//r.GET("/review/:size", route.GetSizeOfReviews)
	//r.POST("/restaurants/:text", route.ParseTextsByKNP)
}
