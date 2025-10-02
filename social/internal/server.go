package internal

import (
	pb "social/pkg/api/social/v1"
)

type Server struct {
	pb.UnimplementedSocialServiceServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}
