package transaction

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"profile/internal/account"
	"profile/internal/event"
)

type Service interface {
	SendPix(ctx context.Context, req *Pix) error
	PixWebhook(ctx context.Context, req *Webhook) error
}

type service struct {
	accountRepository account.Repository
	events            event.Client
}

func (s service) SendPix(ctx context.Context, req *Pix) error {
	pixEvent := PixEvent{
		PixData: req,
	}

	payload, err := json.Marshal(pixEvent)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := s.events.Publish(ctx, payload); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s service) PixWebhook(ctx context.Context, req *Webhook) error {

	sender, err := s.accountRepository.FindAccountById(req.Sender.Name)
	if err != nil {
		return err
	}

	receiver, err := s.accountRepository.FindAccountById(req.Sender.Name)
	if err != nil {
		return err
	}

	switch req.Status {
	case StatusFailed:
		return nil
	case StatusCompleted:
		sender.Balance = sender.Balance.Sub(req.Amount)
		receiver.Balance = receiver.Balance.Add(req.Amount)
		err = s.accountRepository.UpdateAccounts([]*account.Account{sender, receiver})
		if err != nil {
			return err
		}
	}

	return nil
}

func NewService(event event.Client, accountRepository account.Repository) Service {
	return &service{
		events:            event,
		accountRepository: accountRepository,
	}
}
