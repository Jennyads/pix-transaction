package main

import (
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"transaction/internal/cfg"
	"transaction/internal/event"
	"transaction/internal/transactions"
	"transaction/platform/dynamo"
	"transaction/platform/kafka"
	v1 "transaction/proto"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config, err := cfg.Load()
	if err != nil {
		return
	}
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen port 9080 %v", err)
	}

	db := dynamo.NewClient().Connect()

	// repositories
	transactionRepository := transactions.NewRepository(db, config)

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

	//kafka
	kafkaConn := kafka.NewClient(config).Connect()

	eventTransaction := event.NewEvent(kafkaConn, "transaction_events_topic",
		event.WithAttempts(4), event.WithBroker("localhost:9094"))

	eventTransaction.Publish(context.Background(), []byte("test"))
}
