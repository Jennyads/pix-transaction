package pix

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"transaction/internal/errutils"
	"transaction/internal/event"
	"transaction/internal/profile"
	"transaction/internal/transactions"
	"transaction/internal/utils"
)

type Service interface {
	Handler(ctx context.Context, payload []byte) ([]byte, error)
}

type service struct {
	transaction transactions.Service
	profile     profile.Service
	events      event.Client
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
	receiver, err := s.profile.FindReceiver(ctx, pixEvent.PixData.Key)
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

	err = s.Validations(ctx, receiver, sender, pixEvent.PixData)
	if err != nil {
		transaction.Status = transactions.StatusFailed
		transaction.ErrMessage = err.Error()
		err = s.transaction.UpdateTransactionStatus(transaction)
		if err != nil {
			return err
		}
		return err
	}

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

	err = s.profile.SendWebhook(ctx, &profile.Webhook{
		AccountID:  sender.Id,
		ReceiverID: receiver.Id,
		Amount:     pixEvent.PixData.Amount,
		Status:     profile.StatusCompleted,
	})
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
	receiver.Balance = receiver.Balance.Add(pix.Amount)
	err := s.profile.UpdateAccountBalance(ctx, receiver)
	if err != nil {
		return err
	}

	sender.Balance = sender.Balance.Sub(pix.Amount)
	err = s.profile.UpdateAccountBalance(ctx, sender)
	if err != nil {
		return err
	}

	return nil
}
