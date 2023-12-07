package transaction

import (
	"fmt"
	"github.com/shopspring/decimal"
	pb "profile/proto/v1"
)

type Pix struct {
	UserID    string          `gorm:"type:varchar(36);column:user_id"`
	AccountID string          `gorm:"type:varchar(36);column:account_id"`
	Key       string          `gorm:"type:varchar(255);column:key"`
	Receiver  string          `gorm:"type:varchar(255);column:receiver"`
	Amount    decimal.Decimal `gorm:"type:decimal(15,6);column:amount"`
	Status    string          `gorm:"type:varchar(50);column:status"`
}

func ToProto(pix *Pix) *pb.PixTransaction {
	amount, _ := pix.Amount.Float64()
	return &pb.PixTransaction{
		UserId:      pix.UserID,
		SenderKey:   pix.Key,
		ReceiverKey: pix.Receiver,
		Amount:      amount,
		Status:      pix.Status,
	}
}

func ProtoToPix(pix *pb.PixTransaction) *Pix {
	amountStr := fmt.Sprintf("%.2f", pix.Amount)
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		return nil
	}

	return &Pix{
		UserID:   pix.UserId,
		Key:      pix.SenderKey,
		Receiver: pix.ReceiverKey,
		Amount:   amount,
		Status:   pix.Status,
	}
}

type PixEvent struct {
	PixData *Pix
}
