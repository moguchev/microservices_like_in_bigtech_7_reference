package internal

import (
	"context"
	"log"

	pb "chat/pkg/api/chat/v1"
)

func (s *Server) SendMessage(_ context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	log.Println("SendMessage")

	return &pb.SendMessageResponse{}, nil
}
