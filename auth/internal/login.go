package internal

import (
	"context"
	"log"

	pb "auth/pkg/api/auth/v1"
)

func (s *Server) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("Login")

	return &pb.LoginResponse{}, nil
}
