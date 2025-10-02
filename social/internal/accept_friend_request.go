package internal

import (
	"context"
	"log"

	pb "social/pkg/api/social/v1"
)

func (s *Server) AcceptFriendRequest(_ context.Context, req *pb.AcceptFriendRequestRequest) (*pb.AcceptFriendRequestResponse, error) {
	log.Println("AcceptFriendRequest")

	return &pb.AcceptFriendRequestResponse{}, nil
}
