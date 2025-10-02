package internal

import (
	"context"
	"log"

	pb "users/pkg/api/users/v1"
)

func (s *Server) SearchByNickname(_ context.Context, req *pb.SearchByNicknameRequest) (*pb.SearchByNicknameResponse, error) {
	log.Println("SearchByNickname")

	return &pb.SearchByNicknameResponse{}, nil
}
