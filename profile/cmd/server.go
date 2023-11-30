package main

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"profile/internal/account"
	"profile/internal/key"
	"profile/internal/transaction"
	"profile/internal/user"
	"profile/internal/utils"
	"profile/proto/v1"
)

type ProfileServer struct {
	user               user.Service
	account            account.Service
	keys               key.Service
	transactionService transaction.Service
	v1.UnimplementedUserServiceServer
	v1.UnimplementedAccountServiceServer
	v1.UnimplementedKeysServiceServer
}

func (p ProfileServer) SendPix(ctx context.Context, pixEvent *v1.PixTransaction) (*empty.Empty, error) {
	err := p.transactionService.SendPix(ctx, transaction.ProtoToPix(pixEvent))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil

}

func (p ProfileServer) CreateAccount(ctx context.Context, ac *proto.Account) (*proto.Account, error) {
	if ac.UserId == "" {
		return nil, errors.New("user_id is required")
	}

	created, err := p.account.CreateAccount(ctx, account.ProtoToAccount(ac))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return account.ToProto(created), nil
}

func (p ProfileServer) FindAccount(ctx context.Context, ac *proto.AccountRequest) (*proto.Account, error) {
	if ac.UserId == "" && ac.AccountId == "" {
		return nil, errors.New("id and userId are required")
	}

	found, err := p.account.FindAccountById(ctx, account.ProtoToAccountRequest(ac))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return account.ToProto(found), nil
}
func (p ProfileServer) UpdateAccount(ctx context.Context, request *proto.Account) (*empty.Empty, error) {
	if request.UserId == "" {
		return nil, errors.New("userId is required")
	}

	id := utils.ReadMetadata(ctx, "id")
	if id == "" {
		return nil, errors.New("id is required")
	}

	toUpdate := account.ProtoToAccount(request)
	toUpdate.Id = id
	_, err := p.account.UpdateAccount(ctx, toUpdate)
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (p ProfileServer) ListAccounts(ctx context.Context, request *proto.ListAccountRequest) (*proto.ListAccount, error) {
	if len(request.AccountId) == 0 {
		return nil, errors.New("account_ids is required")
	}
	accounts, err := p.account.ListAccounts(ctx, account.ProtoToAccountListRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	findAccounts := make([]*proto.Account, len(accounts))
	for i := range accounts {
		findAccounts[i] = account.ToProto(accounts[i])
	}
	return &proto.ListAccount{Account: findAccounts}, nil
}
func (p ProfileServer) DeleteAccount(ctx context.Context, request *proto.AccountRequest) (*empty.Empty, error) {
	if request.UserId == "" || request.AccountId == "" {
		return nil, status.Error(codes.InvalidArgument, "account_id and user_id are required")
	}
	err := p.account.DeleteAccount(ctx, account.ProtoToAccountRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}
func (p ProfileServer) CreateUser(ctx context.Context, request *proto.User) (*empty.Empty, error) {
	_, err := p.user.CreateUser(user.ProtoToUser(request))
	if err != nil {
		switch err.(type) {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (p ProfileServer) FindUser(ctx context.Context, request *proto.UserRequest) (*proto.User, error) {
	find, err := p.user.FindUserById(user.ProtoToUserRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return user.ToProto(find), nil
}

func (p ProfileServer) UpdateUser(ctx context.Context, request *proto.User) (*empty.Empty, error) {
	id := utils.ReadMetadata(ctx, "id")
	if id == "" {
		return nil, errors.New("id is required")
	}

	toUpdate := user.ProtoToUser(request)
	toUpdate.Id = id
	_, err := p.user.UpdateUser(toUpdate)
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (p ProfileServer) ListUsers(ctx context.Context, request *proto.ListUserRequest) (*proto.ListUser, error) {
	if len(request.Id) == 0 {
		return nil, errors.New("user_ids is required")
	}
	users, err := p.user.ListUsers(user.ProtoToUserListRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	findUsers := make([]*proto.User, len(users))
	for i := range users {
		findUsers[i] = user.ToProto(users[i])
	}
	return &proto.ListUser{Users: findUsers}, nil
}

func (p ProfileServer) DeleteUser(ctx context.Context, request *proto.UserRequest) (*empty.Empty, error) {
	if request.Id == "" {
		return nil, errors.New("account_id is required")
	}
	err := p.user.DeleteUser(user.ProtoToUserRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (p ProfileServer) CreateKey(ctx context.Context, req *proto.Key) (*empty.Empty, error) {
	_, err := p.keys.CreateKey(key.ProtoToKey(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (p ProfileServer) UpdateKey(ctx context.Context, req *proto.Key) (*empty.Empty, error) {
	id := utils.ReadMetadata(ctx, "id")
	if id == "" {
		return nil, errors.New("id is required")
	}

	toUpdate := key.ProtoToKey(req)
	toUpdate.Id = id
	_, err := p.keys.UpdateKey(toUpdate)
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (p ProfileServer) ListKey(ctx context.Context, req *proto.ListKeyRequest) (*proto.ListKeys, error) {
	keys, err := p.keys.ListKey(key.ProtoToKeyListRequest(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	foundKeys := make([]*proto.Key, len(keys))
	for i := range keys {
		foundKeys[i] = key.ToProto(keys[i])
	}
	return &proto.ListKeys{Keys: foundKeys}, nil
}

func (p ProfileServer) DeleteKey(ctx context.Context, req *proto.KeyRequest) (*empty.Empty, error) {
	err := p.keys.DeleteKey(key.ProtoToKeyRequest(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func NewProfileService(userService user.Service, accountService account.Service, keyService key.Service, service transaction.Service) *ProfileServer {
	return &ProfileServer{
		user:    userService,
		account: accountService,
		keys:    keyService,
	}
}
