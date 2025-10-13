package chat

import (
	"context"

	"chat/internal/app/models/types"
	pb "chat/pkg/api/chat/v1"
)

func (c *Controller) CreateDirectChat(ctx context.Context, req *pb.CreateDirectChatRequest) (*pb.CreateDirectChatResponse, error) {
	chat, err := c.ChatUsecase.CreateDirectChat(ctx, types.UserID(req.GetParticipantId()))
	if err != nil {
		return nil, err
	}

	return &pb.CreateDirectChatResponse{
		ChatId: chat.ID.String(),
	}, nil
}
