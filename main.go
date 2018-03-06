package main

import (
	"RestaurantSearcherAPI/config"
	"RestaurantSearcherAPI/server"
	"github.com/olivere/elastic"
)

func main() {
	c := config.LoadConfig()

	client, err := elastic.NewClient(
		elastic.SetURL(c.Elasticsearch.Endpoint),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	server.Run(*client)
}
