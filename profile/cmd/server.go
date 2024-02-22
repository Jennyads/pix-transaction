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
	"profile/proto/profile/v1"
)

type ProfileServer struct {
	user               user.Service
	account            account.Service
	keys               key.Service
	transactionService transaction.Service
	profile.UnimplementedUserServiceServer
	profile.UnimplementedAccountServiceServer
	profile.UnimplementedKeysServiceServer
	profile.UnimplementedPixTransactionServiceServer
}

type EmailAlreadyExist struct {
	Message string
}

func (e *EmailAlreadyExist) Error() string {
	return e.Message
}

type CpfAlreadyExist struct {
	Message string
}

func (e *CpfAlreadyExist) Error() string {
	return e.Message
}

type PhoneAlreadyExist struct {
	Message string
}

func (e *PhoneAlreadyExist) Error() string {
	return e.Message
}

func (p ProfileServer) SendPix(ctx context.Context, pixEvent *profile.PixTransaction) (*empty.Empty, error) {
	err := p.transactionService.SendPix(ctx, transaction.ProtoToPix(pixEvent))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (p ProfileServer) PixWebhook(ctx context.Context, webhook *profile.Webhook) (*empty.Empty, error) {
	err := p.transactionService.PixWebhook(ctx, transaction.ProtoToWebhook(webhook))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (p ProfileServer) CreateAccount(ctx context.Context, ac *profile.Account) (*profile.AccountResponse, error) {
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

func (p ProfileServer) FindAccount(ctx context.Context, ac *profile.AccountRequest) (*profile.AccountResponse, error) {
	if ac.UserId == "" && ac.AccountId == 0 {
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
func (p ProfileServer) UpdateAccount(ctx context.Context, request *profile.Account) (*empty.Empty, error) {
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

func (p ProfileServer) ListAccounts(ctx context.Context, request *profile.ListAccountRequest) (*profile.ListAccount, error) {
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

	findAccounts := make([]*profile.AccountResponse, len(accounts))
	for i := range accounts {
		findAccounts[i] = account.ToProto(accounts[i])
	}
	return &profile.ListAccount{Account: findAccounts}, nil
}

func (p ProfileServer) DeleteAccount(ctx context.Context, request *profile.AccountRequest) (*empty.Empty, error) {
	if request.UserId == "" || request.AccountId == 0 {
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

func (p ProfileServer) FindByKey(ctx context.Context, request *profile.FindByKeyRequest) (*profile.AccountResponse, error) {
	found, err := p.account.FindByKey(ctx, request.Key)
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return account.ToProto(found), nil
}

func (p ProfileServer) CreateUser(ctx context.Context, request *profile.User) (*profile.UserResponse, error) {
	response, err := p.user.CreateUser(user.ProtoToUser(request))
	if err != nil {
		switch err.(type) {
		case *EmailAlreadyExist:
			return nil, status.Error(codes.AlreadyExists, "email already exist")
		case *CpfAlreadyExist:
			return nil, status.Error(codes.AlreadyExists, "cpf already exist")
		case *PhoneAlreadyExist:
			return nil, status.Error(codes.AlreadyExists, "phone already exist")
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return user.ToProtoUserResponse(response), nil
}

func (p ProfileServer) FindUser(ctx context.Context, request *profile.UserRequest) (*profile.UserResponse, error) {
	found, err := p.user.FindUserById(user.ProtoToUserRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return user.ToProtoUserResponse(found), nil
}

func (p ProfileServer) UpdateUser(ctx context.Context, request *profile.User) (*empty.Empty, error) {
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

func (p ProfileServer) ListUsers(ctx context.Context, request *profile.ListUserRequest) (*profile.ListUser, error) {
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

	findUsers := make([]*profile.User, len(users))
	for i := range users {
		findUsers[i] = user.ToProto(users[i])
	}
	return &profile.ListUser{Users: findUsers}, nil
}

func (p ProfileServer) DeleteUser(ctx context.Context, request *profile.UserRequest) (*empty.Empty, error) {
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

func (p ProfileServer) CreateKey(ctx context.Context, req *profile.Key) (*profile.KeyResponse, error) {
	created, err := p.keys.CreateKey(key.ProtoToKey(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return key.ToProto(created), nil
}

func (p ProfileServer) UpdateKey(ctx context.Context, req *profile.Key) (*empty.Empty, error) {
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

func (p ProfileServer) ListKey(ctx context.Context, req *profile.ListKeyRequest) (*profile.ListKeys, error) {
	keys, err := p.keys.ListKey(key.ProtoToKeyListRequest(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	foundKeys := make([]*profile.KeyResponse, len(keys))
	for i := range keys {
		foundKeys[i] = key.ToProto(keys[i])
	}
	return &profile.ListKeys{Keys: foundKeys}, nil
}

func (p ProfileServer) DeleteKey(ctx context.Context, req *profile.KeyRequest) (*empty.Empty, error) {
	err := p.keys.DeleteKey(key.ProtoToKeyRequest(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func NewProfileService(userService user.Service, accountService account.Service, keyService key.Service, transactionService transaction.Service) *ProfileServer {
	return &ProfileServer{
		user:               userService,
		account:            accountService,
		keys:               keyService,
		transactionService: transactionService,
	}
}
