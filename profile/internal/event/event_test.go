package event

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"profile/platform/kafka"
	"testing"
)

type mockKafkaClient struct {
	kafka.Client
	mock.Mock
}

func (m *mockKafkaClient) Publish(ctx context.Context, payload []byte) error {
	args := m.Called(ctx, payload)
	return args.Error(0)
}

func TestEventPublish(t *testing.T) {
	ctx := context.Background()
	topic := "test-topic"
	payload := []byte("test-payload")

	cases := []struct {
		name         string
		client       Client
		mockFunc     func(client *mockKafkaClient)
		expectedErr  error
		expectedLogs []string
	}{
		{
			name: "successful publish",
			client: &event{
				topic:       topic,
				maxAttempts: 3,
				kafka:       &mockKafkaClient{},
				brokers:     []string{"broker-1", "broker-2"},
			},
			mockFunc: func(client *mockKafkaClient) {
				client.On("Publish", ctx, payload).Return(nil)
			},
			expectedErr:  nil,
			expectedLogs: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockClient := c.client.(*event).kafka.(*mockKafkaClient)
			c.mockFunc(mockClient)

			err := c.client.Publish(ctx, payload)

			assert.Equal(t, c.expectedErr, err)

		})
	}
}
