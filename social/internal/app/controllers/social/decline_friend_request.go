package social

import (
	"context"

	"social/internal/app/models/types"
	pb "social/pkg/api/social/v1"
)

func (c *Controller) DeclineFriendRequest(ctx context.Context, req *pb.DeclineFriendRequestRequest) (*pb.DeclineFriendRequestResponse, error) {
	request, err := c.SocialUsecase.DeclineFriendRequest(ctx, types.RequestID(req.GetRequestId()))
	if err != nil {
		return nil, err
	}

	return &pb.DeclineFriendRequestResponse{
		Request: modelsFriendRequestToPb(request),
	}, nil
}
