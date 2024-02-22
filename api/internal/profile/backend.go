package profile

import (
	proto "api/proto/v1"
	"context"
)

type Backend interface {
	Webhook(ctx context.Context, webhook Webhook) error
	CreateUser(ctx context.Context, user User) (string, error)
	ListUsers(ctx context.Context, user []string) ([]*User, error)

	FindAccount(ctx context.Context, userId string) error
	PixWebhook(ctx context.Context, webhook Webhook) error
	CreateAccount(ctx context.Context, userId string, account Account) (string, error)
	SendPix(ctx context.Context, pix PixTransaction) error
	CreateKey(ctx context.Context, key Key) error
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
			Name: webhook.Sender.Name, Agency: webhook.Sender.Agency, Bank: webhook.Sender.Bank},
		Receiver: &proto.WebhookAccount{
			Name:   webhook.Receiver.Name,
			Agency: webhook.Receiver.Agency, Bank: webhook.Receiver.Bank},
		Amount: webhook.Amount.InexactFloat64(),
		Status: proto.Status(proto.Status_value[string(webhook.Status)]),
	})
	if err != nil {
		return err
	}
	return nil
}
func (g *grpc) CreateUser(ctx context.Context, user User) (string, error) {
	response, err := g.user.CreateUser(ctx, &proto.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Cpf:     user.Cpf,
		Phone:   user.Phone,
	})
	if err != nil {
		return "", err
	}
	return response.Id, nil
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

var lastAccountId int64 = 0

func (g *grpc) CreateAccount(ctx context.Context, userId string, account Account) (string, error) {
	account.Balance = 100.0
	lastAccountId++
	newAccountId := lastAccountId
	response, err := g.account.CreateAccount(ctx, &proto.Account{
		Id:      newAccountId,
		UserId:  userId,
		Agency:  account.Agency,
		Bank:    account.Bank,
		Balance: account.Balance,
	})
	if err != nil {
		return "", err
	}
	return response.Id, nil
}

func (g *grpc) FindAccount(ctx context.Context, userId string) error {
	_, err := g.account.FindAccount(ctx, &proto.AccountRequest{
		UserId: userId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *grpc) PixWebhook(ctx context.Context, webhook Webhook) error {
	_, err := g.pix.PixWebhook(ctx, &proto.Webhook{
		Sender: &proto.WebhookAccount{
			Agency: webhook.Sender.Agency, Bank: webhook.Sender.Bank},
		Receiver: &proto.WebhookAccount{
			Agency: webhook.Receiver.Agency, Bank: webhook.Receiver.Bank},
		Amount: webhook.Amount.InexactFloat64(),
		Status: proto.Status(proto.Status_value[string(webhook.Status)]),
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *grpc) SendPix(ctx context.Context, pix PixTransaction) error {
	_, err := g.pix.SendPix(ctx, &proto.PixTransaction{
		ReceiverKey: pix.ReceiverKey,
		Amount:      pix.Amount,
		AccountId:   pix.AccountId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *grpc) CreateKey(ctx context.Context, key Key) error {
	_, err := g.keys.CreateKey(ctx, &proto.Key{
		AccountId: key.AccountId,
		Name:      key.Name,
		Type:      proto.Type(proto.Type_value[string(key.Type)]),
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
