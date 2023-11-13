package pix

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "profile/proto/v1"
)

type PixTransaction struct {
	ID        string
	UserID    string
	Key       string
	Amount    float64
	Timestamp *timestamp.Timestamp
	Status    string
}

type CreatePixTransactionRequest struct {
	PixTransaction *PixTransaction
}

type ListPixTransactionsRequest struct {
	PixTransactionIDs []string
}

func ProtoToPixTransaction(transaction *pb.PixTransaction) *PixTransaction {
	return &PixTransaction{
		ID:        transaction.Id,
		UserID:    transaction.UserId,
		Key:       transaction.Key,
		Amount:    transaction.Amount,
		Timestamp: transaction.Timestamp,
		Status:    transaction.Status,
	}
}

func ToProto(pix *PixTransaction) *pb.PixTransaction {
	return &pb.PixTransaction{
		Id:        pix.ID,
		UserId:    pix.UserID,
		Key:       pix.Key,
		Amount:    pix.Amount,
		Timestamp: pix.Timestamp,
		Status:    pix.Status,
	}
}
