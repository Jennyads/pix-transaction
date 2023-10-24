package account

import (
	pb "profile/proto/v1"
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
}

type ListAccountRequest struct {
	AccountIDs []string
}

func ProtoToAccount(account *pb.Account) *Account {
	return &Account{
		UserID:  account.UserId,
		Balance: account.Balance,
		Agency:  account.Agency,
		Bank:    account.Bank,
		Key:     account.Key,
	}
}

func ProtoToAccountRequest(request *pb.AccountRequest) *AccountRequest {
	return &AccountRequest{
		AccountID: request.AccountId,
	}
}

func ProtoToAccountListRequest(request *pb.ListAccountRequest) *ListAccountRequest {
	return &ListAccountRequest{
		AccountIDs: request.AccountId,
	}
}
