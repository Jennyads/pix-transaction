package transaction

import (
	"time"
)

type Pix struct {
	UserID    string
	AccountId string
	Key       string
	Amount    float64
	Date      time.Time
	Status    string
}

type PixEvent struct {
	PixData *Pix
}
