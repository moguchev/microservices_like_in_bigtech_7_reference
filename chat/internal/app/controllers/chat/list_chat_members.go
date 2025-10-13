package chat

import (
	"context"

	"chat/internal/app/models/types"
	pb "chat/pkg/api/chat/v1"
)

func (c *Controller) ListChatMembers(ctx context.Context, req *pb.ListChatMembersRequest) (*pb.ListChatMembersResponse, error) {
	userIDs, err := c.ChatUsecase.ListChatMembers(ctx, types.ChatID(req.GetChatId()))
	if err != nil {
		return nil, err
	}

	return &pb.ListChatMembersResponse{
		UserIds: userIDsToStrings(userIDs),
	}, nil
}
