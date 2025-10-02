package internal

import (
	"context"
	"log"

	pb "users/pkg/api/users/v1"
)

func (s *Server) GetProfileByID(_ context.Context, req *pb.GetProfileByIDRequest) (*pb.GetProfileByIDResponse, error) {
	log.Println("GetProfileByID")

	return &pb.GetProfileByIDResponse{}, nil
}
