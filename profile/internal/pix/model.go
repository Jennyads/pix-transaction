package pix

import (
	"time"

	//account2 "profile/internal/account"
	//user2 "profile/internal/user"
	//key2 "profile/internal/key"
	pb "profile/proto/v1"
)

//	type PixTransaction struct {
//		ID        string
//		UserID    user2.User
//		AccountId account2.Account
//		Key       key2.Key
//		Amount    float64
//		Timestamp *timestamp.Timestamp
//		Status    string
//	}

type PixTransaction struct {
	ID           string    `validate:"required"`
	UserID       string    `validate:"required"`
	AccountId    string    `validate:"required"`
	SenderKey    string    `validate:"required,validatePixKey"`
	RecipientKey string    `validate:"required,validatePixKey"`
	Amount       float64   `validate:"required,gte=0,validateSenderBalance"`
	Hour         time.Time `validate:"required"`
	Status       string    `validate:"required"`
}

type CreatePixTransactionRequest struct {
	PixTransaction *PixTransaction
}

type ListPixTransactionsRequest struct {
	PixTransactionIDs []string
}

func ProtoToPixTransaction(transaction *pb.PixTransaction) *PixTransaction {
	return &PixTransaction{
		ID:           transaction.Id,
		UserID:       transaction.UserId,
		SenderKey:    transaction.SenderKey,
		RecipientKey: transaction.ReceiverKey,
		Amount:       transaction.Amount,
		Status:       transaction.Status,
	}
}

func ToProto(pix *PixTransaction) *pb.PixTransaction {
	return &pb.PixTransaction{
		Id:          pix.ID,
		UserId:      pix.UserID,
		SenderKey:   pix.SenderKey,
		ReceiverKey: pix.RecipientKey,
		Amount:      pix.Amount,
		Status:      pix.Status,
	}
}
