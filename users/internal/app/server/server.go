package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"lib/grpc_utils"
	pb "users/pkg/api/users/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Config - server config
type Config struct {
	GRPCPort string

	ChainUnaryInterceptors []grpc.UnaryServerInterceptor
	UnaryInterceptors      []grpc.UnaryServerInterceptor
}

type Contollers struct {
	pb.UserServiceServer
}

type Server struct {
	Contollers

	grpc struct {
		lis    net.Listener
		server *grpc.Server
	}
}

// New - returns *Server
func New(ctx context.Context, cfg Config, svcs Contollers) (*Server, error) {
	srv := &Server{Contollers: svcs}

	// middlewares
	grpcServerOptions := grpc_utils.UnaryInterceptorsToGrpcServerOptions(cfg.UnaryInterceptors...)
	grpcServerOptions = append(grpcServerOptions,
		grpc.ChainUnaryInterceptor(cfg.ChainUnaryInterceptors...),
	)

	grpcServer := grpc.NewServer(grpcServerOptions...)
	// router
	pb.RegisterUserServiceServer(grpcServer, srv)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		return nil, fmt.Errorf("server: failed to listen: %v", err)
	}

	srv.grpc.server = grpcServer
	srv.grpc.lis = lis

	return srv, nil
}

// Run - serve
func (s *Server) Run(ctx context.Context) error {
	log.Println("start serve grpc", s.grpc.lis.Addr())

	if err := s.grpc.server.Serve(s.grpc.lis); err != nil {
		return fmt.Errorf("server: serve grpc: %v", err)
	}

	return nil
}
