package internal

import (
	pb "users/pkg/api/users/v1"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}
