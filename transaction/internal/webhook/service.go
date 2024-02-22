package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"transaction/internal/transactions"
)

type Service interface {
	Send(ctx context.Context, data Webhook, url string) error
}

type TransactionRepository interface {
	FindTransactionById(id string) (*transactions.Transaction, error)
	UpdateTransactionStatus(transaction *transactions.Transaction) error
}

type service struct {
	client      http.Client
	transaction TransactionRepository
}

func (s *service) Send(ctx context.Context, data Webhook, url string) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return err
	}

	hreq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return err
	}

	hreq.Header.Set("Content-Type", "application/json")

	result, err := s.client.Do(hreq)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	if result.StatusCode < 200 || result.StatusCode >= 300 {
		return errors.New("error when send webhook")
	}

	if data.Status == StatusCompleted {
		transaction, err := s.transaction.FindTransactionById(data.TransactionId)
		if err != nil {
			return err
		}
		transaction.ProcessedAt = time.Now()
		err = s.transaction.UpdateTransactionStatus(transaction)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewService(transactionRepo TransactionRepository) Service {
	return &service{
		client:      http.Client{},
		transaction: transactionRepo,
	}
}
