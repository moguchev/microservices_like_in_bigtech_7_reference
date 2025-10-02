package internal

import (
	"context"
	"log"

	pb "chat/pkg/api/chat/v1"
)

func (s *Server) CreateDirectChat(_ context.Context, req *pb.CreateDirectChatRequest) (*pb.CreateDirectChatResponse, error) {
	log.Println("CreateDirectChat")

	return &pb.CreateDirectChatResponse{}, nil
}
