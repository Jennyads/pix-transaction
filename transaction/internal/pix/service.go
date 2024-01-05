package pix

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"transaction/internal/errutils"
	keys "transaction/internal/keys"
	"transaction/internal/profile"
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

	err = s.TransactionWorkflow(ctx, &pixEvent)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *service) TransactionWorkflow(ctx context.Context, pixEvent *PixEvent) error {
	receiver, err := s.keysRepo.FindKey(ctx, pixEvent.PixData.Receiver)
	if err != nil {
		if errors.Is(errutils.ErrKeyNotFound, err) {
			log.Println("error key not found")
			return errutils.ErrInvalidKey
		}
		return nil
	}

	sender, err := s.profile.FindAccount(ctx, pixEvent.PixData.AccountId)
	if err != nil {
		return err
	}

	transaction, err := s.transaction.CreateTransaction(&transactions.Transaction{
		AccountID: pixEvent.PixData.AccountId,
		Receiver:  receiver.Id,
		Value:     utils.ToFloat(pixEvent.PixData.Amount),
		Status:    transactions.StatusPending,
	})
	if err != nil {
		return err
	}

	//err = s.Validations(ctx, receiver, sender, pixEvent.PixData)
	//if err != nil {
	//	transaction.Status = transactions.StatusFailed
	//	transaction.ErrMessage = err.Error()
	//	err = s.transaction.UpdateTransactionStatus(transaction)
	//	if err != nil {
	//		return err
	//	}
	//	return err
	//}

	err = s.Transaction(ctx, receiver, sender, pixEvent.PixData)
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

func (s *service) Validations(ctx context.Context, receiver *profile.Account, sender *profile.Account, pix *Pix) error {
	isActive, err := s.profile.IsAccountActive(ctx, pix.AccountId)
	if err != nil {
		return err
	}

	if !isActive {
		return errutils.ErrInactiveAccount
	}

	if sender.Balance.LessThan(pix.Amount) {
		return errutils.ErrInsufficientBalance
	}

	if receiver.BlockedAt == nil {
		return errutils.ErrReceiverAccountBlocked
	}

	return nil
}

func (s *service) Transaction(ctx context.Context, receiver *profile.Account, sender *profile.Account, pix *Pix) error {
	err := s.webhook.Send(ctx, &webhook.Webhook{})
	if err != nil {
		return err
	}
	return nil
}

func NewService(transaction transactions.Service, profile profile.Service) Service {
	return &service{
		transaction: transaction,
		profile:     profile,
	}
}
