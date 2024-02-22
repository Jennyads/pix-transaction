package transaction

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type mockEventClient struct {
	mock.Mock
}
type MockAccountRepository struct {
	mock.Mock
}

func (m *mockEventClient) Publish(ctx context.Context, payload []byte) error {
	args := m.Called(ctx, payload)
	return args.Error(0)
}

func (m *mockEventClient) SendPix(ctx context.Context, req *Pix) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

//func TestSendPix(t *testing.T) {
//	ctx := context.Background()
//
//	cases := []struct {
//		name          string
//		pix           *Pix
//		mockFunc      func(client *mockEventClient)
//		expectedError error
//	}{
//		{
//			name: "successful send pix",
//			pix: &Pix{
//				UserID:    "user-1",
//				AccountID: "account-1",
//				Key:       "key-1",
//				Receiver:  "receiver-1",
//				Amount:    decimal.NewFromFloat(100.50),
//				Status:    "success",
//			},
//			mockFunc: func(client *mockEventClient) {
//				client.On("Publish", ctx, mock.Anything).Return(nil)
//			},
//			expectedError: nil,
//		},
//		{
//			name: "failed to marshal pix event",
//			pix: &Pix{
//				UserID:    "user-1",
//				AccountID: "account-1",
//				Key:       "key-1",
//				Receiver:  "receiver-1",
//				Amount:    decimal.NewFromFloat(100.50),
//				Status:    "success",
//			},
//			mockFunc: func(client *mockEventClient) {
//				client.On("Publish", mock.Anything, mock.Anything).Return(errors.New("json: error calling MarshalJSON for type *transaction.PixEvent: json: error calling MarshalJSON for type transaction.Pix: invalid character 'd' looking for beginning of object key"))
//			},
//			expectedError: status.Error(
//				codes.Internal,
//				"json: error calling MarshalJSON for type *transaction.PixEvent: json: error calling MarshalJSON for type transaction.Pix: invalid character 'd' looking for beginning of object key",
//			),
//		},
//		{
//			name: "failed to publish pix event",
//			pix: &Pix{
//				UserID:    "user-1",
//				AccountID: "account-1",
//				Key:       "key-1",
//				Receiver:  "receiver-1",
//				Amount:    decimal.NewFromFloat(100.50),
//				Status:    "success",
//			},
//			mockFunc: func(client *mockEventClient) {
//				client.On("Publish", ctx, mock.AnythingOfType("[]uint8")).Return(errors.New("publish-error"))
//			},
//			expectedError: status.Error(codes.Internal, "publish-error"),
//		},
//	}
//
//	for _, c := range cases {
//		t.Run(c.name, func(t *testing.T) {
//			mockClient := new(mockEventClient)
//			c.mockFunc(mockClient)
//
//			service := NewService(mockClient, nil)
//			err := service.SendPix(ctx, c.pix)
//
//			assert.Equal(t, c.expectedError, err)
//
//			mockClient.AssertExpectations(t)
//		})
//	}
//}
