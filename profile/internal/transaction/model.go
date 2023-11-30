package transaction

import (
	pb "profile/proto/v1"
)

type Pix struct {
	UserID    string
	AccountId string
	Key       string
	Receiver  string
	Amount    float64
	Status    string
}

func ToProto(pix *Pix) *pb.PixTransaction {
	return &pb.PixTransaction{
		UserId:      pix.UserID,
		SenderKey:   pix.Key,
		ReceiverKey: pix.Receiver,
		Amount:      pix.Amount,
		Status:      pix.Status,
	}
}

func ProtoToPix(pix *pb.PixTransaction) *Pix {
	return &Pix{
		UserID:   pix.UserId,
		Key:      pix.SenderKey,
		Receiver: pix.ReceiverKey,
		Amount:   pix.Amount,
		Status:   pix.Status,
	}
}

type PixEvent struct {
	PixData *Pix
}
