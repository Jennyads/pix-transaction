package event

import (
	"context"
	"encoding/json"
	kafkago "github.com/segmentio/kafka-go"
	"log"
	"profile/internal/transaction"
	"profile/platform/kafka"
)

type Client interface {
	Publish(ctx context.Context, payload []byte) error
}

type PixEvent struct {
	PixData *transaction.Pix
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

type event struct {
	topic       string
	maxAttempts int
	kafka       kafka.Client
	brokers     []string
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

func (e *event) PublishPix(ctx context.Context, pixEvent PixEvent) error {
	payload, err := json.Marshal(pixEvent)
	if err != nil {
		return err
	}
	return e.Publish(ctx, payload)
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
