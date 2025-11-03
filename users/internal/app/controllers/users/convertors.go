package users

import (
	"strings"

	"users/internal/app/models"
	"users/internal/app/models/types"
	users_models "users/internal/app/usecases/users/models"
	pb "users/pkg/api/users/v1"
)

func newProfileFromPbCreateOrderRequest(req *pb.CreateProfileRequest) *users_models.CreateProfileInfo {
	return &users_models.CreateProfileInfo{
		UserID:    types.UserID(req.GetUserId()),
		Name:      req.GetNickname(),
		Email:     req.GetEmail(),
		Bio:       strings.TrimSpace(req.GetBio()),
		AvatarURL: strings.TrimSpace(req.GetAvatarUrl()),
	}
}

func modelsProfileToPb(profile *models.UserProfile) *pb.UserProfile {
	if profile == nil {
		return nil
	}

	return &pb.UserProfile{
		UserId:    profile.ID.String(),
		Nickname:  profile.Name,
		Bio:       profile.Bio,
		AvatarUrl: profile.AvatarURL,
	}
}

func modelsProfilesToPb(profiles []*models.UserProfile) []*pb.UserProfile {
	result := make([]*pb.UserProfile, len(profiles))
	for i, profile := range profiles {
		result[i] = modelsProfileToPb(profile)
	}
	return result
}
