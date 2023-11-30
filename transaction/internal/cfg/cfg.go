package cfg

import (
	"os"
	"strings"
)

type ProfileConfig struct {
	Host string
}

type DynamodbConfig struct {
	TransactionTable string
}

type KafkaConfig struct {
	Brokers []string
}

type Config struct {
	DynamodbConfig DynamodbConfig
	KafkaConfig    KafkaConfig
	ProfileConfig  ProfileConfig
}

func Load() (*Config, error) {
	return &Config{
		DynamodbConfig{
			TransactionTable: Getenv("TRANSACTION_TABLE", "transaction"),
		},
		KafkaConfig{
			Brokers: strings.Split(Getenv("KAFKA_ADVERTISED_LISTENERS", "localhost:9092"), ","),
		},
		ProfileConfig{
			Host: Getenv("PROFILE_HOST", "localhost:9080"),
		},
	}, nil
}

func Getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
