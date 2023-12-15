package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transaction/internal/transactions"
	proto "transaction/proto/v1"
)

type TransactionServer struct {
	transaction transactions.Service
	proto.UnimplementedTransactionServiceServer
}

func (t TransactionServer) CreateTransaction(ctx context.Context, request *proto.Transaction) (*proto.Transaction, error) {
	created, err := t.transaction.CreateTransaction(transactions.ProtoToTransaction(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return transactions.ToProto(created), nil
}

func (t TransactionServer) FindTransactionById(ctx context.Context, transaction *proto.TransactionRequest) (*proto.Transaction, error) {
	_, err := t.transaction.FindTransactionById(transactions.ProtoToTransactionRequest(transaction))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (t TransactionServer) ListTransactions(ctx context.Context, transaction *proto.ListTransactionRequest) (*proto.ListTransaction, error) {
	_, err := t.transaction.ListTransactions(transactions.ProtoToListTransactionRequest(transaction))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}
func newTransactionService(transactionService transactions.Service) *TransactionServer {
	return &TransactionServer{
		transaction: transactionService,
	}
}
