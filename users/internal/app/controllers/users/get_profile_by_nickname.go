package users

import (
	"context"

	pb "users/pkg/api/users/v1"
)

func (c *Controller) GetProfileByNickname(ctx context.Context, req *pb.GetProfileByNicknameRequest) (*pb.GetProfileByNicknameResponse, error) {
	profile, err := c.UsersUsecase.GetProfileByNickname(ctx, req.GetNickname())
	if err != nil {
		return nil, err
	}

	return &pb.GetProfileByNicknameResponse{
		Profile: modelsProfileToPb(profile),
	}, nil
}
