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

func TestFindAccountById(t *testing.T) {
	cases := []struct {
		name     string
		req      *AccountRequest
		mockFunc func(repo *mockRepo)
		want     *Account
		err      error
	}{
		{
			name: "success",
			req: &AccountRequest{
				AccountID: "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("FindAccountById", "1").
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
			name: "failed because error in find account by id",
			req: &AccountRequest{
				AccountID: "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("FindAccountById", "1").
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

			got, err := s.FindAccountById(context.Background(), c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestUpdateAccount(t *testing.T) {
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
				Id:      "1",
				UserID:  "1",
				Balance: decimal.Zero,
				Agency:  "1",
				Bank:    "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("UpdateAccount", &Account{
					Id:      "1",
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
			//mockFunc: func(repo *mockRepo) {
			//	updatedData := &Account{
			//		Id:      "1",
			//		UserID:  "1",
			//		Balance: decimal.NewFromFloat(100.0),
			//		Agency:  "2",
			//		Bank:    "2",
			//	}
			//
			//	repo.On("UpdateAccount", updatedData).
			//		Return(updatedData, nil)
			//},
			//want: &Account{
			//	Id:      "1",
			//	UserID:  "1",
			//	Balance: decimal.NewFromFloat(100.0),
			//	Agency:  "2",
			//	Bank:    "2",
			//},
			//err: nil,

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
			name: "failed because error in update account",
			req: &Account{
				Id:      "1",
				UserID:  "1",
				Balance: decimal.Zero,
				Agency:  "1",
				Bank:    "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("UpdateAccount", &Account{
					Id:      "1",
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

			got, err := s.UpdateAccount(context.Background(), c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestListAccounts(t *testing.T) {
	cases := []struct {
		name     string
		req      *ListAccountRequest
		mockFunc func(repo *mockRepo)
		want     []*Account
		err      error
	}{
		{
			name: "success",
			req: &ListAccountRequest{
				AccountIDs: []string{"1", "2"},
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("ListAccount", []string{"1", "2"}).
					Return([]*Account{
						{
							Id:      "1",
							UserID:  "1",
							Balance: decimal.Zero,
							Agency:  "1",
							Bank:    "1",
						},
						{
							Id:      "2",
							UserID:  "1",
							Balance: decimal.Zero,
							Agency:  "1",
							Bank:    "1",
						},
					}, nil)
			},
			want: []*Account{
				{
					Id:      "1",
					UserID:  "1",
					Balance: decimal.Zero,
					Agency:  "1",
					Bank:    "1",
				},
				{
					Id:      "2",
					UserID:  "1",
					Balance: decimal.Zero,
					Agency:  "1",
					Bank:    "1",
				},
			},
			err: nil,
		},
		{
			name: "failed because error in list accounts",
			req: &ListAccountRequest{
				AccountIDs: []string{"1", "2"},
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("ListAccount", []string{"1", "2"}).
					Return(([]*Account)(nil), errors.New("MOCK-ERROR"))
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

			got, err := s.ListAccounts(context.Background(), c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	cases := []struct {
		name     string
		req      *AccountRequest
		mockFunc func(repo *mockRepo)
		err      error
	}{
		{
			name: "success",
			req: &AccountRequest{
				AccountID: "1",
				UserID:    "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("DeleteAccount", "1").
					Return(nil)
			},
			err: nil,
		},
		{
			name: "failed because error in delete account",
			req: &AccountRequest{
				AccountID: "2",
				UserID:    "1",
			},
			mockFunc: func(repo *mockRepo) {
				repo.On("DeleteAccount", "2").
					Return(errors.New("MOCK-ERROR"))
			},
			err: errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			err := s.DeleteAccount(context.Background(), c.req)
			assert.Equal(t, err, c.err)
		})
	}
}

func TestIsAccountActive(t *testing.T) {
	cases := []struct {
		name     string
		id       string
		mockFunc func(repo *mockRepo)
		want     bool
		err      error
	}{
		{
			name: "active account",
			id:   "1",
			mockFunc: func(repo *mockRepo) {
				repo.On("IsAccountActive", "1").
					Return(true, nil).Times(1)
			},
			want: true,
			err:  nil,
		},
		{
			name: "inactive account",
			id:   "2",
			mockFunc: func(repo *mockRepo) {
				repo.On("IsAccountActive", "2").
					Return(false, nil).Times(1)
			},
			want: false,
			err:  nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			got, err := s.IsAccountActive(context.Background(), c.id)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestFindByKey(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		mockFunc func(repo *mockRepo)
		want     *Account
		err      error
	}{
		{
			name: "success",
			key:  "some_key",
			mockFunc: func(repo *mockRepo) {
				repo.On("FindByKey", "some_key").
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
			name: "failed because error in find by key",
			key:  "invalid_key",
			mockFunc: func(repo *mockRepo) {
				repo.On("FindByKey", "invalid_key").
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

			got, err := s.FindByKey(context.Background(), c.key)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}
