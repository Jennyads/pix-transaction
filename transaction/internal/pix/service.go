package pix

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"transaction/internal/errutils"
	keys "transaction/internal/keys"
	"transaction/internal/transactions"
	"transaction/internal/utils"
	"transaction/internal/webhook"
)

type Service interface {
	Handler(ctx context.Context, payload []byte) ([]byte, error)
}

type service struct {
	transaction transactions.Service
	keysRepo    keys.Repository
	webhook     webhook.Service
}

func (s *service) Handler(ctx context.Context, payload []byte) ([]byte, error) {

	var pixEvent PixEvent
	err := json.Unmarshal(payload, &pixEvent)
	if err != nil {
		log.Println("error unmarshalling pix event")
		return nil, err
	}

	err = s.Transaction(ctx, &pixEvent)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *service) Transaction(ctx context.Context, pixEvent *PixEvent) error {
	receiver, err := s.keysRepo.FindKey(ctx, pixEvent.PixData.Receiver)
	if err != nil {
		if errors.Is(errutils.ErrKeyNotFound, err) {
			log.Println("error key not found")
			return errutils.ErrInvalidKey
		}
		return nil
	}

	transaction, err := s.transaction.CreateTransaction(&transactions.Transaction{
		AccountID: pixEvent.PixData.AccountName,
		Receiver:  receiver.Id,
		Value:     utils.ToFloat(pixEvent.PixData.Amount),
		Status:    transactions.StatusPending,
	})
	if err != nil {
		return err
	}

	err = s.webhook.Send(ctx, webhook.Webhook{
		Sender: webhook.Account{
			Name:   pixEvent.PixData.AccountName,
			Agency: pixEvent.PixData.AccountAgency,
			Bank:   pixEvent.PixData.AccountBank,
			Cpf:    pixEvent.PixData.AccountCpf,
		},
		Receiver: webhook.Account{
			Name:   receiver.Name,
			Agency: receiver.Agency,
			Bank:   receiver.Bank,
			Cpf:    receiver.Cpf,
		},
		Amount: pixEvent.PixData.Amount,
		Status: webhook.StatusCompleted,
	}, pixEvent.PixData.WebhookUrl)
	if err != nil {
		transaction.Status = transactions.StatusFailed
		transaction.ErrMessage = err.Error()
		err = s.transaction.UpdateTransactionStatus(transaction)
		if err != nil {
			return err
		}
		return err
	}

	transaction.Status = transactions.StatusCompleted
	err = s.transaction.UpdateTransactionStatus(transaction)
	if err != nil {
		return err
	}

	return nil
}

func NewService(transaction transactions.Service) Service {
	return &service{
		transaction: transaction,
	}
}
