package social

import (
	"social/internal/app/usecases/social"
)

type Repository struct{}

var (
	_ social.SocialRepository = (*Repository)(nil)
)

func NewRepository() *Repository {
	return &Repository{}
}
