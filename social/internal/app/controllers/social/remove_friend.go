package social

import (
	"context"

	"social/internal/app/models/types"
	pb "social/pkg/api/social/v1"
)

func (c *Controller) RemoveFriend(ctx context.Context, req *pb.RemoveFriendRequest) (*pb.RemoveFriendResponse, error) {
	err := c.SocialUsecase.RemoveFriend(ctx, types.UserID(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &pb.RemoveFriendResponse{}, nil
}
