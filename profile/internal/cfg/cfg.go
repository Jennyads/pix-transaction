package cfg

import (
	"os"
	"strings"
)

type SqlServerConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type KafkaConfig struct {
	Brokers []string
}

type RedisConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

// Config is the struct that holds all the configuration for the application
type Config struct {
	SqlServerConfig SqlServerConfig
	KafkaConfig     KafkaConfig
	RedisConfig     RedisConfig
}

func Load() (*Config, error) {
	return &Config{
		SqlServerConfig{
			Host:     Getenv("DB_HOST", "localhost"),
			Port:     Getenv("DB_PORT", "1433"),
			User:     Getenv("DB_USER", "sa"),
			Password: Getenv("DB_PASS", "SqlServer2019!"),
			Database: Getenv("DB_NAME", "profile"),
		},
		KafkaConfig{
			Brokers: strings.Split(Getenv("KAFKA_ADVERTISED_LISTENERS", "localhost:9092"), ","),
		},
		RedisConfig{
			Host:     Getenv("REDIS_HOST", "localhost"),
			Port:     Getenv("REDIS_PORT", "6379"),
			User:     Getenv("REDIS_USER", "default"),
			Password: Getenv("REDIS_PASSWORD", "redis"),
		},
	}, nil
}

func Getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
