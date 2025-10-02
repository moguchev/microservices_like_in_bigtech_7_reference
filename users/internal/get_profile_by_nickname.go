package internal

import (
	"context"
	"log"

	pb "users/pkg/api/users/v1"
)

func (s *Server) GetProfileByNickname(_ context.Context, req *pb.GetProfileByNicknameRequest) (*pb.GetProfileByNicknameResponse, error) {
	log.Println("GetProfileByNickname")

	return &pb.GetProfileByNicknameResponse{}, nil
}
