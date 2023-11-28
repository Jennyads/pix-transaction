package event

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"log"
	"transaction/platform/kafka"
)

type Client interface {
	Publish(ctx context.Context, payload []byte) error
	RegisterHandler(ctx context.Context, topic string, handler Function) error
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

//func (client *kafkaClient) handleIncomingMessages(ctx context.Context, topic string, handler Function) {
//	r := client.newReader(topic)
//
//	log.WithContext(ctx).WithField("topic", topic).Info("Listener registered")
//
//	for {
//		select {
//		case <-ctx.Done():
//			client.Close()
//			return
//		case <-client.stopSignal:
//			if err := r.Close(); err != nil {
//				log.WithContext(ctx).WithError(err).Warn("Failed to close reader")
//			}
//		default:
//			msg, err := r.FetchMessage(ctx)
//
//			if err != nil {
//				log.WithContext(ctx).WithError(err).Error("Failed to fetch message")
//			}
//
//			if msg.Value != nil {
//				log.WithContext(ctx).WithField("topic", topic).Info("Message received")
//
//				if topic == "transaction_events_topic" {
//
//					pixHandler(ctx, msg.Value)
//				} else {
//					_, err = handler.Handle(ctx, msg.Value)
//					if err != nil {
//						client.reprocess <- reprocess{msg, handler, ctx}
//					}
//
//					if err = r.CommitMessages(ctx, msg); err != nil {
//						log.WithError(err).Error("Failed to commit messages")
//					}
//				}
//			}
//		}
//	}
//}

func (e *event) RegisterHandler(ctx context.Context, topic string, handler Function) error {
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
