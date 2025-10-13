package chat

import (
	"context"

	"chat/internal/app/models/types"
	pb "chat/pkg/api/chat/v1"
)

func (c *Controller) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	msg, err := c.ChatUsecase.SendMessage(ctx, types.ChatID(req.GetChatId()), req.GetText())
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{
		Message: modelsMessageToPb(msg),
	}, nil
}
