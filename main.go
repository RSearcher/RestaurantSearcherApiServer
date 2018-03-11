package main

import (
	"RestaurantSearcherAPI/config"
	"RestaurantSearcherAPI/server"
	"github.com/olivere/elastic"
	"RestaurantSearcherAPI/ml"
	"github.com/go-redis/redis"
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

	rsClient := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Endpoint,
		Password: "",
		DB: 0,
	})

	routes := server.SetupServer(conf, esClient, mlClient, rsClient)

	server.Run(routes)
}
