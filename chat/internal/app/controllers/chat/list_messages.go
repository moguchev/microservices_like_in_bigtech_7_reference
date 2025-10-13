package chat

import (
	"context"

	pb "chat/pkg/api/chat/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *Controller) ListMessages(ctx context.Context, req *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	res, err := c.ChatUsecase.ListMessages(ctx, pbListMessagesRequestToModels(req))
	if err != nil {
		return nil, err
	}

	return &pb.ListMessagesResponse{
		Messages:        modelsMessagesToPb(res.Messages),
		LastMessageTime: timestamppb.New(res.NextLastMessageTime),
	}, nil
}
