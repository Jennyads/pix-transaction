package kafka

import (
	"github.com/segmentio/kafka-go"
	"os"
)

type Client interface {
	Connect() Client
}

type client struct {
	conn *kafka.Conn
}

func (c *client) Conn() *kafka.Conn {
	return c.conn
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) Connect() Client {
	kafkaBrokers := os.Getenv("TRANSACTION_SERVICE_KAFKA_ADVERTISED_LISTENERS")
	if kafkaBrokers == "" {
		kafkaBrokers = "localhost:9092"
	}
	conn, err := kafka.Dial("tcp", kafkaBrokers)
	if err != nil {
		panic(err.Error())
	}
	c.conn = conn
	return c
}

func NewClient() Client {
	return &client{}
}
