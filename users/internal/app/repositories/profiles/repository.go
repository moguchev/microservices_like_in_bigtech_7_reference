package profiles

import (
	"users/internal/app/usecases/users"
)

type Repository struct{}

var (
	_ users.ProfileRepository = (*Repository)(nil)
)

func NewRepository() *Repository {
	return &Repository{}
}
