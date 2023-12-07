package account

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRepo struct {
	Repository
	mock.Mock
}

func (m *mockRepo) CreateAccount(account *Account) (*Account, error) {
	args := m.Called(account)
	return args.Get(0).(*Account), args.Error(1)
}

func (m *mockRepo) FindAccountById(id string) (*Account, error) {
	args := m.Called(id)
	return args.Get(0).(*Account), args.Error(1)
}

func (m *mockRepo) UpdateAccount(account *Account) (*Account, error) {
	args := m.Called(account)
	return args.Get(0).(*Account), args.Error(1)
}

func (m *mockRepo) ListAccount(ids []string) ([]*Account, error) {
	args := m.Called(ids)
	return args.Get(0).([]*Account), args.Error(1)
}

func (m *mockRepo) DeleteAccount(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockRepo) IsAccountActive(ctx context.Context, id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *mockRepo) FindByKey(key string) (*Account, error) {
	args := m.Called(key)
	return args.Get(0).(*Account), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	cases := []struct {
		name     string
		req      *Account
		mockFunc func(repo *mockRepo)
		want     *Account
		err      error
	}{
		{
			name: "success",
			req: &Account{
				UserID:  "1",
				Balance: decimal.Zero,
				Agency:  "1",
				Bank:    "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("CreateAccount", &Account{
					UserID:  "1",
					Balance: decimal.Zero,
					Agency:  "1",
					Bank:    "1",
				}).
					Return(&Account{
						Id:      "1",
						UserID:  "1",
						Balance: decimal.Zero,
						Agency:  "1",
						Bank:    "1",
					}, nil)
			},
			want: &Account{
				Id:      "1",
				UserID:  "1",
				Balance: decimal.Zero,
				Agency:  "1",
				Bank:    "1",
			},
			err: nil,
		},
		{
			name: "failed because error in create account",
			req: &Account{
				UserID:  "1",
				Balance: decimal.Zero,
				Agency:  "1",
				Bank:    "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("CreateAccount", &Account{
					UserID:  "1",
					Balance: decimal.Zero,
					Agency:  "1",
					Bank:    "1",
				}).
					Return((*Account)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			got, err := s.CreateAccount(context.Background(), c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}
