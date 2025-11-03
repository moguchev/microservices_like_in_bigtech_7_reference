package profiles

import (
	"lib/postgres"
	"users/internal/app/usecases/users"

	"github.com/Masterminds/squirrel"
)

var _ users.ProfileRepository = (*Repository)(nil)

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
