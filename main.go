package main

import (
	"RestaurantSearcherAPI/config"
	"RestaurantSearcherAPI/server"
	"github.com/olivere/elastic"
)

func main() {
	conf := config.LoadConfig()

	esClient, err := elastic.NewClient(
		elastic.SetURL(conf.Elasticsearch.Endpoint),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	defer esClient.Stop()

	routes := server.SetupServer(conf, esClient)

	server.Run(routes)
}
