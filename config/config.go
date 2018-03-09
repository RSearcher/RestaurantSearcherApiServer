package config

import "github.com/BurntSushi/toml"

type Config struct {
	Elasticsearch ElasticsearchConfig
}

type ElasticsearchConfig struct {
	Endpoint string
	ReviewsIndexName string
	ReviewsTypeName string
	RestaurantsIndexName string
	RestaurantsTypeName string
}

func LoadConfig() *Config {
	c := &Config{}
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		panic(err)
	}
	return c
}