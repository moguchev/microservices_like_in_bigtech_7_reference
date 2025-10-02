package internal

import (
	"context"
	"log"

	pb "auth/pkg/api/auth/v1"
)

func (s *Server) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Println("Register")

	return &pb.RegisterResponse{}, nil
}
