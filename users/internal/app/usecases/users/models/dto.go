package models

import (
	lib_types "lib/types"
	"users/internal/app/models/types"
)

type (
	// CreateProfileInfo информация для создания профиля пользователя
	CreateProfileInfo struct {
		UserID    types.UserID // required
		Name      string       // required
		Email     string       // required
		Bio       string       // optional
		AvatarURL string       // optional
	}

	// ProfileUpdateFields поля по обновлению профиля пользователя
	ProfileUpdateFields struct {
		Name      lib_types.Field[string]
		Bio       lib_types.Field[string]
		AvatarURL lib_types.Field[string]
	}

	// ProfileSelector структура для подбора и поиска профилей пользователей
	ProfileSelector struct {
		// Тип поиска (OR/AND)
		SelectorType

		ID   lib_types.Field[types.UserID]
		Name lib_types.Field[string]
	}
)

// SelectorType тип селектора
type SelectorType int

const (
	SelectorTypeOR SelectorType = iota
	SelectorTypeAND
)
