package main

import (
	"log"
	"net"

	"chat/internal"
	pb "chat/pkg/api/chat/v1"
	"lib/grpc_utils"
	grpc_middleware "lib/middleware/grpc"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const address = ":8080"

func main() {
	server, err := internal.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	validator, err := protovalidate.New(protovalidate.WithDisableLazy(false))
	if err != nil {
		log.Fatalf("server: failed to initialize validator: %s", err)
	}

	grpcServerOptions := grpc_utils.UnaryInterceptorsToGrpcServerOptions(
		grpc_middleware.ValidateUnaryServerInterceptor(validator),
	)

	grpcServer := grpc.NewServer(grpcServerOptions...)
	pb.RegisterChatServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
