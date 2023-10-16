package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"transaction/internal/transactions"
	"transaction/platform/dynamo"
	v1 "transaction/proto"
)

func main() {
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen port 9080 %v", err)
	}

	db := dynamo.NewClient().Connect()

	// repositories
	transactionRepository := transactions.NewRepository(db)

	// services
	transactionService := transactions.NewService(transactionRepository)

	//server
	transactionServer := newTransactionService(transactionService)
	server := grpc.NewServer()
	v1.RegisterTransactionServiceServer(server, transactionServer)

	log.Printf("Serve is running  on port: %v", "9080")
	if err := server.Serve(list); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9080: %v", err)
	}
}
