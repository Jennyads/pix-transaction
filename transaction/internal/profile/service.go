package profile

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"transaction/internal/cfg"
	"transaction/proto"
)

type Service interface {
	FindAccountBalance(ctx context.Context, id string, userId string) (float64, error)
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

func NewUserService(config cfg.Config) Service {
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
