package cfg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name     string
		want     *Config
		mockFunc func(t *testing.T)
		err      error
	}{
		{
			name: "success",
			mockFunc: func(t *testing.T) {
				t.Setenv("DB_HOST", "example.net")
				//t.Setenv("DB_PORT", "") // expected fallback
				t.Setenv("DB_USER", "user")
				t.Setenv("DB_PASS", "pass")
				t.Setenv("DB_NAME", "name")

				t.Setenv("KAFKA_ADVERTISED_LISTENERS", "localhost:9092,localhost:9093,localhost:9094")
			},
			want: &Config{
				SqlServerConfig: SqlServerConfig{
					Host:     "example.net",
					Port:     "1433",
					User:     "user",
					Password: "pass",
					Database: "name",
				},
				KafkaConfig: KafkaConfig{
					Brokers: []string{"localhost:9092", "localhost:9093", "localhost:9094"},
				},
				RedisConfig: RedisConfig{
					Host:     "localhost",
					Port:     "6379",
					User:     "default",
					Password: "redis",
				},
			},
			err: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.mockFunc(t)
			got, err := Load()
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}
