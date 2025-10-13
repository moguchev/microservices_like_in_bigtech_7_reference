package users

import (
	"context"
	"errors"

	lib_types "lib/types"
	"users/internal/app/models/types"
	"users/internal/app/usecases/users/models"
	pb "users/pkg/api/users/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrOneOfFieldsMustBeSet = errors.New("one of fields must be set")
)

func (c *Controller) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	fields, err := getUpdateUserFields(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	profile, err := c.UsersUsecase.UpdateProfile(ctx, types.UserID(req.UserId), fields)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfileResponse{
		Profile: modelsProfileToPb(profile),
	}, nil
}

func getUpdateUserFields(req *pb.UpdateProfileRequest) (models.ProfileUpdateFields, error) {
	var (
		fields             models.ProfileUpdateFields
		f                  = req.GetUpdateFields()
		fieldPaths         = req.GetUpdateMask().GetPaths()
		atLeastOneFieldSet bool
	)

	for _, path := range fieldPaths {
		switch path {
		case "nickname":
			if f.Nickname != nil {
				fields.Name = lib_types.WrapField[string](*f.Nickname)
				atLeastOneFieldSet = true
			}
		case "bio":
			if f.Bio != nil {
				fields.Bio = lib_types.WrapField[string](*f.Bio)
				atLeastOneFieldSet = true
			}
		case "avatar_url":
			if f.AvatarUrl != nil {
				fields.AvatarURL = lib_types.WrapField[string](*f.AvatarUrl)
				atLeastOneFieldSet = true
			}
		}
	}

	if !atLeastOneFieldSet {
		return models.ProfileUpdateFields{}, ErrOneOfFieldsMustBeSet
	}

	return fields, nil
}
