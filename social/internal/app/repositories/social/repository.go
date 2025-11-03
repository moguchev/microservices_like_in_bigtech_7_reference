package social

import (
	"lib/postgres"
	"social/internal/app/usecases/social"

	"github.com/Masterminds/squirrel"
)

var (
	_ social.SocialRepository = (*Repository)(nil)
)

type Repository struct {
	db postgres.QueryEngineProvider
	qb squirrel.StatementBuilderType
}

func NewRepository(p postgres.QueryEngineProvider) *Repository {
	return &Repository{
		db: p,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
