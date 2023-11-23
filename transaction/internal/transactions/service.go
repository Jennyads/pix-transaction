package transactions

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	FindTransaction(id *TransactionRequest) (*Transaction, error)
	ListTransactions(transactionIDs *ListTransactionRequest) ([]*Transaction, error)
}

type service struct {
	repo Repository
}

func (s service) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	transaction.CreatedAt = time.Now()
	transaction.ProcessedAt = time.Now()
	transaction.ID = uuid.New().String()
	return s.repo.CreateTransaction(transaction)
}

func (s service) FindTransaction(id *TransactionRequest) (*Transaction, error) {
	return s.repo.FindTransactionById(id.transactionID)
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
