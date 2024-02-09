package transaction

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"profile/internal/account"
	"profile/internal/errutils"
	"profile/internal/event"
	"profile/internal/user"
	transpb "profile/proto/transactions/v1"
)

type Service interface {
	SendPix(ctx context.Context, req *Pix) error
	PixWebhook(ctx context.Context, req *Webhook) error
	CreateKey(ctx context.Context, req *Key) error
}

type service struct {
	accountRepository account.Repository
	userRepository    user.Repository
	events            event.Client
	keysBackend       transpb.KeysServiceClient
}

func (s service) SendPix(ctx context.Context, req *Pix) error {

	accountModel, err := s.accountRepository.FindAccountById(req.AccountID)
	if err != nil {
		return err
	}

	userModel, err := s.userRepository.FindUserById(accountModel.UserID)
	if err != nil {
		return err
	}

	if accountModel == nil {
		return errutils.ErrInactiveAccount
	}

	if accountModel.Balance.LessThan(req.Amount) {
		return errutils.ErrInsufficientBalance
	}

	if accountModel.BlockedAt != nil {
		return errutils.ErrReceiverAccountBlocked
	}

	payload, err := json.Marshal(PixEvent{
		AccountName:   accountModel.Id,
		AccountCpf:    userModel.Cpf,
		AccountAgency: accountModel.Agency,
		AccountBank:   accountModel.Bank,
		Receiver:      req.Receiver,
		Amount:        req.Amount,
		WebhookUrl:    "http://localhost:9060/profile/v1/webhook",
	})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err = s.events.Publish(ctx, payload); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s service) PixWebhook(ctx context.Context, req *Webhook) error {

	sender, err := s.accountRepository.FindAccountById(req.Sender.Name)
	if err != nil {
		return err
	}

	receiver, err := s.accountRepository.FindAccountById(req.Receiver.Name)
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

func (s service) CreateKey(ctx context.Context, req *Key) error {
	accountModel, err := s.accountRepository.FindAccountById(req.Account)
	if err != nil {
		return err
	}

	userModel, err := s.userRepository.FindUserById(accountModel.UserID)
	if err != nil {
		return err
	}

	_, err = s.keysBackend.CreateKey(ctx, &transpb.Key{
		Account: &transpb.Account{
			Cpf:    userModel.Cpf,
			Name:   accountModel.Id,
			Agency: accountModel.Agency,
			Bank:   accountModel.Bank,
		},
		Name: req.Name,
		Type: transpb.Type(transpb.Type_value[string(req.Type)]),
	})
	if err != nil {
		return err
	}
	return nil
}

func NewService(event event.Client, accountRepository account.Repository, keysBackend transpb.KeysServiceClient, userRepository user.Repository) Service {
	return &service{
		events:            event,
		accountRepository: accountRepository,
		keysBackend:       keysBackend,
		userRepository:    userRepository,
	}
}
