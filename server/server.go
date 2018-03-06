package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"net/http"
	"github.com/wantedly/yashima-news-timeline/server/util"
	"github.com/wantedly/yashima-news-timeline/server"
)

func Run(r *gin.Engine) {
	httpServer := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: util.NewMethodOverrider(r),
		ReadTimeout: server.ServerReadTimeout,
		WriteTimeout: server.ServerWriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
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