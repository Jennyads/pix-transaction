package main

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"profile/internal/account"
	"profile/internal/keys"
	"profile/internal/user"
	"profile/internal/utils"
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

func (p ProfileServer) CreateAccount(ctx context.Context, ac *v1.Account) (*v1.Account, error) {
	if ac.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	_, err := p.account.CreateAccount(ctx, account.ProtoToAccount(ac))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return nil, nil
}

func (p ProfileServer) FindAccount(ctx context.Context, ac *v1.Account) (*v1.Account, error) {
	if ac.UserId == 0 {
		return nil, errors.New("id is required")
	}

	_, err := p.account.CreateAccount(ctx, account.ProtoToAccount(ac))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}
func (p ProfileServer) UpdateAccount(ctx context.Context, request *v1.Account) (*empty.Empty, error) {
	if request.UserId == 0 {
		return nil, errors.New("id is required")
	}
	_, err := p.account.UpdateAccount(ctx, account.ProtoToAccount(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) ListAccounts(ctx context.Context, request *v1.ListAccountRequest) (*v1.ListAccount, error) {
	if len(request.AccountId) == 0 {
		return nil, errors.New("account_ids is required")
	}
	_, err := p.account.ListAccounts(ctx, account.ProtoToAccountListRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) DeleteAccount(ctx context.Context, request *v1.AccountRequest) (*empty.Empty, error) {
	if request.AccountId == 0 {
		return nil, errors.New("account_id is required")
	}
	err := p.account.DeleteAccount(ctx, account.ProtoToAccountRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) CreateUser(ctx context.Context, request *v1.User) (*empty.Empty, error) {
	_, err := p.user.CreateUser(user.ProtoToUser(request))
	if err != nil {
		switch err.(type) {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) FindUser(ctx context.Context, request *v1.UserRequest) (*v1.User, error) {
	_, err := p.user.FindUserById(user.ProtoToUserRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) UpdateUser(ctx context.Context, request *v1.User) (*empty.Empty, error) {
	id := utils.ReadMetadata(ctx, "id")
	if id == "" {
		return nil, errors.New("id is required")
	}

	_, err := p.user.UpdateUser(user.ProtoToUser(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) ListUsers(ctx context.Context, request *v1.ListUserRequest) (*v1.ListUser, error) {
	if len(request.Id) == 0 {
		return nil, errors.New("user_ids is required")
	}
	_, err := p.user.ListUsers(user.ProtoToUserListRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) DeleteUser(ctx context.Context, request *v1.UserRequest) (*empty.Empty, error) {
	if request.Id == 0 {
		return nil, errors.New("account_id is required")
	}
	err := p.user.DeleteUser(user.ProtoToUserRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) CreateKey(ctx context.Context, req *v1.Key) (*empty.Empty, error) {
	_, err := p.keys.CreateKey(keys.ProtoToKey(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) UpdateKey(ctx context.Context, req *v1.Key) (*empty.Empty, error) {
	_, err := p.keys.UpdateKey(keys.ProtoToKey(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) ListKey(ctx context.Context, req *v1.ListKeyRequest) (*v1.ListKeys, error) {
	_, err := p.keys.ListKey(keys.ProtoToKeyListRequest(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (p ProfileServer) DeleteKey(ctx context.Context, req *v1.KeyRequest) (*empty.Empty, error) {
	err := p.keys.DeleteKey(keys.ProtoToKeyRequest(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func NewProfileService(userService user.Service, accountService account.Service, keyService keys.Service) *ProfileServer {
	return &ProfileServer{}
}
