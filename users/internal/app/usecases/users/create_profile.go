package users

import (
	"context"
	"fmt"

	lib_types "lib/types"
	"users/internal/app/models"
	users_models "users/internal/app/usecases/users/models"
)

func (uc *usecase) CreateProfile(ctx context.Context, userInfo *users_models.CreateProfileInfo) (*models.UserProfile, error) {
	const api = "users.usecase.CreateProfile"

	userID, err := uc.GetUserIDFromIncomingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, models.ErrUnauthenticated)
	}

	if userID != userInfo.UserID.String() {
		return nil, fmt.Errorf("%s: %w", api, models.ErrPermissionDenied)
	}

	profiles, err := uc.ProfilesRepository.GetProfiles(ctx, users_models.ProfileSelector{
		SelectorType: users_models.SelectorTypeOR,
		ID:           lib_types.WrapField(userInfo.UserID),
		Name:         lib_types.WrapField(userInfo.Name),
	})
	if err != nil {
		return nil, err
	}

	for _, profile := range profiles {
		if profile.ID == userInfo.UserID {
			return nil, fmt.Errorf("%s: profile with id '%s': %w", api, userInfo.UserID, models.ErrAlreadyExists)
		}
		if profile.Name == userInfo.Name {
			return nil, fmt.Errorf("%s: profile with name '%s': %w", api, userInfo.Name, models.ErrAlreadyExists)
		}
	}

	newProfile := &models.UserProfile{
		ID:        userInfo.UserID,
		Email:     userInfo.Email,
		Name:      userInfo.Name,
		Bio:       userInfo.Bio,
		AvatarURL: userInfo.AvatarURL,
	}

	if err = uc.ProfilesRepository.CreateProfile(ctx, newProfile); err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return newProfile, nil
}
