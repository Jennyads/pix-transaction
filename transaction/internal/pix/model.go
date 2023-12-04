package pix

import "errors"

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrKeyNotFound         = errors.New("key not found in the database")
	ErrInvalidKey          = errors.New("invalid key")
	ErrInactiveAccount     = errors.New("inactive account")
)

type Pix struct {
	UserID    string
	AccountId string
	Key       string
	Receiver  string
	Amount    float64
	Status    string
}

type PixEvent struct {
	PixData *Pix
}
