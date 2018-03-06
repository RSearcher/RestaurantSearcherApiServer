package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/olivere/elastic"
	"context"
)

type IndexRequest struct {
	Index string `json:"index"`
}

type resp struct {
	Exist	bool `json:"exist"`
}

func ExistIndex(c *gin.Context) {
	var req IndexRequest

	client := c.MustGet("ESClient").(*elastic.Client)

	if err := c.BindJSON(&req); err == nil {
		exists, err := client.IndexExists(req.Index).Do(context.Background())
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		resp := &resp{exists}
		c.JSON(http.StatusOK, resp)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
