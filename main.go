package main

import (
	"RestaurantSearcherAPI/config"
	"RestaurantSearcherAPI/server"
	"github.com/BurntSushi/toml"
	"github.com/olivere/elastic"
)

func loadConfig() *config.Config {
	c := &config.Config{}
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		panic(err)
	}
	return c
}

func main() {
	c := loadConfig()

	client, err := elastic.NewClient(
		elastic.SetURL(c.ElasticSearchURL),
		elastic.SetScheme("https"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	server.Run(*client)
}
