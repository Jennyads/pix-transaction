package account

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	proto "profile/proto/profile/v1"
	"time"
)

type Account struct {
	Id        string          `gorm:"primaryKey;type:varchar(36);column:id"`
	UserID    string          `gorm:"foreignKey;type:varchar(36);column:user_id"`
	Balance   decimal.Decimal `gorm:"type:decimal(10,2);column:balance"`
	Agency    string          `gorm:"type:varchar(100);column:agency"`
	Bank      string          `gorm:"type:varchar(100);column:bank"`
	CreatedAt time.Time       `gorm:"type:datetime;column:created_at"`
	UpdatedAt time.Time       `gorm:"type:datetime;column:updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"index;type:datetime;column:deleted_at"`
	BlockedAt *time.Time      `gorm:"type:datetime;column:blocked_at"`
}

type AccountRequest struct {
	AccountID string
	UserID    string
}

type ListAccountRequest struct {
	AccountIDs []string
}

func ProtoToAccount(account *proto.Account) *Account {
	balanceStr := fmt.Sprintf("%.2f", account.Balance)
	balance, err := decimal.NewFromString(balanceStr)
	if err != nil {
		return nil
	}

	return &Account{
		UserID:  account.UserId,
		Balance: balance,
		Agency:  account.Agency,
		Bank:    account.Bank,
	}
}

func ToProto(account *Account) *proto.AccountResponse {
	balance, _ := account.Balance.Float64()
	return &proto.AccountResponse{
		Id:      account.Id,
		UserId:  account.UserID,
		Balance: balance,
		Agency:  account.Agency,
		Bank:    account.Bank,
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
