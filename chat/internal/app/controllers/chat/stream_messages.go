package chat

import (
	"chat/internal/app/models/types"
	pb "chat/pkg/api/chat/v1"
)

func (c *Controller) StreamMessages(req *pb.StreamMessagesRequest, stream pb.ChatService_StreamMessagesServer) error {
	msgCh, err := c.ChatUsecase.StreamMessages(
		stream.Context(),
		types.ChatID(req.GetChatId()),
		req.GetSinceMessageTime().AsTime().UTC(),
	)
	if err != nil {
		return err
	}

	for msg := range msgCh {
		if err = stream.Send(&pb.StreamMessagesResponse{Message: modelsMessageToPb(msg)}); err != nil {
			return err
		}
	}

	return nil
}
