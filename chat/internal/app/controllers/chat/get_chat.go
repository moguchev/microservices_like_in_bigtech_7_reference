package chat

import (
	"context"

	"chat/internal/app/models/types"
	pb "chat/pkg/api/chat/v1"
)

func (c *Controller) GetChat(ctx context.Context, req *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	chat, err := c.ChatUsecase.GetChat(ctx, types.ChatID(req.GetChatId()))
	if err != nil {
		return nil, err
	}

	return &pb.GetChatResponse{
		Chat: modelsChatToPb(chat),
	}, nil
}
