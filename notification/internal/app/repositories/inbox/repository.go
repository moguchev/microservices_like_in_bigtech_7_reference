package inbox

import (
	"lib/postgres"

	"github.com/Masterminds/squirrel"
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
