package pix

import (
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	CreatePixTransaction(transaction *PixTransaction) (*PixTransaction, error)
	ListPixTransactions(request *ListPixTransactionsRequest) ([]*PixTransaction, error)
}
type service struct {
	repo Repository
}

func (s service) CreatePixTransaction(pixTransaction *PixTransaction) (*PixTransaction, error) {
	pixTransaction.ID = uuid.New().String()
	pixTransaction.Timestamp = &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
	}
	return s.repo.CreatePixTransaction(pixTransaction)
}

func (s service) ListPixTransactions(request *ListPixTransactionsRequest) ([]*PixTransaction, error) {
	if len(request.PixTransactionIDs) == 0 {
		return nil, errors.New("PixTransactionIDs is required")
	}
	return s.repo.ListPixTransactions(request)
}
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
