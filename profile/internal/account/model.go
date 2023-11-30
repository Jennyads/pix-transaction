package account

import (
	"profile/proto/v1"
	"time"
)

type Account struct {
	Id        string `dynamodbav:"PK"`
	UserID    string `dynamodbav:"SK"`
	Balance   float64
	Agency    string
	Bank      string
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type AccountRequest struct {
	AccountID string
	UserID    string
}

type ListAccountRequest struct {
	AccountIDs []string
}

func ProtoToAccount(account *proto.Account) *Account {
	return &Account{
		UserID:  account.UserId,
		Balance: account.Balance,
		Agency:  account.Agency,
		Bank:    account.Bank,
		Key:     account.Key,
	}
}

func ToProto(account *Account) *proto.Account {
	return &proto.Account{
		UserId:  account.UserID,
		Balance: account.Balance,
		Agency:  account.Agency,
		Bank:    account.Bank,
		Key:     account.Key,
	}
}

func ProtoToAccountRequest(request *proto.AccountRequest) *AccountRequest {
	return &AccountRequest{
		UserID:    request.UserId,
		AccountID: request.AccountId,
	}
}

func ProtoToAccountListRequest(request *proto.ListAccountRequest) *ListAccountRequest {
	return &ListAccountRequest{
		AccountIDs: request.AccountId,
	}
}
