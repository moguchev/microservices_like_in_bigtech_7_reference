package internal

import (
	"context"
	"log"

	pb "chat/pkg/api/chat/v1"
)

func (s *Server) ListMessages(_ context.Context, req *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	log.Println("ListMessages")

	return &pb.ListMessagesResponse{}, nil
}
