package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func AliveElastic(c *gin.Context) {
	client, ok := c.Get("Elasticsearch")
	if !ok {
		c.String(http.StatusBadGateway, "error1")
	}

	fmt.Println(client)
}
