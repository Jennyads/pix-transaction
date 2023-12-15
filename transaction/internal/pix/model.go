package pix

import "github.com/shopspring/decimal"

type Pix struct {
	UserID    string
	AccountId string
	Key       string
	Receiver  string
	Amount    decimal.Decimal
	Status    string
}

type PixEvent struct {
	PixData *Pix
}
