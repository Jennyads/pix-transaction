package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"profile/internal/account"
	"profile/internal/keys"
	"profile/internal/user"
	"profile/platform/dynamo"
	v1 "profile/proto/v1"
)

func main() {
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen port 9080 %v", err)
	}

	db := dynamo.NewClient().Connect()

	// repositories
	userRepository := user.NewRepository(db)
	keyRepository := keys.NewRepository(db)
	accountRepository := account.NewRepository(db)

	// services
	userService := user.NewService(userRepository)
	keyService := keys.NewService(keyRepository)
	accountService := account.NewService(accountRepository)

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
