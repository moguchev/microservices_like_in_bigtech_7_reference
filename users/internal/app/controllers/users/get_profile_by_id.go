package users

import (
	"context"

	"users/internal/app/models/types"
	pb "users/pkg/api/users/v1"
)

func (c *Controller) GetProfileByID(ctx context.Context, req *pb.GetProfileByIDRequest) (*pb.GetProfileByIDResponse, error) {
	profile, err := c.UsersUsecase.GetProfileByID(ctx, types.UserID(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.GetProfileByIDResponse{
		Profile: modelsProfileToPb(profile),
	}, nil
}
