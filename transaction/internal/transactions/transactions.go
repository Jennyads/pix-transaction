package internal

import (
	"time"
)

type Transaction struct {
	PK          int
	AccountID   int `dynamodbav:"SK"`
	Receiver    int
	Value       float64
	CreatedAt   time.Time
	ProcessedAt time.Time
}
