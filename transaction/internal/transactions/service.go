package transactions

import (
	"errors"
	"time"
)

type Service interface {
	CreateTransaction(transaction *Transaction) error
	FindTransaction(id *TransactionRequest) (*Transaction, error)
	ListTransactions(transactionIDs *ListTransactionRequest) ([]*Transaction, error)
}

type service struct {
	repo Repository
}

func (s service) CreateTransaction(transaction *Transaction) error {
	transaction.CreatedAt = time.Now()
	transaction.ProcessedAt = time.Now()
	return s.repo.CreateTransaction(transaction)
}

func (s service) FindTransaction(id *TransactionRequest) (*Transaction, error) {
	return s.repo.FindTransaction(id.transactionID)
}

func (s service) ListTransactions(req *ListTransactionRequest) ([]*Transaction, error) {
	if len(req.transactionIDs) == 0 {
		return nil, errors.New("transaction_ids is required")
	}
	return s.repo.ListTransactions(req.transactionIDs)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
