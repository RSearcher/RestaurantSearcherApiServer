package server

import (
	"RestaurantSearcherAPI/route"
	"github.com/gin-gonic/gin"
)

func setRoutes(r *gin.Engine) {
	r.GET("/ping", route.Ping)
	r.GET("/restaurants/:id", route.GetRestaurantById)
}
