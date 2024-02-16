package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"profile/internal/account"
	"profile/internal/cfg"
	"profile/internal/event"
	"profile/internal/key"
	"profile/internal/transaction"
	"profile/internal/user"
	"profile/platform/kafka"
	"profile/platform/sqlserver"
	proto "profile/proto/profile/v1"
	transpb "profile/proto/transactions/v1"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config, err := cfg.Load()
	if err != nil {
		return
	}

	client, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%s", "localhost", "9090"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start connection with grpc on %s:%s: %v", "localhost", "9090", err)
	}

	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen port 9080 %v", err)
	}

	err = sqlserver.RunMigrations(config)
	if err != nil {
		log.Fatalf("Failed to run migrations %v", err)
	}

	db, err := sqlserver.Start(config)
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	//kafka
	kafkaConn := kafka.NewClient(config).Connect()

	events := event.NewEvent(kafkaConn, "transaction_events_topic",
		event.WithAttempts(4), event.WithBroker("localhost:9092"))

	transClient := transpb.NewKeysServiceClient(client)

	// repositories
	userRepository := user.NewRepository(db, config)
	keyRepository := key.NewRepository(db, config)
	accountRepository := account.NewRepository(db, config)

	// services
	transactionService := transaction.NewService(events, accountRepository, transClient, userRepository)
	userService := user.NewService(userRepository)
	keyService := key.NewService(keyRepository, transactionService)
	accountService := account.NewService(accountRepository)

	//server
	profileServer := NewProfileService(userService, accountService, keyService, transactionService)
	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, profileServer)
	proto.RegisterAccountServiceServer(server, profileServer)
	proto.RegisterKeysServiceServer(server, profileServer)
	proto.RegisterPixTransactionServiceServer(server, profileServer)

	log.Printf("Serve is running  on port: %v", "9080")
	if err := server.Serve(list); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9080: %v", err)
	}
}
