package transactions

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"
	"transaction/internal/event"
)

type Service interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	FindTransaction(id *TransactionRequest) (*Transaction, error)
	ListTransactions(transactionIDs *ListTransactionRequest) ([]*Transaction, error)
}

type service struct {
	repo   Repository
	events event.Client
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

func (s service) Handler(ctx context.Context, payload []byte) ([]byte, error) {
	log.Println(payload)
	return nil, nil
}

func (s service) StartListener(ctx context.Context) error {
	err := s.events.RegisterHandler(ctx, "pix-topic", s.Handler)
	if err != nil {
		return nil
	}
	return nil
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
