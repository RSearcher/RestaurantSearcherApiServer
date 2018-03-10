package main

import (
	"RestaurantSearcherAPI/config"
	"RestaurantSearcherAPI/server"
	"github.com/olivere/elastic"
	"RestaurantSearcherAPI/ml"
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

	mlClient, err := ml.NewClient(conf.MLServer.Endpoint)
	if err != nil {
		panic(err)
	}

	routes := server.SetupServer(conf, esClient, mlClient)

	server.Run(routes)
}
