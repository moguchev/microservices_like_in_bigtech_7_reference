package profiles

import (
	"database/sql"
	"time"
)

const tableProfiles = "public.user_profiles"

const (
	columnID        = "id"
	columnEmail     = "email"
	columnNickname  = "nickname"
	columnBio       = "bio"
	columnAvatarURL = "avatar_url"
	columnCreatedAt = "created_at"
)

type profileRow struct {
	ID        string         `db:"id"`         // Уникальный идентификатор пользователя
	Email     string         `db:"email"`      // Email пользователя
	Nickname  string         `db:"nickname"`   // Имя пользователя
	Bio       sql.NullString `db:"bio"`        // Био пользователя (nullable)
	AvatarURL sql.NullString `db:"avatar_url"` // URL аватарки (nullable)
	CreatedAt time.Time      `db:"created_at"` // Когда создан профиль
}

var tableProfilesColumns = []string{
	columnID,
	columnEmail,
	columnNickname,
	columnBio,
	columnAvatarURL,
	columnCreatedAt,
}

func (p *profileRow) mapFields() map[string]any {
	return map[string]any{
		columnID:        p.ID,
		columnEmail:     p.Email,
		columnNickname:  p.Nickname,
		columnBio:       p.Bio,
		columnAvatarURL: p.AvatarURL,
		columnCreatedAt: p.CreatedAt,
	}
}

func (p *profileRow) Values(columns ...string) []any {
	mapFields := p.mapFields()
	values := make([]any, 0, len(columns))
	for i := range columns {
		if v, ok := mapFields[columns[i]]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}
