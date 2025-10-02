package internal

import (
	"context"
	"log"

	pb "social/pkg/api/social/v1"
)

func (s *Server) RemoveFriend(_ context.Context, req *pb.RemoveFriendRequest) (*pb.RemoveFriendResponse, error) {
	log.Println("RemoveFriend")

	return &pb.RemoveFriendResponse{}, nil
}
