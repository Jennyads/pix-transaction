package account

import (
	"time"
)

type Account struct {
	PK        int
	AccountID int `dynamodbav:"PK"`
	Balance   float64
	Agency    string
	Bank      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
