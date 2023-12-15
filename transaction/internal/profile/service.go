package profile

import (
	"context"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"transaction/internal/cfg"
	"transaction/internal/errutils"
	"transaction/internal/utils"
	proto "transaction/proto/v1"
)

type Service interface {
	FindAccount(ctx context.Context, id string) (*Account, error)
	FindReceiver(ctx context.Context, key string) (*Account, error)
	IsAccountActive(ctx context.Context, id string) (bool, error)
	UpdateAccountBalance(ctx context.Context, account *Account) error
	SendWebhook(ctx context.Context, webhook *Webhook) error
}

type service struct {
	account proto.AccountServiceClient
	user    proto.UserServiceClient
	keys    proto.KeysServiceClient
	pix     proto.PixTransactionServiceClient
}

func (s *service) UpdateAccountBalance(ctx context.Context, account *Account) error {
	ctxWithMetadata := metadata.NewOutgoingContext(ctx, metadata.Pairs("id", account.Id))
	_, err := s.account.UpdateAccount(ctxWithMetadata, &proto.Account{UserId: account.UserID,
		Balance: utils.ToFloat(account.Balance), Agency: account.Agency, Bank: account.Bank})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) FindAccount(ctx context.Context, id string) (*Account, error) {
	account, err := s.account.FindAccount(ctx, &proto.AccountRequest{AccountId: id})
	if err != nil {
		return nil, err
	}
	return &Account{
		Id:      id,
		UserID:  account.UserId,
		Balance: decimal.NewFromFloat(account.Balance),
		Agency:  account.Agency,
		Bank:    account.Bank,
	}, nil
}

func (s *service) FindReceiver(ctx context.Context, key string) (*Account, error) {
	account, err := s.account.FindByKey(ctx, &proto.FindByKeyRequest{Key: key})
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			return nil, errutils.ErrInvalidKey
		}

		return nil, err
	}
	return ProtoToAccount(account), nil
}

func (s *service) IsAccountActive(ctx context.Context, id string) (bool, error) {
	isActive, err := s.account.IsAccountActive(ctx, &proto.AccountRequest{AccountId: id})
	if err != nil {
		return false, err
	}
	return isActive.GetValue(), nil
}

func (s *service) SendWebhook(ctx context.Context, webhook *Webhook) error {
	_, err := s.pix.PixWebhook(ctx, &proto.Webhook{AccountId: webhook.AccountID, ReceiverId: webhook.ReceiverID,
		Amount: utils.ToFloat(webhook.Amount), Status: proto.Status(proto.Status_value[string(webhook.Status)])})
	if err != nil {
		return err
	}
	return nil
}

func NewService(config cfg.Config) Service {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(config.ProfileConfig.Host, opts...)
	if err != nil {
		panic(err)
	}

	return &service{
		account: proto.NewAccountServiceClient(conn),
		user:    proto.NewUserServiceClient(conn),
		keys:    proto.NewKeysServiceClient(conn),
		pix:     proto.NewPixTransactionServiceClient(conn),
	}
}
