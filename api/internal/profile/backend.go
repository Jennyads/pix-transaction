package profile

import (
	proto "api/proto/v1"
	"context"
)

type Backend interface {
	Webhook(ctx context.Context, webhook Webhook) error
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

func NewBackend(user proto.UserServiceClient, account proto.AccountServiceClient, keys proto.KeysServiceClient, pix proto.PixTransactionServiceClient) Backend {
	return &grpc{
		user:    user,
		account: account,
		keys:    keys,
		pix:     pix,
	}
}
