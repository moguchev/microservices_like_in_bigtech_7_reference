package internal

import (
	"context"
	"log"

	pb "social/pkg/api/social/v1"
)

func (s *Server) ListFriends(_ context.Context, req *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	log.Println("ListFriends")

	return &pb.ListFriendsResponse{}, nil
}
