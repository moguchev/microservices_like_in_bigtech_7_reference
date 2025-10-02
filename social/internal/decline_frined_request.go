package internal

import (
	"context"
	"log"

	pb "social/pkg/api/social/v1"
)

func (s *Server) DeclineFriendRequest(_ context.Context, req *pb.DeclineFriendRequestRequest) (*pb.DeclineFriendRequestResponse, error) {
	log.Println("DeclineFriendRequest")

	return &pb.DeclineFriendRequestResponse{}, nil
}
