package cfg

import (
	"os"
	"strings"
)

type DynamodbConfig struct {
	AccountTable string
	UserTable    string
	KeyTable     string
	PixTable     string
}

type KafkaConfig struct {
	Brokers []string
}

// Config is the struct that holds all the configuration for the application
type Config struct {
	DynamodbConfig DynamodbConfig
	KafkaConfig    KafkaConfig
}

func Load() (*Config, error) {
	return &Config{
		DynamodbConfig{
			AccountTable: Getenv("ACCOUNT_TABLE", "account"),
			UserTable:    Getenv("USER_TABLE", "user"),
			KeyTable:     Getenv("KEY_TABLE", "key"),
			PixTable:     Getenv("PIX_TABLE", "pix"),
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
