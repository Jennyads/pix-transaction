package cfg

import (
	"os"
	"strings"
)

type DynamodbConfig struct {
	TransactionTable string
}

type KafkaConfig struct {
	Brokers []string
}

type Config struct {
	DynamodbConfig DynamodbConfig
	KafkaConfig    KafkaConfig
}

func Load() (*Config, error) {
	return &Config{
		DynamodbConfig{
			TransactionTable: Getenv("TRANSACTION_TABLE", "transaction"),
		},
		KafkaConfig{
			Brokers: strings.Split(Getenv("KAFKA_ADVERTISED_LISTENERS", "localhost:9092"), ","),
		},
	}, nil
}

func Getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
