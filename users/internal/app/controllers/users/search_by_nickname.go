package users

import (
	"context"

	pb "users/pkg/api/users/v1"
)

func (c *Controller) SearchByNickname(ctx context.Context, req *pb.SearchByNicknameRequest) (*pb.SearchByNicknameResponse, error) {
	profiles, err := c.UsersUsecase.SearchByNickname(ctx, req.GetQuery(), req.GetLimit())
	if err != nil {
		return nil, err
	}

	return &pb.SearchByNicknameResponse{
		Results: modelsProfilesToPb(profiles),
	}, nil
}
