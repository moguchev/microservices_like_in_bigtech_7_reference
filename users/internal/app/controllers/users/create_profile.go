package users

import (
	"context"

	pb "users/pkg/api/users/v1"
)

func (c *Controller) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	profileInfo := newProfileFromPbCreateOrderRequest(req)

	newProfile, err := c.UsersUsecase.CreateProfile(ctx, profileInfo)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProfileResponse{
		Profile: modelsProfileToPb(newProfile),
	}, nil
}
