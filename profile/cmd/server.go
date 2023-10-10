package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"profile/internal/account"
	"profile/internal/keys"
	"profile/internal/user"
	v1 "profile/proto/v1"
)

type ProfileServer struct {
	user    user.Service
	account account.Service
	keys    keys.Service
	v1.UnimplementedUserServiceServer
	v1.UnimplementedAccountServiceServer
	v1.UnimplementedKeysServiceServer
}

func (p ProfileServer) CreateAccount(ctx context.Context, ac *v1.Account) (*empty.Empty, error) {
	_, err := p.account.CreateAccount(account.ProtoToAccount(ac))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) FindAccount(ctx context.Context, account *v1.Account) (*v1.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) UpdateAccount(ctx context.Context, account *v1.Account) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) ListAccounts(ctx context.Context, request *v1.ListAccountRequest) (*v1.ListAccount, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) DeleteAccount(ctx context.Context, request *v1.AccountRequest) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) CreateUser(ctx context.Context, user *v1.User) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) FindUser(ctx context.Context, request *v1.UserRequest) (*v1.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) UpdateUser(ctx context.Context, user *v1.User) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) ListUsers(ctx context.Context, request *v1.ListUserRequest) (*v1.ListUser, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) DeleteUser(ctx context.Context, request *v1.UserRequest) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) CreateKey(ctx context.Context, keys *v1.Keys) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) UpdateKey(ctx context.Context, keys *v1.Keys) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) ListKey(ctx context.Context, keys *v1.Keys) (*v1.ListKeys, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProfileServer) DeleteKey(ctx context.Context, keys *v1.Keys) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func NewProfileService(userService user.Service, accountService account.Service, keyService keys.Service) *ProfileServer {
	return &ProfileServer{}
}
