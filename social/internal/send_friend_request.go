package internal

import (
	"context"
	"log"

	pb "social/pkg/api/social/v1"
)

func (s *Server) SendFriendRequest(_ context.Context, req *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	log.Println("SendFriendRequest")

	return &pb.SendFriendRequestResponse{}, nil
}
