package fileSystemRepo

import (
	"context"
	"errors"
	"file-service/internal/domain"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func mapRepoError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {

		case "23505":
			return domain.ErrConflict

		case "23503":
			return domain.ErrInvalidParent

		case "23514":
			return domain.ErrInvalidState
		}
	}

	return fmt.Errorf("repository error: %w", err)

}
