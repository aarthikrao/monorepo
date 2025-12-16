package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type KycConfig struct {
	PostgresURL string
	MongoURL    string
	RedisURL    string
}

func LoadConfig(file string) (*KycConfig, error) {
	// Load configuration from environment variables or files
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg KycConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
