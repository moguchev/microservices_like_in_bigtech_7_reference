package profiles

import (
	"database/sql"
	"time"

	"users/internal/app/models"
	"users/internal/app/models/types"
)

func toModel(r *profileRow) *models.UserProfile {
	if r == nil {
		return nil
	}

	return &models.UserProfile{
		ID:        types.UserID(r.ID),
		Email:     r.Email,
		Name:      r.Nickname,
		Bio:       r.Bio.String,
		AvatarURL: r.AvatarURL.String,
	}
}

func fromModel(m *models.UserProfile) profileRow {
	if m == nil {
		return profileRow{}
	}

	return profileRow{
		ID:        m.ID.String(),
		Email:     m.Email,
		Nickname:  m.Name,
		Bio:       sql.NullString{String: m.Bio, Valid: m.Bio != ""},
		AvatarURL: sql.NullString{String: m.AvatarURL, Valid: m.AvatarURL != ""},
		CreatedAt: time.Now().UTC(),
	}
}
