package internal

import (
	"log"

	pb "chat/pkg/api/chat/v1"
)

func (s *Server) StreamMessages(req *pb.StreamMessagesRequest, stream pb.ChatService_StreamMessagesServer) error {
	log.Println("StreamMessages")

	return nil
}
