package main

import (
	"RestaurantSearcherAPI/config"
	"github.com/olivere/elastic"
	"RestaurantSearcherAPI/ml"
	"github.com/go-redis/redis"
	"RestaurantSearcherAPI/route"
)

func main() {
	conf := config.LoadConfig()

	esClient := setupElastic(conf)

	defer clearElastic(esClient)

	mlClient := setupML(conf)

	rsClient := setupRedis(conf)

	defer clearRedis(rsClient)

	routes := route.SetupContext(route.SetupRouter(), conf, esClient, mlClient, rsClient)

	routes.Run(":8080")
}

func clearRedis(rsClient *redis.Client) error {
	return rsClient.Close()
}

func setupRedis(conf *config.Config) *redis.Client {
	rsClient := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Endpoint,
		Password: "",
		DB:       0,
	})
	return rsClient
}

func setupML(conf *config.Config) *ml.Client {
	mlClient, err := ml.NewClient(conf.MLServer.Endpoint)
	if err != nil {
		panic(err)
	}
	return mlClient
}

func clearElastic(esClient *elastic.Client) {
	esClient.Stop()
}

func setupElastic(conf *config.Config) *elastic.Client {
	esClient, err := elastic.NewClient(
		elastic.SetURL(conf.Elasticsearch.Endpoint),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	return esClient
}
