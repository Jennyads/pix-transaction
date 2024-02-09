package main

import (
	"api/cmd/handlers"
	"api/internal/config"
	"api/internal/middleware"
	"api/internal/profile"
	proto "api/proto/v1"
	"context"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	cfg := config.Load()

	client, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%s", cfg.Profile.Host, cfg.Profile.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start connection with grpc on %s:%s: %v", cfg.Profile.Host, cfg.Profile.Port, err)
	}

	user := proto.NewUserServiceClient(client)
	account := proto.NewAccountServiceClient(client)
	keys := proto.NewKeysServiceClient(client)
	pix := proto.NewPixTransactionServiceClient(client)

	profileBackend := profile.NewBackend(user, account, keys, pix)

	profileHandler := handlers.NewProfileHandler(profileBackend)

	routes := router.New()

	logger := middleware.NewLogger(true)

	routes = handlers.ProfileRoutes(routes, profileHandler, logger)

	log.Printf("Serve is running on port: %s\n", cfg.Api.Port)
	if err = fasthttp.ListenAndServe(fmt.Sprintf(":%s", cfg.Api.Port), routes.Handler); err != nil {
		log.Fatalf("Failed to listen server on port %s: %v", cfg.Api.Port, err)
	}
}
