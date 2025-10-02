package internal

import (
	"context"
	"log"

	pb "auth/pkg/api/auth/v1"
)

func (s *Server) Refresh(_ context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	log.Println("Refresh")

	return &pb.RefreshResponse{}, nil
}
