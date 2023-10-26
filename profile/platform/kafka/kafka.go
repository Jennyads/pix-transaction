package kafka

import "github.com/segmentio/kafka-go"

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
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	c.conn = conn
	return c
}

func NewClient() Client {
	return &client{}
}
