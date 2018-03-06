package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c := LoadConfig()

	assert.Equal(
		t,
		c.Elasticsearch.Endpoint,
		"vpc-rsearcher-q6wqiureoqz3oaj6udaihubs6y.us-east-1.es.amazonaws.com:9200",
		)
}
