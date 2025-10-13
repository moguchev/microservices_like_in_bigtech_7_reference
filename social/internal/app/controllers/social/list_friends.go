package social

import (
	"context"

	pb "social/pkg/api/social/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *Controller) ListFriends(ctx context.Context, req *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	res, err := c.SocialUsecase.ListFriends(ctx, pbListFriendsRequestToModels(req))
	if err != nil {
		return nil, err
	}

	return &pb.ListFriendsResponse{
		FriendUserIds:         userIDsToStrings(res.UserIDs),
		LastAcceptRequestTime: timestamppb.New(res.NextLastAcceptRequestTime),
	}, nil
}
