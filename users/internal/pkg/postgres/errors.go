package postgres

import (
	"errors"
	"fmt"

	"users/internal/app/models"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func ConvertPGError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("%s: %w", err, models.ErrNotFound)
	}

	// https://github.com/jackc/pgx/wiki/Error-Handling
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return fmt.Errorf("%s: %w", pgErr.Message, models.ErrAlreadyExists)
		default:
			return err
		}
	}
	return err
}
