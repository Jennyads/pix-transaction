package cfg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name     string
		want     *Config
		mockFunc func()
		err      error
	}{
		{
			name: "success",
			mockFunc: func() {
				t.Setenv("TRANSACTION_TABLE", "test_transaction_table")
				t.Setenv("KAFKA_ADVERTISED_LISTENERS", "kafka1:9092,kafka2:9092")
				t.Setenv("PROFILE_HOST", "test_profile_host")
			},
			want: &Config{
				DynamodbConfig: DynamodbConfig{
					TransactionTable: "test_transaction_table",
				},
				KafkaConfig: KafkaConfig{
					Brokers: []string{"kafka1:9092", "kafka2:9092"},
				},
				ProfileConfig: ProfileConfig{
					Host: "test_profile_host",
				},
			},
			err: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.mockFunc()
			got, err := Load()
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}
