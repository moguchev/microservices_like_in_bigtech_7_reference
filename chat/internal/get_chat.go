package internal

import (
	"context"
	"log"

	pb "chat/pkg/api/chat/v1"
)

func (s *Server) GetChat(_ context.Context, req *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	log.Println("GetChat")

	return &pb.GetChatResponse{}, nil
}
