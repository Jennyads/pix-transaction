package cfg

import "os"

type DynamodbConfig struct {
	TransactionTable string
}

type Config struct {
	DynamodbConfig DynamodbConfig
}

func Get() (*Config, error) {
	return &Config{
		DynamodbConfig{
			TransactionTable: os.Getenv("TRANSACTION_TABLE"),
		},
	}, nil
}
