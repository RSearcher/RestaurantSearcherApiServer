package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/olivere/elastic"
	"context"
)

func GetRestaurantById(c *gin.Context) {
	client := c.MustGet("ESClient").(*elastic.Client)

	req := struct {
		Index	string	`json:"index"`
		Type	string	`json:"type"`
		Id		string	`json:"id"`
	}{}

	if err := c.BindJSON(&req); err == nil {
		resp, err := client.Get().
			Index(req.Index).
			Type(req.Type).
			Id(req.Id).
			Do(context.Background())

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}

		c.JSON(http.StatusOK, resp)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
