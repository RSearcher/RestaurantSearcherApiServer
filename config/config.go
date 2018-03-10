package config

import "github.com/BurntSushi/toml"

type Config struct {
	Elasticsearch ElasticsearchConfig
	MLServer MLServerConfig
}

type ElasticsearchConfig struct {
	Endpoint string
	ReviewsIndexName string
	ReviewsTypeName string
	RestaurantsIndexName string
	RestaurantsTypeName string
}

type MLServerConfig struct {
	Endpoint string
}

func LoadConfig() *Config {
	c := &Config{}
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		panic(err)
	}
	return c
}