package main

import (
	"RestaurantSearcherAPI/config"
	"RestaurantSearcherAPI/server"
	"github.com/olivere/elastic"
)

func main() {
	c := config.LoadConfig()

	esClient, err := elastic.NewClient(
		elastic.SetURL(c.Elasticsearch.Endpoint),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	defer esClient.Stop()

	routes := server.SetupServer(esClient)

	server.Run(routes)
}
