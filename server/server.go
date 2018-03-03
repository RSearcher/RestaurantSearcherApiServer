package server

import "github.com/gin-gonic/gin"

func Run() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	setRoutes(r)
	return r
}
