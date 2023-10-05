package internal

import (
	account "profile/internal/account"
	"time"
)

type Type string

const (
	Cpf    Type = "cpf"
	Phone  Type = "phone"
	Email  Type = "email"
	Random Type = "random"
)

type User struct {
	PK        int
	SK        account.Account
	Name      string
	Type      Type
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
