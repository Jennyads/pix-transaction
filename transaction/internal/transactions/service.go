package transactions

import (
	"errors"
	"github.com/google/uuid"
	"time"
	"transaction/internal/event"
)

type Service interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	FindTransactionById(id *TransactionRequest) (*Transaction, error)
	ListTransactions(transactionIDs *ListTransactionRequest) ([]*Transaction, error)
	UpdateTransactionStatus(transaction *Transaction) error
}

type service struct {
	repo   Repository
	events event.Client
}

func (s service) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	transaction.CreatedAt = time.Now()
	transaction.ID = uuid.New().String()
	return s.repo.CreateTransaction(transaction)
}

func (s service) FindTransactionById(id *TransactionRequest) (*Transaction, error) {
	return s.repo.FindTransactionById(id.TransactionID)
}

func (s service) ListTransactions(req *ListTransactionRequest) ([]*Transaction, error) {
	if len(req.transactionIDs) == 0 {
		return nil, errors.New("transaction_ids is required")
	}
	return s.repo.ListTransactions(req.transactionIDs)
}

func (s service) UpdateTransactionStatus(transaction *Transaction) error {
	transaction.UpdatedAt = time.Now()
	return s.repo.UpdateTransactionStatus(transaction)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
