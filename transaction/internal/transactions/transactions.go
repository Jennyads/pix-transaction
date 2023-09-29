package internal

import (
	"os/user"
	"time"
	//user "transaction/internal/user"
)

type Transaction struct {
	PK          int
	SK          user.User
	Receiver    int
	Value       float64
	CreatedAt   time.Time
	ProcessedAt time.Time
}
