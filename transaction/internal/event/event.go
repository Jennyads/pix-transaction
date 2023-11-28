package event

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"log"
	"transaction/platform/kafka"
)

type Client interface {
	CreateTopic() error
	Publish(ctx context.Context, payload []byte) error
	RegisterHandler(ctx context.Context, handler Function) error
}

type Options func(*event)

func WithAttempts(attempts int) Options {
	return func(e *event) {
		e.maxAttempts = attempts
	}
}

func WithBroker(broker string) Options {
	return func(e *event) {
		e.brokers = append(e.brokers, broker)
	}
}

type Function func(ctx context.Context, payload []byte) ([]byte, error)

type event struct {
	topic       string
	maxAttempts int
	kafka       kafka.Client
	brokers     []string
}

func (e *event) CreateTopic() error {
	return e.kafka.Conn().CreateTopics(kafkago.TopicConfig{Topic: e.topic, NumPartitions: 1, ReplicationFactor: 1})
}

func (e *event) Publish(ctx context.Context, payload []byte) error {
	w := &kafkago.Writer{
		Addr:                   kafkago.TCP(e.brokers...),
		Topic:                  e.topic,
		MaxAttempts:            e.maxAttempts,
		Transport:              kafkago.DefaultTransport,
		AllowAutoTopicCreation: true,
	}

	err := w.WriteMessages(ctx, kafkago.Message{Value: payload})
	if err != nil {
		log.Print("failed to write messages:", err)
	}

	if err = w.Close(); err != nil {
		log.Print("failed to close writer:", err)
	}
	return nil
}

func (e *event) handleMessages(ctx context.Context, handler Function) {
	r := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: e.brokers,
		Topic:   e.topic,
		GroupID: e.topic + "_handler",
	})
	log.Printf("listener registered for topic [%s]\n", e.topic)

	for {
		select {
		case <-ctx.Done():
			if err := r.Close(); err != nil {
				log.Print("failed to close reader:", err)
			}
			return
		default:
			msg, err := r.FetchMessage(ctx)
			if err != nil {
				log.Println("Failed to fetch message")
			}

			if msg.Value != nil {
				log.Printf("Message received: [%s]\n", e.topic)
				_, err = handler(ctx, msg.Value)
				if err != nil {
					log.Print("failed to handle message:", err)
				}

				if err = r.CommitMessages(ctx, msg); err != nil {
					log.Println("Failed to commit messages")
				}
			}
		}
	}
}

func (e *event) RegisterHandler(ctx context.Context, handler Function) error {
	go e.handleMessages(ctx, handler)
	return nil
}

func NewEvent(client kafka.Client, topic string, opts ...Options) Client {
	e := &event{
		kafka: client,
		topic: topic,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}
