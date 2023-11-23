package kafka

import (
	"github.com/segmentio/kafka-go"
	"transaction/internal/cfg"
)

type Client interface {
	Connect() Client
}

type client struct {
	conn   *kafka.Conn
	config *cfg.Config
}

func (c *client) Conn() *kafka.Conn {
	return c.conn
}

func (c *client) Close() {
	c.conn.Close()

}

func (c *client) Connect() Client {
	conn, err := kafka.Dial("tcp", c.config.KafkaConfig.Brokers[0])
	if err != nil {
		panic(err.Error())
	}
	c.conn = conn
	return c
}

func NewClient(config *cfg.Config) Client {
	return &client{
		config: config,
	}
}
