package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/olivere/elastic"
	"context"
)

type Request struct {
	Index	string	`json:"index"`
	Type	string	`json:"type"`
	Id		string	`json:"id"`
}

func GetRestaurantById(c *gin.Context) {
	client := c.MustGet("ESClient").(*elastic.Client)

	var req Request

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	resp, err := client.Get().
		Index(req.Index).
		Type(req.Type).
		Id(req.Id).
		Do(context.Background())

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, resp)
}
