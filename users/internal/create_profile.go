package internal

import (
	"context"
	"log"

	pb "users/pkg/api/users/v1"
)

func (s *Server) CreateProfile(_ context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	log.Println("CreateProfile")

	return &pb.CreateProfileResponse{}, nil
}
