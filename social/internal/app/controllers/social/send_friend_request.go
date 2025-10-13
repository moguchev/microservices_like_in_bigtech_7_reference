package social

import (
	"context"

	"social/internal/app/models/types"
	pb "social/pkg/api/social/v1"
)

func (c *Controller) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	request, err := c.SocialUsecase.SendFriendRequest(ctx, types.UserID(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &pb.SendFriendRequestResponse{
		Request: modelsFriendRequestToPb(request),
	}, nil
}
