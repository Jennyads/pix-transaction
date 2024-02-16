package cfg

import (
	"os"
	"strings"
)

type DynamodbConfig struct {
	Host             string
	Port             string
	TransactionTable string
	KeysTable        string
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
			Host:             Getenv("DB_HOST", "localhost"),
			Port:             Getenv("DB_PORT", "8000"),
			TransactionTable: Getenv("TRANSACTION_TABLE", "transaction"),
			KeysTable:        Getenv("KEYS_TABLE", "keys"),
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
