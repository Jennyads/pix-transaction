package account

import (
	"gorm.io/gorm"
	proto "profile/proto/v1"
	"time"
)

type Account struct {
	Id        string `gorm:"primarykey;type:varchar(36)"`
	UserID    string `gorm:"foreignKey:ID;references:users;type:varchar(36)"`
	Balance   float64
	Agency    string
	Bank      string
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
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
