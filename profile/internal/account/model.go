package account

import (
	"time"
)

type Account struct {
	Id        int `dynamodbav:"PK"`
	UserID    int `dynamodbav:"SK"`
	Balance   float64
	Agency    string
	Bank      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type AccountRequest struct {
	AccountID int
}

type ListAccountRequest struct {
	AccountIDs []int
}
