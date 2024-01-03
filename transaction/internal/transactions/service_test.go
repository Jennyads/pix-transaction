package transactions

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockTransactionRepo struct {
	Repository
	mock.Mock
}

func (m *mockTransactionRepo) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(*Transaction), args.Error(1)
}

func (m *mockTransactionRepo) FindTransactionById(id string) (*Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*Transaction), args.Error(1)
}

func (m *mockTransactionRepo) ListTransactions(ids []string) ([]*Transaction, error) {
	args := m.Called(ids)
	return args.Get(0).([]*Transaction), args.Error(1)
}

func (m *mockTransactionRepo) UpdateTransactionStatus(transaction *Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransaction(t *testing.T) {
	cases := []struct {
		name     string
		req      *Transaction
		mockFunc func(repo *mockTransactionRepo)
		want     *Transaction
		err      error
	}{
		{
			name: "success",
			req: &Transaction{
				AccountID: "1",
				Receiver:  "2",
				Value:     100.0,
				Status:    StatusPending,
			},
			mockFunc: func(repo *mockTransactionRepo) {
				repo.On("CreateTransaction", &Transaction{
					AccountID: "1",
					Receiver:  "2",
					Value:     100.0,
					Status:    StatusPending,
				}).
					Return(&Transaction{
						ID:        "1",
						AccountID: "1",
						Receiver:  "2",
						Value:     100.0,
						Status:    StatusPending,
					}, nil)
			},
			want: &Transaction{
				ID:        "1",
				AccountID: "1",
				Receiver:  "2",
				Value:     100.0,
				Status:    StatusPending,
			},
			err: nil,
		},
		{
			name: "failed because error in create transaction",
			req: &Transaction{
				AccountID: "1",
				Receiver:  "2",
				Value:     100.0,
				Status:    StatusPending,
			},
			mockFunc: func(repo *mockTransactionRepo) {
				repo.On("CreateTransaction", &Transaction{
					AccountID: "1",
					Receiver:  "2",
					Value:     100.0,
					Status:    StatusPending,
				}).
					Return((*Transaction)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockTransactionRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			got, err := s.CreateTransaction(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestFindTransactionById(t *testing.T) {
	cases := []struct {
		name     string
		req      *TransactionRequest
		mockFunc func(repo *mockTransactionRepo)
		want     *Transaction
		err      error
	}{
		{
			name: "success",
			req: &TransactionRequest{
				TransactionID: "1",
			},
			mockFunc: func(repo *mockTransactionRepo) {
				repo.On("FindTransactionById", "1").
					Return(&Transaction{
						ID:        "1",
						AccountID: "1",
						Receiver:  "2",
						Value:     100.0,
						Status:    StatusPending,
					}, nil)
			},
			want: &Transaction{
				ID:        "1",
				AccountID: "1",
				Receiver:  "2",
				Value:     100.0,
				Status:    StatusPending,
			},
			err: nil,
		},
		{
			name: "failed because error in find transaction by id",
			req: &TransactionRequest{
				TransactionID: "1",
			},
			mockFunc: func(repo *mockTransactionRepo) {
				repo.On("FindTransactionById", "1").
					Return((*Transaction)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockTransactionRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			got, err := s.FindTransactionById(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestListTransactions(t *testing.T) {
	cases := []struct {
		name     string
		req      *ListTransactionRequest
		mockFunc func(repo *mockTransactionRepo)
		want     []*Transaction
		err      error
	}{
		{
			name: "success",
			req: &ListTransactionRequest{
				transactionIDs: []string{"1", "2"},
			},
			mockFunc: func(repo *mockTransactionRepo) {
				repo.On("ListTransactions", []string{"1", "2"}).
					Return([]*Transaction{
						{
							ID:        "1",
							AccountID: "1",
							Receiver:  "2",
							Value:     100.0,
							Status:    StatusPending,
						},
						{
							ID:        "2",
							AccountID: "1",
							Receiver:  "2",
							Value:     200.0,
							Status:    StatusCompleted,
						},
					}, nil)
			},
			want: []*Transaction{
				{
					ID:        "1",
					AccountID: "1",
					Receiver:  "2",
					Value:     100.0,
					Status:    StatusPending,
				},
				{
					ID:        "2",
					AccountID: "1",
					Receiver:  "2",
					Value:     200.0,
					Status:    StatusCompleted,
				},
			},
			err: nil,
		},
		{
			name: "failed because error in list transactions",
			req: &ListTransactionRequest{
				transactionIDs: []string{"1", "2"},
			},
			mockFunc: func(repo *mockTransactionRepo) {
				repo.On("ListTransactions", []string{"1", "2"}).
					Return(([]*Transaction)(nil), errors.New("MOCK-ERROR"))
			},
			want: nil,
			err:  errors.New("MOCK-ERROR"),
		},
		{
			name: "failed because empty transaction IDs",
			req: &ListTransactionRequest{
				transactionIDs: nil,
			},
			mockFunc: func(repo *mockTransactionRepo) {},
			want:     nil,
			err:      errors.New("transaction_ids is required"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockTransactionRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			got, err := s.ListTransactions(c.req)
			assert.Equal(t, err, c.err)
			assert.Equal(t, got, c.want)
		})
	}
}

func TestUpdateTransactionStatus(t *testing.T) {
	cases := []struct {
		name     string
		req      *Transaction
		mockFunc func(repo *mockTransactionRepo)
		err      error
	}{
		{
			name: "success",
			req: &Transaction{
				ID:     "1",
				Status: StatusCompleted,
			},
			mockFunc: func(repo *mockTransactionRepo) {
				repo.On("UpdateTransactionStatus", &Transaction{
					ID:     "1",
					Status: StatusCompleted,
				}).Return(nil)
			},
			err: nil,
		},
		{
			name: "failed because error in update transaction status",
			req: &Transaction{
				ID:     "1",
				Status: StatusFailed,
			},
			mockFunc: func(repo *mockTransactionRepo) {
				repo.On("UpdateTransactionStatus", &Transaction{
					ID:     "1",
					Status: StatusFailed,
				}).Return(errors.New("MOCK-ERROR"))
			},
			err: errors.New("MOCK-ERROR"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := new(mockTransactionRepo)

			c.mockFunc(repo)

			s := NewService(repo)

			err := s.UpdateTransactionStatus(c.req)
			assert.Equal(t, err, c.err)
		})
	}
}
