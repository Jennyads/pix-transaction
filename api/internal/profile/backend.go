package profile

import (
	proto "api/proto/v1"
	"context"
)

type Backend interface {
	Webhook(ctx context.Context, webhook Webhook) error
	CreateUser(ctx context.Context, user User) error
	ListUsers(ctx context.Context, user []string) ([]*User, error)

	FindAccount(ctx context.Context, account Account) error
	PixWebhook(ctx context.Context, webhook Webhook) error
	CreateAccount(ctx context.Context, account Account) error
	SendPix(ctx context.Context, pix PixTransaction) error
}

type grpc struct {
	user    proto.UserServiceClient
	account proto.AccountServiceClient
	keys    proto.KeysServiceClient
	pix     proto.PixTransactionServiceClient
}

func (g *grpc) Webhook(ctx context.Context, webhook Webhook) error {
	_, err := g.pix.PixWebhook(ctx, &proto.Webhook{
		Sender: &proto.WebhookAccount{
			Agency: webhook.Sender.Agency, Bank: webhook.Sender.Bank, Name: webhook.Sender.Name},
		Receiver: &proto.WebhookAccount{
			Agency: webhook.Receiver.Agency, Bank: webhook.Receiver.Bank, Name: webhook.Receiver.Name},
		Amount: webhook.Amount,
		Status: proto.Status(proto.Status_value[string(webhook.Status)]),
	})
	if err != nil {
		return err
	}
	return nil
}
func (g *grpc) CreateUser(ctx context.Context, user User) error {
	_, err := g.user.CreateUser(ctx, &proto.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Cpf:     user.Cpf,
		Phone:   user.Phone,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *grpc) ListUsers(ctx context.Context, user []string) ([]*User, error) {
	_, err := g.user.ListUsers(ctx, &proto.ListUserRequest{
		Id: user,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (g *grpc) CreateAccount(ctx context.Context, account Account) error {
	_, err := g.account.CreateAccount(ctx, &proto.Account{
		UserId: account.Name,
		Agency: account.Agency,
		Bank:   account.Bank,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *grpc) FindAccount(ctx context.Context, account Account) error {
	_, err := g.account.FindAccount(ctx, &proto.AccountRequest{
		UserId: account.Name,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *grpc) PixWebhook(ctx context.Context, webhook Webhook) error {
	_, err := g.pix.PixWebhook(ctx, &proto.Webhook{
		Sender: &proto.WebhookAccount{
			Agency: webhook.Sender.Agency, Bank: webhook.Sender.Bank, Name: webhook.Sender.Name},
		Receiver: &proto.WebhookAccount{
			Agency: webhook.Receiver.Agency, Bank: webhook.Receiver.Bank, Name: webhook.Receiver.Name},
		Amount: webhook.Amount,
		Status: proto.Status(proto.Status_value[string(webhook.Status)]),
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *grpc) SendPix(ctx context.Context, pix PixTransaction) error {
	_, err := g.pix.SendPix(ctx, &proto.PixTransaction{
		UserId:      pix.UserId,
		SenderKey:   pix.SenderKey,
		ReceiverKey: pix.ReceiverKey,
		Amount:      pix.Amount,
		Hour:        pix.Hour,
		Status:      string(proto.Status(proto.Status_value[string(pix.Status)])),
	})
	if err != nil {
		return err
	}
	return nil
}
func NewBackend(user proto.UserServiceClient, account proto.AccountServiceClient, keys proto.KeysServiceClient, pix proto.PixTransactionServiceClient) Backend {
	return &grpc{
		user:    user,
		account: account,
		keys:    keys,
		pix:     pix,
	}
}
