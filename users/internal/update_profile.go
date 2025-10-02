package internal

import (
	"context"
	"log"

	pb "users/pkg/api/users/v1"
)

func (s *Server) UpdateProfile(_ context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	log.Println("UpdateProfile")

	return &pb.UpdateProfileResponse{}, nil
}
