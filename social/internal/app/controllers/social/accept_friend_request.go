package social

import (
	"context"

	"social/internal/app/models/types"
	pb "social/pkg/api/social/v1"
)

func (c *Controller) AcceptFriendRequest(ctx context.Context, req *pb.AcceptFriendRequestRequest) (*pb.AcceptFriendRequestResponse, error) {
	request, err := c.SocialUsecase.AcceptFriendRequest(ctx, types.RequestID(req.GetRequestId()))
	if err != nil {
		return nil, err
	}

	return &pb.AcceptFriendRequestResponse{
		Request: modelsFriendRequestToPb(request),
	}, nil
}
