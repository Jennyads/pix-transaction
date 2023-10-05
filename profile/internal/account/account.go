package internal

import (
	user "profile/internal/user"
	"time"
)

type nullableTime struct {
	Time  time.Time
	Valid bool
}

type Account struct {
	PK        int
	SK        user.User
	Balance   float64
	Agency    string
	Bank      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *nullableTime
}
