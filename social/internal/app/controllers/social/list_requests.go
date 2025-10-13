package social

import (
	"context"

	"social/internal/app/models/types"
	pb "social/pkg/api/social/v1"
)

func (c *Controller) ListRequests(ctx context.Context, req *pb.ListRequestsRequest) (*pb.ListRequestsResponse, error) {
	requests, err := c.SocialUsecase.ListRequests(ctx, types.UserID(req.UserId))
	if err != nil {
		return nil, err
	}

	return &pb.ListRequestsResponse{
		Requests: modelsFriendRequestsToPb(requests),
	}, nil
}
