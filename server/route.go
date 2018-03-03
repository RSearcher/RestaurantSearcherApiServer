package server

import (
	"github.com/gin-gonic/gin"
	"RestaurantSearcherAPI/route"
)

func setRoutes(r *gin.Engine) {
	r.GET("/ping", route.Ping)
}
