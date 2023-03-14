package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "KITCHEN"

type Config struct {
	DbDsnEntrypoint string `split_words:"true" default:"host=localhost port=6432 dbname=kitchen-test user=test password=test sslmode=disable"`

	AppHost        string `split_words:"true" default:"localhost"`
	AppGRPCPort    int    `split_words:"true" default:"7002"`
	AppHTTPPort    int    `split_words:"true" default:"7000"`
	AppHTTPK8SPort int    `split_words:"true" default:"8080"`

	KafkaVersion       string        `split_words:"true" default:"2.5.0"`
	KafkaBrokers       string        `split_words:"true" default:"127.0.0.1:9092"`
	SessionTimeout     time.Duration `split_words:"true" default:"10s"`
	SessionRestartWait time.Duration `split_words:"true" default:"1s"`
	HeartbeatInterval  time.Duration `split_words:"true" default:"3s"`
	RetryBackoff       time.Duration `split_words:"true" default:"15s"`

	GroupName string `envconfig:"GROUP_NAME" default:"kitchen"`

	RetriesCount        int64         `split_words:"true" default:"5"`
	RetryWait           time.Duration `split_words:"true" default:"100ms"`
	ExternalCallTimeout time.Duration `split_words:"true" default:"5s"`

	ShopOrderEventTopic    string `split_words:"true" default:"shop_order_event"`
	KitchenOrderEventTopic string `split_words:"true" default:"kitchen_order_event"`
}

// FromEnv gets config from env vars
func FromEnv() (*Config, error) {
	config := &Config{}
	err := envconfig.Process(envPrefix, config)
	return config, err
}
