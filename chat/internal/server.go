package internal

import (
	pb "chat/pkg/api/chat/v1"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}
