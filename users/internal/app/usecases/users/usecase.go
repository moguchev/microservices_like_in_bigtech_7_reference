package users

import (
	"context"

	"users/internal/app/models"
	"users/internal/app/models/types"
	users_models "users/internal/app/usecases/users/models"
)

type Usecase interface {
	// CreateProfile - создание профиля пользователя
	//
	// errors:  models.ErrAlreadyExists
	CreateProfile(ctx context.Context, userInfo *users_models.CreateProfileInfo) (*models.UserProfile, error)
	// UpdateProfile - обновление профиля пользователя
	UpdateProfile(ctx context.Context, id types.UserID, f users_models.ProfileUpdateFields) (*models.UserProfile, error)
	// GetProfileByID - получение профиля пользователя по ID
	//
	// errors:  models.ErrNotFound
	GetProfileByID(ctx context.Context, id types.UserID) (*models.UserProfile, error)
	// GetProfileByNickname - получение профиля пользователя по никунейму
	//
	// errors:  models.ErrNotFound
	GetProfileByNickname(ctx context.Context, name string) (*models.UserProfile, error)
	// SearchByNickname - поиск профилей пользователей по части имени
	SearchByNickname(ctx context.Context, query string, limit uint32) ([]*models.UserProfile, error)
}

type (
	ProfileRepository interface {
		// CreateProfile - создание профиля пользователя
		//
		// errors: models.ErrAlreadyExists
		CreateProfile(ctx context.Context, profile *models.UserProfile) error
		// UpdateProfile - обновление профиля пользователя
		//
		// errors: models.ErrAlreadyExists,  models.ErrNotFound
		UpdateProfile(ctx context.Context, id types.UserID, f users_models.ProfileUpdateFields) (*models.UserProfile, error)
		// SearchByNickname - поиск профилей пользователей по части имени
		SearchByNickname(ctx context.Context, query string, limit uint32) ([]*models.UserProfile, error)
		// GetProfiles - поиск профилей пользователей
		GetProfiles(ctx context.Context, selector users_models.ProfileSelector) ([]*models.UserProfile, error)
	}

	UserIDProvider interface {
		GetUserIDFromIncomingContext(ctx context.Context) (string, error)
	}
)

type Deps struct {
	ProfilesRepository ProfileRepository
	UserIDProvider
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
	Deps
}

func NewUsecase(d Deps) *usecase {
	return &usecase{Deps: d}
}
