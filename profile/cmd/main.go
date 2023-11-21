package main

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"profile/internal/account"
	"profile/internal/cfg"
	"profile/internal/event"
	"profile/internal/key"
	"profile/internal/user"
	"profile/platform/dynamo"
	"profile/platform/kafka"
	v1 "profile/proto/v1"
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
	userRepository := user.NewRepository(db, config)
	keyRepository := key.NewRepository(db, config)
	accountRepository := account.NewRepository(db, config)

	// services
	userService := user.NewService(userRepository)
	keyService := key.NewService(keyRepository)
	accountService := account.NewService(accountRepository)

	//kafka
	kafkaConn := kafka.NewClient().Connect()
	eventTransaction := event.NewEvent(kafkaConn, "transaction_events_topic",
		event.WithAttempts(4), event.WithBroker("localhost:9094"))
	eventTransaction.Publish(context.Background(), []byte("test"))

	//pix kafka
	repo := transaction.NewRepository(db, config)
	dynamicData, err := repo.CreatePixTransaction(&transaction.PixTransaction{})
	if err != nil {
		log.Fatalf("Failed to create PIX transaction: %v", err)
	}
	protoPix := transaction.ToProto(dynamicData)
	protoPixBytes, err := proto.Marshal(protoPix)
	if err != nil {
		log.Fatalf("Failed to marshal PIX transaction to bytes: %v", err)
	}
	err = eventTransaction.Publish(context.Background(), protoPixBytes)
	if err != nil {
		log.Fatalf("Failed to publish PIX transaction to Kafka: %v", err)
	}

	//server
	profileServer := NewProfileService(userService, accountService, keyService)
	server := grpc.NewServer()
	v1.RegisterUserServiceServer(server, profileServer)
	v1.RegisterAccountServiceServer(server, profileServer)
	v1.RegisterKeysServiceServer(server, profileServer)

	log.Printf("Serve is running  on port: %v", "9080")
	if err := server.Serve(list); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9080: %v", err)
	}
}
