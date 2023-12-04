package profile

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"transaction/internal/cfg"
	proto "transaction/proto/v1"
)

type Service interface {
	FindAccountBalance(ctx context.Context, id string, userId string) (float64, error)
	FindKey(ctx context.Context, key string, accountId string) (float64, error)
	IsAccountActive(ctx context.Context, id string) (bool, error)
}

type service struct {
	account proto.AccountServiceClient
	user    proto.UserServiceClient
	keys    proto.KeysServiceClient
}

func (s *service) FindAccountBalance(ctx context.Context, id string, userId string) (float64, error) {
	account, err := s.account.FindAccount(ctx, &proto.AccountRequest{AccountId: id, UserId: userId})
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

// FindKey encontra a chave do usuário para buscar usuário a ser enviado o pix
func (s *service) FindKey(ctx context.Context, key string, accountId string) (float64, error) {
	keys, err := s.keys.FindKey(ctx, &proto.KeyRequest{KeyId: key})
	if err != nil {
		return 0, err
	}
	return keys.Value, nil
}

func (s *service) IsAccountActive(ctx context.Context, id string) (bool, error) {
	_, err := s.account.IsAccountActive(ctx, &proto.AccountRequest{AccountId: id})
	if err != nil {
		return false, err
	}
	return true, nil
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
	}
}
