package chat

import (
	"context"

	"chat/internal/app/models/types"
	pb "chat/pkg/api/chat/v1"
)

func (c *Controller) ListUserChats(ctx context.Context, req *pb.ListUserChatsRequest) (*pb.ListUserChatsResponse, error) {
	chats, err := c.ChatUsecase.ListUserChats(ctx, types.UserID(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &pb.ListUserChatsResponse{
		Chats: modelsChatsToPb(chats),
	}, nil
}
