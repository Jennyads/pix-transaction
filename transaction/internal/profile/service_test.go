package profile_test

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"testing"
	"transaction/internal/cfg"
	"transaction/internal/profile"
	"transaction/internal/utils"
	proto "transaction/proto/v1"
)

type MockAccountService struct {
	proto.AccountServiceClient
	mock.Mock
}

func (m *MockAccountService) UpdateAccount(ctx context.Context, account *proto.Account, opts ...grpc.CallOption) error {
	args := m.Called(ctx, account)
	return args.Error(0)
}

func (m *MockAccountService) FindAccount(ctx context.Context, req *proto.AccountRequest, opts ...grpc.CallOption) (*proto.Account, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*proto.Account), args.Error(1)
}

func (m *MockAccountService) FindByKey(ctx context.Context, req *proto.FindByKeyRequest, opts ...grpc.CallOption) (*proto.Account, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*proto.Account), args.Error(1)
}

func (m *MockAccountService) IsAccountActive(ctx context.Context, in *proto.AccountRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	args := m.Called(ctx, in)
	return &wrappers.BoolValue{Value: args.Bool(0)}, args.Error(1)
}

func NewMockAccountService() *MockAccountService {
	return &MockAccountService{}
}

type MockPixTransactionService struct {
	proto.PixTransactionServiceClient
	mock.Mock
}

func (m *MockPixTransactionService) PixWebhook(ctx context.Context, in *proto.Webhook, opts ...grpc.CallOption) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func NewMockPixTransactionService() *MockPixTransactionService {
	return &MockPixTransactionService{}
}
func TestUpdateAccountBalance(t *testing.T) {
	cases := []struct {
		name     string
		account  *profile.Account
		mockFunc func(service *MockAccountService)
		wantErr  error
	}{
		{
			name: "success",
			account: &profile.Account{
				Id:      "1",
				UserID:  "1",
				Balance: decimal.Zero,
				Agency:  "1",
				Bank:    "1",
			},
			mockFunc: func(service *MockAccountService) {
				service.On("UpdateAccount", mock.Anything, &proto.Account{
					UserId:  "1",
					Balance: utils.ToFloat(decimal.Zero),
					Agency:  "1",
					Bank:    "1",
				}).Return(nil, nil)
			},
			wantErr: nil,
		},
		{
			name: "failure",
			account: &profile.Account{
				Id:      "1",
				UserID:  "1",
				Balance: decimal.Zero,
				Agency:  "1",
				Bank:    "1",
			},
			mockFunc: func(service *MockAccountService) {
				service.On("UpdateAccount", mock.Anything, &proto.Account{
					UserId:  "1",
					Balance: utils.ToFloat(decimal.Zero),
					Agency:  "1",
					Bank:    "1",
				}).Return(nil, errors.New("MOCK-ERROR"))
			},
			wantErr: errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			config := &cfg.Config{ProfileConfig: cfg.ProfileConfig{Host: "localhost:50051"}}
			service := profile.NewService(config)

			mockService := NewMockAccountService()
			c.mockFunc(mockService)

			err := service.UpdateAccountBalance(context.Background(), c.account)
			assert.Error(t, err, c.wantErr)
		})
	}
}

func TestFindAccount(t *testing.T) {
	var cases []struct {
		name        string
		accountID   string
		mockFunc    func(service *MockAccountService)
		expected    *profile.Account
		expectedErr error
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			config := &cfg.Config{ProfileConfig: cfg.ProfileConfig{Host: "localhost:50051"}}
			service := profile.NewService(config)
			mockService := NewMockAccountService()
			c.mockFunc(mockService)

			account, err := service.FindAccount(context.Background(), c.accountID)
			assert.Equal(t, c.expected, account)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}

func TestFindReceiver(t *testing.T) {
	var cases []struct {
		name        string
		receiverKey string
		mockFunc    func(service *MockAccountService)
		expected    *profile.Account
		expectedErr error
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			config := &cfg.Config{ProfileConfig: cfg.ProfileConfig{Host: "localhost:50051"}}
			service := profile.NewService(config)
			mockService := NewMockAccountService()
			c.mockFunc(mockService)

			account, err := service.FindReceiver(context.Background(), c.receiverKey)

			assert.Equal(t, c.expected, account)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}

func TestIsAccountActive(t *testing.T) {
	var cases []struct {
		name        string
		accountID   string
		mockFunc    func(service *MockAccountService)
		expected    bool
		expectedErr error
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			config := &cfg.Config{ProfileConfig: cfg.ProfileConfig{Host: "localhost:50051"}}
			service := profile.NewService(config)
			mockService := NewMockAccountService()
			c.mockFunc(mockService)
			isActive, err := service.IsAccountActive(context.Background(), c.accountID)
			assert.Equal(t, c.expected, isActive)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}

func TestSendWebhook(t *testing.T) {
	var cases []struct {
		name        string
		webhook     *profile.Webhook
		mockFunc    func(service *MockPixTransactionService)
		expectedErr error
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			config := &cfg.Config{ProfileConfig: cfg.ProfileConfig{Host: "localhost:50051"}}
			service := profile.NewService(config)
			mockService := NewMockPixTransactionService()
			c.mockFunc(mockService)

			err := service.SendWebhook(context.Background(), c.webhook)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}
