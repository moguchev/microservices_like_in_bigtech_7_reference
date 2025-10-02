package internal

import (
	"context"
	"log"

	pb "chat/pkg/api/chat/v1"
)

func (s *Server) ListChatMembers(_ context.Context, req *pb.ListChatMembersRequest) (*pb.ListChatMembersResponse, error) {
	log.Println("ListChatMembers")

	return &pb.ListChatMembersResponse{}, nil
}
