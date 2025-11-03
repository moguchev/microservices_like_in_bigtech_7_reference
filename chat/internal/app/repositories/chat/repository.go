package chat

import (
	"chat/internal/app/usecases/chat"
	"lib/postgres"
	"lib/postgres/transaction_manager"

	"github.com/Masterminds/squirrel"
)

var (
	_ chat.ChatRepository = (*Repository)(nil)
)

type Repository struct {
	qb squirrel.StatementBuilderType
	db postgres.QueryEngineProvider
	tm *transaction_manager.TransactionManager
}

func NewRepository(db postgres.QueryEngineProvider, tm *transaction_manager.TransactionManager) *Repository {
	return &Repository{
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db: db,
		tm: tm,
	}
}
