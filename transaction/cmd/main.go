package main

import (
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"transaction/internal/cfg"
	"transaction/internal/event"
	"transaction/internal/pix"
	"transaction/internal/transactions"
	"transaction/platform/dynamo"
	"transaction/platform/kafka"
	proto "transaction/proto/v1"
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

	db := dynamo.NewClient().Connect()

	// repositories
	transactionRepository := transactions.NewRepository(db, config)

	// services
	transactionService := transactions.NewService(transactionRepository)

	//server
	transactionServer := newTransactionService(transactionService)

	pixService := pix.NewService(transactionService)

	kafkaConn := kafka.NewClient(config).Connect()

	eventTransaction := event.NewEvent(kafkaConn, "transaction_events_topic",
		event.WithAttempts(4), event.WithBroker("localhost:9092"))

	err = eventTransaction.RegisterHandler(context.Background(), pixService.Handler)
	if err != nil {
		panic(err)
	}

	list, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen port 9080 %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterTransactionServiceServer(server, transactionServer)

	log.Printf("Serve is running  on port: %v", "9090")
	if err := server.Serve(list); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9090: %v", err)
	}
}
