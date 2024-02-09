package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	keys "transaction/internal/keys"
	"transaction/internal/transactions"
	proto "transaction/proto/v1"
)

type TransactionServer struct {
	transaction transactions.Service
	keys        keys.Service
	proto.UnimplementedTransactionServiceServer
	proto.UnimplementedKeysServiceServer
}

func (t *TransactionServer) CreateTransaction(ctx context.Context, request *proto.Transaction) (*proto.Transaction, error) {
	created, err := t.transaction.CreateTransaction(transactions.ProtoToTransaction(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return transactions.ToProto(created), nil
}

func (t *TransactionServer) FindTransactionById(ctx context.Context, transaction *proto.TransactionRequest) (*proto.Transaction, error) {
	_, err := t.transaction.FindTransactionById(transactions.ProtoToTransactionRequest(transaction))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (t *TransactionServer) ListTransactions(ctx context.Context, transaction *proto.ListTransactionRequest) (*proto.ListTransaction, error) {
	_, err := t.transaction.ListTransactions(transactions.ProtoToListTransactionRequest(transaction))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (t *TransactionServer) CreateKey(ctx context.Context, key *proto.Key) (*proto.KeyResponse, error) {
	created, err := t.keys.CreateKey(ctx, keys.ProtoToKey(key))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return keys.ToProto(created), nil
}

func (t *TransactionServer) UpdateKey(ctx context.Context, key *proto.Key) (*proto.KeyResponse, error) {
	created, err := t.keys.UpdateKey(ctx, keys.ProtoToKey(key))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return keys.ToProto(created), nil
}

func (t *TransactionServer) ListKey(ctx context.Context, req *proto.ListKeyRequest) (*proto.ListKeys, error) {
	accounts, err := t.keys.ListKey(ctx, keys.ProtoToKeyListRequest(req))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	findKeys := make([]*proto.KeyResponse, len(accounts))
	for i := range accounts {
		findKeys[i] = keys.ToProto(accounts[i])
	}
	return &proto.ListKeys{Keys: findKeys}, nil
}
func (t *TransactionServer) DeleteKey(ctx context.Context, request *proto.KeyRequest) (*empty.Empty, error) {
	err := t.keys.DeleteKey(ctx, keys.ProtoToKeyRequest(request))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (t *TransactionServer) FindKey(ctx context.Context, key *proto.KeyRequest) (*proto.KeyResponse, error) {
	foundKey, err := t.keys.FindKey(ctx, keys.ProtoToKeyRequest(key))
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return keys.ToProto(foundKey), nil
}

func newTransactionService(transactionService transactions.Service, keysService keys.Service) *TransactionServer {
	return &TransactionServer{
		transaction: transactionService,
		keys:        keysService,
	}
}
