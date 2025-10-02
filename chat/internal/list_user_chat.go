package internal

import (
	"context"
	"log"

	pb "chat/pkg/api/chat/v1"
)

func (s *Server) ListUserChats(_ context.Context, req *pb.ListUserChatsRequest) (*pb.ListUserChatsResponse, error) {
	log.Println("ListUserChats")

	return &pb.ListUserChatsResponse{}, nil
}
