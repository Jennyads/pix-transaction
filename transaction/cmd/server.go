package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transaction/internal/transactions"
	v1 "transaction/proto"
)

type TransactionServer struct {
	transaction transactions.Service
	v1.UnimplementedTransactionServiceServer
}

func (t TransactionServer) CreateTransaction(ctx context.Context, request *v1.Transaction) (*empty.Empty, error) {
	err := t.transaction.CreateTransaction(transactions.ProtoToTransaction(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (t TransactionServer) FindTransaction(ctx context.Context, transaction *v1.TransactionRequest) (*v1.Transaction, error) {
	_, err := t.transaction.FindTransaction(transactions.ProtoToTransactionRequest(transaction))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (t TransactionServer) ListTransactions(ctx context.Context, transaction *v1.ListTransactionRequest) (*v1.ListTransaction, error) {
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
	return &TransactionServer{}
}
