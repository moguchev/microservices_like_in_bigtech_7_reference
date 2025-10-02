package internal

import (
	"context"
	"log"

	pb "social/pkg/api/social/v1"
)

func (s *Server) ListRequests(_ context.Context, req *pb.ListRequestsRequest) (*pb.ListRequestsResponse, error) {
	log.Println("ListRequests")

	return &pb.ListRequestsResponse{}, nil
}
