package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authpb "auth/pkg/api/auth/v1"
	chatpb "chat/pkg/api/chat/v1"
	socialpb "social/pkg/api/social/v1"
	userspb "users/pkg/api/users/v1"
)

const address = ":8080"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gWmux := runtime.NewServeMux()

	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, gWmux, "auth:8081", grpcOpts); err != nil {
		log.Fatalf("failed to register auth service: %v", err)
	}

	if err := chatpb.RegisterChatServiceHandlerFromEndpoint(ctx, gWmux, "chat:8082", grpcOpts); err != nil {
		log.Fatalf("failed to register chat service: %v", err)
	}

	if err := socialpb.RegisterSocialServiceHandlerFromEndpoint(ctx, gWmux, "social:8083", grpcOpts); err != nil {
		log.Fatalf("failed to register social service: %v", err)
	}

	if err := userspb.RegisterUserServiceHandlerFromEndpoint(ctx, gWmux, "users:8084", grpcOpts); err != nil {
		log.Fatalf("failed to register users service: %v", err)
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("../swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	mux.Handle("/", gWmux)

	httpServer := &http.Server{
		Handler: mux,
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Gateway listening at %v", lis.Addr())
	if err := httpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
