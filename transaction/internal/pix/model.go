package pix

import "errors"

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
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
